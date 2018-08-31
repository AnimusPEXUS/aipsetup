package buildercollection

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"text/template"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["wpa_supplicant"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_wpa_supplicant(bs), nil
	}
}

type Builder_wpa_supplicant struct {
	*Builder_std

	src_dir_p_sep string
}

func NewBuilder_wpa_supplicant(bs basictypes.BuildingSiteCtlI) *Builder_wpa_supplicant {

	self := new(Builder_wpa_supplicant)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.src_dir_p_sep = path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "wpa_supplicant")

	self.EditConfigureDirCB = func(log *logger.Logger, ret string) (string, error) {
		return self.src_dir_p_sep, nil
	}

	self.EditConfigureWorkingDirCB = self.EditConfigureDirCB

	self.EditDistributeArgsCB = self.EditDistributeArgs

	return self
}

func (self *Builder_wpa_supplicant) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("build")

	ret.ReplaceShort("configure", self.BuilderActionConfigure)

	return ret, nil
}

func (self *Builder_wpa_supplicant) BuilderActionConfigure(log *logger.Logger) error {

	install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return err
	}

	cfg_path := path.Join(self.src_dir_p_sep, ".config")
	def_cfg_path := path.Join(self.src_dir_p_sep, "defconfig")

	err = filetools.CopyWithInfo(def_cfg_path, cfg_path, log)
	if err != nil {
		return err
	}

	options, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateAutotoolsAllOptionsMap()
	if err != nil {
		return err
	}

	options_str := ""
	for k, v := range options {
		options_str += fmt.Sprintf("%s=%s\n", k, v)
	}

	config_tpl, err := template.New("config").Parse(`
CONFIG_BACKEND=file
CONFIG_CTRL_IFACE=y
CONFIG_DEBUG_FILE=y
CONFIG_DEBUG_SYSLOG=y
CONFIG_DEBUG_SYSLOG_FACILITY=LOG_DAEMON
CONFIG_DRIVER_NL80211=y
CONFIG_DRIVER_WEXT=y
CONFIG_DRIVER_WIRED=y
CONFIG_EAP_GTC=y
CONFIG_EAP_LEAP=y
CONFIG_EAP_MD5=y
CONFIG_EAP_MSCHAPV2=y
CONFIG_EAP_OTP=y
CONFIG_EAP_PEAP=y
CONFIG_EAP_TLS=y
CONFIG_EAP_TTLS=y
CONFIG_IEEE8021X_EAPOL=y
CONFIG_IPV6=y
CONFIG_LIBNL32=y
CONFIG_PEERKEY=y
CONFIG_PKCS12=y
CONFIG_READLINE=y
CONFIG_SMARTCARD=y
CONFIG_WPS=y
CFLAGS += -I{{.Hmd}}/include/libnl3
{{.Generated}}
`)
	if err != nil {
		return err
	}

	b := &bytes.Buffer{}

	err = config_tpl.Execute(
		b,
		struct {
			Hmd       string
			Generated string
		}{
			Hmd:       install_prefix,
			Generated: options_str,
		},
	)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(cfg_path)
	if err != nil {
		return err
	}

	data_lines := strings.Split(string(data), "\n")

	data_lines = append(data_lines, strings.Split(b.String(), "\n")...)

	data = []byte(strings.Join(data_lines, "\n"))

	err = ioutil.WriteFile(cfg_path, data, 0700)
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_wpa_supplicant) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"all",
			"install",
			"LIBDIR=" + path.Join(install_prefix, "lib"),
			"BINDIR=" + path.Join(install_prefix, "bin"),
			"PN531_PATH=" + path.Join(install_prefix, "src", "nfc"),
		}...,
	)
	return ret, nil
}

//    def builder_action_copy_manpages(self, called_as, log):
//        log.info("Copying manuals")

//        src_dir_p_sep = self.custom_data['src_dir_p_sep']

//        os.makedirs(wayround_i2p.utils.path.join(
//            self.get_dst_host_dir(), 'man', 'man8')
//            )
//        os.makedirs(wayround_i2p.utils.path.join(
//            self.get_dst_host_dir(), 'man', 'man5')
//            )

//        m8 = glob.glob(
//            wayround_i2p.utils.path.join(
//                src_dir_p_sep,
//                'doc',
//                'docbook',
//                '*.8')
//            )
//        m5 = glob.glob(
//            wayround_i2p.utils.path.join(
//                src_dir_p_sep,
//                'doc',
//                'docbook',
//                '*.5')
//            )

//        for i in m8:
//            bn = os.path.basename(i)
//            shutil.copyfile(
//                i,
//                wayround_i2p.utils.path.join(
//                    self.calculate_dst_install_prefix(),
//                    'man',
//                    'man8',
//                    bn
//                    )
//                )
//            log.info("    {}".format(i))

//        for i in m5:
//            bn = os.path.basename(i)
//            shutil.copyfile(
//                i,
//                wayround_i2p.utils.path.join(
//                    self.calculate_dst_install_prefix(),
//                    'man',
//                    'man5',
//                    bn)
//                )
//            log.info("    {}".format(i))
//        return 0
