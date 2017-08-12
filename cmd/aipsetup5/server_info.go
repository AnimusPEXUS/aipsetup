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

	"github.com/go-ini/ini"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"

	"github.com/AnimusPEXUS/aipsetup"

	"github.com/AnimusPEXUS/aipsetup/cmd/aipsetup5/templates"
)

type InfoServer struct {
	//path    string
	host    string
	port    int
	prefix  string //TODO: decide if this field is needed
	infodir string

	srv     *http.Server
	handler *mux.Router
}

var INFO_SERVER_CONFIG = []byte(`
[main]
host = localhost
port = 8080
prefix =
infodir = info
`)

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

func NewInfoServer() (*InfoServer, error) {

	var err error

	ret := new(InfoServer)

	cfg_path := filepath.Join("/etc", "aipsetup5.info_server.ini")

	cfg, err := ini.Load(INFO_SERVER_CONFIG)
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(cfg_path)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if !os.IsNotExist(err) {
		cfg.Append(cfg_path)
	} else {
		fmt.Println("[w] config file not found")
	}

	sect, err := cfg.GetSection("main")
	if err != nil {
		return nil, err
	}

	ret.host = sect.Key("host").MustString("localhost")
	ret.port = sect.Key("port").MustInt(8080)
	ret.prefix = sect.Key("prefix").MustString("")
	ret.infodir = sect.Key("infodir").MustString("info")

	ret.infodir, err = ReworkInfoDir(cfg_path, ret.infodir)
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
		self.infodir,
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

	dir, err = ioutil.ReadDir(self.infodir)
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
	w.Write(([]byte)("404. info not found or some other error\n"))
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
		// w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF8")
	w.WriteHeader(200)

	w.Write([]byte(templates.InfoServerHtmlInfoPage(name, info)))
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
