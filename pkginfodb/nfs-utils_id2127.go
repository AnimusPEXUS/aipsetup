package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_nfs_utils = &basictypes.PackageInfo{

	Description: `write something here, please`,
	HomePage:    "https://sourceforge.net/projects/dgen",

	BuilderName: "std",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"sf_hosted:aa-project", "sf_hosted:acpid", "sf_hosted:acpitool", "sf_hosted:alleg", "sf_hosted:alsamodular", "sf_hosted:amule", "sf_hosted:aoetools", "sf_hosted:aqsis", "sf_hosted:aria2", "sf_hosted:armagetronad", "sf_hosted:arss", "sf_hosted:asciidoc", "sf_hosted:assimp", "sf_hosted:astyle", "sf_hosted:atari800", "sf_hosted:audacity", "sf_hosted:avidemux", "sf_hosted:azureus", "sf_hosted:bastard", "sf_hosted:beecrypt", "sf_hosted:beye", "sf_hosted:blackboxwm", "sf_hosted:bogofilter", "sf_hosted:boost", "sf_hosted:bridge", "sf_hosted:btmgr", "sf_hosted:bzflag", "sf_hosted:cdemu", "sf_hosted:cdrdao", "sf_hosted:cdrtools", "sf_hosted:cedet", "sf_hosted:cel", "sf_hosted:cfv", "sf_hosted:cgkit", "sf_hosted:check", "sf_hosted:clamav", "sf_hosted:cppunit", "sf_hosted:cracklib", "sf_hosted:crayzedsgui", "sf_hosted:crengine", "sf_hosted:crystal", "sf_hosted:csound", "sf_hosted:cups", "sf_hosted:curlftpfs", "sf_hosted:dgen", "sf_hosted:djvu", "sf_hosted:docbook", "sf_hosted:docbook2x", "sf_hosted:dom4j", "sf_hosted:dosbox", "sf_hosted:doxygen", "sf_hosted:dvdstyler", "sf_hosted:e2fsprogs", "sf_hosted:ecb", "sf_hosted:enlightenment", "sf_hosted:epydoc", "sf_hosted:eric-ide", "sf_hosted:espeak", "sf_hosted:espgs", "sf_hosted:evms", "sf_hosted:expat", "sf_hosted:expect", "sf_hosted:faac", "sf_hosted:fakerootng", "sf_hosted:fileschanged", "sf_hosted:flac", "sf_hosted:flex", "sf_hosted:fluidsynth", "sf_hosted:fontforge", "sf_hosted:freeassociation", "sf_hosted:freeciv", "sf_hosted:freeglut", "sf_hosted:freeimage", "sf_hosted:freepascal", "sf_hosted:freetype", "sf_hosted:ftgl", "sf_hosted:gens", "sf_hosted:geshi", "sf_hosted:giflib", "sf_hosted:gimp-print", "sf_hosted:gkernel", "sf_hosted:glew", "sf_hosted:gmerlin", "sf_hosted:gns-3", "sf_hosted:gnuplot", "sf_hosted:gphoto", "sf_hosted:gsoc-impos", "sf_hosted:gswitchit", "sf_hosted:gtk2-perl", "sf_hosted:gtkglext", "sf_hosted:gtkspell", "sf_hosted:guichan", "sf_hosted:gwenview", "sf_hosted:hdparm", "sf_hosted:heirloom", "sf_hosted:heroines", "sf_hosted:hplip", "sf_hosted:hte", "sf_hosted:hylafax", "sf_hosted:infozip", "sf_hosted:iodbc", "sf_hosted:ipband", "sf_hosted:iperf", "sf_hosted:iperf2", "sf_hosted:ipsec-tools", "sf_hosted:irrlicht", "sf_hosted:jack-rack", "sf_hosted:jackit", "sf_hosted:jagoclient", "sf_hosted:junit", "sf_hosted:kuickshow", "sf_hosted:lame", "sf_hosted:lazarus", "sf_hosted:lcms", "sf_hosted:libass", "sf_hosted:libavc1394", "sf_hosted:libclc", "sf_hosted:libdv", "sf_hosted:libebook", "sf_hosted:libexif", "sf_hosted:libircclient", "sf_hosted:libjpeg", "sf_hosted:libjpeg-turbo", "sf_hosted:libjson", "sf_hosted:libmng", "sf_hosted:liboauth", "sf_hosted:libpng", "sf_hosted:libquicktime", "sf_hosted:libraw1394", "sf_hosted:libseccomp", "sf_hosted:libsieve", "sf_hosted:libtirpc", "sf_hosted:libtorrent", "sf_hosted:libusb", "sf_hosted:libvncserver", "sf_hosted:libwpd", "sf_hosted:linux-igd", "sf_hosted:linux-ntfs", "sf_hosted:linux-udf", "sf_hosted:linux-usb", "sf_hosted:linuxwacom", "sf_hosted:lives", "sf_hosted:log4c", "sf_hosted:lprng", "sf_hosted:mad", "sf_hosted:madwifi", "sf_hosted:mantisbt", "sf_hosted:marathon", "sf_hosted:meanwhile", "sf_hosted:mediatomb", "sf_hosted:mednafen", "sf_hosted:mesa3d", "sf_hosted:meshlab", "sf_hosted:mhash", "sf_hosted:ming", "sf_hosted:mjpeg", "sf_hosted:mldonkey", "sf_hosted:mlt", "sf_hosted:mpg123", "sf_hosted:ms-sys", "sf_hosted:multivalent", "sf_hosted:nas", "sf_hosted:nc110", "sf_hosted:net-snmp", "sf_hosted:net-tools", "sf_hosted:netcat", "sf_hosted:nethack", "sf_hosted:nfs", "sf_hosted:nrappkit", "sf_hosted:ogl-math", "sf_hosted:ogre", "sf_hosted:ol2mbox", "sf_hosted:omxil", "sf_hosted:open1x", "sf_hosted:opencvlibrary", "sf_hosted:opende", "sf_hosted:openil", "sf_hosted:openjade", "sf_hosted:openjpeg.mirror", "sf_hosted:opennx", "sf_hosted:openobex", "sf_hosted:oscaf", "sf_hosted:p7zip", "sf_hosted:palomino-sim", "sf_hosted:paps", "sf_hosted:pcmcia-cs", "sf_hosted:pdfedit", "sf_hosted:pdfshuffler", "sf_hosted:pdftohtml", "sf_hosted:phpbb", "sf_hosted:phpdocu", "sf_hosted:pidgin", "sf_hosted:plib", "sf_hosted:podofo", "sf_hosted:poopmup", "sf_hosted:poptop", "sf_hosted:pptpclient", "sf_hosted:procps-ng", "sf_hosted:psi", "sf_hosted:psmisc", "sf_hosted:pupnp", "sf_hosted:pyode", "sf_hosted:pyopengl", "sf_hosted:pyqt", "sf_hosted:pywebsvcs", "sf_hosted:pyxml", "sf_hosted:qdvdauthor", "sf_hosted:qimageblitz", "sf_hosted:qpdf", "sf_hosted:qucs", "sf_hosted:quesoglc", "sf_hosted:rdesktop", "sf_hosted:recordmydesktop", "sf_hosted:retroshare", "sf_hosted:rhash", "sf_hosted:rp-l2tp", "sf_hosted:scintilla", "sf_hosted:scons", "sf_hosted:scribus", "sf_hosted:sf-xpaint", "sf_hosted:slrn", "sf_hosted:smartmontools", "sf_hosted:sonik", "sf_hosted:soprano", "sf_hosted:sox", "sf_hosted:spacenav", "sf_hosted:springrts", "sf_hosted:squashfs", "sf_hosted:squirrel-sql", "sf_hosted:stereograph", "sf_hosted:strace", "sf_hosted:strigi", "sf_hosted:subcomposer", "sf_hosted:sv1", "sf_hosted:swig", "sf_hosted:synfig", "sf_hosted:tcl", "sf_hosted:tcllib", "sf_hosted:tftp-server", "sf_hosted:tintin", "sf_hosted:tkgate", "sf_hosted:tkimg", "sf_hosted:tls", "sf_hosted:trn", "sf_hosted:trousers", "sf_hosted:tuxracer", "sf_hosted:unimediaserver", "sf_hosted:vamp", "sf_hosted:virtualgl", "sf_hosted:vnc-tight", "sf_hosted:warzone2100", "sf_hosted:webcamstudio", "sf_hosted:wgois", "sf_hosted:wine", "sf_hosted:winetools", "sf_hosted:worldforge", "sf_hosted:wvware", "sf_hosted:wxpython", "sf_hosted:wxwindows", "sf_hosted:x3270", "sf_hosted:xdvi", "sf_hosted:xfce", "sf_hosted:xindy", "sf_hosted:xine", "sf_hosted:xmpppy", "sf_hosted:xvidcap", "sf_hosted:zbar", "sf_hosted:zint", "sf_hosted:zziplib"},

	TarballVersionTool: "std",

	Filters:               []string{},
	TarballName:           "nfs-utils",
	TarballFileNameParser: "std",
	TarballProvider:       "sf",
	TarballProviderArguments: []string{
		"dgen"},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}
