package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"

	"github.com/AnimusPEXUS/aipsetup"

	"github.com/AnimusPEXUS/aipsetup/cmd/aipsetup4/templates"
)

type InfoServerConfig struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Prefix  string `yaml:"prefix"`
	InfoDir string `yaml:"info_dir"`
}

type InfoServer struct {
	path     string
	host     string
	port     int
	prefix   string //TODO: decide if this field is needed
	info_dir string

	srv     *http.Server
	handler *mux.Router
}

func NewInfoServerConfig() *InfoServerConfig {
	ret := &InfoServerConfig{
		Host:    "localhost",
		Port:    8080,
		Prefix:  "",
		InfoDir: "info",
	}
	return ret
}

func ReworkInfoDir(cfg_filename, info_dir string) (string, error) {

	var (
		ret string
		err error
	)

	if filepath.IsAbs(info_dir) {
		return info_dir, nil
	}

	cfg_filename, err = filepath.Abs(cfg_filename)

	if err != nil {
		return "", err
	}

	cfg_filename_dir := filepath.Dir(cfg_filename)

	ret, err = filepath.Abs(filepath.Join(cfg_filename_dir, info_dir))

	if err != nil {
		return "", err
	}

	return ret, nil
}

func (self *InfoServerConfig) YAMLString() (string, error) {
	ret, err := yaml.Marshal(self)
	return string(ret), err
}

func (self *InfoServerConfig) LoadFromFile(filename string) error {

	yaml_file, err := ioutil.ReadFile(filename)

	if err == nil {
		if yaml.Unmarshal(yaml_file, self) != nil {
			return err
		}
	}

	return nil
}

func NewInfoServer(path string) (*InfoServer, error) {

	var err error

	ret := new(InfoServer)

	ret.path = path

	cfg_file := filepath.Join(path, "aipsetup_info_server.cfg.yaml")

	cfg := NewInfoServerConfig()

	err = cfg.LoadFromFile(cfg_file)

	if err != nil {
		return nil, err
	}

	ret.host = cfg.Host
	ret.port = cfg.Port
	ret.prefix = cfg.Prefix

	ret.info_dir, err = ReworkInfoDir(cfg_file, cfg.InfoDir)
	if err != nil {
		return nil, err
	}

	ret.handler = mux.NewRouter()

	// ret.handler.PathPrefix(ret.prefix)

	ret.handler.HandleFunc("/", ret.RenderIndexPage).Methods("GET")

	ret.handler.HandleFunc("/all", ret.RenderAllPage).Methods("GET")

	ret.handler.HandleFunc("/goto", ret.RenderGoToRedirectPage).Methods("GET")

	ret.handler.HandleFunc("/json/{name}", ret.RenderInfoJSON).Methods("GET")
	ret.handler.HandleFunc("/yaml/{name}", ret.RenderInfoYAML).Methods("GET")
	ret.handler.HandleFunc("/text/{name}", ret.RenderInfoText).Methods("GET")

	ret.handler.HandleFunc("/css.css", ret.RenderCSS).Methods("GET")

	ret.srv = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", ret.host, ret.port),
		Handler:        ret.handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return ret, err
}

func (self *InfoServer) infoFileName(name string) string {
	return filepath.Join(
		self.info_dir,
		fmt.Sprintf("%s.json", name),
	)
}

func (self *InfoServer) IsExistsInfo(name string) bool {
	file, err := os.Open(self.infoFileName(name))
	if err != nil {
		return false
	}
	file.Close()
	return true
}

func (self *InfoServer) LoadInfo(name string) (
	*aipsetup.CompletePackageInfo,
	error,
) {

	var (
		err  error
		text []byte
	)

	filename := self.infoFileName(name)

	text, err = ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	ret := &aipsetup.CompletePackageInfo{}

	err = json.Unmarshal(text, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *InfoServer) ListInfo() (
	[]string,
	error,
) {

	var (
		err error
		dir []os.FileInfo
		ret []string
	)

	dir, err = ioutil.ReadDir(self.info_dir)
	if err != nil {
		return []string{}, err
	}

	for _, i := range dir {

		name := i.Name()

		if !i.IsDir() && strings.HasSuffix(name, ".json") {
			ret = append(ret, name[:len(name)-len(".json")])
		}
	}

	return ret, nil
}

func (self *InfoServer) requestGetName(req *http.Request) string {
	splitted_path := strings.Split(req.URL.Path, "/")
	name := splitted_path[len(splitted_path)-1]
	return name
}

func (self *InfoServer) printInfoNotFoundError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain;charset=UTF8")
	w.WriteHeader(404)
	w.Write(([]byte)("404. not found or some other error\n"))
	return
}

func (self *InfoServer) printInfoMarshalError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain;charset=UTF8")
	w.WriteHeader(500)
	w.Write(
		([]byte)("500. info marshal error. " +
			"If You see this error for too long, please file a bugreport\n"),
	)
	return
}

func (self *InfoServer) RenderCSS(
	w http.ResponseWriter,
	req *http.Request,
) {

	b := bytes.Buffer{}
	b.Write([]byte(templates.OveralCSS()))

	w.Header().Set("Content-Type", "text/css;charset=UTF8")
	w.WriteHeader(200)

	b.WriteTo(w)
}

func (self *InfoServer) RenderInfoJSON(
	w http.ResponseWriter,
	req *http.Request,
) {
	var (
		info *aipsetup.CompletePackageInfo
		data []byte
		err  error
	)

	name := self.requestGetName(req)

	info, err = self.LoadInfo(name)
	if err != nil || info == nil {
		fmt.Println(err.Error())
		self.printInfoNotFoundError(w)
		return
	}

	data, err = json.Marshal(info)
	if err != nil {
		self.printInfoMarshalError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF8")
	w.WriteHeader(200)
	w.Write(data)
}

func (self *InfoServer) RenderInfoYAML(
	w http.ResponseWriter, req *http.Request,
) {
	var (
		info *aipsetup.CompletePackageInfo
		data []byte
		err  error
	)

	name := self.requestGetName(req)

	info, err = self.LoadInfo(name)
	if err != nil || info == nil {
		self.printInfoNotFoundError(w)
		return
	}

	data, err = yaml.Marshal(info)
	if err != nil {
		self.printInfoMarshalError(w)
		return
	}

	// TODO: YAML doesn't have own media tipe at time of this writting,
	//       so i think text/plain will go
	w.Header().Set("Content-Type", "text/plain;charset=UTF8")
	w.WriteHeader(200)
	w.Write(data)
}

func (self *InfoServer) RenderInfoText(
	w http.ResponseWriter,
	req *http.Request,
) {

	var (
		info *aipsetup.CompletePackageInfo
		err  error
	)

	name := self.requestGetName(req)

	info, err = self.LoadInfo(name)
	if err != nil || info == nil {
		self.printInfoNotFoundError(w)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF8")
	w.WriteHeader(200)

	//templates.InfoServerIndexPage()

	return
}

func (self *InfoServer) RenderIndexPage(
	w http.ResponseWriter,
	req *http.Request,
) {

	b := bytes.Buffer{}
	b.Write([]byte(templates.InfoServerIndexPage()))

	w.Header().Set("Content-Type", "text/html;charset=UTF8")
	w.WriteHeader(200)

	b.WriteTo(w)

	return
}

func (self *InfoServer) RenderAllPage(
	w http.ResponseWriter,
	req *http.Request,
) {

	lst, err := self.ListInfo()

	if err != nil {
		w.Header().Set("Content-Type", "text/plain;charset=UTF8")
		w.WriteHeader(500)
		w.Write([]byte("server error: Can't get Info List"))
		return
	}

	b := bytes.Buffer{}
	b.Write([]byte(templates.InfoServerAllInfoPage(lst)))

	w.Header().Set("Content-Type", "text/html;charset=UTF8")
	w.WriteHeader(200)

	b.WriteTo(w)

	return
}

func (self *InfoServer) RenderGoToRedirectPage(
	w http.ResponseWriter,
	req *http.Request,
) {
	var (
		name   string
		format string
	)
	{
		v := req.URL.Query()
		format = v.Get("format")
		name = v.Get("name")
	}

	if name == "" {
		w.Header().Set("Content-Type", "text/plain;charset=UTF8")
		w.WriteHeader(400)
		w.Write([]byte("name field value is required."))
	}

	if format == "" {
		format = "yaml"
	}

	w.Header().Set("Content-Type", "text/plain;charset=UTF8")
	w.Header().Set("Location", fmt.Sprintf("/%s/%s", format, name))
	w.WriteHeader(302)

}

func (self *InfoServer) Run() error {
	return self.srv.ListenAndServe()
}
