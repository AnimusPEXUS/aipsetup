package infoeditor

import "github.com/AnimusPEXUS/aipsetup/basictypes"

func (self *InfoEditor) ApplyBuilders(index map[string]*basictypes.PackageInfo) error {
	for k, v := range index {
		if k == "3DLDF" {
			v.BuilderName = "std"
		}

		if k == "A3Com" {
			v.BuilderName = "std"
		}

		if k == "AfterStep" {
			v.BuilderName = "None"
		}

		if k == "Alchemy" {
			v.BuilderName = ""
		}

		if k == "Archive-Zip" {
			v.BuilderName = "perl_mod"
		}

		if k == "Atlas-C++" {
			v.BuilderName = "std"
		}

		if k == "Authen-PAM" {
			v.BuilderName = "std"
		}

		if k == "BitTornado" {
			v.BuilderName = ""
		}

		if k == "Botan" {
			v.BuilderName = "botan"
		}

		if k == "CAPA" {
			v.BuilderName = "std"
		}

		if k == "CEGUI" {
			v.BuilderName = "std"
		}

		if k == "CSSC" {
			v.BuilderName = "std"
		}

		if k == "CVS+Linux-PAM" {
			v.BuilderName = "std"
		}

		if k == "CVS+SimplePAMApps" {
			v.BuilderName = "std"
		}

		if k == "Cg" {
			v.BuilderName = "Cg"
		}

		if k == "Coin" {
			v.BuilderName = ""
		}

		if k == "ConsoleKit" {
			v.BuilderName = "ConsoleKit"
		}

		if k == "Cython" {
			v.BuilderName = ""
		}

		if k == "DBDesigner" {
			v.BuilderName = ""
		}

		if k == "DevIL" {
			v.BuilderName = "std"
		}

		if k == "Device-Cdio-v" {
			v.BuilderName = "std"
		}

		if k == "DeviceKit-disks" {
			v.BuilderName = "std"
		}

		if k == "DeviceKit-power" {
			v.BuilderName = "std"
		}

		if k == "DeviceKit" {
			v.BuilderName = "std"
		}

		if k == "Django" {
			v.BuilderName = ""
		}

		if k == "ETL" {
			v.BuilderName = ""
		}

		if k == "EchoVNC" {
			v.BuilderName = ""
		}

		if k == "FreeImage" {
			v.BuilderName = "FreeImage"
		}

		if k == "FreeRDP" {
			v.BuilderName = ""
		}

		if k == "GConf-dbus" {
			v.BuilderName = "std"
		}

		if k == "GConf" {
			v.BuilderName = "gconf"
		}

		if k == "GSAPI" {
			v.BuilderName = "std"
		}

		if k == "Gens" {
			v.BuilderName = ""
		}

		if k == "Genshi" {
			v.BuilderName = "None"
		}

		if k == "GeoIP" {
			v.BuilderName = "std"
		}

		if k == "GnomeHello" {
			v.BuilderName = "std"
		}

		if k == "Gtk2-Ex-Utils" {
			v.BuilderName = "perl_mod"
		}

		if k == "Gtk2" {
			v.BuilderName = "perl_mod"
		}

		if k == "Guppi" {
			v.BuilderName = "std"
		}

		if k == "HTML-Parser" {
			v.BuilderName = "perl_mod"
		}

		if k == "InstantVNC" {
			v.BuilderName = ""
		}

		if k == "IronPython" {
			v.BuilderName = ""
		}

		if k == "Jython" {
			v.BuilderName = ""
		}

		if k == "LDasm" {
			v.BuilderName = ""
		}

		if k == "LPRng" {
			v.BuilderName = ""
		}

		if k == "LibVNCServer" {
			v.BuilderName = "std"
		}

		if k == "Linux-PAM" {
			v.BuilderName = "linux_pam"
		}

		if k == "Mail-SpamAssassin" {
			v.BuilderName = "perl_mod"
		}

		if k == "Maverik" {
			v.BuilderName = "std"
		}

		if k == "MaverikDemos" {
			v.BuilderName = "std"
		}

		if k == "Mesa" {
			v.BuilderName = "std"
		}

		if k == "MesaDemos" {
			v.BuilderName = "std"
		}

		if k == "MesaGLUT" {
			v.BuilderName = "mesaglut"
		}

		if k == "MesaLib" {
			v.BuilderName = "mesalib"
		}

		if k == "ModemManager" {
			v.BuilderName = "ModemManager"
		}

		if k == "Mon" {
			v.BuilderName = "std"
		}

		if k == "Net-DNS" {
			v.BuilderName = "perl_mod"
		}

		if k == "NetAddr-IP" {
			v.BuilderName = "perl_mod"
		}

		if k == "NetworkManager-iodine" {
			v.BuilderName = "std"
		}

		if k == "NetworkManager-openconnect" {
			v.BuilderName = "std"
		}

		if k == "NetworkManager-openswan" {
			v.BuilderName = "std"
		}

		if k == "NetworkManager-openvpn" {
			v.BuilderName = "std"
		}

		if k == "NetworkManager-pptp" {
			v.BuilderName = "std"
		}

		if k == "NetworkManager-vpnc" {
			v.BuilderName = "std"
		}

		if k == "NetworkManager" {
			v.BuilderName = "NetworkManager"
		}

		if k == "ORBit" {
			v.BuilderName = "std"
		}

		if k == "ORBit2" {
			v.BuilderName = "orbit2"
		}

		if k == "OpenSP" {
			v.BuilderName = ""
		}

		if k == "PPSkit" {
			v.BuilderName = "std"
		}

		if k == "Parallel-ForkManager" {
			v.BuilderName = "perl_mod"
		}

		if k == "PolicyKit-gnome" {
			v.BuilderName = "std"
		}

		if k == "PolicyKit" {
			v.BuilderName = "std"
		}

		if k == "PyChecker" {
			v.BuilderName = ""
		}

		if k == "Pyrex" {
			v.BuilderName = ""
		}

		if k == "Python2" {
			v.BuilderName = "python"
		}

		if k == "Python3" {
			v.BuilderName = "python"
		}

		if k == "PythonCAD" {
			v.BuilderName = ""
		}

		if k == "QScintilla" {
			v.BuilderName = ""
		}

		if k == "QScintilla2" {
			v.BuilderName = "qscintilla"
		}

		if k == "RCS+Linux-PAM" {
			v.BuilderName = "std"
		}

		if k == "RCS+SimplePAMApps" {
			v.BuilderName = "std"
		}

		if k == "SDL" {
			v.BuilderName = "SDL"
		}

		if k == "SDL2" {
			v.BuilderName = "SDL2"
		}

		if k == "SDL_image" {
			v.BuilderName = "std"
		}

		if k == "SDL_mixer" {
			v.BuilderName = "std"
		}

		if k == "SDL_net" {
			v.BuilderName = "std"
		}

		if k == "SDL_sound" {
			v.BuilderName = "std"
		}

		if k == "SDL_ttf" {
			v.BuilderName = "std"
		}

		if k == "SQLAlchemy" {
			v.BuilderName = ""
		}

		if k == "Search-Xapian" {
			v.BuilderName = ""
		}

		if k == "SimplePAMApps" {
			v.BuilderName = "std"
		}

		if k == "SoQt" {
			v.BuilderName = ""
		}

		if k == "TeXLive" {
			v.BuilderName = ""
		}

		if k == "Terminal" {
			v.BuilderName = ""
		}

		if k == "Thunar" {
			v.BuilderName = ""
		}

		if k == "TiMidity++" {
			v.BuilderName = ""
		}

		if k == "Twisted" {
			v.BuilderName = ""
		}

		if k == "Type" {
			v.BuilderName = "std"
		}

		if k == "UPower" {
			v.BuilderName = "std"
		}

		if k == "URI" {
			v.BuilderName = "perl_mod"
		}

		if k == "UltraVNC" {
			v.BuilderName = ""
		}

		if k == "WFUT" {
			v.BuilderName = "std"
		}

		if k == "WindowMaker" {
			v.BuilderName = "std"
		}

		if k == "XML-Namespace" {
			v.BuilderName = "perl_mod"
		}

		if k == "XML-NamespaceSupport" {
			v.BuilderName = "perl_mod"
		}

		if k == "XML-Parser" {
			v.BuilderName = "perl_mod"
		}

		if k == "XML-SAX-Base" {
			v.BuilderName = "perl_mod"
		}

		if k == "XML-SAX" {
			v.BuilderName = "perl_mod"
		}

		if k == "XML-Simple" {
			v.BuilderName = "perl_mod"
		}

		if k == "Xnee" {
			v.BuilderName = "std"
		}

		if k == "Xvnc" {
			v.BuilderName = ""
		}

		if k == "YafRay" {
			v.BuilderName = ""
		}

		if k == "Zope" {
			v.BuilderName = ""
		}

		if k == "a2ps" {
			v.BuilderName = "std"
		}

		if k == "a52dec" {
			v.BuilderName = "a52dec"
		}

		if k == "abi-tools" {
			v.BuilderName = "std"
		}

		if k == "abi" {
			v.BuilderName = "std"
		}

		if k == "abispell" {
			v.BuilderName = "std"
		}

		if k == "abiword" {
			v.BuilderName = "std"
		}

		if k == "abuse" {
			v.BuilderName = "std"
		}

		if k == "accerciser" {
			v.BuilderName = "std"
		}

		if k == "accounts-qt" {
			v.BuilderName = ""
		}

		if k == "accountsdialog" {
			v.BuilderName = "std"
		}

		if k == "accountsservice" {
			v.BuilderName = "std"
		}

		if k == "acct" {
			v.BuilderName = "std"
		}

		if k == "acl" {
			v.BuilderName = "xfs"
		}

		if k == "acm" {
			v.BuilderName = "std"
		}

		if k == "acme" {
			v.BuilderName = "std"
		}

		if k == "acpica-unix" {
			v.BuilderName = ""
		}

		if k == "acpica-unix2" {
			v.BuilderName = ""
		}

		if k == "acpid" {
			v.BuilderName = "acpid"
		}

		if k == "activation" {
			v.BuilderName = "std"
		}

		if k == "adns" {
			v.BuilderName = "std"
		}

		if k == "adwaita-icon-theme" {
			v.BuilderName = "std"
		}

		if k == "aegisub" {
			v.BuilderName = ""
		}

		if k == "agg" {
			v.BuilderName = "std"
		}

		if k == "aisleriot" {
			v.BuilderName = "std"
		}

		if k == "akonadi" {
			v.BuilderName = ""
		}

		if k == "alacarte" {
			v.BuilderName = "std"
		}

		if k == "alive" {
			v.BuilderName = "std"
		}

		if k == "allegro" {
			v.BuilderName = "std_cmake"
		}

		if k == "alleyoop" {
			v.BuilderName = "std"
		}

		if k == "almanah" {
			v.BuilderName = "std"
		}

		if k == "alsa-firmware" {
			v.BuilderName = "std"
		}

		if k == "alsa-lib" {
			v.BuilderName = "alsa_lib"
		}

		if k == "alsa-oss" {
			v.BuilderName = "std"
		}

		if k == "alsa-plugins" {
			v.BuilderName = "std"
		}

		if k == "alsa-tools" {
			v.BuilderName = "std"
		}

		if k == "alsa-utils" {
			v.BuilderName = "std"
		}

		if k == "ammonite" {
			v.BuilderName = "std"
		}

		if k == "ams" {
			v.BuilderName = ""
		}

		if k == "analitza" {
			v.BuilderName = ""
		}

		if k == "anjal" {
			v.BuilderName = "std"
		}

		if k == "anjuta-extras" {
			v.BuilderName = "std"
		}

		if k == "anjuta" {
			v.BuilderName = "std"
		}

		if k == "anubis" {
			v.BuilderName = "std"
		}

		if k == "apache-ant" {
			v.BuilderName = "apache_ant"
		}

		if k == "apache-couchdb" {
			v.BuilderName = "couchdb"
		}

		if k == "apl" {
			v.BuilderName = "std"
		}

		if k == "apparmor" {
			v.BuilderName = "std"
		}

		if k == "appdata-tools" {
			v.BuilderName = "std"
		}

		if k == "applewmproto" {
			v.BuilderName = "std"
		}

		if k == "appres" {
			v.BuilderName = "std"
		}

		if k == "apr-util" {
			v.BuilderName = "apr_util"
		}

		if k == "apr" {
			v.BuilderName = "apr"
		}

		if k == "aravis" {
			v.BuilderName = "std"
		}

		if k == "archimedes" {
			v.BuilderName = "std"
		}

		if k == "ardour" {
			v.BuilderName = ""
		}

		if k == "aria2" {
			v.BuilderName = "aria2"
		}

		if k == "aris" {
			v.BuilderName = "std"
		}

		if k == "ark" {
			v.BuilderName = ""
		}

		if k == "asciidoc" {
			v.BuilderName = "std"
		}

		if k == "aspell-af" {
			v.BuilderName = "std"
		}

		if k == "aspell-bg" {
			v.BuilderName = "std"
		}

		if k == "aspell-br" {
			v.BuilderName = "std"
		}

		if k == "aspell-ca" {
			v.BuilderName = "std"
		}

		if k == "aspell-cs" {
			v.BuilderName = "std"
		}

		if k == "aspell-cy" {
			v.BuilderName = "std"
		}

		if k == "aspell-da" {
			v.BuilderName = "std"
		}

		if k == "aspell-de" {
			v.BuilderName = "std"
		}

		if k == "aspell-el" {
			v.BuilderName = "std"
		}

		if k == "aspell-en" {
			v.BuilderName = "std"
		}

		if k == "aspell-eo" {
			v.BuilderName = "std"
		}

		if k == "aspell-es" {
			v.BuilderName = "std"
		}

		if k == "aspell-fo" {
			v.BuilderName = "std"
		}

		if k == "aspell-fr" {
			v.BuilderName = "std"
		}

		if k == "aspell-ga" {
			v.BuilderName = "std"
		}

		if k == "aspell-gd" {
			v.BuilderName = "std"
		}

		if k == "aspell-gl" {
			v.BuilderName = "std"
		}

		if k == "aspell-gv" {
			v.BuilderName = "std"
		}

		if k == "aspell-hr" {
			v.BuilderName = "std"
		}

		if k == "aspell-ia" {
			v.BuilderName = "std"
		}

		if k == "aspell-id" {
			v.BuilderName = "std"
		}

		if k == "aspell-is" {
			v.BuilderName = "std"
		}

		if k == "aspell-it" {
			v.BuilderName = "std"
		}

		if k == "aspell-lang" {
			v.BuilderName = "std"
		}

		if k == "aspell-mi" {
			v.BuilderName = "std"
		}

		if k == "aspell-mk" {
			v.BuilderName = "std"
		}

		if k == "aspell-ms" {
			v.BuilderName = "std"
		}

		if k == "aspell-mt" {
			v.BuilderName = "std"
		}

		if k == "aspell-nb" {
			v.BuilderName = "std"
		}

		if k == "aspell-nl" {
			v.BuilderName = "std"
		}

		if k == "aspell-nn" {
			v.BuilderName = "std"
		}

		if k == "aspell-no" {
			v.BuilderName = "std"
		}

		if k == "aspell-pl" {
			v.BuilderName = "std"
		}

		if k == "aspell-pt" {
			v.BuilderName = "std"
		}

		if k == "aspell-ro" {
			v.BuilderName = "std"
		}

		if k == "aspell-ru" {
			v.BuilderName = "std"
		}

		if k == "aspell-rw" {
			v.BuilderName = "std"
		}

		if k == "aspell-sk" {
			v.BuilderName = "std"
		}

		if k == "aspell-sl" {
			v.BuilderName = "std"
		}

		if k == "aspell-sv" {
			v.BuilderName = "std"
		}

		if k == "aspell-sw" {
			v.BuilderName = "std"
		}

		if k == "aspell-tn" {
			v.BuilderName = "std"
		}

		if k == "aspell-tr" {
			v.BuilderName = "std"
		}

		if k == "aspell-uk" {
			v.BuilderName = "std"
		}

		if k == "aspell-wa" {
			v.BuilderName = "std"
		}

		if k == "aspell-zu" {
			v.BuilderName = "std"
		}

		if k == "aspell" {
			v.BuilderName = "std"
		}

		if k == "aspell5-be" {
			v.BuilderName = "std"
		}

		if k == "aspell5-bg" {
			v.BuilderName = "std"
		}

		if k == "aspell5-da" {
			v.BuilderName = "std"
		}

		if k == "aspell5-en" {
			v.BuilderName = "std"
		}

		if k == "aspell5-fo" {
			v.BuilderName = "std"
		}

		if k == "aspell5-ga" {
			v.BuilderName = "std"
		}

		if k == "aspell5-gd" {
			v.BuilderName = "std"
		}

		if k == "aspell5-hil" {
			v.BuilderName = "std"
		}

		if k == "aspell5-id" {
			v.BuilderName = "std"
		}

		if k == "aspell5-ku" {
			v.BuilderName = "std"
		}

		if k == "aspell5-mg" {
			v.BuilderName = "std"
		}

		if k == "aspell5-ny" {
			v.BuilderName = "std"
		}

		if k == "aspell5-pl" {
			v.BuilderName = "std"
		}

		if k == "aspell5-ro" {
			v.BuilderName = "std"
		}

		if k == "aspell5-ru" {
			v.BuilderName = "std"
		}

		if k == "aspell5-sc" {
			v.BuilderName = "std"
		}

		if k == "aspell5-tet" {
			v.BuilderName = "std"
		}

		if k == "aspell5-tk" {
			v.BuilderName = "std"
		}

		if k == "aspell5-tl" {
			v.BuilderName = "std"
		}

		if k == "aspell5-tn" {
			v.BuilderName = "std"
		}

		if k == "aspell6-am" {
			v.BuilderName = "std"
		}

		if k == "aspell6-ar" {
			v.BuilderName = "std"
		}

		if k == "aspell6-ast" {
			v.BuilderName = "std"
		}

		if k == "aspell6-az" {
			v.BuilderName = "std"
		}

		if k == "aspell6-bg" {
			v.BuilderName = "std"
		}

		if k == "aspell6-bn" {
			v.BuilderName = "std"
		}

		if k == "aspell6-ca" {
			v.BuilderName = "std"
		}

		if k == "aspell6-cs" {
			v.BuilderName = "std"
		}

		if k == "aspell6-csb" {
			v.BuilderName = "std"
		}

		if k == "aspell6-de-alt" {
			v.BuilderName = "std"
		}

		if k == "aspell6-de" {
			v.BuilderName = "std"
		}

		if k == "aspell6-en" {
			v.BuilderName = "std"
		}

		if k == "aspell6-eo" {
			v.BuilderName = "std"
		}

		if k == "aspell6-es" {
			v.BuilderName = "std"
		}

		if k == "aspell6-et" {
			v.BuilderName = "std"
		}

		if k == "aspell6-fa" {
			v.BuilderName = "std"
		}

		if k == "aspell6-fi" {
			v.BuilderName = "std"
		}

		if k == "aspell6-fy" {
			v.BuilderName = "std"
		}

		if k == "aspell6-gl" {
			v.BuilderName = "std"
		}

		if k == "aspell6-grc" {
			v.BuilderName = "std"
		}

		if k == "aspell6-gu" {
			v.BuilderName = "std"
		}

		if k == "aspell6-he" {
			v.BuilderName = "std"
		}

		if k == "aspell6-hi" {
			v.BuilderName = "std"
		}

		if k == "aspell6-hsb" {
			v.BuilderName = "std"
		}

		if k == "aspell6-hu" {
			v.BuilderName = "std"
		}

		if k == "aspell6-hus" {
			v.BuilderName = "std"
		}

		if k == "aspell6-hy" {
			v.BuilderName = "std"
		}

		if k == "aspell6-it" {
			v.BuilderName = "std"
		}

		if k == "aspell6-kn" {
			v.BuilderName = "std"
		}

		if k == "aspell6-ky" {
			v.BuilderName = "std"
		}

		if k == "aspell6-la" {
			v.BuilderName = "std"
		}

		if k == "aspell6-lt" {
			v.BuilderName = "std"
		}

		if k == "aspell6-lv" {
			v.BuilderName = "std"
		}

		if k == "aspell6-ml" {
			v.BuilderName = "std"
		}

		if k == "aspell6-mn" {
			v.BuilderName = "std"
		}

		if k == "aspell6-mr" {
			v.BuilderName = "std"
		}

		if k == "aspell6-nds" {
			v.BuilderName = "std"
		}

		if k == "aspell6-or" {
			v.BuilderName = "std"
		}

		if k == "aspell6-pa" {
			v.BuilderName = "std"
		}

		if k == "aspell6-pl" {
			v.BuilderName = "std"
		}

		if k == "aspell6-pt_BR" {
			v.BuilderName = "std"
		}

		if k == "aspell6-pt_PT" {
			v.BuilderName = "std"
		}

		if k == "aspell6-qq" {
			v.BuilderName = "std"
		}

		if k == "aspell6-qu" {
			v.BuilderName = "std"
		}

		if k == "aspell6-ru" {
			v.BuilderName = "std"
		}

		if k == "aspell6-sk" {
			v.BuilderName = "std"
		}

		if k == "aspell6-sr" {
			v.BuilderName = "std"
		}

		if k == "aspell6-ta" {
			v.BuilderName = "std"
		}

		if k == "aspell6-te" {
			v.BuilderName = "std"
		}

		if k == "aspell6-uk" {
			v.BuilderName = "std"
		}

		if k == "aspell6-uz" {
			v.BuilderName = "std"
		}

		if k == "aspell6-vi" {
			v.BuilderName = "std"
		}

		if k == "aspell6-yi" {
			v.BuilderName = "std"
		}

		if k == "assimp" {
			v.BuilderName = "std_cmake"
		}

		if k == "asterisk" {
			v.BuilderName = ""
		}

		if k == "astyle" {
			v.BuilderName = "astyle"
		}

		if k == "at-poke" {
			v.BuilderName = "std"
		}

		if k == "at-spi" {
			v.BuilderName = "std"
		}

		if k == "at-spi2-atk" {
			v.BuilderName = "std"
		}

		if k == "at-spi2-core" {
			v.BuilderName = "std"
		}

		if k == "atk" {
			v.BuilderName = "std"
		}

		if k == "atkmm" {
			v.BuilderName = "std"
		}

		if k == "atomix" {
			v.BuilderName = "std"
		}

		if k == "atrack" {
			v.BuilderName = ""
		}

		if k == "atril" {
			v.BuilderName = "std"
		}

		if k == "attica" {
			v.BuilderName = ""
		}

		if k == "attr" {
			v.BuilderName = "xfs"
		}

		if k == "auctex" {
			v.BuilderName = "std"
		}

		if k == "audacity" {
			v.BuilderName = "audacity"
		}

		if k == "audiofile" {
			v.BuilderName = "std"
		}

		if k == "autoconf-archive" {
			v.BuilderName = "std"
		}

		if k == "autoconf" {
			v.BuilderName = "std"
		}

		if k == "autoconf2.13" {
			v.BuilderName = "std"
		}

		if k == "autofs" {
			v.BuilderName = "std"
		}

		if k == "autofs4" {
			v.BuilderName = "std"
		}

		if k == "autogen" {
			v.BuilderName = "autogen"
		}

		if k == "automake" {
			v.BuilderName = "std"
		}

		if k == "automoc4" {
			v.BuilderName = ""
		}

		if k == "avahi" {
			v.BuilderName = "avahi"
		}

		if k == "avidemux" {
			v.BuilderName = ""
		}

		if k == "avl" {
			v.BuilderName = "std"
		}

		if k == "awesome" {
			v.BuilderName = "std_cmake"
		}

		if k == "babl" {
			v.BuilderName = "std"
		}

		if k == "bakery" {
			v.BuilderName = "std"
		}

		if k == "ballandpaddle" {
			v.BuilderName = "std"
		}

		if k == "balsa" {
			v.BuilderName = "std"
		}

		if k == "bandersnatch" {
			v.BuilderName = ""
		}

		if k == "banshee-1" {
			v.BuilderName = "std"
		}

		if k == "banshee" {
			v.BuilderName = "std"
		}

		if k == "banter" {
			v.BuilderName = "std"
		}

		if k == "baobab" {
			v.BuilderName = "std"
		}

		if k == "barcode" {
			v.BuilderName = "std"
		}

		if k == "bash-completion" {
			v.BuilderName = "std"
		}

		if k == "bash-doc" {
			v.BuilderName = "std"
		}

		if k == "bash" {
			v.BuilderName = "bash"
		}

		if k == "battfink" {
			v.BuilderName = "std"
		}

		if k == "bayonne" {
			v.BuilderName = "std"
		}

		if k == "bayonne2" {
			v.BuilderName = "std"
		}

		if k == "bc" {
			v.BuilderName = "bc"
		}

		if k == "bdflush" {
			v.BuilderName = "std"
		}

		if k == "bdftopcf" {
			v.BuilderName = "std"
		}

		if k == "beagle-xesam" {
			v.BuilderName = "std"
		}

		if k == "beagle" {
			v.BuilderName = "std"
		}

		if k == "beast" {
			v.BuilderName = ""
		}

		if k == "beforelight" {
			v.BuilderName = "std"
		}

		if k == "beye" {
			v.BuilderName = ""
		}

		if k == "biew" {
			v.BuilderName = "std"
		}

		if k == "bigboard" {
			v.BuilderName = "std"
		}

		if k == "bigreqsproto" {
			v.BuilderName = "std"
		}

		if k == "bijiben" {
			v.BuilderName = "std"
		}

		if k == "billreminder" {
			v.BuilderName = "std"
		}

		if k == "bind10" {
			v.BuilderName = "bind10"
		}

		if k == "bind9" {
			v.BuilderName = "std"
		}

		if k == "binutils" {
			v.BuilderName = "binutils"
		}

		if k == "bison" {
			v.BuilderName = "std"
		}

		if k == "bitcoin" {
			v.BuilderName = "bitcoin"
		}

		if k == "bitcoins" {
			v.BuilderName = ""
		}

		if k == "bitlbee" {
			v.BuilderName = "bitlbee"
		}

		if k == "bitmap" {
			v.BuilderName = "std"
		}

		if k == "blam" {
			v.BuilderName = "std"
		}

		if k == "blender" {
			v.BuilderName = "blender"
		}

		if k == "blink-qt" {
			v.BuilderName = ""
		}

		if k == "blinken" {
			v.BuilderName = ""
		}

		if k == "bluez-hcidump" {
			v.BuilderName = "std"
		}

		if k == "bluez" {
			v.BuilderName = "bluez"
		}

		if k == "bochs" {
			v.BuilderName = ""
		}

		if k == "bogofilter" {
			v.BuilderName = "std"
		}

		if k == "bokken" {
			v.BuilderName = "std"
		}

		if k == "bombermaze" {
			v.BuilderName = "std"
		}

		if k == "bonobo-activation" {
			v.BuilderName = "std"
		}

		if k == "bonobo-conf" {
			v.BuilderName = "std"
		}

		if k == "bonobo-config" {
			v.BuilderName = "std"
		}

		if k == "bonobo" {
			v.BuilderName = "std"
		}

		if k == "bool" {
			v.BuilderName = "std"
		}

		if k == "boost" {
			v.BuilderName = "boost"
		}

		if k == "bootp.monitor" {
			v.BuilderName = "std"
		}

		if k == "boswars" {
			v.BuilderName = ""
		}

		if k == "bouml" {
			v.BuilderName = ""
		}

		if k == "bpel2owfn" {
			v.BuilderName = "std"
		}

		if k == "brasero" {
			v.BuilderName = "std"
		}

		if k == "bridge-utils" {
			v.BuilderName = "std"
		}

		if k == "brlcad" {
			v.BuilderName = ""
		}

		if k == "bsdgames" {
			v.BuilderName = ""
		}

		if k == "bti" {
			v.BuilderName = "std"
		}

		if k == "btrfs-progs" {
			v.BuilderName = "std"
		}

		if k == "bug-buddy" {
			v.BuilderName = "std"
		}

		if k == "bugzilla" {
			v.BuilderName = ""
		}

		if k == "buildtree" {
			v.BuilderName = "std"
		}

		if k == "bullet" {
			v.BuilderName = "std_cmake"
		}

		if k == "bwbar" {
			v.BuilderName = "std"
		}

		if k == "bwidget" {
			v.BuilderName = ""
		}

		if k == "bwmonitor" {
			v.BuilderName = "std"
		}

		if k == "byzanz" {
			v.BuilderName = "std"
		}

		if k == "bzflag" {
			v.BuilderName = "std"
		}

		if k == "bzip2" {
			v.BuilderName = "bzip2"
		}

		if k == "bzr-explorer" {
			v.BuilderName = ""
		}

		if k == "bzr" {
			v.BuilderName = "std"
		}

		if k == "c-graph" {
			v.BuilderName = "std"
		}

		if k == "cabextract" {
			v.BuilderName = "std"
		}

		if k == "cairo-5c" {
			v.BuilderName = "std"
		}

		if k == "cairo-java" {
			v.BuilderName = "std"
		}

		if k == "cairo" {
			v.BuilderName = "cairo"
		}

		if k == "cairomm" {
			v.BuilderName = "std"
		}

		if k == "caja-dropbox" {
			v.BuilderName = "std"
		}

		if k == "caja-extensions" {
			v.BuilderName = "std"
		}

		if k == "caja" {
			v.BuilderName = "std"
		}

		if k == "cal3d" {
			v.BuilderName = "std_cmake"
		}

		if k == "california" {
			v.BuilderName = "std"
		}

		if k == "camorama" {
			v.BuilderName = "std"
		}

		if k == "cantarell-fonts" {
			v.BuilderName = "std"
		}

		if k == "cantor" {
			v.BuilderName = ""
		}

		if k == "capuchin-glib" {
			v.BuilderName = "std"
		}

		if k == "capuchin" {
			v.BuilderName = "std"
		}

		if k == "cargo" {
			v.BuilderName = "std"
		}

		if k == "caribou" {
			v.BuilderName = "std"
		}

		if k == "catcodec" {
			v.BuilderName = ""
		}

		if k == "cb-binutils-i686-pc-linux-gnu" {
			v.BuilderName = "binutils"
		}

		if k == "cb-binutils-x86-pc-linux-gnu" {
			v.BuilderName = "binutils"
		}

		if k == "cb-binutils-x86_64-pc-linux-gnu" {
			v.BuilderName = "binutils"
		}

		if k == "cb-gcc-i686-pc-linux-gnu" {
			v.BuilderName = "gcc"
		}

		if k == "cb-gcc-x86-pc-linux-gnu" {
			v.BuilderName = "gcc"
		}

		if k == "cb-gcc-x86_64-pc-linux-gnu" {
			v.BuilderName = "gcc"
		}

		if k == "cb-glibc-i686-pc-linux-gnu" {
			v.BuilderName = "glibc"
		}

		if k == "cb-glibc-x86-pc-linux-gnu" {
			v.BuilderName = "glibc"
		}

		if k == "cb-glibc-x86_64-pc-linux-gnu" {
			v.BuilderName = "glibc"
		}

		if k == "cb-linux-headers-i686-pc-linux-gnu" {
			v.BuilderName = "linux"
		}

		if k == "cb-linux-headers-x86-pc-linux-gnu" {
			v.BuilderName = "linux"
		}

		if k == "cb-linux-headers-x86_64-pc-linux-gnu" {
			v.BuilderName = "linux"
		}

		if k == "ccache" {
			v.BuilderName = "std"
		}

		if k == "ccaudio" {
			v.BuilderName = "std"
		}

		if k == "ccaudio2" {
			v.BuilderName = "std"
		}

		if k == "ccd2cue" {
			v.BuilderName = "std"
		}

		if k == "ccrtp" {
			v.BuilderName = "std"
		}

		if k == "ccscript" {
			v.BuilderName = "std"
		}

		if k == "ccscript3" {
			v.BuilderName = "std"
		}

		if k == "cdemu-client" {
			v.BuilderName = "std_cmake"
		}

		if k == "cdemu-daemon" {
			v.BuilderName = "std_cmake"
		}

		if k == "cdparanoia-III" {
			v.BuilderName = "std"
		}

		if k == "cdrkit" {
			v.BuilderName = ""
		}

		if k == "cdrtools" {
			v.BuilderName = "cdrtools"
		}

		if k == "cederqvist" {
			v.BuilderName = "std"
		}

		if k == "cedit" {
			v.BuilderName = ""
		}

		if k == "cegui8" {
			v.BuilderName = "std_cmake"
		}

		if k == "cel" {
			v.BuilderName = ""
		}

		if k == "celt" {
			v.BuilderName = "std"
		}

		if k == "cf2html" {
			v.BuilderName = "std"
		}

		if k == "cfe" {
			v.BuilderName = "llvm_components"
		}

		if k == "cfengine" {
			v.BuilderName = "std"
		}

		if k == "cflow" {
			v.BuilderName = "std"
		}

		if k == "cgicc" {
			v.BuilderName = "std"
		}

		if k == "cgkit-py2k" {
			v.BuilderName = "None"
		}

		if k == "cgkit-py3k" {
			v.BuilderName = "cgkit_py3k"
		}

		if k == "check" {
			v.BuilderName = "std"
		}

		if k == "checkpolicy" {
			v.BuilderName = "std"
		}

		if k == "cheese" {
			v.BuilderName = "std"
		}

		if k == "cherokee" {
			v.BuilderName = "cherokee"
		}

		if k == "chicken_export" {
			v.BuilderName = ""
		}

		if k == "chronojump" {
			v.BuilderName = "std"
		}

		if k == "cim" {
			v.BuilderName = "std"
		}

		if k == "cinilerra" {
			v.BuilderName = ""
		}

		if k == "clamav" {
			v.BuilderName = "std"
		}

		if k == "clang-tools-extra" {
			v.BuilderName = "llvm_components"
		}

		if k == "clang" {
			v.BuilderName = "llvm_components"
		}

		if k == "classpath" {
			v.BuilderName = "std"
		}

		if k == "clisp-2.27-PowerMacintosh-powerpc-Darwin" {
			v.BuilderName = "std"
		}

		if k == "clisp-2.27-i686-unknown-Linux" {
			v.BuilderName = "std"
		}

		if k == "clisp-2.28-i686-unknown-cygwin_me-4.90" {
			v.BuilderName = "std"
		}

		if k == "clisp-2.28-i686-unknown-linux" {
			v.BuilderName = "std"
		}

		if k == "clisp-2.29-i386-netbsd" {
			v.BuilderName = "std"
		}

		if k == "clisp-2.29-i686-unknown-cygwin_nt-5.0" {
			v.BuilderName = "std"
		}

		if k == "clisp-2.29-i686-unknown-linux" {
			v.BuilderName = "std"
		}

		if k == "clisp-2.30-i686-unknown-cygwin32" {
			v.BuilderName = "std"
		}

		if k == "clisp-2.30-i686-unknown-linux" {
			v.BuilderName = "std"
		}

		if k == "clisp-alpha-dec-osf" {
			v.BuilderName = "std"
		}

		if k == "clisp-hppa" {
			v.BuilderName = "std"
		}

		if k == "clisp-i" {
			v.BuilderName = "std"
		}

		if k == "clisp-i386-solaris" {
			v.BuilderName = "std"
		}

		if k == "clisp-ia" {
			v.BuilderName = "std"
		}

		if k == "clisp-mips-sgi-irix" {
			v.BuilderName = "std"
		}

		if k == "clisp-powerpc-unknown-linuxlibc" {
			v.BuilderName = "std"
		}

		if k == "clisp-rs6000-ibm-aix" {
			v.BuilderName = "std"
		}

		if k == "clisp" {
			v.BuilderName = "std"
		}

		if k == "cloog" {
			v.BuilderName = "cloog"
		}

		if k == "clutter-box2dmm" {
			v.BuilderName = "std"
		}

		if k == "clutter-cairomm" {
			v.BuilderName = "std"
		}

		if k == "clutter-gst" {
			v.BuilderName = "std"
		}

		if k == "clutter-gtk" {
			v.BuilderName = "std"
		}

		if k == "clutter-gtkmm" {
			v.BuilderName = "std"
		}

		if k == "clutter" {
			v.BuilderName = "clutter"
		}

		if k == "cluttermm" {
			v.BuilderName = "std"
		}

		if k == "cluttermm_tutorial" {
			v.BuilderName = "std"
		}

		if k == "clx.sept" {
			v.BuilderName = "std"
		}

		if k == "cmake" {
			v.BuilderName = "cmake"
		}

		if k == "codeblocks" {
			v.BuilderName = "std"
		}

		if k == "cogl" {
			v.BuilderName = "cogl"
		}

		if k == "cohoba" {
			v.BuilderName = "std"
		}

		if k == "colord-gtk" {
			v.BuilderName = "std"
		}

		if k == "colord" {
			v.BuilderName = "colord"
		}

		if k == "combine" {
			v.BuilderName = "std"
		}

		if k == "comm" {
			v.BuilderName = "std"
		}

		if k == "commoncpp" {
			v.BuilderName = ""
		}

		if k == "commoncpp2" {
			v.BuilderName = "std"
		}

		if k == "compiler-rt" {
			v.BuilderName = "llvm_components"
		}

		if k == "compiz" {
			v.BuilderName = "std"
		}

		if k == "complexity" {
			v.BuilderName = "std"
		}

		if k == "compositeproto" {
			v.BuilderName = "std"
		}

		if k == "conduit" {
			v.BuilderName = "std"
		}

		if k == "connman" {
			v.BuilderName = "std"
		}

		if k == "constype" {
			v.BuilderName = "std"
		}

		if k == "contacts" {
			v.BuilderName = "std"
		}

		if k == "control-center-plus" {
			v.BuilderName = "std"
		}

		if k == "control-center" {
			v.BuilderName = "std"
		}

		if k == "coreutils" {
			v.BuilderName = "std"
		}

		if k == "cortado" {
			v.BuilderName = "std"
		}

		if k == "couchdb-glib" {
			v.BuilderName = "std"
		}

		if k == "cpio" {
			v.BuilderName = "std"
		}

		if k == "cppi" {
			v.BuilderName = "std"
		}

		if k == "cppunit" {
			v.BuilderName = "std"
		}

		if k == "cpufrequtils" {
			v.BuilderName = "std"
		}

		if k == "cracklib" {
			v.BuilderName = "std"
		}

		if k == "crda" {
			v.BuilderName = "std"
		}

		if k == "cronie" {
			v.BuilderName = "std"
		}

		if k == "crosstool-ng" {
			v.BuilderName = "std"
		}

		if k == "crux" {
			v.BuilderName = "std"
		}

		if k == "cryptoapi" {
			v.BuilderName = "std"
		}

		if k == "cryptsetup-luks" {
			v.BuilderName = "std"
		}

		if k == "cryptsetup" {
			v.BuilderName = "std"
		}

		if k == "crystalspace" {
			v.BuilderName = ""
		}

		if k == "csl" {
			v.BuilderName = ""
		}

		if k == "csound" {
			v.BuilderName = ""
		}

		if k == "cuneyform-linux" {
			v.BuilderName = ""
		}

		if k == "cupid" {
			v.BuilderName = "std"
		}

		if k == "cups-filters" {
			v.BuilderName = "std"
		}

		if k == "cups" {
			v.BuilderName = "cups"
		}

		if k == "curl" {
			v.BuilderName = "std"
		}

		if k == "curlftpfs" {
			v.BuilderName = "std"
		}

		if k == "cursynth" {
			v.BuilderName = "std"
		}

		if k == "cvs" {
			v.BuilderName = "std"
		}

		if k == "cyrus-sasl" {
			v.BuilderName = "std"
		}

		if k == "d-feet" {
			v.BuilderName = "std"
		}

		if k == "damageproto" {
			v.BuilderName = "std"
		}

		if k == "dap" {
			v.BuilderName = "std"
		}

		if k == "darcs" {
			v.BuilderName = ""
		}

		if k == "dasher" {
			v.BuilderName = "std"
		}

		if k == "datamash" {
			v.BuilderName = "std"
		}

		if k == "dates" {
			v.BuilderName = "std"
		}

		if k == "db" {
			v.BuilderName = "db"
		}

		if k == "dbmail" {
			v.BuilderName = "std"
		}

		if k == "dbus-glib" {
			v.BuilderName = "std"
		}

		if k == "dbus-java" {
			v.BuilderName = "std"
		}

		if k == "dbus-python" {
			v.BuilderName = "std_py23"
		}

		if k == "dbus" {
			v.BuilderName = "dbus"
		}

		if k == "dconf-editor" {
			v.BuilderName = "std"
		}

		if k == "dconf" {
			v.BuilderName = "std"
		}

		if k == "ddd" {
			v.BuilderName = "std"
		}

		if k == "ddrescue" {
			v.BuilderName = "std"
		}

		if k == "default-icon-theme" {
			v.BuilderName = "std"
		}

		if k == "dejagnu" {
			v.BuilderName = "std"
		}

		if k == "deluge" {
			v.BuilderName = "std_cmake"
		}

		if k == "denemo" {
			v.BuilderName = "std"
		}

		if k == "deskbar-applet" {
			v.BuilderName = "std"
		}

		if k == "deskscribe" {
			v.BuilderName = "std"
		}

		if k == "desktop-data-model" {
			v.BuilderName = "std"
		}

		if k == "desktop-file-utils" {
			v.BuilderName = "std"
		}

		if k == "dev86" {
			v.BuilderName = "dev86"
		}

		if k == "devfsd-v" {
			v.BuilderName = "std"
		}

		if k == "devhelp" {
			v.BuilderName = "std"
		}

		if k == "device-mapper" {
			v.BuilderName = "std"
		}

		if k == "dfarc" {
			v.BuilderName = "std"
		}

		if k == "dgen-sdl" {
			v.BuilderName = ""
		}

		if k == "dhcp" {
			v.BuilderName = "dhcp"
		}

		if k == "dhcpcd-dbus" {
			v.BuilderName = "std"
		}

		if k == "dhcpcd" {
			v.BuilderName = "std"
		}

		if k == "dia" {
			v.BuilderName = "std"
		}

		if k == "dialog" {
			v.BuilderName = "std"
		}

		if k == "dico" {
			v.BuilderName = "std"
		}

		if k == "diction" {
			v.BuilderName = "std"
		}

		if k == "diethotplug" {
			v.BuilderName = "std"
		}

		if k == "dietlibc" {
			v.BuilderName = "std"
		}

		if k == "diffuse" {
			v.BuilderName = ""
		}

		if k == "diffutils" {
			v.BuilderName = "std"
		}

		if k == "dionysus" {
			v.BuilderName = "std"
		}

		if k == "dired" {
			v.BuilderName = ""
		}

		if k == "direvent" {
			v.BuilderName = "std"
		}

		if k == "discident-glib" {
			v.BuilderName = "std"
		}

		if k == "dissy" {
			v.BuilderName = ""
		}

		if k == "djview" {
			v.BuilderName = "std"
		}

		if k == "djvulibre" {
			v.BuilderName = "std"
		}

		if k == "dmapi" {
			v.BuilderName = "std"
		}

		if k == "dmxproto" {
			v.BuilderName = "std"
		}

		if k == "docbook-sgml3" {
			v.BuilderName = "docbook_sgml3"
		}

		if k == "docbook-sgml4" {
			v.BuilderName = "docbook_sgml4"
		}

		if k == "docbook-xml4" {
			v.BuilderName = "docbook_xml4"
		}

		if k == "docbook-xsl" {
			v.BuilderName = "docbook_xsl"
		}

		if k == "docbook2X" {
			v.BuilderName = "std"
		}

		if k == "dogtail" {
			v.BuilderName = "std"
		}

		if k == "dominion" {
			v.BuilderName = "std"
		}

		if k == "dosbox" {
			v.BuilderName = "dosbox"
		}

		if k == "doschk" {
			v.BuilderName = "std"
		}

		if k == "dosfstools" {
			v.BuilderName = "std"
		}

		if k == "dots" {
			v.BuilderName = "std"
		}

		if k == "dovecot" {
			v.BuilderName = "dovecot"
		}

		if k == "doxygen" {
			v.BuilderName = "doxygen"
		}

		if k == "dracut" {
			v.BuilderName = "std"
		}

		if k == "dragonegg" {
			v.BuilderName = "std"
		}

		if k == "dri2proto" {
			v.BuilderName = "std"
		}

		if k == "dri3proto" {
			v.BuilderName = "std"
		}

		if k == "drwright" {
			v.BuilderName = "std"
		}

		if k == "dryad" {
			v.BuilderName = "std"
		}

		if k == "dstat" {
			v.BuilderName = ""
		}

		if k == "dtc" {
			v.BuilderName = "std"
		}

		if k == "dtquery" {
			v.BuilderName = "std"
		}

		if k == "e" {
			v.BuilderName = "std"
		}

		if k == "e2fsprogs-libs" {
			v.BuilderName = "e2fsprogs"
		}

		if k == "e2fsprogs" {
			v.BuilderName = "e2fsprogs"
		}

		if k == "ease" {
			v.BuilderName = "std"
		}

		if k == "easejs" {
			v.BuilderName = "std"
		}

		if k == "easytag" {
			v.BuilderName = "std"
		}

		if k == "ecb" {
			v.BuilderName = ""
		}

		if k == "ecj" {
			v.BuilderName = ""
		}

		if k == "ed" {
			v.BuilderName = "std"
		}

		if k == "editres" {
			v.BuilderName = "std"
		}

		if k == "ee" {
			v.BuilderName = "std"
		}

		if k == "eel" {
			v.BuilderName = "std"
		}

		if k == "efilinux" {
			v.BuilderName = "std"
		}

		if k == "eggcups" {
			v.BuilderName = "std"
		}

		if k == "eggdbus" {
			v.BuilderName = "std"
		}

		if k == "eigen" {
			v.BuilderName = ""
		}

		if k == "eiskaltdc" {
			v.BuilderName = ""
		}

		if k == "ejabberd" {
			v.BuilderName = "ejabberd"
		}

		if k == "eject" {
			v.BuilderName = ""
		}

		if k == "ekiga" {
			v.BuilderName = "std"
		}

		if k == "electric" {
			v.BuilderName = "std"
		}

		if k == "elfutils" {
			v.BuilderName = "elfutils"
		}

		if k == "elisp-manual-21" {
			v.BuilderName = "std"
		}

		if k == "em" {
			v.BuilderName = "std"
		}

		if k == "emacs-lisp-intro" {
			v.BuilderName = "std"
		}

		if k == "emacs" {
			v.BuilderName = "emacs"
		}

		if k == "ember" {
			v.BuilderName = "std"
		}

		if k == "emerillon" {
			v.BuilderName = "std"
		}

		if k == "emms" {
			v.BuilderName = "std"
		}

		if k == "empathy" {
			v.BuilderName = "std"
		}

		if k == "emscripten" {
			v.BuilderName = "llvm_components"
		}

		if k == "enchant" {
			v.BuilderName = "std"
		}

		if k == "encodings" {
			v.BuilderName = "std"
		}

		if k == "engrampa" {
			v.BuilderName = "std"
		}

		if k == "enscript" {
			v.BuilderName = "std"
		}

		if k == "eog-plugins" {
			v.BuilderName = "std"
		}

		if k == "eog" {
			v.BuilderName = "std"
		}

		if k == "eom" {
			v.BuilderName = "std"
		}

		if k == "epiphany-extensions" {
			v.BuilderName = "std"
		}

		if k == "epiphany" {
			v.BuilderName = "epiphany"
		}

		if k == "epydoc" {
			v.BuilderName = ""
		}

		if k == "eric4" {
			v.BuilderName = "eric"
		}

		if k == "eric5" {
			v.BuilderName = "eric"
		}

		if k == "eris" {
			v.BuilderName = "std"
		}

		if k == "erlang" {
			v.BuilderName = "erlang"
		}

		if k == "esound" {
			v.BuilderName = "std"
		}

		if k == "espeak" {
			v.BuilderName = ""
		}

		if k == "etc" {
			v.BuilderName = ""
		}

		if k == "ethtool" {
			v.BuilderName = "std"
		}

		if k == "evieext" {
			v.BuilderName = "std"
		}

		if k == "evince" {
			v.BuilderName = "evince"
		}

		if k == "evolution-activesync" {
			v.BuilderName = "std"
		}

		if k == "evolution-caldav" {
			v.BuilderName = "std"
		}

		if k == "evolution-couchdb" {
			v.BuilderName = "std"
		}

		if k == "evolution-data-server-dbus" {
			v.BuilderName = "std"
		}

		if k == "evolution-data-server" {
			v.BuilderName = "evolution_data_server"
		}

		if k == "evolution-ews" {
			v.BuilderName = "std"
		}

		if k == "evolution-exchange" {
			v.BuilderName = "std"
		}

		if k == "evolution-groupwise" {
			v.BuilderName = "std"
		}

		if k == "evolution-jescs" {
			v.BuilderName = "std"
		}

		if k == "evolution-kolab" {
			v.BuilderName = "std"
		}

		if k == "evolution-mapi" {
			v.BuilderName = "std"
		}

		if k == "evolution-scalix" {
			v.BuilderName = "std"
		}

		if k == "evolution-sharp" {
			v.BuilderName = "std"
		}

		if k == "evolution-webcal" {
			v.BuilderName = "std"
		}

		if k == "evolution" {
			v.BuilderName = "evolution"
		}

		if k == "exempi" {
			v.BuilderName = "std"
		}

		if k == "exif" {
			v.BuilderName = "std"
		}

		if k == "exim" {
			v.BuilderName = "exim"
		}

		if k == "exiv2" {
			v.BuilderName = "std"
		}

		if k == "exo" {
			v.BuilderName = ""
		}

		if k == "expat" {
			v.BuilderName = "std"
		}

		if k == "expect" {
			v.BuilderName = "expect"
		}

		if k == "ezstream" {
			v.BuilderName = "std"
		}

		if k == "f-spot" {
			v.BuilderName = "std"
		}

		if k == "faac" {
			v.BuilderName = "std"
		}

		if k == "faad2" {
			v.BuilderName = "std"
		}

		if k == "fakeroot-ng" {
			v.BuilderName = "std"
		}

		if k == "fantasy-island" {
			v.BuilderName = ""
		}

		if k == "farsight" {
			v.BuilderName = "std"
		}

		if k == "farsight2" {
			v.BuilderName = "std"
		}

		if k == "farstream" {
			v.BuilderName = "std"
		}

		if k == "farstream0.1" {
			v.BuilderName = "std"
		}

		if k == "fast-user-switch-applet" {
			v.BuilderName = "std"
		}

		if k == "fatattr" {
			v.BuilderName = ""
		}

		if k == "fcpackage" {
			v.BuilderName = "std"
		}

		if k == "fcron" {
			v.BuilderName = "fcron"
		}

		if k == "fdisk" {
			v.BuilderName = "std"
		}

		if k == "ferret" {
			v.BuilderName = "std"
		}

		if k == "ffmpeg" {
			v.BuilderName = "ffmpeg"
		}

		if k == "fftw" {
			v.BuilderName = "std"
		}

		if k == "file-roller" {
			v.BuilderName = "std"
		}

		if k == "file" {
			v.BuilderName = "std"
		}

		if k == "filelight" {
			v.BuilderName = ""
		}

		if k == "findutils" {
			v.BuilderName = "std"
		}

		if k == "firefox" {
			v.BuilderName = "firefox"
		}

		if k == "firestarter" {
			v.BuilderName = "std"
		}

		if k == "fisicalab-mac" {
			v.BuilderName = "std"
		}

		if k == "fisicalab" {
			v.BuilderName = "std"
		}

		if k == "five-or-more" {
			v.BuilderName = "std"
		}

		if k == "fixesproto" {
			v.BuilderName = "std"
		}

		if k == "flac" {
			v.BuilderName = "std"
		}

		if k == "flex" {
			v.BuilderName = "std"
		}

		if k == "flock" {
			v.BuilderName = "std"
		}

		if k == "fltk" {
			v.BuilderName = "std_cmake"
		}

		if k == "fluidsynth" {
			v.BuilderName = "std"
		}

		if k == "folks" {
			v.BuilderName = "folks"
		}

		if k == "font-adobe-100dpi" {
			v.BuilderName = "std"
		}

		if k == "font-adobe-75dpi" {
			v.BuilderName = "std"
		}

		if k == "font-adobe-utopia-100dpi" {
			v.BuilderName = "std"
		}

		if k == "font-adobe-utopia-75dpi" {
			v.BuilderName = "std"
		}

		if k == "font-adobe-utopia-type1" {
			v.BuilderName = "std"
		}

		if k == "font-alias" {
			v.BuilderName = "std"
		}

		if k == "font-arabic-misc" {
			v.BuilderName = "std"
		}

		if k == "font-bh-100dpi" {
			v.BuilderName = "std"
		}

		if k == "font-bh-75dpi" {
			v.BuilderName = "std"
		}

		if k == "font-bh-lucidatypewriter-100dpi" {
			v.BuilderName = "std"
		}

		if k == "font-bh-lucidatypewriter-75dpi" {
			v.BuilderName = "std"
		}

		if k == "font-bh-ttf" {
			v.BuilderName = "std"
		}

		if k == "font-bh-type1" {
			v.BuilderName = "std"
		}

		if k == "font-bitstream-100dpi" {
			v.BuilderName = "std"
		}

		if k == "font-bitstream-75dpi" {
			v.BuilderName = "std"
		}

		if k == "font-bitstream-speedo" {
			v.BuilderName = "std"
		}

		if k == "font-bitstream-type1" {
			v.BuilderName = "std"
		}

		if k == "font-cronyx-cyrillic" {
			v.BuilderName = "std"
		}

		if k == "font-cursor-misc" {
			v.BuilderName = "std"
		}

		if k == "font-daewoo-misc" {
			v.BuilderName = "std"
		}

		if k == "font-dec-misc" {
			v.BuilderName = "std"
		}

		if k == "font-ibm-type1" {
			v.BuilderName = "std"
		}

		if k == "font-isas-misc" {
			v.BuilderName = "std"
		}

		if k == "font-jis-misc" {
			v.BuilderName = "std"
		}

		if k == "font-micro-misc" {
			v.BuilderName = "std"
		}

		if k == "font-misc-cyrillic" {
			v.BuilderName = "std"
		}

		if k == "font-misc-ethiopic" {
			v.BuilderName = "std"
		}

		if k == "font-misc-meltho" {
			v.BuilderName = "std"
		}

		if k == "font-misc-misc" {
			v.BuilderName = "std"
		}

		if k == "font-mutt-misc" {
			v.BuilderName = "std"
		}

		if k == "font-schumacher-misc" {
			v.BuilderName = "std"
		}

		if k == "font-screen-cyrillic" {
			v.BuilderName = "std"
		}

		if k == "font-sony-misc" {
			v.BuilderName = "std"
		}

		if k == "font-sun-misc" {
			v.BuilderName = "std"
		}

		if k == "font-util" {
			v.BuilderName = "std"
		}

		if k == "font-winitzki-cyrillic" {
			v.BuilderName = "std"
		}

		if k == "font-xfree86-type1" {
			v.BuilderName = "std"
		}

		if k == "fontcacheproto" {
			v.BuilderName = "std"
		}

		if k == "fontconfig" {
			v.BuilderName = "std"
		}

		if k == "fontforge" {
			v.BuilderName = "std"
		}

		if k == "fontilus" {
			v.BuilderName = "std"
		}

		if k == "fontsproto" {
			v.BuilderName = "std"
		}

		if k == "fonttosfnt" {
			v.BuilderName = "std"
		}

		if k == "fontutils" {
			v.BuilderName = "std"
		}

		if k == "foomatic-db-engine" {
			v.BuilderName = "std"
		}

		if k == "foomatic-db" {
			v.BuilderName = "std"
		}

		if k == "foomatic-filters" {
			v.BuilderName = "std"
		}

		if k == "fop" {
			v.BuilderName = "fop"
		}

		if k == "fossil" {
			v.BuilderName = "fossil"
		}

		if k == "four-in-a-row" {
			v.BuilderName = "std"
		}

		if k == "fpc" {
			v.BuilderName = "fpc"
		}

		if k == "fping" {
			v.BuilderName = "std"
		}

		if k == "fpm" {
			v.BuilderName = "std"
		}

		if k == "free-cad" {
			v.BuilderName = ""
		}

		if k == "freealut" {
			v.BuilderName = "std"
		}

		if k == "freeciv" {
			v.BuilderName = ""
		}

		if k == "freedink-data" {
			v.BuilderName = "std"
		}

		if k == "freedink" {
			v.BuilderName = "std"
		}

		if k == "freefont-otf" {
			v.BuilderName = "std"
		}

		if k == "freefont-sfd" {
			v.BuilderName = "std"
		}

		if k == "freefont-src" {
			v.BuilderName = "std"
		}

		if k == "freefont-ttf" {
			v.BuilderName = "std"
		}

		if k == "freeglut" {
			v.BuilderName = "std_cmake"
		}

		if k == "freeipmi" {
			v.BuilderName = "std"
		}

		if k == "freenx-server" {
			v.BuilderName = ""
		}

		if k == "freeorion" {
			v.BuilderName = "freeorion"
		}

		if k == "freetype" {
			v.BuilderName = "freetype"
		}

		if k == "frei0r-plugins" {
			v.BuilderName = "std"
		}

		if k == "fribidi" {
			v.BuilderName = "std"
		}

		if k == "frogr" {
			v.BuilderName = "std"
		}

		if k == "fslsfonts" {
			v.BuilderName = "std"
		}

		if k == "fstobdf" {
			v.BuilderName = "std"
		}

		if k == "ftgl" {
			v.BuilderName = "ftgl"
		}

		if k == "fuse" {
			v.BuilderName = "std"
		}

		if k == "fxload" {
			v.BuilderName = "std"
		}

		if k == "g++" {
			v.BuilderName = "std"
		}

		if k == "g-print" {
			v.BuilderName = "std"
		}

		if k == "g-wrap" {
			v.BuilderName = "std"
		}

		if k == "gASQL-0.6" {
			v.BuilderName = "std"
		}

		if k == "gDesklets" {
			v.BuilderName = "std"
		}

		if k == "gabber" {
			v.BuilderName = "std"
		}

		if k == "gail" {
			v.BuilderName = "std"
		}

		if k == "gajim" {
			v.BuilderName = "std"
		}

		if k == "gal" {
			v.BuilderName = "std"
		}

		if k == "gal2-0" {
			v.BuilderName = "std"
		}

		if k == "galeon" {
			v.BuilderName = "std"
		}

		if k == "gama" {
			v.BuilderName = "std"
		}

		if k == "gamin" {
			v.BuilderName = "std"
		}

		if k == "garcon" {
			v.BuilderName = ""
		}

		if k == "garnet" {
			v.BuilderName = "std"
		}

		if k == "garnome" {
			v.BuilderName = "std"
		}

		if k == "garpd" {
			v.BuilderName = "std"
		}

		if k == "gaupol" {
			v.BuilderName = "std_pythons"
		}

		if k == "gavl" {
			v.BuilderName = "std"
		}

		if k == "gawk-doc" {
			v.BuilderName = "std"
		}

		if k == "gawk-ps" {
			v.BuilderName = "std"
		}

		if k == "gawk" {
			v.BuilderName = "std"
		}

		if k == "gazpacho" {
			v.BuilderName = "std"
		}

		if k == "gb" {
			v.BuilderName = "std"
		}

		if k == "gc" {
			v.BuilderName = "std"
		}

		if k == "gcab" {
			v.BuilderName = "std"
		}

		if k == "gcal" {
			v.BuilderName = "std"
		}

		if k == "gcalctool" {
			v.BuilderName = "std"
		}

		if k == "gcc-ada" {
			v.BuilderName = "std"
		}

		if k == "gcc-chill" {
			v.BuilderName = "std"
		}

		if k == "gcc-core" {
			v.BuilderName = "std"
		}

		if k == "gcc-fortran" {
			v.BuilderName = "std"
		}

		if k == "gcc-g++" {
			v.BuilderName = "std"
		}

		if k == "gcc-g77" {
			v.BuilderName = "std"
		}

		if k == "gcc-go" {
			v.BuilderName = "std"
		}

		if k == "gcc-java" {
			v.BuilderName = "std"
		}

		if k == "gcc-objc" {
			v.BuilderName = "std"
		}

		if k == "gcc-testsuite" {
			v.BuilderName = "std"
		}

		if k == "gcc-vms" {
			v.BuilderName = "std"
		}

		if k == "gcc" {
			v.BuilderName = "gcc"
		}

		if k == "gccmakedep" {
			v.BuilderName = "std"
		}

		if k == "gcdemu" {
			v.BuilderName = "std_cmake"
		}

		if k == "gcide" {
			v.BuilderName = "std"
		}

		if k == "gcl" {
			v.BuilderName = "std"
		}

		if k == "gcompris" {
			v.BuilderName = "std"
		}

		if k == "gconf-editor" {
			v.BuilderName = "std"
		}

		if k == "gconfmm" {
			v.BuilderName = "std"
		}

		if k == "gcr" {
			v.BuilderName = "std"
		}

		if k == "gdb" {
			v.BuilderName = "gdb"
		}

		if k == "gdbm" {
			v.BuilderName = "gdbm"
		}

		if k == "gdk-pixbuf" {
			v.BuilderName = "gdk_pixbuf"
		}

		if k == "gdl" {
			v.BuilderName = "std"
		}

		if k == "gdlmm" {
			v.BuilderName = "std"
		}

		if k == "gdm" {
			v.BuilderName = "std"
		}

		if k == "geany" {
			v.BuilderName = "std"
		}

		if k == "geary" {
			v.BuilderName = "std"
		}

		if k == "gecb" {
			v.BuilderName = "std"
		}

		if k == "geda-gaf" {
			v.BuilderName = ""
		}

		if k == "gedit-code-assistance" {
			v.BuilderName = "std"
		}

		if k == "gedit-collaboration" {
			v.BuilderName = "std"
		}

		if k == "gedit-cossa" {
			v.BuilderName = "std"
		}

		if k == "gedit-latex" {
			v.BuilderName = "std"
		}

		if k == "gedit-plugins" {
			v.BuilderName = "std"
		}

		if k == "gedit" {
			v.BuilderName = "std"
		}

		if k == "gedit2" {
			v.BuilderName = "std"
		}

		if k == "gegl-gtk" {
			v.BuilderName = ""
		}

		if k == "gegl-qt" {
			v.BuilderName = ""
		}

		if k == "gegl-vala" {
			v.BuilderName = "std"
		}

		if k == "gegl_0_2" {
			v.BuilderName = "gegl"
		}

		if k == "gegl_0_3" {
			v.BuilderName = "gegl"
		}

		if k == "geglmm" {
			v.BuilderName = "std"
		}

		if k == "gen-hosts" {
			v.BuilderName = "std"
		}

		if k == "generator" {
			v.BuilderName = ""
		}

		if k == "gengen" {
			v.BuilderName = "std"
		}

		if k == "gengetopt" {
			v.BuilderName = "std"
		}

		if k == "genius" {
			v.BuilderName = "std"
		}

		if k == "geoclue" {
			v.BuilderName = "std"
		}

		if k == "geocode-glib" {
			v.BuilderName = "std"
		}

		if k == "gerberv" {
			v.BuilderName = ""
		}

		if k == "gerwin" {
			v.BuilderName = "std"
		}

		if k == "gettext" {
			v.BuilderName = "std"
		}

		if k == "gevice" {
			v.BuilderName = "std"
		}

		if k == "gexiv2" {
			v.BuilderName = "std"
		}

		if k == "gfax" {
			v.BuilderName = "std"
		}

		if k == "gfbgraph" {
			v.BuilderName = "std"
		}

		if k == "gfloppy" {
			v.BuilderName = "std"
		}

		if k == "gforth" {
			v.BuilderName = "std"
		}

		if k == "gget" {
			v.BuilderName = "std"
		}

		if k == "ggradebook" {
			v.BuilderName = "std"
		}

		if k == "ggv" {
			v.BuilderName = "std"
		}

		if k == "ghc" {
			v.BuilderName = ""
		}

		if k == "ghex" {
			v.BuilderName = "std"
		}

		if k == "ghfaxviewer" {
			v.BuilderName = "std"
		}

		if k == "ghostscript-fonts-other" {
			v.BuilderName = "std"
		}

		if k == "ghostscript-fonts-std" {
			v.BuilderName = "std"
		}

		if k == "ghostscript" {
			v.BuilderName = "ghostscript"
		}

		if k == "gide" {
			v.BuilderName = "std"
		}

		if k == "gidfwizard" {
			v.BuilderName = "std"
		}

		if k == "giflib" {
			v.BuilderName = "std"
		}

		if k == "giflib4" {
			v.BuilderName = "std"
		}

		if k == "gift" {
			v.BuilderName = "std"
		}

		if k == "giggle" {
			v.BuilderName = "std"
		}

		if k == "gimagereader" {
			v.BuilderName = ""
		}

		if k == "gimp-gap" {
			v.BuilderName = ""
		}

		if k == "gimp" {
			v.BuilderName = "gimp"
		}

		if k == "gio-standalone" {
			v.BuilderName = "std"
		}

		if k == "gir-repository" {
			v.BuilderName = "std"
		}

		if k == "girl" {
			v.BuilderName = "std"
		}

		if k == "git-core" {
			v.BuilderName = "std"
		}

		if k == "git-htmldocs" {
			v.BuilderName = "std"
		}

		if k == "git-manpages" {
			v.BuilderName = "git_manpages"
		}

		if k == "git" {
			v.BuilderName = "git"
		}

		if k == "gitg" {
			v.BuilderName = "std"
		}

		if k == "gjdoc" {
			v.BuilderName = "std"
		}

		if k == "gjs" {
			v.BuilderName = "gjs"
		}

		if k == "glabels" {
			v.BuilderName = "std"
		}

		if k == "glade-3" {
			v.BuilderName = "std"
		}

		if k == "glade" {
			v.BuilderName = "std"
		}

		if k == "glade3" {
			v.BuilderName = "std"
		}

		if k == "glamor-egl" {
			v.BuilderName = "std"
		}

		if k == "gleem" {
			v.BuilderName = "std"
		}

		if k == "glew" {
			v.BuilderName = "glew"
		}

		if k == "glib-java" {
			v.BuilderName = "std"
		}

		if k == "glib-networking" {
			v.BuilderName = "glib_networking"
		}

		if k == "glib" {
			v.BuilderName = "glib"
		}

		if k == "glibc-crypt" {
			v.BuilderName = "std"
		}

		if k == "glibc-libidn" {
			v.BuilderName = "std"
		}

		if k == "glibc-linuxthreads" {
			v.BuilderName = "std"
		}

		if k == "glibc-localedata" {
			v.BuilderName = "std"
		}

		if k == "glibc-ports" {
			v.BuilderName = "glibc"
		}

		if k == "glibc" {
			v.BuilderName = "glibc"
		}

		if k == "glibmm" {
			v.BuilderName = "std"
		}

		if k == "glibwww" {
			v.BuilderName = "std"
		}

		if k == "glick2" {
			v.BuilderName = "std"
		}

		if k == "glimmer" {
			v.BuilderName = "std"
		}

		if k == "glm" {
			v.BuilderName = "std_cmake"
		}

		if k == "glob2" {
			v.BuilderName = ""
		}

		if k == "global" {
			v.BuilderName = "std"
		}

		if k == "glom-postgresql-setup" {
			v.BuilderName = "std"
		}

		if k == "glom" {
			v.BuilderName = "std"
		}

		if k == "glpk" {
			v.BuilderName = "std"
		}

		if k == "glproto" {
			v.BuilderName = "std"
		}

		if k == "glu" {
			v.BuilderName = "std"
		}

		if k == "glw" {
			v.BuilderName = "std"
		}

		if k == "gmdns" {
			v.BuilderName = "std"
		}

		if k == "gmerlin-avdecoder" {
			v.BuilderName = "std"
		}

		if k == "gmerlin-effectv" {
			v.BuilderName = "std"
		}

		if k == "gmerlin-encoders" {
			v.BuilderName = "std"
		}

		if k == "gmerlin-mozilla" {
			v.BuilderName = ""
		}

		if k == "gmerlin" {
			v.BuilderName = "std"
		}

		if k == "gmf" {
			v.BuilderName = "std"
		}

		if k == "gmime" {
			v.BuilderName = "gmime"
		}

		if k == "gmp" {
			v.BuilderName = "gmp"
		}

		if k == "gnash" {
			v.BuilderName = "gnash"
		}

		if k == "gnat-gpl" {
			v.BuilderName = ""
		}

		if k == "gnats" {
			v.BuilderName = "std"
		}

		if k == "gnatsweb" {
			v.BuilderName = "std"
		}

		if k == "gnet" {
			v.BuilderName = "std"
		}

		if k == "gnome-2048" {
			v.BuilderName = "std"
		}

		if k == "gnome-admin" {
			v.BuilderName = "std"
		}

		if k == "gnome-applets" {
			v.BuilderName = "std"
		}

		if k == "gnome-audio" {
			v.BuilderName = "std"
		}

		if k == "gnome-backgrounds" {
			v.BuilderName = "std"
		}

		if k == "gnome-battery-bench" {
			v.BuilderName = "std"
		}

		if k == "gnome-blog" {
			v.BuilderName = "std"
		}

		if k == "gnome-bluetooth" {
			v.BuilderName = "std"
		}

		if k == "gnome-boxes-nonfree" {
			v.BuilderName = "std"
		}

		if k == "gnome-boxes" {
			v.BuilderName = "std"
		}

		if k == "gnome-braille" {
			v.BuilderName = "std"
		}

		if k == "gnome-build" {
			v.BuilderName = "std"
		}

		if k == "gnome-builder" {
			v.BuilderName = "std"
		}

		if k == "gnome-calculator" {
			v.BuilderName = "std"
		}

		if k == "gnome-calendar" {
			v.BuilderName = "std"
		}

		if k == "gnome-canvose" {
			v.BuilderName = "std"
		}

		if k == "gnome-characters" {
			v.BuilderName = "std"
		}

		if k == "gnome-chart" {
			v.BuilderName = "std"
		}

		if k == "gnome-chess" {
			v.BuilderName = "std"
		}

		if k == "gnome-clocks" {
			v.BuilderName = "std"
		}

		if k == "gnome-code-assistance" {
			v.BuilderName = "std"
		}

		if k == "gnome-color-manager" {
			v.BuilderName = "std"
		}

		if k == "gnome-commander" {
			v.BuilderName = "std"
		}

		if k == "gnome-common" {
			v.BuilderName = "std"
		}

		if k == "gnome-contacts" {
			v.BuilderName = "std"
		}

		if k == "gnome-control-center" {
			v.BuilderName = "std"
		}

		if k == "gnome-core" {
			v.BuilderName = "std"
		}

		if k == "gnome-crash" {
			v.BuilderName = "std"
		}

		if k == "gnome-cups-manager" {
			v.BuilderName = "std"
		}

		if k == "gnome-db" {
			v.BuilderName = "std"
		}

		if k == "gnome-debug" {
			v.BuilderName = "std"
		}

		if k == "gnome-desktop-sharp" {
			v.BuilderName = "std"
		}

		if k == "gnome-desktop-testing" {
			v.BuilderName = "std"
		}

		if k == "gnome-desktop" {
			v.BuilderName = "std"
		}

		if k == "gnome-devel-docs" {
			v.BuilderName = "std"
		}

		if k == "gnome-device-manager" {
			v.BuilderName = "std"
		}

		if k == "gnome-dictionary" {
			v.BuilderName = "std"
		}

		if k == "gnome-directory-thumbnailer" {
			v.BuilderName = "std"
		}

		if k == "gnome-disk-utility" {
			v.BuilderName = "std"
		}

		if k == "gnome-doc-utils" {
			v.BuilderName = "std"
		}

		if k == "gnome-documents" {
			v.BuilderName = "std"
		}

		if k == "gnome-dvb-daemon" {
			v.BuilderName = "std"
		}

		if k == "gnome-epub-thumbnailer" {
			v.BuilderName = "std"
		}

		if k == "gnome-file-selector" {
			v.BuilderName = "std"
		}

		if k == "gnome-flashback" {
			v.BuilderName = "std"
		}

		if k == "gnome-font-viewer" {
			v.BuilderName = "std"
		}

		if k == "gnome-games-extra-data" {
			v.BuilderName = "std"
		}

		if k == "gnome-games" {
			v.BuilderName = "std"
		}

		if k == "gnome-getting-started-docs" {
			v.BuilderName = "std"
		}

		if k == "gnome-gpg" {
			v.BuilderName = "std"
		}

		if k == "gnome-guile" {
			v.BuilderName = "std"
		}

		if k == "gnome-hello" {
			v.BuilderName = "std"
		}

		if k == "gnome-icon-theme-extras" {
			v.BuilderName = "std"
		}

		if k == "gnome-icon-theme-symbolic" {
			v.BuilderName = "std"
		}

		if k == "gnome-icon-theme" {
			v.BuilderName = "std"
		}

		if k == "gnome-initial-setup" {
			v.BuilderName = "std"
		}

		if k == "gnome-jabber" {
			v.BuilderName = "std"
		}

		if k == "gnome-js-common" {
			v.BuilderName = "std"
		}

		if k == "gnome-keyring-manager" {
			v.BuilderName = "std"
		}

		if k == "gnome-keyring" {
			v.BuilderName = "gnomekeyring"
		}

		if k == "gnome-klotski" {
			v.BuilderName = "std"
		}

		if k == "gnome-kra-ora-thumbnailer" {
			v.BuilderName = "std"
		}

		if k == "gnome-launch-box" {
			v.BuilderName = "std"
		}

		if k == "gnome-libs" {
			v.BuilderName = "std"
		}

		if k == "gnome-linuxconf" {
			v.BuilderName = "std"
		}

		if k == "gnome-lirc-properties" {
			v.BuilderName = "std"
		}

		if k == "gnome-logs" {
			v.BuilderName = "std"
		}

		if k == "gnome-lokkit" {
			v.BuilderName = "std"
		}

		if k == "gnome-mag" {
			v.BuilderName = "std"
		}

		if k == "gnome-mahjongg" {
			v.BuilderName = "std"
		}

		if k == "gnome-main-menu" {
			v.BuilderName = "std"
		}

		if k == "gnome-maps" {
			v.BuilderName = "std"
		}

		if k == "gnome-media" {
			v.BuilderName = "std"
		}

		if k == "gnome-menus" {
			v.BuilderName = "std"
		}

		if k == "gnome-mime-data" {
			v.BuilderName = "std"
		}

		if k == "gnome-mines" {
			v.BuilderName = "std"
		}

		if k == "gnome-mount" {
			v.BuilderName = "std"
		}

		if k == "gnome-mud" {
			v.BuilderName = "std"
		}

		if k == "gnome-multi-writer" {
			v.BuilderName = "std"
		}

		if k == "gnome-music" {
			v.BuilderName = "std"
		}

		if k == "gnome-nds-thumbnailer" {
			v.BuilderName = "std"
		}

		if k == "gnome-netinfo" {
			v.BuilderName = "std"
		}

		if k == "gnome-netstatus" {
			v.BuilderName = "std"
		}

		if k == "gnome-nettool" {
			v.BuilderName = "std"
		}

		if k == "gnome-network" {
			v.BuilderName = "std"
		}

		if k == "gnome-nibbles" {
			v.BuilderName = "std"
		}

		if k == "gnome-objc" {
			v.BuilderName = "std"
		}

		if k == "gnome-online-accounts" {
			v.BuilderName = "gnome_online_accounts"
		}

		if k == "gnome-online-miners" {
			v.BuilderName = "std"
		}

		if k == "gnome-packagekit" {
			v.BuilderName = "std"
		}

		if k == "gnome-panel" {
			v.BuilderName = "std"
		}

		if k == "gnome-perfmeter" {
			v.BuilderName = "std"
		}

		if k == "gnome-phone-manager" {
			v.BuilderName = "std"
		}

		if k == "gnome-photos" {
			v.BuilderName = "std"
		}

		if k == "gnome-pilot-conduits" {
			v.BuilderName = "std"
		}

		if k == "gnome-pilot" {
			v.BuilderName = "std"
		}

		if k == "gnome-pim" {
			v.BuilderName = "std"
		}

		if k == "gnome-power-manager" {
			v.BuilderName = "std"
		}

		if k == "gnome-print" {
			v.BuilderName = "std"
		}

		if k == "gnome-python-desktop" {
			v.BuilderName = "std"
		}

		if k == "gnome-python-extras" {
			v.BuilderName = "std"
		}

		if k == "gnome-python" {
			v.BuilderName = "std"
		}

		if k == "gnome-reset" {
			v.BuilderName = "std"
		}

		if k == "gnome-robots" {
			v.BuilderName = "std"
		}

		if k == "gnome-scan" {
			v.BuilderName = "std"
		}

		if k == "gnome-screensaver" {
			v.BuilderName = "std"
		}

		if k == "gnome-screenshot" {
			v.BuilderName = "std"
		}

		if k == "gnome-search-tool" {
			v.BuilderName = "std"
		}

		if k == "gnome-session" {
			v.BuilderName = "gnomesession"
		}

		if k == "gnome-settings-daemon" {
			v.BuilderName = "gnome_settings_daemon"
		}

		if k == "gnome-sharp" {
			v.BuilderName = "std"
		}

		if k == "gnome-shell-extensions" {
			v.BuilderName = "std"
		}

		if k == "gnome-shell" {
			v.BuilderName = "std"
		}

		if k == "gnome-software" {
			v.BuilderName = "std"
		}

		if k == "gnome-sound-recorder" {
			v.BuilderName = "std"
		}

		if k == "gnome-specimen" {
			v.BuilderName = "std"
		}

		if k == "gnome-speech" {
			v.BuilderName = "std"
		}

		if k == "gnome-spell" {
			v.BuilderName = "std"
		}

		if k == "gnome-sudoku" {
			v.BuilderName = "std"
		}

		if k == "gnome-system-log" {
			v.BuilderName = "std"
		}

		if k == "gnome-system-monitor" {
			v.BuilderName = "std"
		}

		if k == "gnome-system-tools" {
			v.BuilderName = "std"
		}

		if k == "gnome-taquin" {
			v.BuilderName = "std"
		}

		if k == "gnome-terminal" {
			v.BuilderName = "gnome_terminal"
		}

		if k == "gnome-tetravex" {
			v.BuilderName = "std"
		}

		if k == "gnome-themes-extras" {
			v.BuilderName = "std"
		}

		if k == "gnome-themes-standard" {
			v.BuilderName = "std"
		}

		if k == "gnome-themes" {
			v.BuilderName = "std"
		}

		if k == "gnome-tweak-tool" {
			v.BuilderName = "std"
		}

		if k == "gnome-user-docs" {
			v.BuilderName = "std"
		}

		if k == "gnome-user-share" {
			v.BuilderName = "std"
		}

		if k == "gnome-utils" {
			v.BuilderName = "std"
		}

		if k == "gnome-vfs-extras" {
			v.BuilderName = "std"
		}

		if k == "gnome-vfs-monikers" {
			v.BuilderName = "std"
		}

		if k == "gnome-vfs-obexftp" {
			v.BuilderName = "std"
		}

		if k == "gnome-vfs" {
			v.BuilderName = "std"
		}

		if k == "gnome-vfsmm" {
			v.BuilderName = "std"
		}

		if k == "gnome-video-arcade" {
			v.BuilderName = "std"
		}

		if k == "gnome-video-effects" {
			v.BuilderName = "std"
		}

		if k == "gnome-volume-manager" {
			v.BuilderName = "std"
		}

		if k == "gnome-weather" {
			v.BuilderName = "std"
		}

		if k == "gnome-web-photo" {
			v.BuilderName = "std"
		}

		if k == "gnome-xcf-thumbnailer" {
			v.BuilderName = "std"
		}

		if k == "gnome" {
			v.BuilderName = "std"
		}

		if k == "gnome2-user-docs" {
			v.BuilderName = "std"
		}

		if k == "gnomebaker" {
			v.BuilderName = "std"
		}

		if k == "gnomeicu" {
			v.BuilderName = "std"
		}

		if k == "gnomemeeting" {
			v.BuilderName = "std"
		}

		if k == "gnomemm-all" {
			v.BuilderName = "std"
		}

		if k == "gnomemm" {
			v.BuilderName = "std"
		}

		if k == "gnomemm_hello" {
			v.BuilderName = "std"
		}

		if k == "gnomoku" {
			v.BuilderName = "std"
		}

		if k == "gnonlin" {
			v.BuilderName = "std"
		}

		if k == "gnopernicus" {
			v.BuilderName = "std"
		}

		if k == "gnorpm" {
			v.BuilderName = "std"
		}

		if k == "gnote" {
			v.BuilderName = "std"
		}

		if k == "gnotepad+" {
			v.BuilderName = "std"
		}

		if k == "gnu-c-manual" {
			v.BuilderName = "std"
		}

		if k == "gnu-cobol" {
			v.BuilderName = "std"
		}

		if k == "gnu-crypto" {
			v.BuilderName = "std"
		}

		if k == "gnu-ghostscript" {
			v.BuilderName = "std"
		}

		if k == "gnu-gs-fonts-other" {
			v.BuilderName = "std"
		}

		if k == "gnu-gs-fonts-std" {
			v.BuilderName = "std"
		}

		if k == "gnu-gs" {
			v.BuilderName = "std"
		}

		if k == "gnu-objc-issues" {
			v.BuilderName = "std"
		}

		if k == "gnu-pw-mgr" {
			v.BuilderName = "std"
		}

		if k == "gnu-radius" {
			v.BuilderName = "std"
		}

		if k == "gnubatch" {
			v.BuilderName = "std"
		}

		if k == "gnubik" {
			v.BuilderName = "std"
		}

		if k == "gnucap" {
			v.BuilderName = "std"
		}

		if k == "gnucash" {
			v.BuilderName = "std"
		}

		if k == "gnuchess" {
			v.BuilderName = "std"
		}

		if k == "gnudos" {
			v.BuilderName = "std"
		}

		if k == "gnue-appserver" {
			v.BuilderName = "std"
		}

		if k == "gnue-common" {
			v.BuilderName = "std"
		}

		if k == "gnue-forms" {
			v.BuilderName = "std"
		}

		if k == "gnue-navigator" {
			v.BuilderName = "std"
		}

		if k == "gnue-reports" {
			v.BuilderName = "std"
		}

		if k == "gnuedma" {
			v.BuilderName = "std"
		}

		if k == "gnufdisk" {
			v.BuilderName = "std"
		}

		if k == "gnugo" {
			v.BuilderName = "std"
		}

		if k == "gnuhealth" {
			v.BuilderName = "std"
		}

		if k == "gnuhealth_patchset" {
			v.BuilderName = "std"
		}

		if k == "gnuit" {
			v.BuilderName = "std"
		}

		if k == "gnujump" {
			v.BuilderName = "std"
		}

		if k == "gnumach" {
			v.BuilderName = "std"
		}

		if k == "gnumeric" {
			v.BuilderName = "std"
		}

		if k == "gnun" {
			v.BuilderName = "std"
		}

		if k == "gnunet-fuse" {
			v.BuilderName = "std"
		}

		if k == "gnunet-gtk" {
			v.BuilderName = "std"
		}

		if k == "gnunet-java" {
			v.BuilderName = "std"
		}

		if k == "gnunet-qt" {
			v.BuilderName = "std"
		}

		if k == "gnunet" {
			v.BuilderName = "std"
		}

		if k == "gnupg1" {
			v.BuilderName = "std"
		}

		if k == "gnupg2" {
			v.BuilderName = "std"
		}

		if k == "gnuplot" {
			v.BuilderName = "std"
		}

		if k == "gnupod" {
			v.BuilderName = "std"
		}

		if k == "gnuprolog-java" {
			v.BuilderName = "std"
		}

		if k == "gnuradio-core" {
			v.BuilderName = "std"
		}

		if k == "gnuradio-examples" {
			v.BuilderName = "std"
		}

		if k == "gnuradio" {
			v.BuilderName = "std"
		}

		if k == "gnurobots" {
			v.BuilderName = "std"
		}

		if k == "gnuschool" {
			v.BuilderName = "std"
		}

		if k == "gnushogi" {
			v.BuilderName = "std"
		}

		if k == "gnusound" {
			v.BuilderName = "std"
		}

		if k == "gnuspool" {
			v.BuilderName = "std"
		}

		if k == "gnustep-back" {
			v.BuilderName = ""
		}

		if k == "gnustep-base" {
			v.BuilderName = ""
		}

		if k == "gnustep-examples" {
			v.BuilderName = ""
		}

		if k == "gnustep-gui" {
			v.BuilderName = ""
		}

		if k == "gnustep-make" {
			v.BuilderName = ""
		}

		if k == "gnustep-startup" {
			v.BuilderName = ""
		}

		if k == "gnutls" {
			v.BuilderName = "gnutls"
		}

		if k == "gnutrition" {
			v.BuilderName = "std"
		}

		if k == "go" {
			v.BuilderName = "go"
		}

		if k == "gob" {
			v.BuilderName = "std"
		}

		if k == "gob2" {
			v.BuilderName = "std"
		}

		if k == "gobject-introspection" {
			v.BuilderName = "std"
		}

		if k == "goffice" {
			v.BuilderName = "std"
		}

		if k == "gok" {
			v.BuilderName = "std"
		}

		if k == "gom" {
			v.BuilderName = "std"
		}

		if k == "goobox" {
			v.BuilderName = "std"
		}

		if k == "goocanvas" {
			v.BuilderName = "std"
		}

		if k == "goocanvasmm" {
			v.BuilderName = "std"
		}

		if k == "googletest" {
			v.BuilderName = "std"
		}

		if k == "googlizer" {
			v.BuilderName = "std"
		}

		if k == "gopersist" {
			v.BuilderName = "std"
		}

		if k == "goptical" {
			v.BuilderName = "std"
		}

		if k == "gossip" {
			v.BuilderName = "std"
		}

		if k == "gource" {
			v.BuilderName = "std"
		}

		if k == "gpa" {
			v.BuilderName = "std"
		}

		if k == "gpc" {
			v.BuilderName = ""
		}

		if k == "gpdf" {
			v.BuilderName = "std"
		}

		if k == "gperf" {
			v.BuilderName = "std"
		}

		if k == "gpgme" {
			v.BuilderName = "gpgme"
		}

		if k == "gphoto2" {
			v.BuilderName = "std"
		}

		if k == "gphotofs" {
			v.BuilderName = "std"
		}

		if k == "gprolog" {
			v.BuilderName = "std"
		}

		if k == "gr-audio-alsa" {
			v.BuilderName = "std"
		}

		if k == "gr-audio-jack" {
			v.BuilderName = "std"
		}

		if k == "gr-audio-oss" {
			v.BuilderName = "std"
		}

		if k == "gr-audio-portaudio" {
			v.BuilderName = "std"
		}

		if k == "gr-gsm-fr-vocoder" {
			v.BuilderName = "std"
		}

		if k == "gr-howto-write-a-block" {
			v.BuilderName = "std"
		}

		if k == "gr-mc4020" {
			v.BuilderName = "std"
		}

		if k == "gr-usrp" {
			v.BuilderName = "std"
		}

		if k == "gr-wxgui" {
			v.BuilderName = "std"
		}

		if k == "grafx2" {
			v.BuilderName = ""
		}

		if k == "grandr" {
			v.BuilderName = "std"
		}

		if k == "graphviz" {
			v.BuilderName = "std"
		}

		if k == "greg" {
			v.BuilderName = "std"
		}

		if k == "grep" {
			v.BuilderName = "std"
		}

		if k == "grfcodec" {
			v.BuilderName = ""
		}

		if k == "grilo-plugins" {
			v.BuilderName = "std"
		}

		if k == "grilo" {
			v.BuilderName = "std"
		}

		if k == "groff" {
			v.BuilderName = "std"
		}

		if k == "grokmirror" {
			v.BuilderName = "std"
		}

		if k == "groupserver" {
			v.BuilderName = ""
		}

		if k == "grub" {
			v.BuilderName = "std"
		}

		if k == "grub2" {
			v.BuilderName = "grub2"
		}

		if k == "gsasl" {
			v.BuilderName = "std"
		}

		if k == "gsegrafix" {
			v.BuilderName = "std"
		}

		if k == "gsettings-desktop-schemas" {
			v.BuilderName = "std"
		}

		if k == "gsl" {
			v.BuilderName = "std"
		}

		if k == "gsm" {
			v.BuilderName = ""
		}

		if k == "gsound" {
			v.BuilderName = "std"
		}

		if k == "gsrc" {
			v.BuilderName = "std"
		}

		if k == "gss" {
			v.BuilderName = "std"
		}

		if k == "gssdp" {
			v.BuilderName = "std"
		}

		if k == "gst-editor" {
			v.BuilderName = "std"
		}

		if k == "gst-ffmpeg" {
			v.BuilderName = "std"
		}

		if k == "gst-ffmpeg0.10" {
			v.BuilderName = "std"
		}

		if k == "gst-libav" {
			v.BuilderName = "std"
		}

		if k == "gst-libav0.10" {
			v.BuilderName = "std"
		}

		if k == "gst-monkeysaudio" {
			v.BuilderName = "std"
		}

		if k == "gst-omx" {
			v.BuilderName = "std"
		}

		if k == "gst-openmax" {
			v.BuilderName = "std"
		}

		if k == "gst-openmax0.10" {
			v.BuilderName = "std"
		}

		if k == "gst-player" {
			v.BuilderName = "std"
		}

		if k == "gst-plugin" {
			v.BuilderName = "std"
		}

		if k == "gst-plugins-bad" {
			v.BuilderName = "std"
		}

		if k == "gst-plugins-bad0.10" {
			v.BuilderName = "std"
		}

		if k == "gst-plugins-base" {
			v.BuilderName = "gst_plugins_base"
		}

		if k == "gst-plugins-base0.10" {
			v.BuilderName = "std"
		}

		if k == "gst-plugins-farsight" {
			v.BuilderName = "std"
		}

		if k == "gst-plugins-gl" {
			v.BuilderName = "std"
		}

		if k == "gst-plugins-gl0.10" {
			v.BuilderName = "std"
		}

		if k == "gst-plugins-good" {
			v.BuilderName = "std"
		}

		if k == "gst-plugins-good0.10" {
			v.BuilderName = "std"
		}

		if k == "gst-plugins-ugly" {
			v.BuilderName = "std"
		}

		if k == "gst-plugins-ugly0.10" {
			v.BuilderName = "std"
		}

		if k == "gst-plugins" {
			v.BuilderName = "std"
		}

		if k == "gst-plugins0.10" {
			v.BuilderName = "std"
		}

		if k == "gst-python" {
			v.BuilderName = "std"
		}

		if k == "gst-python0.10" {
			v.BuilderName = "std"
		}

		if k == "gst-rtsp-server" {
			v.BuilderName = "std"
		}

		if k == "gst-rtsp" {
			v.BuilderName = "std"
		}

		if k == "gst-rtsp0.10" {
			v.BuilderName = "std"
		}

		if k == "gst-validate" {
			v.BuilderName = "std"
		}

		if k == "gstreamer-editing-services" {
			v.BuilderName = "std"
		}

		if k == "gstreamer-filters" {
			v.BuilderName = "std"
		}

		if k == "gstreamer-sharp" {
			v.BuilderName = "std"
		}

		if k == "gstreamer" {
			v.BuilderName = "std"
		}

		if k == "gstreamer0.10" {
			v.BuilderName = "std"
		}

		if k == "gstreamermm" {
			v.BuilderName = "std"
		}

		if k == "gswitchit-plugins" {
			v.BuilderName = "std"
		}

		if k == "gtetrinet" {
			v.BuilderName = "std"
		}

		if k == "gthumb" {
			v.BuilderName = "std"
		}

		if k == "gtk+" {
			v.BuilderName = "gtk3"
		}

		if k == "gtk+2" {
			v.BuilderName = "std"
		}

		if k == "gtk-css-engine" {
			v.BuilderName = "std"
		}

		if k == "gtk-doc" {
			v.BuilderName = "std"
		}

		if k == "gtk-engines" {
			v.BuilderName = "std"
		}

		if k == "gtk-mac-bundler" {
			v.BuilderName = "std"
		}

		if k == "gtk-mac-integration" {
			v.BuilderName = "std"
		}

		if k == "gtk-sharp" {
			v.BuilderName = "std"
		}

		if k == "gtk-theme-engine-clearlooks" {
			v.BuilderName = "std"
		}

		if k == "gtk-thinice-engine" {
			v.BuilderName = "std"
		}

		if k == "gtk-vnc" {
			v.BuilderName = "std"
		}

		if k == "gtk-xfce-engine" {
			v.BuilderName = ""
		}

		if k == "gtkglarea" {
			v.BuilderName = "std"
		}

		if k == "gtkglext" {
			v.BuilderName = "std"
		}

		if k == "gtkglextmm" {
			v.BuilderName = "std"
		}

		if k == "gtkhtml" {
			v.BuilderName = "std"
		}

		if k == "gtkmm-documentation" {
			v.BuilderName = "std"
		}

		if k == "gtkmm" {
			v.BuilderName = "std"
		}

		if k == "gtkmm2" {
			v.BuilderName = "std"
		}

		if k == "gtkmm_hello" {
			v.BuilderName = "std"
		}

		if k == "gtkmozedit" {
			v.BuilderName = "std"
		}

		if k == "gtkmozembedmm" {
			v.BuilderName = "std"
		}

		if k == "gtksourceview" {
			v.BuilderName = "std"
		}

		if k == "gtksourceviewmm" {
			v.BuilderName = "std"
		}

		if k == "gtkspell" {
			v.BuilderName = "std"
		}

		if k == "gtkspell3" {
			v.BuilderName = "std"
		}

		if k == "gtm" {
			v.BuilderName = "std"
		}

		if k == "gtop" {
			v.BuilderName = "std"
		}

		if k == "gtranslator" {
			v.BuilderName = "std"
		}

		if k == "gturing" {
			v.BuilderName = "std"
		}

		if k == "gtypist" {
			v.BuilderName = "std"
		}

		if k == "guake" {
			v.BuilderName = "std"
		}

		if k == "gucharmap" {
			v.BuilderName = "std"
		}

		if k == "guile-clutter" {
			v.BuilderName = "std"
		}

		if k == "guile-gnome-gstreamer" {
			v.BuilderName = "std"
		}

		if k == "guile-gnome-gtksourceview" {
			v.BuilderName = "std"
		}

		if k == "guile-gnome-platform" {
			v.BuilderName = "std"
		}

		if k == "guile-gobject" {
			v.BuilderName = "std"
		}

		if k == "guile-gtk" {
			v.BuilderName = "std"
		}

		if k == "guile-ncurses" {
			v.BuilderName = "std"
		}

		if k == "guile-oops" {
			v.BuilderName = "std"
		}

		if k == "guile-opengl" {
			v.BuilderName = "std"
		}

		if k == "guile-rpc" {
			v.BuilderName = "std"
		}

		if k == "guile-sdl" {
			v.BuilderName = "std"
		}

		if k == "guile-www" {
			v.BuilderName = "std"
		}

		if k == "guile" {
			v.BuilderName = "guile"
		}

		if k == "gupnp-av" {
			v.BuilderName = "std"
		}

		if k == "gupnp-dlna" {
			v.BuilderName = "std"
		}

		if k == "gupnp-igd" {
			v.BuilderName = "std"
		}

		if k == "gupnp-tools" {
			v.BuilderName = "std"
		}

		if k == "gupnp-ui" {
			v.BuilderName = "std"
		}

		if k == "gupnp-vala" {
			v.BuilderName = "std"
		}

		if k == "gupnp" {
			v.BuilderName = "std"
		}

		if k == "gutenprint" {
			v.BuilderName = "std"
		}

		if k == "gv" {
			v.BuilderName = "std"
		}

		if k == "gvfs" {
			v.BuilderName = "gvfs"
		}

		if k == "gvpe" {
			v.BuilderName = "std"
		}

		if k == "gwenview" {
			v.BuilderName = ""
		}

		if k == "gwget" {
			v.BuilderName = "std"
		}

		if k == "gwt-glom" {
			v.BuilderName = "std"
		}

		if k == "gxmessage" {
			v.BuilderName = "std"
		}

		if k == "gxml" {
			v.BuilderName = "std"
		}

		if k == "gyrus" {
			v.BuilderName = "std"
		}

		if k == "gzip" {
			v.BuilderName = "std"
		}

		if k == "hal-info" {
			v.BuilderName = "std"
		}

		if k == "hal" {
			v.BuilderName = "std"
		}

		if k == "hamcrest-java" {
			v.BuilderName = "hamcrest_java"
		}

		if k == "hamster-applet" {
			v.BuilderName = "std"
		}

		if k == "harfbuzz" {
			v.BuilderName = "std"
		}

		if k == "haskell-platform" {
			v.BuilderName = ""
		}

		if k == "haze" {
			v.BuilderName = "std"
		}

		if k == "hdparm" {
			v.BuilderName = "hdparm"
		}

		if k == "hello" {
			v.BuilderName = "std"
		}

		if k == "help2man" {
			v.BuilderName = "std"
		}

		if k == "hicolor-icon-theme" {
			v.BuilderName = "std"
		}

		if k == "hippo-canvas" {
			v.BuilderName = "std"
		}

		if k == "hitori" {
			v.BuilderName = "std"
		}

		if k == "hostname" {
			v.BuilderName = "hostname"
		}

		if k == "hotplug-ng" {
			v.BuilderName = "std"
		}

		if k == "hotplug" {
			v.BuilderName = "std"
		}

		if k == "hotssh" {
			v.BuilderName = "std"
		}

		if k == "hp2xx" {
			v.BuilderName = "std"
		}

		if k == "hplip" {
			v.BuilderName = "hplip"
		}

		if k == "ht" {
			v.BuilderName = "std"
		}

		if k == "httpd" {
			v.BuilderName = "httpd"
		}

		if k == "httptunnel" {
			v.BuilderName = "std"
		}

		if k == "hurd" {
			v.BuilderName = "std"
		}

		if k == "hyena" {
			v.BuilderName = "std"
		}

		if k == "hyperbole" {
			v.BuilderName = "std"
		}

		if k == "iMule" {
			v.BuilderName = ""
		}

		if k == "iagno" {
			v.BuilderName = "std"
		}

		if k == "iana" {
			v.BuilderName = ""
		}

		if k == "iat" {
			v.BuilderName = "std"
		}

		if k == "ibus-anthy" {
			v.BuilderName = ""
		}

		if k == "ibus-chewing" {
			v.BuilderName = ""
		}

		if k == "ibus-gjs" {
			v.BuilderName = ""
		}

		if k == "ibus-hangul" {
			v.BuilderName = ""
		}

		if k == "ibus-m17n" {
			v.BuilderName = ""
		}

		if k == "ibus-pinyin" {
			v.BuilderName = ""
		}

		if k == "ibus-qt" {
			v.BuilderName = ""
		}

		if k == "ibus-table-chinese" {
			v.BuilderName = ""
		}

		if k == "ibus-table-extraphrase" {
			v.BuilderName = ""
		}

		if k == "ibus-table-others" {
			v.BuilderName = ""
		}

		if k == "ibus-table" {
			v.BuilderName = ""
		}

		if k == "ibus-xkb" {
			v.BuilderName = ""
		}

		if k == "ibus" {
			v.BuilderName = "ibus"
		}

		if k == "iceauth" {
			v.BuilderName = "std"
		}

		if k == "icecast" {
			v.BuilderName = "std"
		}

		if k == "icedtea-sound" {
			v.BuilderName = "icedtea_sound"
		}

		if k == "icedtea-web" {
			v.BuilderName = "icedtea_web"
		}

		if k == "icedtea" {
			v.BuilderName = "icedtea"
		}

		if k == "ices" {
			v.BuilderName = "std"
		}

		if k == "ico" {
			v.BuilderName = "std"
		}

		if k == "icon-naming-utils" {
			v.BuilderName = "std"
		}

		if k == "icu4c" {
			v.BuilderName = "icu4c"
		}

		if k == "id-utils" {
			v.BuilderName = "std"
		}

		if k == "idutils" {
			v.BuilderName = "std"
		}

		if k == "ignuit" {
			v.BuilderName = "std"
		}

		if k == "ijs" {
			v.BuilderName = "std"
		}

		if k == "ilmbase" {
			v.BuilderName = "std"
		}

		if k == "image-analyzer" {
			v.BuilderName = "std_cmake"
		}

		if k == "imagemagick" {
			v.BuilderName = "std"
		}

		if k == "imagick" {
			v.BuilderName = ""
		}

		if k == "imake" {
			v.BuilderName = "std"
		}

		if k == "imap" {
			v.BuilderName = ""
		}

		if k == "imlib" {
			v.BuilderName = "std"
		}

		if k == "include" {
			v.BuilderName = "std"
		}

		if k == "indent" {
			v.BuilderName = "std"
		}

		if k == "indri" {
			v.BuilderName = ""
		}

		if k == "inetlib" {
			v.BuilderName = "std"
		}

		if k == "inetutils" {
			v.BuilderName = "inetutils"
		}

		if k == "inkscape" {
			v.BuilderName = "inkscape"
		}

		if k == "inn" {
			v.BuilderName = "inn"
		}

		if k == "inputproto" {
			v.BuilderName = "std"
		}

		if k == "inspircd" {
			v.BuilderName = "inspircd"
		}

		if k == "intel-gpu-tools" {
			v.BuilderName = "std"
		}

		if k == "intlfonts" {
			v.BuilderName = "std"
		}

		if k == "intltool" {
			v.BuilderName = "std"
		}

		if k == "iogrind" {
			v.BuilderName = "std"
		}

		if k == "iperf" {
			v.BuilderName = "std"
		}

		if k == "iproute2" {
			v.BuilderName = "iproute2"
		}

		if k == "ipsec-tools" {
			v.BuilderName = "ipsec_tools"
		}

		if k == "iptables" {
			v.BuilderName = "std"
		}

		if k == "iptraf" {
			v.BuilderName = ""
		}

		if k == "ipvsadm" {
			v.BuilderName = "std"
		}

		if k == "irrlicht" {
			v.BuilderName = ""
		}

		if k == "isl" {
			v.BuilderName = "std"
		}

		if k == "iso-codes" {
			v.BuilderName = "std"
		}

		if k == "ispell" {
			v.BuilderName = "std"
		}

		if k == "italc" {
			v.BuilderName = ""
		}

		if k == "itstool" {
			v.BuilderName = "itstool"
		}

		if k == "iw" {
			v.BuilderName = "std"
		}

		if k == "jabberd1" {
			v.BuilderName = "std"
		}

		if k == "jabberd2" {
			v.BuilderName = "jabberd2"
		}

		if k == "jacal" {
			v.BuilderName = "std"
		}

		if k == "jack-rack" {
			v.BuilderName = "std"
		}

		if k == "jack0" {
			v.BuilderName = "jack"
		}

		if k == "jack1" {
			v.BuilderName = "jack"
		}

		if k == "jack2" {
			v.BuilderName = "jack"
		}

		if k == "jamboree" {
			v.BuilderName = "std"
		}

		if k == "jasper" {
			v.BuilderName = "std"
		}

		if k == "jato" {
			v.BuilderName = "std"
		}

		if k == "java-access-bridge" {
			v.BuilderName = "std"
		}

		if k == "java-atk-wrapper" {
			v.BuilderName = "std"
		}

		if k == "java-gnome" {
			v.BuilderName = "std"
		}

		if k == "java-libglom" {
			v.BuilderName = "std"
		}

		if k == "jaxp" {
			v.BuilderName = "std"
		}

		if k == "jel" {
			v.BuilderName = "std"
		}

		if k == "jhbuild" {
			v.BuilderName = "std"
		}

		if k == "jitsi" {
			v.BuilderName = ""
		}

		if k == "jojodiff" {
			v.BuilderName = ""
		}

		if k == "jovie" {
			v.BuilderName = ""
		}

		if k == "jp2a" {
			v.BuilderName = "std"
		}

		if k == "jpeg" {
			v.BuilderName = "None"
		}

		if k == "jpegsrc.v" {
			v.BuilderName = "std"
		}

		if k == "js185" {
			v.BuilderName = "js185"
		}

		if k == "jsdoc" {
			v.BuilderName = ""
		}

		if k == "json-c" {
			v.BuilderName = "std"
		}

		if k == "json-glib" {
			v.BuilderName = "std"
		}

		if k == "junit" {
			v.BuilderName = "junit"
		}

		if k == "jwhois" {
			v.BuilderName = "std"
		}

		if k == "kaccessible" {
			v.BuilderName = ""
		}

		if k == "kactivities" {
			v.BuilderName = ""
		}

		if k == "kadu" {
			v.BuilderName = "std_cmake"
		}

		if k == "kalgebra" {
			v.BuilderName = ""
		}

		if k == "kalzium" {
			v.BuilderName = ""
		}

		if k == "kamera" {
			v.BuilderName = ""
		}

		if k == "kanagram" {
			v.BuilderName = ""
		}

		if k == "kate" {
			v.BuilderName = ""
		}

		if k == "kawa-doc" {
			v.BuilderName = "std"
		}

		if k == "kawa" {
			v.BuilderName = "std"
		}

		if k == "kbd" {
			v.BuilderName = "kbd"
		}

		if k == "kbdraw" {
			v.BuilderName = "std"
		}

		if k == "kbproto" {
			v.BuilderName = "std"
		}

		if k == "kbruch" {
			v.BuilderName = ""
		}

		if k == "kcalc" {
			v.BuilderName = ""
		}

		if k == "kcharselect" {
			v.BuilderName = ""
		}

		if k == "kcolorchooser" {
			v.BuilderName = ""
		}

		if k == "kde-baseapps" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-ar" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-bg" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-bs" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-ca" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-ca@valencia" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-cs" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-da" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-de" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-el" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-en_GB" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-es" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-et" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-eu" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-fa" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-fi" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-fr" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-ga" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-gl" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-hr" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-hu" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-ia" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-is" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-it" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-ja" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-kk" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-km" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-ko" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-lt" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-lv" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-nb" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-nds" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-nl" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-nn" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-pa" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-pl" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-pt" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-pt_BR" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-ro" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-ru" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-si" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-sk" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-sl" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-sr" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-sv" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-tg" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-th" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-tr" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-uk" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-vi" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-wa" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-zh_CN" {
			v.BuilderName = ""
		}

		if k == "kde-l10n-zh_TW" {
			v.BuilderName = ""
		}

		if k == "kde-runtime" {
			v.BuilderName = ""
		}

		if k == "kde-wallpapers" {
			v.BuilderName = ""
		}

		if k == "kde-workspace" {
			v.BuilderName = ""
		}

		if k == "kdeadmin" {
			v.BuilderName = ""
		}

		if k == "kdeartwork" {
			v.BuilderName = ""
		}

		if k == "kdegames" {
			v.BuilderName = ""
		}

		if k == "kdegraphics-mobipocket" {
			v.BuilderName = ""
		}

		if k == "kdegraphics-strigi-analyzer" {
			v.BuilderName = ""
		}

		if k == "kdegraphics-thumbnailers" {
			v.BuilderName = ""
		}

		if k == "kdelibs" {
			v.BuilderName = ""
		}

		if k == "kdemultimedia" {
			v.BuilderName = ""
		}

		if k == "kdenetwork" {
			v.BuilderName = ""
		}

		if k == "kdenlive" {
			v.BuilderName = ""
		}

		if k == "kdepim-runtime" {
			v.BuilderName = ""
		}

		if k == "kdepim" {
			v.BuilderName = ""
		}

		if k == "kdepimlibs" {
			v.BuilderName = ""
		}

		if k == "kdeplasma-addons" {
			v.BuilderName = ""
		}

		if k == "kdesdk" {
			v.BuilderName = ""
		}

		if k == "kdetoys" {
			v.BuilderName = ""
		}

		if k == "kdevelop" {
			v.BuilderName = ""
		}

		if k == "kdevplatform" {
			v.BuilderName = ""
		}

		if k == "kdewebdev" {
			v.BuilderName = ""
		}

		if k == "kdf" {
			v.BuilderName = ""
		}

		if k == "kernel-headers" {
			v.BuilderName = "std"
		}

		if k == "kexec-tools" {
			v.BuilderName = "std"
		}

		if k == "kfloppy" {
			v.BuilderName = ""
		}

		if k == "kgamma" {
			v.BuilderName = ""
		}

		if k == "kgeography" {
			v.BuilderName = ""
		}

		if k == "kgpg" {
			v.BuilderName = ""
		}

		if k == "khangman" {
			v.BuilderName = ""
		}

		if k == "kig" {
			v.BuilderName = ""
		}

		if k == "kimono" {
			v.BuilderName = ""
		}

		if k == "kino" {
			v.BuilderName = ""
		}

		if k == "kismet" {
			v.BuilderName = "kismet"
		}

		if k == "kiten" {
			v.BuilderName = ""
		}

		if k == "kiwi" {
			v.BuilderName = "std"
		}

		if k == "klettres" {
			v.BuilderName = ""
		}

		if k == "klibc" {
			v.BuilderName = "std"
		}

		if k == "kmag" {
			v.BuilderName = ""
		}

		if k == "kmod" {
			v.BuilderName = "kmod"
		}

		if k == "kmousetool" {
			v.BuilderName = ""
		}

		if k == "kmouth" {
			v.BuilderName = ""
		}

		if k == "kmplot" {
			v.BuilderName = ""
		}

		if k == "kolourpaint" {
			v.BuilderName = ""
		}

		if k == "konsole" {
			v.BuilderName = ""
		}

		if k == "korundum" {
			v.BuilderName = ""
		}

		if k == "krb5-auth-dialog" {
			v.BuilderName = "std"
		}

		if k == "krb5" {
			v.BuilderName = "krb5"
		}

		if k == "kremotecontrol" {
			v.BuilderName = ""
		}

		if k == "kross-interpreters" {
			v.BuilderName = ""
		}

		if k == "kruler" {
			v.BuilderName = ""
		}

		if k == "ksaneplugin" {
			v.BuilderName = ""
		}

		if k == "ksecrets" {
			v.BuilderName = ""
		}

		if k == "ksnapshot" {
			v.BuilderName = ""
		}

		if k == "kstars" {
			v.BuilderName = ""
		}

		if k == "ksymoops" {
			v.BuilderName = "std"
		}

		if k == "ktimer" {
			v.BuilderName = ""
		}

		if k == "ktouch" {
			v.BuilderName = ""
		}

		if k == "kturtle" {
			v.BuilderName = ""
		}

		if k == "kup" {
			v.BuilderName = "std"
		}

		if k == "kvm" {
			v.BuilderName = ""
		}

		if k == "kwallet" {
			v.BuilderName = ""
		}

		if k == "kwordquiz" {
			v.BuilderName = ""
		}

		if k == "labyrinth" {
			v.BuilderName = "std"
		}

		if k == "ladspa" {
			v.BuilderName = ""
		}

		if k == "lame" {
			v.BuilderName = "lame"
		}

		if k == "lasem" {
			v.BuilderName = "std"
		}

		if k == "latexila" {
			v.BuilderName = "std"
		}

		if k == "lbxproxy" {
			v.BuilderName = "std"
		}

		if k == "lcms" {
			v.BuilderName = "std"
		}

		if k == "lcms2" {
			v.BuilderName = "std"
		}

		if k == "lector" {
			v.BuilderName = ""
		}

		if k == "leim" {
			v.BuilderName = "std"
		}

		if k == "lemuria" {
			v.BuilderName = "std"
		}

		if k == "leptonica" {
			v.BuilderName = ""
		}

		if k == "less" {
			v.BuilderName = "less"
		}

		if k == "lfm" {
			v.BuilderName = ""
		}

		if k == "libAppleWM" {
			v.BuilderName = "std"
		}

		if k == "libFS" {
			v.BuilderName = "std"
		}

		if k == "libICE" {
			v.BuilderName = "std"
		}

		if k == "libIDL" {
			v.BuilderName = "std"
		}

		if k == "libPropList" {
			v.BuilderName = "std"
		}

		if k == "libSM" {
			v.BuilderName = "std"
		}

		if k == "libWindowsWM" {
			v.BuilderName = "std"
		}

		if k == "libX11" {
			v.BuilderName = "std"
		}

		if k == "libXScrnSaver" {
			v.BuilderName = "std"
		}

		if k == "libXTrap" {
			v.BuilderName = "std"
		}

		if k == "libXau" {
			v.BuilderName = "std"
		}

		if k == "libXaw" {
			v.BuilderName = "std"
		}

		if k == "libXaw3d" {
			v.BuilderName = "std"
		}

		if k == "libXcomposite" {
			v.BuilderName = "std"
		}

		if k == "libXcursor" {
			v.BuilderName = "std"
		}

		if k == "libXdamage" {
			v.BuilderName = "std"
		}

		if k == "libXdmcp" {
			v.BuilderName = "std"
		}

		if k == "libXevie" {
			v.BuilderName = "std"
		}

		if k == "libXext" {
			v.BuilderName = "std"
		}

		if k == "libXfixes" {
			v.BuilderName = "std"
		}

		if k == "libXfont" {
			v.BuilderName = "std"
		}

		if k == "libXfont2" {
			v.BuilderName = "std"
		}

		if k == "libXfontcache" {
			v.BuilderName = "std"
		}

		if k == "libXft" {
			v.BuilderName = "std"
		}

		if k == "libXi" {
			v.BuilderName = "std"
		}

		if k == "libXinerama" {
			v.BuilderName = "std"
		}

		if k == "libXmu" {
			v.BuilderName = "std"
		}

		if k == "libXp" {
			v.BuilderName = "std"
		}

		if k == "libXpm" {
			v.BuilderName = "std"
		}

		if k == "libXpresent" {
			v.BuilderName = "std"
		}

		if k == "libXprintAppUtil" {
			v.BuilderName = "std"
		}

		if k == "libXprintUtil" {
			v.BuilderName = "std"
		}

		if k == "libXrandr" {
			v.BuilderName = "std"
		}

		if k == "libXrender" {
			v.BuilderName = "std"
		}

		if k == "libXres" {
			v.BuilderName = "std"
		}

		if k == "libXt" {
			v.BuilderName = "std"
		}

		if k == "libXtst" {
			v.BuilderName = "std"
		}

		if k == "libXv" {
			v.BuilderName = "std"
		}

		if k == "libXvMC" {
			v.BuilderName = "std"
		}

		if k == "libXxf86dga" {
			v.BuilderName = "std"
		}

		if k == "libXxf86misc" {
			v.BuilderName = "std"
		}

		if k == "libXxf86vm" {
			v.BuilderName = "std"
		}

		if k == "liba52" {
			v.BuilderName = ""
		}

		if k == "libaccounts-glib" {
			v.BuilderName = "std"
		}

		if k == "libaio" {
			v.BuilderName = "libaio"
		}

		if k == "libao" {
			v.BuilderName = "std"
		}

		if k == "libarchive" {
			v.BuilderName = "std"
		}

		if k == "libart" {
			v.BuilderName = ""
		}

		if k == "libart_lgpl" {
			v.BuilderName = "std"
		}

		if k == "libasn1" {
			v.BuilderName = "std"
		}

		if k == "libass" {
			v.BuilderName = "std"
		}

		if k == "libassuan" {
			v.BuilderName = "libassuan"
		}

		if k == "libatasmart" {
			v.BuilderName = "std"
		}

		if k == "libatomic_ops" {
			v.BuilderName = "std"
		}

		if k == "libaudit" {
			v.BuilderName = "std"
		}

		if k == "libav" {
			v.BuilderName = ""
		}

		if k == "libavc1394" {
			v.BuilderName = "std"
		}

		if k == "libbeagle" {
			v.BuilderName = "std"
		}

		if k == "libbonobo" {
			v.BuilderName = "std"
		}

		if k == "libbonobomm" {
			v.BuilderName = "std"
		}

		if k == "libbonoboui" {
			v.BuilderName = "std"
		}

		if k == "libbonobouimm" {
			v.BuilderName = "std"
		}

		if k == "libbtctl" {
			v.BuilderName = "std"
		}

		if k == "libc" {
			v.BuilderName = "std"
		}

		if k == "libcanberra" {
			v.BuilderName = "libcanberra"
		}

		if k == "libcap-ng" {
			v.BuilderName = "std"
		}

		if k == "libcap" {
			v.BuilderName = "libcap"
		}

		if k == "libcap2" {
			v.BuilderName = "libcap2"
		}

		if k == "libcapplet" {
			v.BuilderName = "std"
		}

		if k == "libccc" {
			v.BuilderName = "std"
		}

		if k == "libcdio-paranoia" {
			v.BuilderName = "std"
		}

		if k == "libcdio" {
			v.BuilderName = "std"
		}

		if k == "libchamplain" {
			v.BuilderName = "std"
		}

		if k == "libclc" {
			v.BuilderName = "std_configure_py2"
		}

		if k == "libcm" {
			v.BuilderName = "std"
		}

		if k == "libcroco" {
			v.BuilderName = "std"
		}

		if k == "libcryptui" {
			v.BuilderName = "std"
		}

		if k == "libcutl" {
			v.BuilderName = "std"
		}

		if k == "libcxx" {
			v.BuilderName = "llvm_components"
		}

		if k == "libcxxabi" {
			v.BuilderName = "llvm_components"
		}

		if k == "libdaemon" {
			v.BuilderName = "std"
		}

		if k == "libdbus-java" {
			v.BuilderName = "std"
		}

		if k == "libdbusmenu-qt" {
			v.BuilderName = ""
		}

		if k == "libdmx" {
			v.BuilderName = "std"
		}

		if k == "libdrm" {
			v.BuilderName = "libdrm"
		}

		if k == "libdv" {
			v.BuilderName = ""
		}

		if k == "libdvdnav" {
			v.BuilderName = "std"
		}

		if k == "libdvdread" {
			v.BuilderName = "std"
		}

		if k == "libe-book" {
			v.BuilderName = "libe_book"
		}

		if k == "libeXosip2" {
			v.BuilderName = ""
		}

		if k == "libeds-java" {
			v.BuilderName = "std"
		}

		if k == "libepc" {
			v.BuilderName = "std"
		}

		if k == "libepoxy" {
			v.BuilderName = "std"
		}

		if k == "libevdev" {
			v.BuilderName = "std"
		}

		if k == "libevent" {
			v.BuilderName = "std"
		}

		if k == "libexif" {
			v.BuilderName = "std"
		}

		if k == "libextractor-java" {
			v.BuilderName = "std"
		}

		if k == "libextractor-mono" {
			v.BuilderName = "std"
		}

		if k == "libextractor-python" {
			v.BuilderName = "std"
		}

		if k == "libextractor" {
			v.BuilderName = "std"
		}

		if k == "libffi" {
			v.BuilderName = "std"
		}

		if k == "libfishsound" {
			v.BuilderName = "std"
		}

		if k == "libfontenc" {
			v.BuilderName = "std"
		}

		if k == "libgail-gnome" {
			v.BuilderName = "std"
		}

		if k == "libgconf-java" {
			v.BuilderName = "std"
		}

		if k == "libgcrypt" {
			v.BuilderName = "libgcrypt"
		}

		if k == "libgd" {
			v.BuilderName = "std"
		}

		if k == "libgda-uimm" {
			v.BuilderName = "std"
		}

		if k == "libgda" {
			v.BuilderName = "std"
		}

		if k == "libgdamm" {
			v.BuilderName = "std"
		}

		if k == "libgdata" {
			v.BuilderName = "std"
		}

		if k == "libgdiplus" {
			v.BuilderName = "std"
		}

		if k == "libgee" {
			v.BuilderName = "std"
		}

		if k == "libgee0.6" {
			v.BuilderName = "std"
		}

		if k == "libghttp" {
			v.BuilderName = "std"
		}

		if k == "libgif" {
			v.BuilderName = ""
		}

		if k == "libgit2-glib" {
			v.BuilderName = "std"
		}

		if k == "libgit2" {
			v.BuilderName = "std_cmake"
		}

		if k == "libglade-java" {
			v.BuilderName = "std"
		}

		if k == "libglade" {
			v.BuilderName = "std"
		}

		if k == "libglademm" {
			v.BuilderName = "std"
		}

		if k == "libgnetwork" {
			v.BuilderName = "std"
		}

		if k == "libgnome-java" {
			v.BuilderName = "std"
		}

		if k == "libgnome-keyring" {
			v.BuilderName = "std"
		}

		if k == "libgnome-media-profiles" {
			v.BuilderName = "std"
		}

		if k == "libgnome" {
			v.BuilderName = "std"
		}

		if k == "libgnome2" {
			v.BuilderName = "std"
		}

		if k == "libgnomecanvas" {
			v.BuilderName = "std"
		}

		if k == "libgnomecanvas2" {
			v.BuilderName = "std"
		}

		if k == "libgnomecanvasmm" {
			v.BuilderName = "std"
		}

		if k == "libgnomecompat2" {
			v.BuilderName = "std"
		}

		if k == "libgnomecups" {
			v.BuilderName = "std"
		}

		if k == "libgnomedb" {
			v.BuilderName = "std"
		}

		if k == "libgnomedbmm" {
			v.BuilderName = "std"
		}

		if k == "libgnomefilesel" {
			v.BuilderName = "std"
		}

		if k == "libgnomekbd" {
			v.BuilderName = "std"
		}

		if k == "libgnomemm" {
			v.BuilderName = "std"
		}

		if k == "libgnomeprint" {
			v.BuilderName = "std"
		}

		if k == "libgnomeprintmm" {
			v.BuilderName = "std"
		}

		if k == "libgnomeprintui" {
			v.BuilderName = "std"
		}

		if k == "libgnomeprintuimm" {
			v.BuilderName = "std"
		}

		if k == "libgnomesu" {
			v.BuilderName = "std"
		}

		if k == "libgnomeui" {
			v.BuilderName = "std"
		}

		if k == "libgnomeui2" {
			v.BuilderName = "std"
		}

		if k == "libgnomeuimm" {
			v.BuilderName = "std"
		}

		if k == "libgovirt" {
			v.BuilderName = "std"
		}

		if k == "libgpg-error" {
			v.BuilderName = "std"
		}

		if k == "libgphoto2" {
			v.BuilderName = "std"
		}

		if k == "libgsasl" {
			v.BuilderName = "std"
		}

		if k == "libgsf" {
			v.BuilderName = "std"
		}

		if k == "libgssapi" {
			v.BuilderName = "std"
		}

		if k == "libgssglue" {
			v.BuilderName = "std"
		}

		if k == "libgssh" {
			v.BuilderName = "std"
		}

		if k == "libgsystem" {
			v.BuilderName = "std"
		}

		if k == "libgtcpsocket" {
			v.BuilderName = "std"
		}

		if k == "libgtk-java" {
			v.BuilderName = "std"
		}

		if k == "libgtkhtml-java" {
			v.BuilderName = "std"
		}

		if k == "libgtkhtml" {
			v.BuilderName = "std"
		}

		if k == "libgtkmozembed-java" {
			v.BuilderName = "std"
		}

		if k == "libgtkmusic" {
			v.BuilderName = "std"
		}

		if k == "libgtksourceviewmm" {
			v.BuilderName = "std"
		}

		if k == "libgtop" {
			v.BuilderName = "std"
		}

		if k == "libgudev" {
			v.BuilderName = "std"
		}

		if k == "libgusb" {
			v.BuilderName = "std"
		}

		if k == "libgweather" {
			v.BuilderName = "std"
		}

		if k == "libgxps" {
			v.BuilderName = "std"
		}

		if k == "libical-glib" {
			v.BuilderName = "std"
		}

		if k == "libical" {
			v.BuilderName = "std_cmake"
		}

		if k == "libiconv" {
			v.BuilderName = "std"
		}

		if k == "libid3tag" {
			v.BuilderName = "std"
		}

		if k == "libidn" {
			v.BuilderName = "std"
		}

		if k == "libiec61883" {
			v.BuilderName = "std"
		}

		if k == "libinput" {
			v.BuilderName = "std"
		}

		if k == "libiodbc" {
			v.BuilderName = "std"
		}

		if k == "libircclient" {
			v.BuilderName = "std"
		}

		if k == "libjingle" {
			v.BuilderName = "std"
		}

		if k == "libjpeg-turbo" {
			v.BuilderName = "libjpeg_turbo"
		}

		if k == "libjson" {
			v.BuilderName = "libjson"
		}

		if k == "libkdcraw" {
			v.BuilderName = ""
		}

		if k == "libkdeedu" {
			v.BuilderName = ""
		}

		if k == "libkexiv2" {
			v.BuilderName = ""
		}

		if k == "libkipi" {
			v.BuilderName = ""
		}

		if k == "libksane" {
			v.BuilderName = ""
		}

		if k == "libksba" {
			v.BuilderName = "libksba"
		}

		if k == "liblbxutil" {
			v.BuilderName = "std"
		}

		if k == "libmad" {
			v.BuilderName = "std"
		}

		if k == "libmatekbd" {
			v.BuilderName = "std"
		}

		if k == "libmatemixer" {
			v.BuilderName = "std"
		}

		if k == "libmateweather" {
			v.BuilderName = "std"
		}

		if k == "libmatheval" {
			v.BuilderName = "std"
		}

		if k == "libmbim" {
			v.BuilderName = "std"
		}

		if k == "libmediaart" {
			v.BuilderName = "std"
		}

		if k == "libmicrohttpd" {
			v.BuilderName = "std"
		}

		if k == "libmirage" {
			v.BuilderName = "std_cmake"
		}

		if k == "libmng" {
			v.BuilderName = "std"
		}

		if k == "libmnl" {
			v.BuilderName = "std"
		}

		if k == "libmrproject" {
			v.BuilderName = "std"
		}

		if k == "libmusicbrainz" {
			v.BuilderName = "std_cmake"
		}

		if k == "libndp" {
			v.BuilderName = "std"
		}

		if k == "libnfnetlink" {
			v.BuilderName = "std"
		}

		if k == "libnfsidmap" {
			v.BuilderName = "std"
		}

		if k == "libnftnl" {
			v.BuilderName = "std"
		}

		if k == "libnice" {
			v.BuilderName = "std"
		}

		if k == "libnl" {
			v.BuilderName = "std"
		}

		if k == "libnotify" {
			v.BuilderName = "std"
		}

		if k == "libnotifymm" {
			v.BuilderName = "std"
		}

		if k == "liboauth" {
			v.BuilderName = "std"
		}

		if k == "libodb-boost" {
			v.BuilderName = "std"
		}

		if k == "libodb-mssql" {
			v.BuilderName = "std"
		}

		if k == "libodb-mysql" {
			v.BuilderName = "std"
		}

		if k == "libodb-oracle" {
			v.BuilderName = "std"
		}

		if k == "libodb-pgsql" {
			v.BuilderName = "std"
		}

		if k == "libodb-qt" {
			v.BuilderName = "std"
		}

		if k == "libodb-sqlite" {
			v.BuilderName = "std"
		}

		if k == "libodb" {
			v.BuilderName = "std"
		}

		if k == "libogg" {
			v.BuilderName = "std"
		}

		if k == "liboggz" {
			v.BuilderName = "std"
		}

		if k == "liboil" {
			v.BuilderName = "std"
		}

		if k == "liboldX" {
			v.BuilderName = "std"
		}

		if k == "libole2" {
			v.BuilderName = "std"
		}

		if k == "libomxil-bellagio" {
			v.BuilderName = "std"
		}

		if k == "liboobs" {
			v.BuilderName = "std"
		}

		if k == "liboop" {
			v.BuilderName = "std"
		}

		if k == "libopenraw" {
			v.BuilderName = "std"
		}

		if k == "libopenshot-audio" {
			v.BuilderName = "std_cmake"
		}

		if k == "libopenshot" {
			v.BuilderName = "std_cmake"
		}

		if k == "libopts" {
			v.BuilderName = "std"
		}

		if k == "libosip2" {
			v.BuilderName = "std"
		}

		if k == "libotr" {
			v.BuilderName = "std"
		}

		if k == "libpanelappletmm" {
			v.BuilderName = "std"
		}

		if k == "libpcap" {
			v.BuilderName = "std"
		}

		if k == "libpciaccess" {
			v.BuilderName = "std"
		}

		if k == "libpeas" {
			v.BuilderName = "std"
		}

		if k == "libpipeline" {
			v.BuilderName = "std"
		}

		if k == "libpng" {
			v.BuilderName = "std"
		}

		if k == "libpng10" {
			v.BuilderName = "std"
		}

		if k == "libpng12" {
			v.BuilderName = "std"
		}

		if k == "libpng14" {
			v.BuilderName = "std"
		}

		if k == "libpng15" {
			v.BuilderName = "std"
		}

		if k == "libproxy" {
			v.BuilderName = "std_cmake"
		}

		if k == "libpthread-stubs" {
			v.BuilderName = "std"
		}

		if k == "libpwquality" {
			v.BuilderName = "libpwquality"
		}

		if k == "libqmi" {
			v.BuilderName = "std"
		}

		if k == "libquicktime" {
			v.BuilderName = "libquicktime"
		}

		if k == "libraw1394" {
			v.BuilderName = "std"
		}

		if k == "librejs" {
			v.BuilderName = "std"
		}

		if k == "libreoffice" {
			v.BuilderName = "libreoffice"
		}

		if k == "librep" {
			v.BuilderName = "std"
		}

		if k == "librevenge" {
			v.BuilderName = "std"
		}

		if k == "librpcsecgss" {
			v.BuilderName = "std"
		}

		if k == "librsvg" {
			v.BuilderName = "std"
		}

		if k == "librsvgmm" {
			v.BuilderName = "std"
		}

		if k == "libsamplerate" {
			v.BuilderName = "std"
		}

		if k == "libseccomp" {
			v.BuilderName = "libseccomp"
		}

		if k == "libsecret" {
			v.BuilderName = "std"
		}

		if k == "libselinux" {
			v.BuilderName = "std"
		}

		if k == "libsemanage" {
			v.BuilderName = "std"
		}

		if k == "libsepol" {
			v.BuilderName = "std"
		}

		if k == "libshout" {
			v.BuilderName = "std"
		}

		if k == "libsigc++" {
			v.BuilderName = "std"
		}

		if k == "libsignon-glib" {
			v.BuilderName = "std"
		}

		if k == "libsigsegv" {
			v.BuilderName = "std"
		}

		if k == "libslab" {
			v.BuilderName = "std"
		}

		if k == "libsndfile" {
			v.BuilderName = "std"
		}

		if k == "libsocialweb" {
			v.BuilderName = "std"
		}

		if k == "libsodium" {
			v.BuilderName = "std"
		}

		if k == "libsoup" {
			v.BuilderName = "std"
		}

		if k == "libspectre" {
			v.BuilderName = "std"
		}

		if k == "libspnav" {
			v.BuilderName = "libspnav"
		}

		if k == "libssh" {
			v.BuilderName = "libssh"
		}

		if k == "libtasn1" {
			v.BuilderName = "std"
		}

		if k == "libtelepathy" {
			v.BuilderName = "std"
		}

		if k == "libtheora" {
			v.BuilderName = "libtheora"
		}

		if k == "libtirpc" {
			v.BuilderName = "libtirpc"
		}

		if k == "libtool" {
			v.BuilderName = "libtool"
		}

		if k == "libtorrent-rasterbar" {
			v.BuilderName = "std_cmake"
		}

		if k == "libungif" {
			v.BuilderName = "std"
		}

		if k == "libunicap" {
			v.BuilderName = "std"
		}

		if k == "libunicode" {
			v.BuilderName = "std"
		}

		if k == "libunique" {
			v.BuilderName = "std"
		}

		if k == "libunistring" {
			v.BuilderName = "std"
		}

		if k == "libunwind" {
			v.BuilderName = "llvm_components"
		}

		if k == "libupnp" {
			v.BuilderName = "std"
		}

		if k == "libusb-compat" {
			v.BuilderName = "std"
		}

		if k == "libusb" {
			v.BuilderName = "libusb"
		}

		if k == "libv4l" {
			v.BuilderName = ""
		}

		if k == "libvdpau" {
			v.BuilderName = "std"
		}

		if k == "libvorbis" {
			v.BuilderName = "std"
		}

		if k == "libvpx" {
			v.BuilderName = "libvpx"
		}

		if k == "libvte-java" {
			v.BuilderName = "std"
		}

		if k == "libvtemm" {
			v.BuilderName = "std"
		}

		if k == "libwacom" {
			v.BuilderName = "std"
		}

		if k == "libwebp" {
			v.BuilderName = "std"
		}

		if k == "libwfut" {
			v.BuilderName = "std"
		}

		if k == "libwnck" {
			v.BuilderName = "std"
		}

		if k == "libxcb" {
			v.BuilderName = "std"
		}

		if k == "libxfce4ui" {
			v.BuilderName = "std"
		}

		if k == "libxfce4util" {
			v.BuilderName = "std"
		}

		if k == "libxfcegui4" {
			v.BuilderName = ""
		}

		if k == "libxkbcommon" {
			v.BuilderName = "std"
		}

		if k == "libxkbfile" {
			v.BuilderName = "std"
		}

		if k == "libxkbui" {
			v.BuilderName = "std"
		}

		if k == "libxklavier" {
			v.BuilderName = "std"
		}

		if k == "libxmi" {
			v.BuilderName = "std"
		}

		if k == "libxml++" {
			v.BuilderName = "std"
		}

		if k == "libxml" {
			v.BuilderName = "std"
		}

		if k == "libxml2" {
			v.BuilderName = "libxml2"
		}

		if k == "libxshmfence" {
			v.BuilderName = "std"
		}

		if k == "libxslt" {
			v.BuilderName = "libxslt"
		}

		if k == "libyaml" {
			v.BuilderName = "std"
		}

		if k == "libzapojit" {
			v.BuilderName = "std"
		}

		if k == "libzeitgeist" {
			v.BuilderName = "std"
		}

		if k == "libzip" {
			v.BuilderName = "std"
		}

		if k == "libzrtpcpp" {
			v.BuilderName = "std_cmake"
		}

		if k == "libzvt" {
			v.BuilderName = "std"
		}

		if k == "libzvt2" {
			v.BuilderName = "std"
		}

		if k == "lightning" {
			v.BuilderName = "std"
		}

		if k == "lightsoff" {
			v.BuilderName = "std"
		}

		if k == "lightspark" {
			v.BuilderName = "std_cmake"
		}

		if k == "lighttpd" {
			v.BuilderName = ""
		}

		if k == "linc" {
			v.BuilderName = "std"
		}

		if k == "linphone" {
			v.BuilderName = ""
		}

		if k == "linux-firmware" {
			v.BuilderName = "std"
		}

		if k == "linux-libc-headers" {
			v.BuilderName = "std"
		}

		if k == "linux-pre" {
			v.BuilderName = "std"
		}

		if k == "linux-user-chroot" {
			v.BuilderName = "std"
		}

		if k == "linux" {
			v.BuilderName = "linux"
		}

		if k == "linuxwacom" {
			v.BuilderName = "std"
		}

		if k == "liquidwar6" {
			v.BuilderName = "std"
		}

		if k == "lispdebug" {
			v.BuilderName = "std"
		}

		if k == "listres" {
			v.BuilderName = "std"
		}

		if k == "littlecms" {
			v.BuilderName = ""
		}

		if k == "lld" {
			v.BuilderName = "llvm_components"
		}

		if k == "lldb" {
			v.BuilderName = "llvm_components"
		}

		if k == "llvm" {
			v.BuilderName = "llvm"
		}

		if k == "lmms" {
			v.BuilderName = "std_cmake"
		}

		if k == "lndir" {
			v.BuilderName = "std"
		}

		if k == "lock-service" {
			v.BuilderName = "std"
		}

		if k == "log4c" {
			v.BuilderName = "std"
		}

		if k == "loggerhead" {
			v.BuilderName = ""
		}

		if k == "longomatch" {
			v.BuilderName = "std"
		}

		if k == "longrun" {
			v.BuilderName = "std"
		}

		if k == "loudmouth-ruby" {
			v.BuilderName = "std"
		}

		if k == "loudmouth" {
			v.BuilderName = "std"
		}

		if k == "lsh" {
			v.BuilderName = "std"
		}

		if k == "lsof" {
			v.BuilderName = "lsof"
		}

		if k == "lsr" {
			v.BuilderName = "std"
		}

		if k == "lua" {
			v.BuilderName = "lua"
		}

		if k == "luit" {
			v.BuilderName = "std"
		}

		if k == "lvm2" {
			v.BuilderName = "std"
		}

		if k == "lxml" {
			v.BuilderName = ""
		}

		if k == "lynx" {
			v.BuilderName = "lynx"
		}

		if k == "lyx" {
			v.BuilderName = "std"
		}

		if k == "lzo" {
			v.BuilderName = "std"
		}

		if k == "m4" {
			v.BuilderName = "std"
		}

		if k == "macchanger" {
			v.BuilderName = "std"
		}

		if k == "madplay" {
			v.BuilderName = "std"
		}

		if k == "madwifi" {
			v.BuilderName = ""
		}

		if k == "magicdev" {
			v.BuilderName = "std"
		}

		if k == "mail" {
			v.BuilderName = "std"
		}

		if k == "mailman" {
			v.BuilderName = "std"
		}

		if k == "mailutils" {
			v.BuilderName = "std"
		}

		if k == "make" {
			v.BuilderName = "make"
		}

		if k == "makedepend" {
			v.BuilderName = "std"
		}

		if k == "man-db" {
			v.BuilderName = "std"
		}

		if k == "man-pages-posix" {
			v.BuilderName = "std"
		}

		if k == "man-pages" {
			v.BuilderName = "man_pages"
		}

		if k == "man" {
			v.BuilderName = "man"
		}

		if k == "marble" {
			v.BuilderName = ""
		}

		if k == "marco" {
			v.BuilderName = "std"
		}

		if k == "marst" {
			v.BuilderName = "std"
		}

		if k == "mate-applets" {
			v.BuilderName = "std"
		}

		if k == "mate-backgrounds" {
			v.BuilderName = "std"
		}

		if k == "mate-common" {
			v.BuilderName = "std"
		}

		if k == "mate-control-center" {
			v.BuilderName = "std"
		}

		if k == "mate-desktop" {
			v.BuilderName = "std"
		}

		if k == "mate-icon-theme-faenza" {
			v.BuilderName = "std"
		}

		if k == "mate-icon-theme" {
			v.BuilderName = "std"
		}

		if k == "mate-indicator-applet" {
			v.BuilderName = "std"
		}

		if k == "mate-media" {
			v.BuilderName = "std"
		}

		if k == "mate-menus" {
			v.BuilderName = "std"
		}

		if k == "mate-netbook" {
			v.BuilderName = "std"
		}

		if k == "mate-netspeed" {
			v.BuilderName = "std"
		}

		if k == "mate-notification-daemon" {
			v.BuilderName = "std"
		}

		if k == "mate-panel" {
			v.BuilderName = "std"
		}

		if k == "mate-polkit" {
			v.BuilderName = "std"
		}

		if k == "mate-power-manager" {
			v.BuilderName = "std"
		}

		if k == "mate-screensaver" {
			v.BuilderName = "std"
		}

		if k == "mate-sensors-applet" {
			v.BuilderName = "std"
		}

		if k == "mate-session-manager" {
			v.BuilderName = "std"
		}

		if k == "mate-settings-daemon" {
			v.BuilderName = "std"
		}

		if k == "mate-system-monitor" {
			v.BuilderName = "std"
		}

		if k == "mate-terminal" {
			v.BuilderName = "std"
		}

		if k == "mate-themes-gtk3.10" {
			v.BuilderName = "std"
		}

		if k == "mate-themes-gtk3.12" {
			v.BuilderName = "std"
		}

		if k == "mate-themes-gtk3.14" {
			v.BuilderName = "std"
		}

		if k == "mate-themes-gtk3.16" {
			v.BuilderName = "std"
		}

		if k == "mate-themes-gtk3.18" {
			v.BuilderName = "std"
		}

		if k == "mate-themes-gtk3.8" {
			v.BuilderName = "std"
		}

		if k == "mate-user-guide" {
			v.BuilderName = "std"
		}

		if k == "mate-user-share" {
			v.BuilderName = "std"
		}

		if k == "mate-utils" {
			v.BuilderName = "std"
		}

		if k == "maverik-demos" {
			v.BuilderName = "std"
		}

		if k == "maverik" {
			v.BuilderName = "std"
		}

		if k == "maxima" {
			v.BuilderName = "std"
		}

		if k == "mc" {
			v.BuilderName = "mc"
		}

		if k == "mc4020" {
			v.BuilderName = "std"
		}

		if k == "mcabber" {
			v.BuilderName = ""
		}

		if k == "mcron" {
			v.BuilderName = "std"
		}

		if k == "mcsim" {
			v.BuilderName = "std"
		}

		if k == "mdadm" {
			v.BuilderName = "mdadm"
		}

		if k == "mdctl-v" {
			v.BuilderName = "std"
		}

		if k == "mdctl" {
			v.BuilderName = "std"
		}

		if k == "mdk" {
			v.BuilderName = "std"
		}

		if k == "meanwhile" {
			v.BuilderName = ""
		}

		if k == "media-player-info" {
			v.BuilderName = "std"
		}

		if k == "mediatomb" {
			v.BuilderName = "std"
		}

		if k == "medusa" {
			v.BuilderName = "std"
		}

		if k == "meld" {
			v.BuilderName = "std"
		}

		if k == "memprof" {
			v.BuilderName = "std"
		}

		if k == "memtest86+" {
			v.BuilderName = ""
		}

		if k == "mercator" {
			v.BuilderName = "std"
		}

		if k == "mercurial" {
			v.BuilderName = ""
		}

		if k == "mergeant" {
			v.BuilderName = "std"
		}

		if k == "merkaartor" {
			v.BuilderName = ""
		}

		if k == "mesa-demos" {
			v.BuilderName = "std"
		}

		if k == "mesa" {
			v.BuilderName = "mesalib"
		}

		if k == "mess-desktop-entries" {
			v.BuilderName = "std"
		}

		if k == "metacity" {
			v.BuilderName = "std"
		}

		if k == "metahtml" {
			v.BuilderName = "std"
		}

		if k == "metatheme" {
			v.BuilderName = "std"
		}

		if k == "mgm" {
			v.BuilderName = ""
		}

		if k == "mifluz" {
			v.BuilderName = "std"
		}

		if k == "mig" {
			v.BuilderName = "std"
		}

		if k == "mime-support" {
			v.BuilderName = ""
		}

		if k == "minetest" {
			v.BuilderName = ""
		}

		if k == "ming" {
			v.BuilderName = ""
		}

		if k == "miniupnpc" {
			v.BuilderName = "miniupnpc"
		}

		if k == "miniupnpd" {
			v.BuilderName = "miniupnpd"
		}

		if k == "mino" {
			v.BuilderName = "std"
		}

		if k == "minotaur" {
			v.BuilderName = "std"
		}

		if k == "miscfiles" {
			v.BuilderName = "std"
		}

		if k == "mit-scheme-c-20080130-x" {
			v.BuilderName = "std"
		}

		if k == "mit-scheme-c" {
			v.BuilderName = "std"
		}

		if k == "mit-scheme-doc" {
			v.BuilderName = "std"
		}

		if k == "mit-scheme" {
			v.BuilderName = "std"
		}

		if k == "mjpegtools" {
			v.BuilderName = "std"
		}

		if k == "mkcfm" {
			v.BuilderName = "std"
		}

		if k == "mkcomposecache" {
			v.BuilderName = "std"
		}

		if k == "mkfontdir" {
			v.BuilderName = "std"
		}

		if k == "mkfontscale" {
			v.BuilderName = "std"
		}

		if k == "mldonkey" {
			v.BuilderName = ""
		}

		if k == "mlt" {
			v.BuilderName = "mlt"
		}

		if k == "mlview" {
			v.BuilderName = "std"
		}

		if k == "mm-common" {
			v.BuilderName = "std"
		}

		if k == "mobile-broadband-provider-info" {
			v.BuilderName = "std"
		}

		if k == "mod_wsgi" {
			v.BuilderName = "mod_wsgi"
		}

		if k == "model" {
			v.BuilderName = "std"
		}

		if k == "module-init-tools" {
			v.BuilderName = "moduleinittools"
		}

		if k == "modules" {
			v.BuilderName = "std"
		}

		if k == "modutils" {
			v.BuilderName = "std"
		}

		if k == "moe" {
			v.BuilderName = "std"
		}

		if k == "mon-client" {
			v.BuilderName = "std"
		}

		if k == "mon-syslog" {
			v.BuilderName = "std"
		}

		if k == "mon-talk" {
			v.BuilderName = "std"
		}

		if k == "mon.cgi" {
			v.BuilderName = "std"
		}

		if k == "mon" {
			v.BuilderName = "std"
		}

		if k == "mongodb" {
			v.BuilderName = "mongodb"
		}

		if k == "monitoring-plugins" {
			v.BuilderName = "std"
		}

		if k == "mono" {
			v.BuilderName = "std"
		}

		if k == "moserial" {
			v.BuilderName = "std"
		}

		if k == "motti" {
			v.BuilderName = "std"
		}

		if k == "mousetrap" {
			v.BuilderName = "std"
		}

		if k == "mousetweaks" {
			v.BuilderName = "std"
		}

		if k == "mozembed" {
			v.BuilderName = "mozembed"
		}

		if k == "mozjs17" {
			v.BuilderName = "mozjs"
		}

		if k == "mozjs24" {
			v.BuilderName = "mozjs"
		}

		if k == "mozo" {
			v.BuilderName = "std"
		}

		if k == "mpc" {
			v.BuilderName = "std"
		}

		if k == "mpfr" {
			v.BuilderName = "std"
		}

		if k == "mplayer" {
			v.BuilderName = "mplayer"
		}

		if k == "mrproject" {
			v.BuilderName = "std"
		}

		if k == "msigna" {
			v.BuilderName = "std_qmake"
		}

		if k == "msitools" {
			v.BuilderName = "std"
		}

		if k == "msr-tools" {
			v.BuilderName = "std"
		}

		if k == "mtdev" {
			v.BuilderName = "std"
		}

		if k == "mtools" {
			v.BuilderName = "std"
		}

		if k == "mtr" {
			v.BuilderName = "mtr"
		}

		if k == "muine" {
			v.BuilderName = "std"
		}

		if k == "multivalent" {
			v.BuilderName = ""
		}

		if k == "murrine" {
			v.BuilderName = "std"
		}

		if k == "mutter-wayland" {
			v.BuilderName = "std"
		}

		if k == "mutter" {
			v.BuilderName = "mutter"
		}

		if k == "mypaint" {
			v.BuilderName = ""
		}

		if k == "myserver" {
			v.BuilderName = "std"
		}

		if k == "mysql-connector-c++" {
			v.BuilderName = ""
		}

		if k == "mysql-connector-c" {
			v.BuilderName = ""
		}

		if k == "mysql-connector-java" {
			v.BuilderName = ""
		}

		if k == "mysql-connector-odbc" {
			v.BuilderName = ""
		}

		if k == "mysql-workbench-oss" {
			v.BuilderName = ""
		}

		if k == "mysql" {
			v.BuilderName = "mysql"
		}

		if k == "naev" {
			v.BuilderName = ""
		}

		if k == "named" {
			v.BuilderName = ""
		}

		if k == "nanny" {
			v.BuilderName = "std"
		}

		if k == "nano" {
			v.BuilderName = "std"
		}

		if k == "nas" {
			v.BuilderName = "nas"
		}

		if k == "nasm" {
			v.BuilderName = "nasm"
		}

		if k == "nautilus-actions" {
			v.BuilderName = "std"
		}

		if k == "nautilus-cd-burner" {
			v.BuilderName = "std"
		}

		if k == "nautilus-gtkhtml" {
			v.BuilderName = "std"
		}

		if k == "nautilus-image-converter" {
			v.BuilderName = "std"
		}

		if k == "nautilus-media" {
			v.BuilderName = "std"
		}

		if k == "nautilus-mozilla" {
			v.BuilderName = "std"
		}

		if k == "nautilus-open-terminal" {
			v.BuilderName = "std"
		}

		if k == "nautilus-python" {
			v.BuilderName = "std"
		}

		if k == "nautilus-rpm" {
			v.BuilderName = "std"
		}

		if k == "nautilus-sendto" {
			v.BuilderName = "std"
		}

		if k == "nautilus-share" {
			v.BuilderName = "std"
		}

		if k == "nautilus" {
			v.BuilderName = "std"
		}

		if k == "nbxmpp" {
			v.BuilderName = "std_pythons"
		}

		if k == "nclx" {
			v.BuilderName = "std"
		}

		if k == "ncurses" {
			v.BuilderName = "ncurses"
		}

		if k == "neard" {
			v.BuilderName = "std"
		}

		if k == "nemiver" {
			v.BuilderName = "std"
		}

		if k == "neon" {
			v.BuilderName = "neon"
		}

		if k == "nepomuk-music-kio-slave" {
			v.BuilderName = ""
		}

		if k == "nepomukannotation" {
			v.BuilderName = ""
		}

		if k == "nepomukextras" {
			v.BuilderName = ""
		}

		if k == "nepomukshell" {
			v.BuilderName = ""
		}

		if k == "nepomuktvnamer" {
			v.BuilderName = ""
		}

		if k == "nesemu" {
			v.BuilderName = ""
		}

		if k == "net-snmp" {
			v.BuilderName = "net_snmp"
		}

		if k == "net-tools" {
			v.BuilderName = ""
		}

		if k == "netbeans" {
			v.BuilderName = ""
		}

		if k == "netcat" {
			v.BuilderName = ""
		}

		if k == "nethack" {
			v.BuilderName = "nethack"
		}

		if k == "netstat-nat" {
			v.BuilderName = ""
		}

		if k == "nettle" {
			v.BuilderName = "std"
		}

		if k == "nettle2" {
			v.BuilderName = "std"
		}

		if k == "network-manager-applet" {
			v.BuilderName = "std"
		}

		if k == "network-manager-netbook" {
			v.BuilderName = "std"
		}

		if k == "nfs-utils" {
			v.BuilderName = "std"
		}

		if k == "nftables" {
			v.BuilderName = "std"
		}

		if k == "nginx" {
			v.BuilderName = ""
		}

		if k == "nis-utils" {
			v.BuilderName = "std"
		}

		if k == "nistool" {
			v.BuilderName = "std"
		}

		if k == "nmap" {
			v.BuilderName = "std"
		}

		if k == "notification-daemon" {
			v.BuilderName = "std"
		}

		if k == "notify-sharp" {
			v.BuilderName = "std"
		}

		if k == "notion" {
			v.BuilderName = ""
		}

		if k == "npth" {
			v.BuilderName = "std"
		}

		if k == "nspr" {
			v.BuilderName = "nspr"
		}

		if k == "nss" {
			v.BuilderName = "nss"
		}

		if k == "nss_db" {
			v.BuilderName = "std"
		}

		if k == "nss_lwres" {
			v.BuilderName = "std"
		}

		if k == "ntfs-3g_ntfsprogs" {
			v.BuilderName = "ntfs3gntfsprogs"
		}

		if k == "ntp" {
			v.BuilderName = "ntp"
		}

		if k == "nut" {
			v.BuilderName = "std"
		}

		if k == "oaf" {
			v.BuilderName = "std"
		}

		if k == "obconf" {
			v.BuilderName = "std"
		}

		if k == "obexd" {
			v.BuilderName = "std"
		}

		if k == "obexfs" {
			v.BuilderName = ""
		}

		if k == "obexftp" {
			v.BuilderName = ""
		}

		if k == "ocaml" {
			v.BuilderName = ""
		}

		if k == "oclock" {
			v.BuilderName = "std"
		}

		if k == "ocrad" {
			v.BuilderName = "std"
		}

		if k == "ocrfeeder" {
			v.BuilderName = "std"
		}

		if k == "octave" {
			v.BuilderName = "std"
		}

		if k == "odb-examples" {
			v.BuilderName = "std"
		}

		if k == "odb-tests" {
			v.BuilderName = "std"
		}

		if k == "odb" {
			v.BuilderName = "std"
		}

		if k == "ode-python" {
			v.BuilderName = "std"
		}

		if k == "ode" {
			v.BuilderName = "std"
		}

		if k == "office-runner" {
			v.BuilderName = "std"
		}

		if k == "ofono" {
			v.BuilderName = "std"
		}

		if k == "ogre" {
			v.BuilderName = "std_cmake"
		}

		if k == "ogre1.7" {
			v.BuilderName = "std_cmake"
		}

		if k == "ois" {
			v.BuilderName = "std"
		}

		if k == "okular" {
			v.BuilderName = ""
		}

		if k == "oleo" {
			v.BuilderName = "std"
		}

		if k == "online-desktop" {
			v.BuilderName = "std"
		}

		if k == "ontv" {
			v.BuilderName = "std"
		}

		if k == "oolite" {
			v.BuilderName = ""
		}

		if k == "opal" {
			v.BuilderName = "std"
		}

		if k == "open-mpi" {
			v.BuilderName = ""
		}

		if k == "openal-soft" {
			v.BuilderName = "std_cmake"
		}

		if k == "openbox" {
			v.BuilderName = "std"
		}

		if k == "opencascade" {
			v.BuilderName = ""
		}

		if k == "opencde" {
			v.BuilderName = ""
		}

		if k == "opencdk" {
			v.BuilderName = ""
		}

		if k == "opencv2" {
			v.BuilderName = "std_cmake"
		}

		if k == "opencv3" {
			v.BuilderName = "std_cmake"
		}

		if k == "opendias" {
			v.BuilderName = ""
		}

		if k == "openerp" {
			v.BuilderName = ""
		}

		if k == "openexr" {
			v.BuilderName = "std"
		}

		if k == "opengfx" {
			v.BuilderName = ""
		}

		if k == "opengl-glib" {
			v.BuilderName = "std"
		}

		if k == "openimageio" {
			v.BuilderName = "openimageio"
		}

		if k == "openjade" {
			v.BuilderName = ""
		}

		if k == "openjdk8" {
			v.BuilderName = "openjdk8"
		}

		if k == "openjpeg1" {
			v.BuilderName = "openjpeg"
		}

		if k == "openjpeg2" {
			v.BuilderName = "openjpeg"
		}

		if k == "openldap" {
			v.BuilderName = "openldap"
		}

		if k == "openmotif" {
			v.BuilderName = ""
		}

		if k == "openmovieeditor" {
			v.BuilderName = ""
		}

		if k == "openmp" {
			v.BuilderName = "llvm_components"
		}

		if k == "openmpi" {
			v.BuilderName = "std"
		}

		if k == "openmsx" {
			v.BuilderName = ""
		}

		if k == "openobex" {
			v.BuilderName = "std_cmake"
		}

		if k == "opensfx" {
			v.BuilderName = ""
		}

		if k == "openshot-qt" {
			v.BuilderName = ""
		}

		if k == "openshot" {
			v.BuilderName = ""
		}

		if k == "openssh" {
			v.BuilderName = "openssh"
		}

		if k == "openssl" {
			v.BuilderName = "openssl"
		}

		if k == "openttd" {
			v.BuilderName = "openttd"
		}

		if k == "openvpn" {
			v.BuilderName = "openvpn"
		}

		if k == "opus-tools" {
			v.BuilderName = "std"
		}

		if k == "opus" {
			v.BuilderName = "opus"
		}

		if k == "opusfile" {
			v.BuilderName = "std"
		}

		if k == "orage" {
			v.BuilderName = ""
		}

		if k == "orbit-python" {
			v.BuilderName = "std"
		}

		if k == "orbitcpp" {
			v.BuilderName = "std"
		}

		if k == "orc" {
			v.BuilderName = "std"
		}

		if k == "orca" {
			v.BuilderName = "std"
		}

		if k == "orgadoc" {
			v.BuilderName = "std"
		}

		if k == "osip" {
			v.BuilderName = ""
		}

		if k == "ostree-embeddeps" {
			v.BuilderName = "std"
		}

		if k == "ostree" {
			v.BuilderName = "std"
		}

		if k == "oxygen-gtk" {
			v.BuilderName = ""
		}

		if k == "oxygen-gtk2" {
			v.BuilderName = ""
		}

		if k == "oxygen-gtk3" {
			v.BuilderName = ""
		}

		if k == "oxygen-icons" {
			v.BuilderName = ""
		}

		if k == "p11-kit" {
			v.BuilderName = "p11_kit"
		}

		if k == "p7zip" {
			v.BuilderName = "p7zip"
		}

		if k == "pacrunner" {
			v.BuilderName = "std"
		}

		if k == "pam_biomouseplus" {
			v.BuilderName = "std"
		}

		if k == "pam_cap" {
			v.BuilderName = "std"
		}

		if k == "pam_chroot" {
			v.BuilderName = "std"
		}

		if k == "pam_login_alert" {
			v.BuilderName = "std"
		}

		if k == "pam_nw_auth" {
			v.BuilderName = "std"
		}

		if k == "pam_opie" {
			v.BuilderName = "std"
		}

		if k == "pam_pgsql_day" {
			v.BuilderName = "std"
		}

		if k == "pam_skey" {
			v.BuilderName = "std"
		}

		if k == "pan" {
			v.BuilderName = "std"
		}

		if k == "panda3d" {
			v.BuilderName = "panda3d"
		}

		if k == "panelmm" {
			v.BuilderName = "std"
		}

		if k == "pango" {
			v.BuilderName = "pango"
		}

		if k == "pangomm" {
			v.BuilderName = "std"
		}

		if k == "pangox-compat" {
			v.BuilderName = "std"
		}

		if k == "paperbox" {
			v.BuilderName = "std"
		}

		if k == "parallel" {
			v.BuilderName = "std"
		}

		if k == "parley" {
			v.BuilderName = ""
		}

		if k == "parted" {
			v.BuilderName = "std"
		}

		if k == "passepartout" {
			v.BuilderName = "std"
		}

		if k == "patch" {
			v.BuilderName = "std"
		}

		if k == "patchelf" {
			v.BuilderName = "std"
		}

		if k == "pavucontrol" {
			v.BuilderName = "std"
		}

		if k == "pcb" {
			v.BuilderName = ""
		}

		if k == "pciids" {
			v.BuilderName = ""
		}

		if k == "pciutils" {
			v.BuilderName = "pciutils"
		}

		if k == "pcl+clx.sept" {
			v.BuilderName = "std"
		}

		if k == "pcl-gcl" {
			v.BuilderName = "std"
		}

		if k == "pcl.sept" {
			v.BuilderName = "std"
		}

		if k == "pcmciautils" {
			v.BuilderName = "std"
		}

		if k == "pcre" {
			v.BuilderName = "pcre"
		}

		if k == "pcre2" {
			v.BuilderName = "pcre"
		}

		if k == "pdfedit" {
			v.BuilderName = ""
		}

		if k == "pdfmod" {
			v.BuilderName = "std"
		}

		if k == "pdfshuffler" {
			v.BuilderName = ""
		}

		if k == "pdftk" {
			v.BuilderName = "pdftk"
		}

		if k == "pem" {
			v.BuilderName = "std"
		}

		if k == "perl" {
			v.BuilderName = "perl"
		}

		if k == "perlkde" {
			v.BuilderName = ""
		}

		if k == "perlqt" {
			v.BuilderName = ""
		}

		if k == "pessulus" {
			v.BuilderName = "std"
		}

		if k == "pexec" {
			v.BuilderName = "std"
		}

		if k == "phantom_home-beta" {
			v.BuilderName = "std"
		}

		if k == "phantom_home" {
			v.BuilderName = "std"
		}

		if k == "phantom_security-beta" {
			v.BuilderName = "std"
		}

		if k == "phantom_security" {
			v.BuilderName = "std"
		}

		if k == "phodav" {
			v.BuilderName = "std"
		}

		if k == "phonesim" {
			v.BuilderName = "std"
		}

		if k == "phonon" {
			v.BuilderName = ""
		}

		if k == "php" {
			v.BuilderName = "php"
		}

		if k == "phpbb" {
			v.BuilderName = ""
		}

		if k == "physfs" {
			v.BuilderName = "std_cmake"
		}

		if k == "pidgin" {
			v.BuilderName = "pidgin"
		}

		if k == "pies" {
			v.BuilderName = "std"
		}

		if k == "pil" {
			v.BuilderName = "std"
		}

		if k == "pinpoint" {
			v.BuilderName = "std"
		}

		if k == "pioneer" {
			v.BuilderName = "std"
		}

		if k == "pipelight" {
			v.BuilderName = "pipelight"
		}

		if k == "pitivi" {
			v.BuilderName = "std"
		}

		if k == "pixman" {
			v.BuilderName = "std"
		}

		if k == "pjproject" {
			v.BuilderName = "std"
		}

		if k == "pkg-config" {
			v.BuilderName = "std"
		}

		if k == "pkgconfig" {
			v.BuilderName = "std"
		}

		if k == "planb" {
			v.BuilderName = "std"
		}

		if k == "planeshift" {
			v.BuilderName = ""
		}

		if k == "planner" {
			v.BuilderName = "std"
		}

		if k == "plone" {
			v.BuilderName = ""
		}

		if k == "plotutils" {
			v.BuilderName = "std"
		}

		if k == "pluma" {
			v.BuilderName = "std"
		}

		if k == "plymouth" {
			v.BuilderName = "std"
		}

		if k == "pngtoico" {
			v.BuilderName = "std"
		}

		if k == "poco" {
			v.BuilderName = "std_cmake"
		}

		if k == "podofo" {
			v.BuilderName = "std_cmake"
		}

		if k == "podofobrowser" {
			v.BuilderName = ""
		}

		if k == "polari" {
			v.BuilderName = "std"
		}

		if k == "policycoreutils" {
			v.BuilderName = "std"
		}

		if k == "polkit-gnome" {
			v.BuilderName = "std"
		}

		if k == "polkit-kde-agent-1" {
			v.BuilderName = ""
		}

		if k == "polkit-qt-1" {
			v.BuilderName = ""
		}

		if k == "polkit" {
			v.BuilderName = "polkit"
		}

		if k == "polly" {
			v.BuilderName = "llvm_components"
		}

		if k == "polypaudio" {
			v.BuilderName = "std"
		}

		if k == "pong" {
			v.BuilderName = "std"
		}

		if k == "poppler-data" {
			v.BuilderName = "std_cmake"
		}

		if k == "poppler" {
			v.BuilderName = "poppler"
		}

		if k == "popt" {
			v.BuilderName = "popt"
		}

		if k == "portaudio" {
			v.BuilderName = "std"
		}

		if k == "postgresql" {
			v.BuilderName = "postgreSQL"
		}

		if k == "postr" {
			v.BuilderName = "std"
		}

		if k == "ppp" {
			v.BuilderName = "ppp"
		}

		if k == "pptp" {
			v.BuilderName = "pptp"
		}

		if k == "pptpd" {
			v.BuilderName = "pptpd"
		}

		if k == "prefixsuffix" {
			v.BuilderName = "std"
		}

		if k == "present" {
			v.BuilderName = "std"
		}

		if k == "presentproto" {
			v.BuilderName = "std"
		}

		if k == "printer-applet" {
			v.BuilderName = ""
		}

		if k == "printman" {
			v.BuilderName = "std"
		}

		if k == "printproto" {
			v.BuilderName = "std"
		}

		if k == "procman" {
			v.BuilderName = "std"
		}

		if k == "procps-ng" {
			v.BuilderName = "procps_ng"
		}

		if k == "proftpd" {
			v.BuilderName = ""
		}

		if k == "protobuf" {
			v.BuilderName = "std"
		}

		if k == "proxyknife" {
			v.BuilderName = "std"
		}

		if k == "proxymngr" {
			v.BuilderName = "std"
		}

		if k == "psi" {
			v.BuilderName = "psi"
		}

		if k == "psimedia" {
			v.BuilderName = "psimedia"
		}

		if k == "psmisc" {
			v.BuilderName = "psmisc"
		}

		if k == "pspp" {
			v.BuilderName = "std"
		}

		if k == "psychosynth" {
			v.BuilderName = "std"
		}

		if k == "pth" {
			v.BuilderName = "std"
		}

		if k == "ptlib" {
			v.BuilderName = "std"
		}

		if k == "pulseaudio" {
			v.BuilderName = "pulseaudio"
		}

		if k == "pwdutils" {
			v.BuilderName = "std"
		}

		if k == "pwlib" {
			v.BuilderName = "std"
		}

		if k == "py2cairo" {
			v.BuilderName = "py2cairo"
		}

		if k == "pyOpenSSL" {
			v.BuilderName = ""
		}

		if k == "pyalsa" {
			v.BuilderName = "std"
		}

		if k == "pyatspi" {
			v.BuilderName = "std"
		}

		if k == "pybliographer" {
			v.BuilderName = "std"
		}

		if k == "pycairo" {
			v.BuilderName = "pycairo"
		}

		if k == "pycdio" {
			v.BuilderName = "std"
		}

		if k == "pyconfigure" {
			v.BuilderName = "std"
		}

		if k == "pygame" {
			v.BuilderName = ""
		}

		if k == "pygda" {
			v.BuilderName = "std"
		}

		if k == "pygi" {
			v.BuilderName = "std"
		}

		if k == "pygobject" {
			v.BuilderName = "pygobject"
		}

		if k == "pygobject2" {
			v.BuilderName = "pygobject2"
		}

		if k == "pygobject3.4" {
			v.BuilderName = "pygobject"
		}

		if k == "pygobject3.8" {
			v.BuilderName = "pygobject"
		}

		if k == "pygoocanvas" {
			v.BuilderName = "std"
		}

		if k == "pygtk" {
			v.BuilderName = "std"
		}

		if k == "pygtk2reference" {
			v.BuilderName = "std"
		}

		if k == "pygtkglext" {
			v.BuilderName = "std"
		}

		if k == "pygtksourceview" {
			v.BuilderName = "std"
		}

		if k == "pyjamas" {
			v.BuilderName = ""
		}

		if k == "pykde4" {
			v.BuilderName = ""
		}

		if k == "pymsn" {
			v.BuilderName = "std"
		}

		if k == "pyorbit" {
			v.BuilderName = "std"
		}

		if k == "pyphany" {
			v.BuilderName = "std"
		}

		if k == "pypoppler" {
			v.BuilderName = ""
		}

		if k == "pyqt4" {
			v.BuilderName = "pyqt"
		}

		if k == "pyqt5" {
			v.BuilderName = "pyqt"
		}

		if k == "pysvn" {
			v.BuilderName = ""
		}

		if k == "python-caja" {
			v.BuilderName = "std"
		}

		if k == "python-inet_diag" {
			v.BuilderName = "std"
		}

		if k == "python-linux-procfs" {
			v.BuilderName = "std"
		}

		if k == "python-schedutils" {
			v.BuilderName = "std"
		}

		if k == "python-sipsimple" {
			v.BuilderName = ""
		}

		if k == "pyxdg" {
			v.BuilderName = "std_pythons"
		}

		if k == "qbittorrent" {
			v.BuilderName = "std_cmake"
		}

		if k == "qbzr" {
			v.BuilderName = ""
		}

		if k == "qca" {
			v.BuilderName = "qca"
		}

		if k == "qemu" {
			v.BuilderName = "qemu"
		}

		if k == "qimageblitz" {
			v.BuilderName = ""
		}

		if k == "qpdf" {
			v.BuilderName = "std"
		}

		if k == "qrencode" {
			v.BuilderName = "std"
		}

		if k == "qt-gstreamer" {
			v.BuilderName = "std"
		}

		if k == "qt4" {
			v.BuilderName = "qt"
		}

		if k == "qt5" {
			v.BuilderName = "qt"
		}

		if k == "qtcreator" {
			v.BuilderName = "qtcreator"
		}

		if k == "qtruby" {
			v.BuilderName = ""
		}

		if k == "quadrapassel" {
			v.BuilderName = "std"
		}

		if k == "quesoglc" {
			v.BuilderName = ""
		}

		if k == "quick-lounge-applet" {
			v.BuilderName = "std"
		}

		if k == "qyoto" {
			v.BuilderName = ""
		}

		if k == "radare" {
			v.BuilderName = "std"
		}

		if k == "radare2" {
			v.BuilderName = "std"
		}

		if k == "radioactive" {
			v.BuilderName = "std"
		}

		if k == "radius" {
			v.BuilderName = "std"
		}

		if k == "rails" {
			v.BuilderName = ""
		}

		if k == "randrproto" {
			v.BuilderName = "std"
		}

		if k == "ranpwd" {
			v.BuilderName = "std"
		}

		if k == "rapicorn" {
			v.BuilderName = ""
		}

		if k == "rarian" {
			v.BuilderName = "std"
		}

		if k == "raw-thumbnailer" {
			v.BuilderName = "std"
		}

		if k == "rcairo" {
			v.BuilderName = "std"
		}

		if k == "rcs" {
			v.BuilderName = "std"
		}

		if k == "rdesktop" {
			v.BuilderName = "rdesktop"
		}

		if k == "readline-doc" {
			v.BuilderName = "std"
		}

		if k == "readline" {
			v.BuilderName = "std"
		}

		if k == "recording-level-monitor" {
			v.BuilderName = ""
		}

		if k == "recordproto" {
			v.BuilderName = "std"
		}

		if k == "recutils" {
			v.BuilderName = "std"
		}

		if k == "redmine" {
			v.BuilderName = ""
		}

		if k == "reftex" {
			v.BuilderName = "std"
		}

		if k == "regexxer" {
			v.BuilderName = "std"
		}

		if k == "remotecontrol" {
			v.BuilderName = "std"
		}

		if k == "rendercheck" {
			v.BuilderName = "std"
		}

		if k == "renderproto" {
			v.BuilderName = "std"
		}

		if k == "rep-gtk-gnome2" {
			v.BuilderName = "std"
		}

		if k == "rep-gtk" {
			v.BuilderName = "std"
		}

		if k == "resourceproto" {
			v.BuilderName = "std"
		}

		if k == "rest" {
			v.BuilderName = "rest"
		}

		if k == "rfkill" {
			v.BuilderName = "std"
		}

		if k == "rgb" {
			v.BuilderName = "std"
		}

		if k == "rhino" {
			v.BuilderName = ""
		}

		if k == "rhythmbox" {
			v.BuilderName = "std"
		}

		if k == "ristretto" {
			v.BuilderName = ""
		}

		if k == "rng-tools" {
			v.BuilderName = ""
		}

		if k == "rocs" {
			v.BuilderName = ""
		}

		if k == "rogg" {
			v.BuilderName = "std"
		}

		if k == "rottlog" {
			v.BuilderName = "std"
		}

		if k == "rp-l2tp" {
			v.BuilderName = "rpl2tp"
		}

		if k == "rp-pppoe" {
			v.BuilderName = "rp_pppoe"
		}

		if k == "rpcnis-headers" {
			v.BuilderName = "rpcnis_headers"
		}

		if k == "rpge" {
			v.BuilderName = "std"
		}

		if k == "rrdtool" {
			v.BuilderName = "std"
		}

		if k == "rstart" {
			v.BuilderName = "std"
		}

		if k == "rsync" {
			v.BuilderName = "std"
		}

		if k == "rtmpdump" {
			v.BuilderName = "rtmpdump"
		}

		if k == "rubberband" {
			v.BuilderName = ""
		}

		if k == "ruby" {
			v.BuilderName = "ruby"
		}

		if k == "rubygems" {
			v.BuilderName = ""
		}

		if k == "rush" {
			v.BuilderName = "std"
		}

		if k == "rustc" {
			v.BuilderName = "rustc"
		}

		if k == "rustfmt" {
			v.BuilderName = "std"
		}

		if k == "rxload" {
			v.BuilderName = "std"
		}

		if k == "rxvt" {
			v.BuilderName = "rxvt"
		}

		if k == "rygel-gst-0-10-fullscreen-renderer" {
			v.BuilderName = "std"
		}

		if k == "rygel-gst-0-10-media-engine" {
			v.BuilderName = "std"
		}

		if k == "rygel-gst-0-10-plugins" {
			v.BuilderName = "std"
		}

		if k == "rygel" {
			v.BuilderName = "std"
		}

		if k == "sabayon" {
			v.BuilderName = "std"
		}

		if k == "samba" {
			v.BuilderName = "samba"
		}

		if k == "sane-backends" {
			v.BuilderName = "sane_backends"
		}

		if k == "sapwood" {
			v.BuilderName = "std"
		}

		if k == "sather-contrib" {
			v.BuilderName = "std"
		}

		if k == "sather-extra" {
			v.BuilderName = "std"
		}

		if k == "sather-specification" {
			v.BuilderName = "std"
		}

		if k == "sather-tutorial" {
			v.BuilderName = "std"
		}

		if k == "sather" {
			v.BuilderName = "std"
		}

		if k == "sauce" {
			v.BuilderName = "std"
		}

		if k == "sawfish-gnome2" {
			v.BuilderName = "std"
		}

		if k == "sawfish" {
			v.BuilderName = "std"
		}

		if k == "sbc" {
			v.BuilderName = "std"
		}

		if k == "scaffold" {
			v.BuilderName = "std"
		}

		if k == "schedule-oncall" {
			v.BuilderName = "std"
		}

		if k == "schedule-tools" {
			v.BuilderName = "std"
		}

		if k == "schroedinger" {
			v.BuilderName = ""
		}

		if k == "scm" {
			v.BuilderName = "std"
		}

		if k == "scons" {
			v.BuilderName = "scons"
		}

		if k == "screen" {
			v.BuilderName = "std"
		}

		if k == "scribus" {
			v.BuilderName = ""
		}

		if k == "scripts" {
			v.BuilderName = "std"
		}

		if k == "scrnsaverproto" {
			v.BuilderName = "std"
		}

		if k == "scrollkeeper" {
			v.BuilderName = "std"
		}

		if k == "scute" {
			v.BuilderName = "std"
		}

		if k == "seahorse-nautilus" {
			v.BuilderName = "std"
		}

		if k == "seahorse-plugins" {
			v.BuilderName = "std"
		}

		if k == "seahorse-sharing" {
			v.BuilderName = "std"
		}

		if k == "seahorse" {
			v.BuilderName = "std"
		}

		if k == "seamonkey" {
			v.BuilderName = "seamonkey"
		}

		if k == "sed" {
			v.BuilderName = "std"
		}

		if k == "seed" {
			v.BuilderName = "seed"
		}

		if k == "sendmail" {
			v.BuilderName = ""
		}

		if k == "sepolgen" {
			v.BuilderName = "std"
		}

		if k == "serf" {
			v.BuilderName = "std"
		}

		if k == "serveez" {
			v.BuilderName = "std"
		}

		if k == "servletapi" {
			v.BuilderName = "std"
		}

		if k == "sessreg" {
			v.BuilderName = "std"
		}

		if k == "sethdlc" {
			v.BuilderName = "std"
		}

		if k == "setuptools" {
			v.BuilderName = ""
		}

		if k == "setxkbmap" {
			v.BuilderName = "std"
		}

		if k == "sflphone" {
			v.BuilderName = ""
		}

		if k == "sgml-common" {
			v.BuilderName = "sgml_common"
		}

		if k == "shadow" {
			v.BuilderName = "shadow"
		}

		if k == "shake" {
			v.BuilderName = ""
		}

		if k == "shared-desktop-ontologies" {
			v.BuilderName = "std_cmake"
		}

		if k == "shared-mime-info" {
			v.BuilderName = "std"
		}

		if k == "sharutils" {
			v.BuilderName = "std"
		}

		if k == "shishi" {
			v.BuilderName = "std"
		}

		if k == "shmm" {
			v.BuilderName = "std"
		}

		if k == "shotwell" {
			v.BuilderName = "std"
		}

		if k == "showfont" {
			v.BuilderName = "std"
		}

		if k == "shtool" {
			v.BuilderName = "std"
		}

		if k == "shutter" {
			v.BuilderName = ""
		}

		if k == "signon-oauth2" {
			v.BuilderName = "std_qmake"
		}

		if k == "signon" {
			v.BuilderName = "std_qmake"
		}

		if k == "simple-scan" {
			v.BuilderName = "std"
		}

		if k == "simutrans" {
			v.BuilderName = ""
		}

		if k == "siobhan" {
			v.BuilderName = "std"
		}

		if k == "sip-comunicator" {
			v.BuilderName = ""
		}

		if k == "sip" {
			v.BuilderName = "sip"
		}

		if k == "sipwitch" {
			v.BuilderName = "std"
		}

		if k == "skstream" {
			v.BuilderName = "std"
		}

		if k == "slang" {
			v.BuilderName = "slang"
		}

		if k == "slib" {
			v.BuilderName = "std"
		}

		if k == "slrn" {
			v.BuilderName = ""
		}

		if k == "smack" {
			v.BuilderName = ""
		}

		if k == "smail" {
			v.BuilderName = "std"
		}

		if k == "smalltalk" {
			v.BuilderName = "std"
		}

		if k == "smartmontools" {
			v.BuilderName = "std"
		}

		if k == "smokegen" {
			v.BuilderName = ""
		}

		if k == "smokekde" {
			v.BuilderName = ""
		}

		if k == "smokeqt" {
			v.BuilderName = ""
		}

		if k == "smplayer" {
			v.BuilderName = ""
		}

		if k == "smproxy" {
			v.BuilderName = "std"
		}

		if k == "snappy" {
			v.BuilderName = "std"
		}

		if k == "snmpvar.monitor" {
			v.BuilderName = "std"
		}

		if k == "snowy" {
			v.BuilderName = "std"
		}

		if k == "socat" {
			v.BuilderName = "std"
		}

		if k == "sodipodi" {
			v.BuilderName = "std"
		}

		if k == "solang" {
			v.BuilderName = "std"
		}

		if k == "solfege-easybuild" {
			v.BuilderName = "std"
		}

		if k == "solfege" {
			v.BuilderName = "std"
		}

		if k == "sonic-visualiser" {
			v.BuilderName = ""
		}

		if k == "soprano" {
			v.BuilderName = "std_cmake"
		}

		if k == "sound-juicer" {
			v.BuilderName = "std"
		}

		if k == "soup" {
			v.BuilderName = "std"
		}

		if k == "source-highlight" {
			v.BuilderName = "std"
		}

		if k == "sovix" {
			v.BuilderName = "std"
		}

		if k == "sox" {
			v.BuilderName = "sox"
		}

		if k == "spacechart" {
			v.BuilderName = "std"
		}

		if k == "spacenavd" {
			v.BuilderName = "libspnav"
		}

		if k == "spamcan" {
			v.BuilderName = "std"
		}

		if k == "sparse" {
			v.BuilderName = "std"
		}

		if k == "speedx" {
			v.BuilderName = "std"
		}

		if k == "speex" {
			v.BuilderName = "std"
		}

		if k == "speexdsp" {
			v.BuilderName = "std"
		}

		if k == "spell" {
			v.BuilderName = "std"
		}

		if k == "spring" {
			v.BuilderName = ""
		}

		if k == "sqlcipher" {
			v.BuilderName = "sqlcipher"
		}

		if k == "sqlite-amalgamation" {
			v.BuilderName = "std"
		}

		if k == "sqlite-autoconf" {
			v.BuilderName = "sqliteautoconf"
		}

		if k == "sqlitebrowser" {
			v.BuilderName = ""
		}

		if k == "sqlitecpp" {
			v.BuilderName = "std_cmake"
		}

		if k == "sqltutor" {
			v.BuilderName = "std"
		}

		if k == "squashfs" {
			v.BuilderName = "squashfs"
		}

		if k == "squid" {
			v.BuilderName = ""
		}

		if k == "ssh-contact" {
			v.BuilderName = "std"
		}

		if k == "sshfs-fuse" {
			v.BuilderName = "std"
		}

		if k == "startup-notification" {
			v.BuilderName = "std"
		}

		if k == "statd" {
			v.BuilderName = "std"
		}

		if k == "step" {
			v.BuilderName = ""
		}

		if k == "storage.monitor" {
			v.BuilderName = "std"
		}

		if k == "stow" {
			v.BuilderName = "std"
		}

		if k == "strace" {
			v.BuilderName = "std"
		}

		if k == "straw" {
			v.BuilderName = "std"
		}

		if k == "strigi" {
			v.BuilderName = ""
		}

		if k == "strongwind" {
			v.BuilderName = "std"
		}

		if k == "subtitleeditor" {
			v.BuilderName = "std"
		}

		if k == "subversion" {
			v.BuilderName = "subversion"
		}

		if k == "sudo" {
			v.BuilderName = "sudo"
		}

		if k == "superkaramba" {
			v.BuilderName = ""
		}

		if k == "superopt" {
			v.BuilderName = "std"
		}

		if k == "sushi" {
			v.BuilderName = "std"
		}

		if k == "svgpart" {
			v.BuilderName = ""
		}

		if k == "swbis" {
			v.BuilderName = "std"
		}

		if k == "sweeper" {
			v.BuilderName = ""
		}

		if k == "swell-foop" {
			v.BuilderName = "std"
		}

		if k == "swfdec-gnome" {
			v.BuilderName = "std"
		}

		if k == "swig" {
			v.BuilderName = "std"
		}

		if k == "synfig" {
			v.BuilderName = ""
		}

		if k == "synfigstudio" {
			v.BuilderName = ""
		}

		if k == "sysfsutils" {
			v.BuilderName = "std"
		}

		if k == "syslinux" {
			v.BuilderName = "syslinux"
		}

		if k == "system-tools-backends" {
			v.BuilderName = "std"
		}

		if k == "system-tray-applet" {
			v.BuilderName = "std"
		}

		if k == "systemd-ui" {
			v.BuilderName = "std"
		}

		if k == "systemd" {
			v.BuilderName = "systemd"
		}

		if k == "sysutils" {
			v.BuilderName = "std"
		}

		if k == "sysvinit" {
			v.BuilderName = ""
		}

		if k == "t-engine4" {
			v.BuilderName = ""
		}

		if k == "tack" {
			v.BuilderName = "std"
		}

		if k == "tali" {
			v.BuilderName = "std"
		}

		if k == "talloc" {
			v.BuilderName = "std_waf"
		}

		if k == "tango-icon-theme-extras" {
			v.BuilderName = "std"
		}

		if k == "tango-icon-theme" {
			v.BuilderName = "std"
		}

		if k == "tango" {
			v.BuilderName = ""
		}

		if k == "tar" {
			v.BuilderName = "std"
		}

		if k == "tasks" {
			v.BuilderName = "std"
		}

		if k == "tasque" {
			v.BuilderName = "std"
		}

		if k == "tcl" {
			v.BuilderName = "tcltk"
		}

		if k == "tcllib" {
			v.BuilderName = "std"
		}

		if k == "tcp_wrappers" {
			v.BuilderName = "tcp_wrappers"
		}

		if k == "tcpdump" {
			v.BuilderName = "std"
		}

		if k == "tct" {
			v.BuilderName = ""
		}

		if k == "tdb" {
			v.BuilderName = "std_waf"
		}

		if k == "tdb1.2" {
			v.BuilderName = "std_waf"
		}

		if k == "telegnome" {
			v.BuilderName = "std"
		}

		if k == "telepathy-butterfly" {
			v.BuilderName = "std"
		}

		if k == "telepathy-fargo" {
			v.BuilderName = "std"
		}

		if k == "telepathy-farsight" {
			v.BuilderName = "std"
		}

		if k == "telepathy-farstream" {
			v.BuilderName = "std"
		}

		if k == "telepathy-feed" {
			v.BuilderName = "std"
		}

		if k == "telepathy-gabble" {
			v.BuilderName = "telepathy_gabble"
		}

		if k == "telepathy-glib" {
			v.BuilderName = "telepathy_glib"
		}

		if k == "telepathy-haze" {
			v.BuilderName = "std"
		}

		if k == "telepathy-idle" {
			v.BuilderName = "std"
		}

		if k == "telepathy-inspector" {
			v.BuilderName = "std"
		}

		if k == "telepathy-logger" {
			v.BuilderName = "std"
		}

		if k == "telepathy-mission-control" {
			v.BuilderName = "telepathy_mission_control"
		}

		if k == "telepathy-phoenix" {
			v.BuilderName = "std"
		}

		if k == "telepathy-pinocchio" {
			v.BuilderName = "std"
		}

		if k == "telepathy-python" {
			v.BuilderName = "std"
		}

		if k == "telepathy-qt" {
			v.BuilderName = "std"
		}

		if k == "telepathy-qt4-prototype" {
			v.BuilderName = "std"
		}

		if k == "telepathy-qt4-yell" {
			v.BuilderName = "std"
		}

		if k == "telepathy-qt4" {
			v.BuilderName = "std"
		}

		if k == "telepathy-rakia" {
			v.BuilderName = "std"
		}

		if k == "telepathy-ring" {
			v.BuilderName = "std"
		}

		if k == "telepathy-salut" {
			v.BuilderName = "std"
		}

		if k == "telepathy-sofiasip" {
			v.BuilderName = "std"
		}

		if k == "telepathy-spec" {
			v.BuilderName = "std"
		}

		if k == "telepathy-stream-engine" {
			v.BuilderName = "std"
		}

		if k == "telepathy-sunshine" {
			v.BuilderName = "std"
		}

		if k == "telepathy-yell" {
			v.BuilderName = "std"
		}

		if k == "termcap" {
			v.BuilderName = "std"
		}

		if k == "termutils" {
			v.BuilderName = "std"
		}

		if k == "teseq" {
			v.BuilderName = "std"
		}

		if k == "tesseract-gui" {
			v.BuilderName = ""
		}

		if k == "tesseract" {
			v.BuilderName = ""
		}

		if k == "test-suite" {
			v.BuilderName = "std"
		}

		if k == "teximpatient" {
			v.BuilderName = "std"
		}

		if k == "texinfo" {
			v.BuilderName = "std"
		}

		if k == "texlive" {
			v.BuilderName = "texlive"
		}

		if k == "texlivetug" {
			v.BuilderName = ""
		}

		if k == "text-highlight" {
			v.BuilderName = "perl_mod"
		}

		if k == "tftp-hpa" {
			v.BuilderName = "std"
		}

		if k == "thales" {
			v.BuilderName = "std"
		}

		if k == "the-board" {
			v.BuilderName = "std"
		}

		if k == "themus" {
			v.BuilderName = "std"
		}

		if k == "thunar-vfs" {
			v.BuilderName = ""
		}

		if k == "thunderbird" {
			v.BuilderName = "thunderbird"
		}

		if k == "tiff" {
			v.BuilderName = "std"
		}

		if k == "tigase-server" {
			v.BuilderName = ""
		}

		if k == "tigervnc" {
			v.BuilderName = "std_cmake"
		}

		if k == "tightvnc" {
			v.BuilderName = ""
		}

		if k == "tile-forth" {
			v.BuilderName = "std"
		}

		if k == "time" {
			v.BuilderName = "std"
		}

		if k == "tk" {
			v.BuilderName = "tcltk"
		}

		if k == "tkabber-plugins" {
			v.BuilderName = "tkabber"
		}

		if k == "tkabber" {
			v.BuilderName = "tkabber"
		}

		if k == "tkimage" {
			v.BuilderName = ""
		}

		if k == "tls" {
			v.BuilderName = "std"
		}

		if k == "tmw" {
			v.BuilderName = ""
		}

		if k == "tolua++" {
			v.BuilderName = ""
		}

		if k == "tomboy" {
			v.BuilderName = "std"
		}

		if k == "tor" {
			v.BuilderName = "std"
		}

		if k == "torsocks" {
			v.BuilderName = "std"
		}

		if k == "totem-pl-parser" {
			v.BuilderName = "std"
		}

		if k == "totem" {
			v.BuilderName = "std"
		}

		if k == "toutdoux" {
			v.BuilderName = "std"
		}

		if k == "toxcore" {
			v.BuilderName = "std"
		}

		if k == "tracker-miner-media" {
			v.BuilderName = "std"
		}

		if k == "tracker" {
			v.BuilderName = "std"
		}

		if k == "tramp" {
			v.BuilderName = "std"
		}

		if k == "transset" {
			v.BuilderName = "std"
		}

		if k == "trapproto" {
			v.BuilderName = "std"
		}

		if k == "tree" {
			v.BuilderName = "tree"
		}

		if k == "trilobite" {
			v.BuilderName = "std"
		}

		if k == "tritium" {
			v.BuilderName = ""
		}

		if k == "trueprint" {
			v.BuilderName = "std"
		}

		if k == "ttf-bitstream-vera" {
			v.BuilderName = "std"
		}

		if k == "tuna" {
			v.BuilderName = "std"
		}

		if k == "turbovnc" {
			v.BuilderName = "turbovnc"
		}

		if k == "tuxnes" {
			v.BuilderName = ""
		}

		if k == "twm" {
			v.BuilderName = "std"
		}

		if k == "tzcode" {
			v.BuilderName = ""
		}

		if k == "tzdata" {
			v.BuilderName = "tzdata"
		}

		if k == "uClibc" {
			v.BuilderName = "std"
		}

		if k == "uae" {
			v.BuilderName = "std"
		}

		if k == "ucb-toolset-x86_64-pc-linux-gnu" {
			v.BuilderName = "crossbuilder_tc"
		}

		if k == "ucblogo" {
			v.BuilderName = "std"
		}

		if k == "ucl" {
			v.BuilderName = "std"
		}

		if k == "ucommon" {
			v.BuilderName = "std"
		}

		if k == "udev" {
			v.BuilderName = "std"
		}

		if k == "udisks" {
			v.BuilderName = "std"
		}

		if k == "udns" {
			v.BuilderName = "udns"
		}

		if k == "uhttpmock" {
			v.BuilderName = "std"
		}

		if k == "ulogd" {
			v.BuilderName = ""
		}

		if k == "unifont" {
			v.BuilderName = "std"
		}

		if k == "units" {
			v.BuilderName = "std"
		}

		if k == "unittest-cpp" {
			v.BuilderName = "std_cmake"
		}

		if k == "unixODBC" {
			v.BuilderName = "std"
		}

		if k == "unrtf" {
			v.BuilderName = "std"
		}

		if k == "unzip" {
			v.BuilderName = "infozip"
		}

		if k == "update-manager" {
			v.BuilderName = "std"
		}

		if k == "upower" {
			v.BuilderName = "std"
		}

		if k == "upx" {
			v.BuilderName = ""
		}

		if k == "usbutils" {
			v.BuilderName = "std"
		}

		if k == "users-guide" {
			v.BuilderName = "std"
		}

		if k == "userv-utils" {
			v.BuilderName = "std"
		}

		if k == "userv" {
			v.BuilderName = "std"
		}

		if k == "usrp" {
			v.BuilderName = "std"
		}

		if k == "util-linux-ng" {
			v.BuilderName = "std"
		}

		if k == "util-linux" {
			v.BuilderName = "utillinux"
		}

		if k == "util-macros" {
			v.BuilderName = "std"
		}

		if k == "uucp" {
			v.BuilderName = "std"
		}

		if k == "uwsgi" {
			v.BuilderName = "uwsgi"
		}

		if k == "v" {
			v.BuilderName = "std"
		}

		if k == "vacuum" {
			v.BuilderName = "std_cmake"
		}

		if k == "vala" {
			v.BuilderName = "std"
		}

		if k == "valencia" {
			v.BuilderName = "std"
		}

		if k == "valgrind" {
			v.BuilderName = "std"
		}

		if k == "varconf" {
			v.BuilderName = "std"
		}

		if k == "vc-dwim" {
			v.BuilderName = "std"
		}

		if k == "vcdimager" {
			v.BuilderName = "std"
		}

		if k == "vdpauinfo" {
			v.BuilderName = "std"
		}

		if k == "vegastrike-data" {
			v.BuilderName = ""
		}

		if k == "vegastrike-src" {
			v.BuilderName = "std"
		}

		if k == "vera" {
			v.BuilderName = "std"
		}

		if k == "vhba-module" {
			v.BuilderName = "vhbamodule"
		}

		if k == "videoproto" {
			v.BuilderName = "std"
		}

		if k == "viewres" {
			v.BuilderName = "std"
		}

		if k == "vigra" {
			v.BuilderName = "std_cmake"
		}

		if k == "vinagre" {
			v.BuilderName = "std"
		}

		if k == "vino" {
			v.BuilderName = "std"
		}

		if k == "virglrenderer" {
			v.BuilderName = "std"
		}

		if k == "virtme" {
			v.BuilderName = "std"
		}

		if k == "virtualbox" {
			v.BuilderName = "vbox"
		}

		if k == "vlc" {
			v.BuilderName = "vlc"
		}

		if k == "vorbis-tools" {
			v.BuilderName = "std"
		}

		if k == "vte" {
			v.BuilderName = "vte"
		}

		if k == "waf" {
			v.BuilderName = "waf"
		}

		if k == "warzone2100" {
			v.BuilderName = "std"
		}

		if k == "wavpack" {
			v.BuilderName = ""
		}

		if k == "wayland-protocols" {
			v.BuilderName = "std"
		}

		if k == "wayland" {
			v.BuilderName = "wayland"
		}

		if k == "wb" {
			v.BuilderName = "std"
		}

		if k == "wdiff" {
			v.BuilderName = "std"
		}

		if k == "webalizer" {
			v.BuilderName = ""
		}

		if k == "webkitgtk" {
			v.BuilderName = "webkitgtk_cmake"
		}

		if k == "webkitgtk2.4" {
			v.BuilderName = "webkitgtk"
		}

		if k == "websocket4j" {
			v.BuilderName = "std"
		}

		if k == "wesnoth" {
			v.BuilderName = ""
		}

		if k == "weston" {
			v.BuilderName = "weston"
		}

		if k == "wfmath" {
			v.BuilderName = "std"
		}

		if k == "wget" {
			v.BuilderName = "std"
		}

		if k == "which" {
			v.BuilderName = "std"
		}

		if k == "whois" {
			v.BuilderName = ""
		}

		if k == "whoisserver" {
			v.BuilderName = ""
		}

		if k == "win-gerwin" {
			v.BuilderName = "std"
		}

		if k == "windowmaker" {
			v.BuilderName = ""
		}

		if k == "windowswmproto" {
			v.BuilderName = "std"
		}

		if k == "wine" {
			v.BuilderName = "wine"
		}

		if k == "wireless-regdb" {
			v.BuilderName = "std"
		}

		if k == "wireless_tools" {
			v.BuilderName = "wireless_tools"
		}

		if k == "wol" {
			v.BuilderName = ""
		}

		if k == "wordpress" {
			v.BuilderName = ""
		}

		if k == "wpa_supplicant" {
			v.BuilderName = "wpa_supplicant"
		}

		if k == "wvdial" {
			v.BuilderName = "std"
		}

		if k == "wvstreams" {
			v.BuilderName = "std"
		}

		if k == "wxPython" {
			v.BuilderName = "None"
		}

		if k == "wxWidgets" {
			v.BuilderName = "wxwidgets"
		}

		if k == "x11perf" {
			v.BuilderName = "std"
		}

		if k == "x11vnc" {
			v.BuilderName = ""
		}

		if k == "x264" {
			v.BuilderName = "std"
		}

		if k == "xalf" {
			v.BuilderName = "std"
		}

		if k == "xaos" {
			v.BuilderName = "std"
		}

		if k == "xapian-bindings" {
			v.BuilderName = ""
		}

		if k == "xapian-core" {
			v.BuilderName = ""
		}

		if k == "xapian-omega" {
			v.BuilderName = ""
		}

		if k == "xauth" {
			v.BuilderName = "std"
		}

		if k == "xbacklight" {
			v.BuilderName = "std"
		}

		if k == "xbiff" {
			v.BuilderName = "std"
		}

		if k == "xbitmaps" {
			v.BuilderName = "std"
		}

		if k == "xboard" {
			v.BuilderName = "std"
		}

		if k == "xbt" {
			v.BuilderName = ""
		}

		if k == "xcalc" {
			v.BuilderName = "std"
		}

		if k == "xcb-demo" {
			v.BuilderName = "std"
		}

		if k == "xcb-proto" {
			v.BuilderName = "std"
		}

		if k == "xcb-util-cursor" {
			v.BuilderName = "std"
		}

		if k == "xcb-util-errors" {
			v.BuilderName = "std"
		}

		if k == "xcb-util-image" {
			v.BuilderName = "std"
		}

		if k == "xcb-util-keysyms" {
			v.BuilderName = "std"
		}

		if k == "xcb-util-renderutil" {
			v.BuilderName = "std"
		}

		if k == "xcb-util-wm" {
			v.BuilderName = "std"
		}

		if k == "xcb-util" {
			v.BuilderName = "std"
		}

		if k == "xchat-gnome" {
			v.BuilderName = "std"
		}

		if k == "xchat" {
			v.BuilderName = "std"
		}

		if k == "xclipboard" {
			v.BuilderName = "std"
		}

		if k == "xclock" {
			v.BuilderName = "std"
		}

		if k == "xcmiscproto" {
			v.BuilderName = "std"
		}

		if k == "xcmsdb" {
			v.BuilderName = "std"
		}

		if k == "xcompmgr" {
			v.BuilderName = "std"
		}

		if k == "xconsole" {
			v.BuilderName = "std"
		}

		if k == "xcursor-themes" {
			v.BuilderName = "std"
		}

		if k == "xcursorgen" {
			v.BuilderName = "std"
		}

		if k == "xdbedizzy" {
			v.BuilderName = "std"
		}

		if k == "xdebug" {
			v.BuilderName = ""
		}

		if k == "xdg-user-dirs-gtk" {
			v.BuilderName = "std"
		}

		if k == "xdg-utils" {
			v.BuilderName = "std"
		}

		if k == "xditview" {
			v.BuilderName = "std"
		}

		if k == "xdm" {
			v.BuilderName = "std"
		}

		if k == "xdpyinfo" {
			v.BuilderName = "std"
		}

		if k == "xdriinfo" {
			v.BuilderName = "std"
		}

		if k == "xedit" {
			v.BuilderName = "std"
		}

		if k == "xen" {
			v.BuilderName = "xen"
		}

		if k == "xerces-c-tools" {
			v.BuilderName = "std"
		}

		if k == "xerces-c" {
			v.BuilderName = "std"
		}

		if k == "xev" {
			v.BuilderName = "std"
		}

		if k == "xextproto" {
			v.BuilderName = "std"
		}

		if k == "xeyes" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-acecad" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-aiptek" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-calcomp" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-citron" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-digitaledge" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-dmc" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-dynapro" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-elo2300" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-elographics" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-evdev" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-fpit" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-hyperpen" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-jamstudio" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-joystick" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-keyboard" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-libinput" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-magellan" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-magictouch" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-microtouch" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-mouse" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-mutouch" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-palmax" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-penmount" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-spaceorb" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-summa" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-synaptics" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-tek4957" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-ur98" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-vmmouse" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-void" {
			v.BuilderName = "std"
		}

		if k == "xf86-input-wacom" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-amd" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-apm" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-ark" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-ast" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-ati" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-chips" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-cirrus" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-cyrix" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-dummy" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-fbdev" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-freedreno" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-geode" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-glide" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-glint" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-i128" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-i740" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-i810" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-impact" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-imstt" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-intel" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-mach64" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-mga" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-modesetting" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-neomagic" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-newport" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-nouveau" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-nsc" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-nv" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-omap" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-openchrome" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-opentegra" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-qxl" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-r128" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-radeonhd" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-rendition" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-s3" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-s3virge" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-savage" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-siliconmotion" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-sis" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-sisusb" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-sunbw2" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-suncg14" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-suncg3" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-suncg6" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-sunffb" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-sunleo" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-suntcx" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-tdfx" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-tga" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-trident" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-tseng" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-v4l" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-vermilion" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-vesa" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-vga" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-via" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-vmware" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-voodoo" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-wsfb" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-xgi" {
			v.BuilderName = "std"
		}

		if k == "xf86-video-xgixp" {
			v.BuilderName = "std"
		}

		if k == "xf86bigfontproto" {
			v.BuilderName = "std"
		}

		if k == "xf86dga" {
			v.BuilderName = "std"
		}

		if k == "xf86dgaproto" {
			v.BuilderName = "std"
		}

		if k == "xf86driproto" {
			v.BuilderName = "std"
		}

		if k == "xf86miscproto" {
			v.BuilderName = "std"
		}

		if k == "xf86rushproto" {
			v.BuilderName = "std"
		}

		if k == "xf86vidmodeproto" {
			v.BuilderName = "std"
		}

		if k == "xfce-utils" {
			v.BuilderName = ""
		}

		if k == "xfce4-appfinder" {
			v.BuilderName = ""
		}

		if k == "xfce4-clipman-plugin" {
			v.BuilderName = ""
		}

		if k == "xfce4-dev-tools" {
			v.BuilderName = ""
		}

		if k == "xfce4-panel" {
			v.BuilderName = ""
		}

		if k == "xfce4-session" {
			v.BuilderName = ""
		}

		if k == "xfce4-settings" {
			v.BuilderName = ""
		}

		if k == "xfce4-terminal" {
			v.BuilderName = "std"
		}

		if k == "xfce4-xkb-plugin" {
			v.BuilderName = ""
		}

		if k == "xfconf" {
			v.BuilderName = "std"
		}

		if k == "xfd" {
			v.BuilderName = "std"
		}

		if k == "xfdesktop" {
			v.BuilderName = ""
		}

		if k == "xfindproxy" {
			v.BuilderName = "std"
		}

		if k == "xfontsel" {
			v.BuilderName = "std"
		}

		if k == "xfs" {
			v.BuilderName = "std"
		}

		if k == "xfsdump" {
			v.BuilderName = "xfs"
		}

		if k == "xfsinfo" {
			v.BuilderName = "std"
		}

		if k == "xfsprogs" {
			v.BuilderName = "xfs"
		}

		if k == "xfwm4" {
			v.BuilderName = ""
		}

		if k == "xfwp" {
			v.BuilderName = "std"
		}

		if k == "xgamma" {
			v.BuilderName = "std"
		}

		if k == "xgc" {
			v.BuilderName = "std"
		}

		if k == "xhippo" {
			v.BuilderName = "std"
		}

		if k == "xhost" {
			v.BuilderName = "std"
		}

		if k == "ximian-connector" {
			v.BuilderName = "std"
		}

		if k == "ximian-setup-tools" {
			v.BuilderName = "std"
		}

		if k == "xineramaproto" {
			v.BuilderName = "std"
		}

		if k == "xinetd" {
			v.BuilderName = "std"
		}

		if k == "xinit" {
			v.BuilderName = "std"
		}

		if k == "xinput" {
			v.BuilderName = "std"
		}

		if k == "xkbcomp" {
			v.BuilderName = "std"
		}

		if k == "xkbdata" {
			v.BuilderName = "std"
		}

		if k == "xkbevd" {
			v.BuilderName = "std"
		}

		if k == "xkbmapgui" {
			v.BuilderName = ""
		}

		if k == "xkbprint" {
			v.BuilderName = "std"
		}

		if k == "xkbutils" {
			v.BuilderName = "std"
		}

		if k == "xkeyboard-config" {
			v.BuilderName = "std"
		}

		if k == "xkill" {
			v.BuilderName = "std"
		}

		if k == "xload" {
			v.BuilderName = "std"
		}

		if k == "xlogmaster" {
			v.BuilderName = "std"
		}

		if k == "xlogo" {
			v.BuilderName = "std"
		}

		if k == "xlsatoms" {
			v.BuilderName = "std"
		}

		if k == "xlsclients" {
			v.BuilderName = "std"
		}

		if k == "xlsfonts" {
			v.BuilderName = "std"
		}

		if k == "xmag" {
			v.BuilderName = "std"
		}

		if k == "xman" {
			v.BuilderName = "std"
		}

		if k == "xmessage" {
			v.BuilderName = "std"
		}

		if k == "xmh" {
			v.BuilderName = "std"
		}

		if k == "xml-i18n-tools" {
			v.BuilderName = "std"
		}

		if k == "xmlto" {
			v.BuilderName = "std"
		}

		if k == "xmodmap" {
			v.BuilderName = "std"
		}

		if k == "xmonad" {
			v.BuilderName = ""
		}

		if k == "xmore" {
			v.BuilderName = "std"
		}

		if k == "xmpppy" {
			v.BuilderName = ""
		}

		if k == "xnee" {
			v.BuilderName = "std"
		}

		if k == "xorg-cf-files" {
			v.BuilderName = "std"
		}

		if k == "xorg-docs" {
			v.BuilderName = "std"
		}

		if k == "xorg-gtest" {
			v.BuilderName = "std"
		}

		if k == "xorg-server" {
			v.BuilderName = "std"
		}

		if k == "xorg-sgml-doctools" {
			v.BuilderName = "std"
		}

		if k == "xorriso" {
			v.BuilderName = "std"
		}

		if k == "xpenguins_applet" {
			v.BuilderName = "std"
		}

		if k == "xphelloworld" {
			v.BuilderName = "std"
		}

		if k == "xplsprinters" {
			v.BuilderName = "std"
		}

		if k == "xpm2wico" {
			v.BuilderName = "std"
		}

		if k == "xpr" {
			v.BuilderName = "std"
		}

		if k == "xprehashprinterlist" {
			v.BuilderName = "std"
		}

		if k == "xprop" {
			v.BuilderName = "std"
		}

		if k == "xproto" {
			v.BuilderName = "std"
		}

		if k == "xproxymanagementprotocol" {
			v.BuilderName = "std"
		}

		if k == "xpyb" {
			v.BuilderName = "xpyb"
		}

		if k == "xrandr" {
			v.BuilderName = "std"
		}

		if k == "xrdb" {
			v.BuilderName = "std"
		}

		if k == "xrdp" {
			v.BuilderName = ""
		}

		if k == "xrefresh" {
			v.BuilderName = "std"
		}

		if k == "xrx" {
			v.BuilderName = "std"
		}

		if k == "xsane" {
			v.BuilderName = "std"
		}

		if k == "xscope" {
			v.BuilderName = "std"
		}

		if k == "xscreensaver" {
			v.BuilderName = ""
		}

		if k == "xset" {
			v.BuilderName = "std"
		}

		if k == "xsetmode" {
			v.BuilderName = "std"
		}

		if k == "xsetpointer" {
			v.BuilderName = "std"
		}

		if k == "xsetroot" {
			v.BuilderName = "std"
		}

		if k == "xshogi" {
			v.BuilderName = "std"
		}

		if k == "xsm" {
			v.BuilderName = "std"
		}

		if k == "xstdcmap" {
			v.BuilderName = "std"
		}

		if k == "xterm" {
			v.BuilderName = "std"
		}

		if k == "xtrans" {
			v.BuilderName = "std"
		}

		if k == "xtrap" {
			v.BuilderName = "std"
		}

		if k == "xts" {
			v.BuilderName = "std"
		}

		if k == "xulrunner" {
			v.BuilderName = "xulrunner"
		}

		if k == "xvidcap" {
			v.BuilderName = "std"
		}

		if k == "xvidcore" {
			v.BuilderName = "xvidcore"
		}

		if k == "xvidtune" {
			v.BuilderName = "std"
		}

		if k == "xvinfo" {
			v.BuilderName = "std"
		}

		if k == "xwd" {
			v.BuilderName = "std"
		}

		if k == "xwininfo" {
			v.BuilderName = "std"
		}

		if k == "xwud" {
			v.BuilderName = "std"
		}

		if k == "xz" {
			v.BuilderName = "std"
		}

		if k == "yabause" {
			v.BuilderName = ""
		}

		if k == "yajl" {
			v.BuilderName = "std_cmake"
		}

		if k == "yaml" {
			v.BuilderName = "std"
		}

		if k == "yasm" {
			v.BuilderName = "std"
		}

		if k == "yelp-tools" {
			v.BuilderName = "std"
		}

		if k == "yelp-xsl" {
			v.BuilderName = "std"
		}

		if k == "yelp" {
			v.BuilderName = "std"
		}

		if k == "yp-tools" {
			v.BuilderName = "std"
		}

		if k == "ypbind-bsd" {
			v.BuilderName = "std"
		}

		if k == "ypbind-mt" {
			v.BuilderName = "std"
		}

		if k == "ypbind" {
			v.BuilderName = "std"
		}

		if k == "ypmake" {
			v.BuilderName = "std"
		}

		if k == "yppasswd" {
			v.BuilderName = "std"
		}

		if k == "yps" {
			v.BuilderName = "std"
		}

		if k == "ypserv" {
			v.BuilderName = "std"
		}

		if k == "zbar" {
			v.BuilderName = "zbare"
		}

		if k == "zeitgeist" {
			v.BuilderName = "std"
		}

		if k == "zenity" {
			v.BuilderName = "std"
		}

		if k == "zeromq" {
			v.BuilderName = "std"
		}

		if k == "zile" {
			v.BuilderName = "std"
		}

		if k == "zint" {
			v.BuilderName = "std_cmake"
		}

		if k == "zip" {
			v.BuilderName = "infozip"
		}

		if k == "zisofs-tools" {
			v.BuilderName = "std"
		}

		if k == "zlib" {
			v.BuilderName = "zlib"
		}

		if k == "zziplib" {
			v.BuilderName = "zziplib"
		}
	}
	return nil
}
