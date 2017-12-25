package infoeditor

//       sourceforge's project name
//              \_______         tarballs spawned by them
//                      \                  |
//                       ------,        ,--
//                             V       V
var SOURCEFORGE_PROJECTS = map[string]([]string){

	"aa-project": []string{
		"aa3d",
		"aalib",
		"aavga",
		"aview",
		"bb",
		"bb-fnl",
	},
	"acpid": []string{
		"acpid",
	},
	"acpitool": []string{
		"acpitool",
	},
	"alleg": []string{
		"all",
		"allegro",
	},
	"alsamodular": []string{
		"ams",
		"kaconnect",
		"qamix",
		"qarecord",
		"qmidiarp",
		"qmidicontrol",
		"qmidiroute",
	},
	"amule": []string{
		"aMule",
		"amule",
	},
	"aoetools": []string{
		"aoetools",
		"cec",
		"vblade",
	},
	"aqsis": []string{
		"Aqsis",
		"aqsis",
	},
	"aria2": []string{
		"aria2c",
	},
	"armagetronad": []string{
		"armagetronad",
		"armagetronad-dedicated",
		"armagetronad-styct",
		"armagetronad-styct-dedicated",
		"armagetronad-styctap",
		"armagetronad-styctap-dedicated",
		"armagetronad-t2o",
		"fortschool",
		"fortschool-dedicated",
	},
	"arss": []string{
		"arss",
	},
	"asciidoc": []string{
		"asciidoc",
	},
	"assimp": []string{
		"assimp",
	},
	"astyle": []string{
		"AStyle",
		"astyle",
	},
	"atari800": []string{
		"Atari800",
		"atari800",
	},
	"audacity": []string{
		"audacity",
		"audacity-src",
	},
	"avidemux": []string{
		"avidemux",
		"avidemux2_cli",
		"avidemux_sdk",
	},
	"avifile": []string{},
	"azureus": []string{
		"Vuze",
	},
	"bastard": []string{
		"bastard",
		"bastard_bin",
		"bastard_src",
		"libdisasm",
		"libdisasm_bin",
		"libdisasm_src",
		"seer",
		"typhoon",
	},
	"beecrypt": []string{
		"beecrypt",
	},
	"beye": []string{
		"biew",
	},
	"blackboxwm": []string{
		"blackbox",
	},
	"bogofilter": []string{
		"bogofilter",
	},
	"boost": []string{
		"bjam",
		"boost",
	},
	"bridge": []string{
		"bridge-utils",
	},
	"btmgr": []string{
		"btmgr",
	},
	"bzflag": []string{
		"bzflag",
	},
	"cdemu": []string{
		"cdemu-client",
		"cdemu-daemon",
		"gcdemu",
		"image-analyzer",
		"libmirage",
		"vhba-module",
	},
	"cdrdao": []string{
		"cdrdao",
		"cygwin",
	},
	"cdrtools": []string{
		"cdrtools",
	},
	"cedet": []string{
		"COGRE",
		"EDE",
		"cedet",
		"eieio",
		"semantic",
		"speedbar",
	},
	"cel": []string{
		"cel",
		"cel-pseudo-stable",
		"cel-src",
	},
	"cfv": []string{
		"cfv",
	},
	"cgkit": []string{
		"cgkit",
		"maya",
	},
	"check": []string{
		"check",
	},
	"clamav": []string{
		"clamav",
	},
	"cppunit": []string{
		"cppunit",
	},
	"cracklib": []string{
		"cracklib",
	},
	"crayzedsgui": []string{
		"CEED-snapshot",
		"CEGUI",
		"CEGUI-DEMOS",
		"CEGUI-DEPS",
		"CEGUI-DOCS",
		"CEGUI-SDK",
		"CEGUI-Win32-Demos",
		"CEImagesetEditor",
		"CEImagesetEditor-v",
		"CELayoutEditor",
		"CELayoutEditorSetup",
		"QuadraticLook",
		"SILLY",
		"SILLY-DEPS",
		"SILLY-DOCS",
		"SILLY-SDK",
		"ceed",
		"ceed-snapshot",
		"cegui",
		"cegui-deps",
		"cegui-docs",
		"cegui-sdk",
		"cegui_deps_vc",
		"cegui_mk",
		"cegui_mk2",
		"cegui_mk2-apireference",
		"cegui_mk2-deps-vc",
		"cegui_mk2-deps-vc6-stlport",
		"cegui_mk2-deps-vc7.0-stlport",
		"cegui_mk2-source",
		"cegui_mk2-win32",
		"deps",
		"premake",
		"quadratic_look",
	},
	"crengine": []string{
		"V",
		"V3Update_fb2_libs_only_CR",
		"V3update-c",
		"V3update-cr",
		"V3update-full-cr",
		"V3update_full_cr",
		"V3update_full_fb",
		"coolreader3",
		"cr",
		"cr3",
		"cr3-V3-libfb2-so-only",
		"cr3-newui-opengl-win32",
		"cr3-newui-opengl-win32-qt-static-angle",
		"cr3-newui-opengl-win32-static",
		"cr3-oi",
		"cr3-qt-win",
		"cr3-qt-win32",
		"cr3-src",
		"cr3-v3a-bin",
		"cr3-v5-bin",
		"cr3-win",
		"cr3-win32-qt-opengl",
		"cr3qt",
		"crengine-src",
		"keymaps",
		"lbook-fb",
		"new-cr",
		"skindemo-win",
		"v",
		"v3-cr3-only-cr",
		"v3-newsdk-cr",
		"v3-simulator-win32-cr",
		"v3a-cr",
		"v3a-cr3-only",
		"v3update_cr",
		"v5-cr",
		"v5-cr3-only",
	},
	"crystal": []string{
		"cel-pseudo-stable",
		"crystalspace-src",
		"cs",
		"cs-pseudo-stable",
		"cs96_patch",
		"cslinux",
		"cswin",
		"msvc7_libs",
		"msvc_libs",
		"patch",
	},
	"csound": []string{
		"Csound",
		"Csound-PNaCl",
		"Csound-android",
		"Csound-iOS",
		"Csound_manual",
		"DebianCsound",
		"Debian_amd",
		"HRTF-data",
		"Loris_STK_src",
		"csound",
		"csound-OSX-SDK",
		"csound-android",
		"csound-emscripten",
		"csound-iOS",
		"csound~_v",
		"olpcsound",
		"wiiuse_v",
	},
	"cups": []string{
		"cups",
		"cups-drivers-all",
		"kups",
		"qtcups",
		"xpp",
	},
	"curlftpfs": []string{
		"curlftpfs",
	},
	"dgen": []string{
		"dgen-sdl",
	},
	"djvu": []string{
		"DjVu",
		"DjVuGUI",
		"DjVuRefLib",
		"DjVuUnixViewer",
		"djview",
		"djview4",
		"djvulibre",
		"gsdjvu",
	},
	"docbook": []string{
		"FOeditor",
		"assembly-xsl",
		"docbook-dsssl",
		"docbook-dsssl-doc",
		"docbook-epub",
		"docbook-menu",
		"docbook-slides",
		"docbook-slides-demo",
		"docbook-website",
		"docbook-xsl",
		"docbook-xsl-doc",
		"docbook-xsl-ns",
		"docbook-xsl-saxon",
		"docbook-xsl-webhelpindexer",
		"docbook-xsl-xalan",
		"docbook5-xsl",
		"fo-editor",
	},
	"docbook2x": []string{
		"docbook2X",
		"docbook2man-sgmlspl",
	},
	"dom4j": []string{
		"dom4j",
	},
	"dosbox": []string{
		"DOSBox",
		"dosbox",
	},
	"doxygen": []string{
		"doxygen",
		"doxygen_manual",
		"doxygenw",
		"tr",
	},
	"dvdstyler": []string{
		"DVDStyler",
	},
	"e2fsprogs": []string{
		"e2fsprogs",
		"e2fsprogs-1.43~WIP",
		"e2fsprogs-libs",
		"e2fsprogs-libs-1.43~WIP",
	},
	"ecb": []string{
		"ecb",
	},
	"enlightenment": []string{
		"E-CPU2v",
		"E-Dial",
		"E-Mixer",
		"E-Mount",
		"E-PlayCD",
		"E-SD",
		"E-Xmms",
		"e16",
		"e16-docs",
		"e16-theme-BlueSteel",
		"e16-theme-BrushedMetal-Tigert",
		"e16-theme-Ganymede",
		"e16-theme-ShinyMetal",
		"e16-themes",
		"e16keyedit",
		"e16menuedit",
		"e16menuedit2",
		"eApm",
		"ecdplay",
		"enlightenment",
		"enlightenment-docs",
		"enlightenment-theme-BlueSteel",
		"enlightenment-theme-BrushedMetal-Tigert",
		"enlightenment-theme-Ganymede",
		"enlightenment-theme-ShinyMetal",
		"enxamp",
		"epplet-base",
		"epplets",
		"imlib2",
		"imlib2_loaders",
	},
	"epydoc": []string{
		"epydoc",
	},
	"eric-ide": []string{
		"eric6",
		"eric6-i18n-cs",
		"eric6-i18n-de",
		"eric6-i18n-en",
		"eric6-i18n-es",
		"eric6-i18n-fr",
		"eric6-i18n-it",
		"eric6-i18n-pt",
		"eric6-i18n-ru",
		"eric6-i18n-tr",
		"eric6-i18n-zh_CN",
		"eric6-i18n-zh_CN.GB2312",
		"eric6-nolang",
	},
	"espeak": []string{
		"espeak",
		"espeakedit",
		"speak",
		"test",
	},
	"espgs": []string{
		"espgs",
	},
	"evms": []string{
		"evms",
	},
	"expat": []string{
		"expat",
		"expat_win32bin",
	},
	"expect": []string{
		"expect",
	},
	"faac": []string{
		"aacinfo",
		"cool_neromp",
		"faac",
		"faac_src",
		"faad",
		"faad2",
		"faad_src",
	},
	"fakerootng": []string{
		"fakeroot-ng",
	},
	"fileschanged": []string{
		"fileschanged",
	},
	"flac": []string{
		"flac",
	},
	"flex": []string{
		"flex",
	},
	"fluidsynth": []string{
		"fluidsynth",
	},
	"fontforge": []string{
		"FFLibs-intelmac",
		"FFLibs-mac",
		"FontForge",
		"FontForge_intel64macX.6_py",
		"FontForge_intelmacX.4_x86-py",
		"FontForge_intelmacX.5_x86-py",
		"FontForge_intelmac_x",
		"FontForge_intelmac_x86-py",
		"FontForge_intelmac_x86X.4-py",
		"FontForge_macX.4_ppc_py",
		"FontForge_macX.4_py",
		"FontForge_macX.5_ppc-py",
		"FontForge_macX.5_ppc_py",
		"FontForge_mac_ppc",
		"FontForge_mac_ppc_py",
		"FontForge_macunivX",
		"fontforge-gtktoy",
		"fontforge-ml_cygwin",
		"fontforge-solaris",
		"fontforge_cygwin",
		"fontforge_experimental",
		"fontforge_full",
		"fontforge_htdocs",
		"fontforge_ja_htdocs",
		"fontforge_solaris",
		"fontforge_solaris_sparc",
	},
	"freeassociation": []string{
		"libical",
	},
	"freeciv": []string{
		"freeciv",
		"freeciv-graphics-materials",
		"gnuwin",
		"msys2-freeciv-win",
	},
	"freeglut": []string{
		"freeglut",
	},
	"freeimage": []string{
		"FreeImage",
	},
	"freepascal": []string{
		"SVNfiles",
		"arm-gba-fpc",
		"arm-nds-fpc",
		"asldos",
		"baseos",
		"dos",
		"fcmkos",
		"fpc",
		"fpcbuild",
		"fpcjvmlinux-snapshot",
		"fpcjvmmacosx-snapshot",
		"fpcjvmwin",
		"fpcmos",
		"fpdocos",
		"fpmcos",
		"fppkgos",
		"fprcpos",
		"fpresos",
		"gdbos",
		"h",
		"i",
		"idedos-gdb",
		"ideos",
		"makeos",
		"os",
		"p",
		"powerpc-wii-fpc",
		"ppcarm-gba-i",
		"ppcarm-nds-i",
		"rmwos",
		"src",
		"tplyos",
		"ubz",
		"ubzos",
		"uchmos",
		"ucodeos",
		"ufcjsos",
		"ufclbos",
		"ufcldos",
		"ufcleos",
		"ufclios",
		"ufcljos",
		"ufclnos",
		"ufclpos",
		"ufclros",
		"ufclsos",
		"ufcluos",
		"ufclxos",
		"ufcpdos",
		"ufcsdos",
		"ufcstos",
		"ufpgtos",
		"ufpmkos",
		"ufppkos",
		"ufsjsos",
		"ufsndos",
		"ufvos",
		"ugtk",
		"uhashos",
		"uhd",
		"uhermos",
		"uimlbos",
		"ulgdos",
		"ulpngos",
		"ultaros",
		"universal-macosx",
		"uos",
		"uosslos",
		"upjpos",
		"upzlos",
		"uregos",
		"urexxos",
		"urtlcos",
		"urtleos",
		"urtloos",
		"urtluos",
		"usymbos",
		"utclos",
		"utilos",
		"utilsos",
		"ux",
		"uzipos",
		"uzlibos",
		"x",
	},
	"freetype": []string{
		"freetype",
		"freetype-doc",
		"ft",
		"ft2demos",
		"ftdmo",
		"ftdoc",
		"ftdocs",
		"ftjam",
		"ttfautohint",
	},
	"ftgl": []string{
		"ftgl",
	},
	"gens": []string{
		"BeOSGens",
		"Gens",
		"gens",
		"gens-rc",
		"gens-src-multiplatform",
		"gens-win32",
	},
	"geshi": []string{
		"GeSHi",
		"geshi",
	},
	"giflib": []string{
		"giflib",
		"libungif",
	},
	"gimp-print": []string{
		"espgs",
		"gimp-print",
		"gutenprint",
		"uninstall-gutenprint",
		"usbtb",
	},
	"gkernel": []string{
		"blktool",
		"ethtool",
		"kcompat24",
		"rng-tools",
	},
	"glew": []string{
		"glew",
	},
	"gmerlin": []string{
		"gavl",
		"gmerlin",
		"gmerlin-all-in-one",
		"gmerlin-avdecoder",
		"gmerlin-dependencies",
		"gmerlin-effectv",
		"gmerlin-encoders",
		"lemuria",
	},
	"gns-3": []string{
		"GNS",
		"GNS-3",
		"GNS3",
		"QEMU",
		"Unpack",
		"dynamips",
		"gns3-gui",
		"gns3-server",
		"lisa",
		"pemu",
		"qemu",
		"routeros",
		"vboxwrapper",
		"vpcs",
		"vyatta",
		"vyos",
	},
	"gnuplot": []string{
		"gnuplot",
		"gp",
		"gp466-cygwin-x",
		"gp50rc2-cygwin-x",
		"gpdoc",
	},
	"gphoto": []string{
		"api-docs",
		"gphoto-suite",
		"gphoto2",
		"gphoto2-manual-html",
		"gphotofs",
		"gtkam",
		"libgphoto2",
		"libgphoto2-sharp",
	},
	"gsoc-impos": []string{
		"imposition",
	},
	"gswitchit": []string{
		"gswitchit_plugins",
		"libxklavier",
	},
	"gtk2-perl": []string{
		"Cairo",
		"Cairo-GObject",
		"ExtUtils-Depends",
		"ExtUtils-PkgConfig",
		"GStreamer",
		"GStreamer-GConf",
		"GStreamer-Interfaces",
		"Glib",
		"Glib-Object-Introspection",
		"Gnome2",
		"Gnome2-Canvas",
		"Gnome2-Dia",
		"Gnome2-GConf",
		"Gnome2-PanelApplet",
		"Gnome2-Print",
		"Gnome2-Rsvg",
		"Gnome2-VFS",
		"Gnome2-Vte",
		"Gnome2-Wnck",
		"Gtk2",
		"Gtk2-GLExt",
		"Gtk2-GladeXML",
		"Gtk2-Html2",
		"Gtk2-MozEmbed",
		"Gtk2-SourceView",
		"Gtk2-Spell",
		"Gtk2-TrayIcon",
		"Gtk3",
		"Pango",
	},
	"gtkglext": []string{
		"gtkglext",
		"gtkglext-sharp",
		"gtkglext-win32",
		"gtkglextmm",
		"gtkglextmm-win32",
		"pygtkglext",
	},
	"gtkspell": []string{
		"gtkspell",
		"gtkspell3",
		"gtkspellmm",
	},
	"guichan": []string{
		"guichan",
		"guichanffdemo",
		"guichanfpsdemo",
	},
	"gwenview": []string{
		"gwenview",
		"gwenview-de",
		"gwenview-el",
		"gwenview-i18n",
		"gwenview-id",
		"gwenview-it",
		"gwenview-ja",
		"gwenview-ko",
		"gwenview-nl",
		"gwenview-pl",
		"gwenview-ro",
		"gwenview-sv",
		"gwenview_hack",
		"kipi-plugins",
		"kipi-plugins-20040611-i",
		"kipi-plugins-20040801-i",
		"kipi-plugins-20040919-i",
		"libkexif",
		"libkexif-20040611-i",
		"libkexif-20040801-i",
		"libkexif-20040919-i",
		"libkipi",
		"libkipi-20040611-i",
		"libkipi-20040801-i",
		"libkipi-20040919-i",
	},
	"hdparm": []string{
		"hdparm",
		"wiper",
	},
	"heirloom": []string{
		"heirloom",
		"heirloom-devtools",
		"heirloom-doctools",
		"heirloom-pkgtools",
		"heirloom-sh",
		"mailx",
	},
	"heroines": []string{
		"audioloopback",
		"bcast",
		"cinelerra",
		"cinelerra-6-x",
		"cinelerra-beta",
		"cpuusage",
		"firehose",
		"hvirtual",
		"hvirtual-beta",
		"libmpeg",
		"libmpeg3",
		"mix",
		"mix2000",
		"mix2005",
		"quicktime",
		"quicktime4linux",
		"rh",
		"xmovie",
	},
	"hplip": []string{
		"hplip",
	},
	"hte": []string{
		"ht",
	},
	"hylafax": []string{
		"hylafax",
	},
	"infozip": []string{
		"WiZ",
		"maczip",
		"unzip",
		"wiz",
		"zcrypt",
		"zip",
	},
	"iodbc": []string{
		"libiodbc",
	},
	"ipband": []string{
		"ipband",
	},
	"iperf": []string{
		"iperf",
		"jperf",
	},
	"iperf2": []string{
		"iperf",
	},
	"ipsec-tools": []string{
		"ipsec-tools",
		"ipsec-tools-CVS",
	},
	"irrlicht": []string{
		"IrrlichtDemo",
		"irrlicht",
		"irrlicht-SDK",
		"irrlicht-Sources",
		"irrlicht-terrain",
		"irrxml",
	},
	"jack-rack": []string{
		"jack-rack",
	},
	"jackit": []string{
		"jack-audio-connection-kit",
	},
	"jagoclient": []string{
		"jago",
		"jagosrc",
	},
	"junit": []string{
		"junit",
		"junit-dep",
	},
	"kuickshow": []string{
		"kuickshow",
	},
	"lame": []string{
		"lame",
		"py-lame",
	},
	"lazarus": []string{
		"doc-chm-fpc",
		"doc-chm_fpc",
		"fpc",
		"fpc-2.2.2-lazarus",
		"fpc-lazarus",
		"fpc-lazarus-doc-chm",
		"fpc-lazarus-doc-html",
		"fpc-src",
		"fpc_all_amd",
		"fpc_all_i",
		"lazarus",
	},
	"lcms": []string{
		"lcms",
		"lcms2",
	},
	"libass": []string{
		"libass",
	},
	"libavc1394": []string{
		"libavc1394",
	},
	"libclc": []string{
		"libclc",
	},
	"libdv": []string{
		"libdv",
	},
	"libebook": []string{
		"libe-book",
	},
	"libexif": []string{
		"exif",
		"gexif",
		"libexif",
		"libexif-gtk",
	},
	"libircclient": []string{
		"IRCClient.framework",
		"libircclient",
		"libircclient-dochtml",
		"libircclient-win32",
		"libircclient-win32-vc",
	},
	"libjpeg": []string{
		"jpegsr",
		"jpegsrc.v",
	},
	"libjpeg-turbo": []string{
		"libjpeg-turbo",
		"libjpeg-turbo-official",
	},
	"libjson": []string{
		"libJSON",
		"libjson",
	},
	"libmng": []string{
		"JNGsuite",
		"MNGsuite",
		"TNGImage",
		"fixmng-win32-gui",
		"gtk-mngview",
		"kolMNG",
		"libmng",
		"lm",
		"lmng",
		"mng-win",
		"mngview-linux",
		"mngview-win",
		"sdl-mngplay",
	},
	"liboauth": []string{
		"liboauth",
	},
	"libpng": []string{
		"libpng",
		"lp",
		"lpng",
		"zlib",
	},
	"libquicktime": []string{
		"libquicktime",
	},
	"libraw1394": []string{
		"libiec61883",
		"libraw1394",
	},
	"libseccomp": []string{
		"libseccomp",
	},
	"libsieve": []string{
		"libsieve",
	},
	"libtirpc": []string{
		"libtirpc",
	},
	"libtorrent": []string{
		"libtorrent",
		"libtorrent-rasterbar",
	},
	"libusb": []string{
		"libusb",
		"libusb-compat",
	},
	"libvncserver": []string{
		"LibVNCServer",
		"libvncserver",
		"x11vnc",
	},
	"libwpd": []string{
		"libodfgen",
		"librevenge",
		"libwpd",
		"libwpd-bindings",
		"libwpd-devel",
		"libwpd-tools",
		"wpd2sxw",
		"writerperfect",
	},
	"linux-igd": []string{
		"gateway",
		"linuxigd",
	},
	"linux-ntfs": []string{
		"ldmdoc",
		"ldmdocpdf",
		"linux-ldm",
		"ntfsdoc",
		"ntfsprogs",
	},
	"linux-udf": []string{
		"udf",
		"udftools",
	},
	"linux-usb": []string{
		"USB-guide",
		"USBMon",
		"comtest",
		"speedbundle",
		"speedtouch",
		"usb_flood-V",
		"usbtest_fw",
	},
	"linuxwacom": []string{
		"input-wacom",
		"libwacom",
		"linuxwacom",
		"wdaemon",
		"xf86-input-wacom",
	},
	"lives": []string{
		"lives",
	},
	"log4c": []string{
		"log",
		"log4c",
	},
	"lprng": []string{
		"LPRng",
		"LPRngTool",
		"ifhp",
		"lprng",
	},
	"mad": []string{
		"libid3tag",
		"libmad",
		"madplay",
	},
	"madwifi": []string{
		"madwifi",
	},
	"mantisbt": []string{
		"InstantMantis",
		"mantis",
		"mantisbt",
	},
	"marathon": []string{
		"AlephOne",
		"AlephOne-OSX",
		"AlephOne-OSXNIBs",
		"AlephOne-Windows",
		"AlephOne-m",
		"Marathon",
		"MarathonInfinity",
		"aleph",
		"alephone",
		"marathon",
	},
	"meanwhile": []string{
		"gaim-meanwhile",
		"gaim-meanwhile-win32",
		"meanwhile",
		"meanwhile-gaim",
	},
	"mediatomb": []string{
		"mediatomb",
		"mediatomb-static",
	},
	"mednafen": []string{
		"mednafen",
		"mednafen-server",
	},
	"mesa3d": []string{
		"MesaDemos",
		"MesaGLUT",
		"MesaLib",
		"MesaWinBinaries",
	},
	"meshlab": []string{
		"MeshLab",
		"MeshLabSrc_AllInc_v",
		"MeshLabSrc_v",
		"MeshLab_v",
		"Meshlab",
		"Meshlab-v",
		"Meshlab_v",
		"meshlab_v",
	},
	"mhash": []string{
		"mhash",
	},
	"ming": []string{
		"ming",
		"ming-ch",
		"ming-ch-win",
		"ming-fonts",
		"ming-java",
		"ming-perl",
		"ming-php",
		"ming-py",
		"ming-rb",
		"ming-tcl",
		"ttf2fft",
	},
	"mjpeg": []string{
		"MA-Zoran",
		"ac",
		"doc-package",
		"driver-zoran",
		"gap_vid_enc.v",
		"mjpegtools",
		"mpeg2dec",
		"zoran-driver",
	},
	"mldonkey": []string{
		"mldonkey",
		"mldonkey-tools",
	},
	"mlt": []string{
		"mlt",
		"mlt++",
	},
	"mpg123": []string{
		"mpg123",
	},
	"ms-sys": []string{
		"ms-sys",
	},
	"multivalent": []string{
		"DVI",
		"Multivalent",
	},
	"nas": []string{
		"nas",
	},
	"nc110": []string{
		"nc",
	},
	"net-snmp": []string{
		"net-snmp",
		"ucd-snmp",
	},
	"net-tools": []string{
		"net-tools",
	},
	"netcat": []string{
		"netcat",
	},
	"nethack": []string{
		"NetHack",
		"nethack",
		"nethack-343-tiles",
		"nethack-360-win-x",
		"nh",
		"nh34os",
		"nh34x11os",
		"tiles",
	},
	"nfs": []string{
		"dhiggen_merge",
		"kernel-nfs-dhiggen_merge",
		"nfs-utils",
	},
	"nrappkit": []string{
		"nrappkit",
	},
	"ogl-math": []string{
		"glm",
	},
	"ogre": []string{
		"AndroidDependencies",
		"Dagon_Mac_Dependencies",
		"EihortDependenciesUniversal",
		"Eihort_Universal_Dependencies_beta",
		"MacFrameworks",
		"Ogre3DSExporter",
		"Ogre3DSMaxExport",
		"OgreBlenderExport_v",
		"OgreBlenderExporter",
		"OgreCEGuiMeshViewer",
		"OgreCommandLineTools",
		"OgreCommandLineToolsMac",
		"OgreDemos",
		"OgreDependencies",
		"OgreDependenciesOSX",
		"OgreDependencies_CBMinGW_Eihort",
		"OgreDependencies_CBMinGW_Shoggoth",
		"OgreDependencies_CBMingW",
		"OgreDependencies_CBMingW_STLP",
		"OgreDependencies_Eihort_OSX",
		"OgreDependencies_GLES",
		"OgreDependencies_MSVC",
		"OgreDependencies_MinGW",
		"OgreDependencies_OSX",
		"OgreDependencies_OSX_Eihort",
		"OgreDependencies_OSX_libc++",
		"OgreDependencies_VC",
		"OgreDependencies_VC6",
		"OgreDependencies_VC70",
		"OgreDependencies_VC71",
		"OgreDependencies_VC8",
		"OgreEihortToolsOSX",
		"OgreExporters",
		"OgreLexiExportInstall",
		"OgreLightwaveConverter",
		"OgreMaya6Exporter",
		"OgreMayaExporter",
		"OgreMayaExporter_maya2008",
		"OgreMayaExporter_maya70",
		"OgreMayaExporter_maya80",
		"OgreMayaExporter_maya85",
		"OgreMeshViewer",
		"OgreMilkshapeExporter",
		"OgreParticleEditor",
		"OgreSDKSetup",
		"OgreSDK_Android_v",
		"OgreWin32Dependencies",
		"OgreWings3DExporter",
		"OgreXSIExporter_v",
		"OgreXcodeTemplates",
		"Ogre_PDBs_vc",
		"Ogre_PDBs_vc8_v",
		"Ogre_Xcode",
		"Ogre_Xcode_Templates",
		"ParticleEditor",
		"ogre",
		"ogre-linux-osx-v",
		"ogre-linux_osx-v",
		"ogre-v",
		"ogre-win32-tools",
		"ogre-win32-v",
		"ogreExporter_bin_maya",
		"ogreExporter_bin_maya6.0",
		"ogreExporter_bin_maya6.5",
		"ogreExporter_bin_maya7.0",
		"ogreMayaExporter_maya",
		"ogreMayaExporter_maya6.0",
		"ogreMayaExporter_maya6.5",
		"ogreMayaExporter_maya7.0",
		"ogreMayaExporter_maya8.0",
		"ogreMayaExporter_maya8.5",
		"ogre_src_v",
	},
	"ol2mbox": []string{
		"LibDBX-v",
		"kdbx",
		"libdbx",
		"libpst",
	},
	"omxil": []string{
		"libomxalsa",
		"libomxaudiotemplates",
		"libomxcamera",
		"libomxfbdevsink",
		"libomxffmpegdist",
		"libomxffmpegnonfree",
		"libomxil-B",
		"libomxil-bellagio",
		"libomxjpeg",
		"libomxmad",
		"libomxvideosrc",
		"libomxvorbis",
		"libomxxvideo",
		"omxBellagioExamples",
		"omxil-bellagio",
	},
	"open1x": []string{
		"XSupplicant",
		"Xsupplicant",
		"xsupplicant",
		"xsupplicant-src",
		"xsupplicant-ui",
	},
	"opencvlibrary": []string{
		"OpenCV",
		"chopencv",
		"ippicv_linux",
		"ippicv_macosx",
		"ippicv_windows",
		"opencv",
	},
	"opende": []string{
		"Ode.NET",
		"ode",
		"ode-src",
		"ode-win32",
	},
	"openil": []string{
		"DevIL",
		"DevIL-Docs",
		"DevIL-EndUser",
		"DevIL-EndUser-Unicode",
		"DevIL-EndUser-x64",
		"DevIL-EndUser-x64-Unicode",
		"DevIL-EndUser-x86",
		"DevIL-EndUser-x86-Unicode",
		"DevIL-Manual",
		"DevIL-SDK",
		"DevIL-SDK-x64",
		"DevIL-SDK-x86",
		"DevIL-Windows-SDK",
		"Devil",
		"Devil-SDK",
		"LibCompiled-vc",
		"devil",
	},
	"openjade": []string{
		"OpenSP",
		"openjade",
	},
	"openjpeg.mirror": []string{
		"openjp3d_v",
		"openjpeg",
		"openjpeg_v",
		"openjpip_v",
	},
	"opennx": []string{
		"nxclient",
		"opennx",
	},
	"openobex": []string{
		"ircp",
		"obexfs",
		"obexftp",
		"openobex",
		"openobex-apps",
	},
	"oscaf": []string{
		"shared-desktop-ontologies",
	},
	"p7zip": []string{
		"J7zip",
		"java_lzma",
		"p7zip",
	},
	"palomino-sim": []string{
		"palomino_data_misc",
		"palomino_data_models",
		"palomino_data_sounds",
		"palomino_data_terrain",
		"palomino_src",
	},
	"paps": []string{
		"paps",
	},
	"pcmcia-cs": []string{
		"pcmcia-cs",
	},
	"pdfedit": []string{
		"pdfedit",
		"tools-Win32",
	},
	"pdfshuffler": []string{
		"pdfshuffler",
	},
	"pdftohtml": []string{
		"pdftohtml",
	},
	"phpbb": []string{
		"conv_YaBB",
		"conv_convert_bb",
		"conv_convert_thwb",
		"conv_convert_wbb",
		"conv_convert_xmb",
		"conv_convert_yabbse",
		"conv_ikonboard",
		"conv_ipb",
		"conv_ubb",
		"conv_vb",
		"conv_xmb",
		"phpBB",
		"phpBB-3.0.RC7_to",
		"phpBB-3.0.RC8_to",
		"subsilver2",
	},
	"phpdocu": []string{
		"PhpDocumentor",
		"phpDocumentor",
		"phpdoc",
		"phpdocumentor",
	},
	"pidgin": []string{
		"gtk",
		"gtk-dev",
		"gtk-runtime",
		"pidgin",
	},
	"plib": []string{
		"plib",
	},
	"podofo": []string{
		"example_helloworld",
		"podofo",
		"podofobrowser",
	},
	"poopmup": []string{
		"poopmup",
	},
	"poptop": []string{
		"pptpd",
	},
	"pptpclient": []string{
		"pptp",
		"pptp-extras",
		"pptp-linux",
		"pptpconfig",
	},
	"procps-ng": []string{
		"procps-ng",
	},
	"psi": []string{
		"Psi",
		"psi",
		"psi-win",
		"psi-win32",
		"qssl",
		"qt",
		"qtlinguist",
		"sw10-psi",
		"sw10-qca",
		"sw10-qca-tls",
	},
	"psmisc": []string{
		"psmisc",
	},
	"pupnp": []string{
		"libupnp",
		"libupnp-doc",
		"upnpsdk",
	},
	"pyode": []string{
		"PyODE",
		"PyODE-snapshot",
	},
	"pyopengl": []string{
		"OpenGL",
		"OpenGLContext",
		"PyOpenGL",
		"PyOpenGL-Demo",
		"PyOpenGL-VCPP",
		"PyOpenGL-accelerate",
		"gle",
		"py",
	},
	"pyqt": []string{
		"PyKDE",
		"PyQt-gpl",
		"PyQt-mac-gpl",
		"PyQt-win-gpl",
		"PyQt-x11-gpl",
		"PyQt5_gpl",
		"PyQtChart_gpl",
		"PyQtDataVisualization_gpl",
		"PyQtMobility-gpl",
		"PyQtPurchasing_gpl",
		"QScintilla-1.71-gpl",
		"QScintilla-gpl",
		"QScintilla_gpl",
		"sip",
	},
	"pywebsvcs": []string{
		"SOAPpy",
		"ZSI",
	},
	"pyxml": []string{
		"PyXML",
	},
	"qdvdauthor": []string{
		"qdvdauthor",
		"qdvdauthor-templates",
	},
	"qimageblitz": []string{
		"blitz",
		"qimageblitz",
	},
	"qpdf": []string{
		"pcre",
		"qpdf",
		"qpdf_vc",
	},
	"qucs": []string{
		"asco",
		"freehdl",
		"qucs",
		"qucs-doc",
	},
	"quesoglc": []string{
		"quesoglc",
	},
	"rdesktop": []string{
		"rdesktop",
	},
	"recordmydesktop": []string{
		"gtk-recordMyDesktop",
		"gtk-recordmydesktop",
		"qt-recordmydesktop",
		"recordmydesktop",
	},
	"retroshare": []string{
		"RetroLocal",
		"RetroShare",
		"RetroShare-v",
		"RetroShare_src_v",
		"RetroShare_v",
		"Rs-Linguist-Pack-v",
		"retroShare-LGPL-Source-v",
		"retroshare",
		"retroshare-pkg-linux-src-v",
		"v",
	},
	"rhash": []string{
		"librhash",
		"rhash",
		"rhash-bindings",
	},
	"rp-l2tp": []string{
		"rp-l2tp",
	},
	"scintilla": []string{
		"PentacleG",
		"PentacleJ",
		"PentacleN",
		"PentacleW",
		"SciTE",
		"Scintilla.framework_OSX_Universal",
		"TentacleG",
		"TentacleJ",
		"TentacleN",
		"TentacleW",
		"gscite",
		"icons",
		"scintilla",
		"scintillahaiku",
		"scite",
		"sinkworld",
		"tentacle",
		"wscite",
	},
	"scons": []string{
		"scons",
		"scons-local",
		"scons-src",
	},
	"scribus": []string{
		"presentation_templates",
		"scribus",
		"scribus-common-libs",
	},
	"sf-xpaint": []string{
		"GTKstereograph",
		"deskwrite",
		"dvipgm",
		"dvippm",
		"extractimage",
		"geg",
		"gv-xft",
		"libXaw3dXft",
		"libpgf",
		"libxaw3dxft",
		"netwmpager",
		"networktablet+",
		"pdfshuffler",
		"seshat",
		"sphereEversion",
		"xaw95",
		"xdiary-xft",
		"xdu-xft",
		"xfig-xft",
		"xless-xft",
		"xlupe",
		"xpaint",
	},
	"slrn": []string{
		"slrn",
	},
	"smartmontools": []string{
		"smartctl",
		"smartmontools",
	},
	"sonik": []string{
		"sonik",
	},
	"soprano": []string{
		"soprano",
	},
	"sox": []string{
		"sox",
	},
	"spacenav": []string{
		"libspnav",
		"libspnav_java",
		"spacenav_win32",
		"spacenavd",
		"spnavcfg",
	},
	"springrts": []string{
		"spring",
		"taspring",
	},
	"squashfs": []string{
		"squashfs",
	},
	"squirrel-sql": []string{
		"SquirrelsQLMac",
		"jedit",
		"mysql",
		"oracle",
		"osx-squirrel-sql",
		"sessionscript",
		"sqlval",
		"squirrel-sql",
		"squirrel-sql-snapshot",
		"squirrel-sql-src",
		"squirrelsql",
		"squirrelsql-macosx-installer",
		"squirrelsql-other-installer",
		"squirrelsql-plainzip",
		"squirrelsqlfx-snapshot",
		"win9598-squirrel-sql",
	},
	"stereograph": []string{
		"GTKstereograph",
		"gtk_gui",
		"stereograph",
	},
	"strace": []string{
		"strace",
	},
	"strigi": []string{
		"strigi",
		"strigiapplet",
	},
	"subcomposer": []string{
		"subtitlecomposer",
	},
	"sv1": []string{
		"sonic-annotator",
		"sonic-visualiser",
	},
	"swig": []string{
		"Swig",
		"swig",
		"swigwin",
	},
	"synfig": []string{
		"ETL",
		"synfig",
		"synfigstudio",
	},
	"tcl": []string{
		"TclOO",
		"itcl",
		"msys_mingw",
		"newclock",
		"sqlite",
		"tcl",
		"tcl-core",
		"tcltk",
		"tdbc",
		"tdbcmysql",
		"tdbcodbc",
		"tdbcpostgres",
		"tdbcsqlite3",
		"thread",
		"tk",
		"vclibs",
	},
	"tcllib": []string{
		"BWidget",
		"bwidget",
		"tcllib",
		"tklib",
	},
	"tftp-server": []string{
		"OpenTFTPServerMTSourceV",
		"OpenTFTPServerSPSourceV",
		"opentftpmtV",
		"opentftpspV",
	},
	"tintin": []string{
		"tintin",
	},
	"tkgate": []string{
		"tkgate",
	},
	"tkimg": []string{
		"Img",
		"Img-Darwin64",
		"Img-Linux32",
		"Img-Linux64",
		"Img-Source",
		"Img-win32",
		"Img-win64",
		"img",
		"tkimg",
	},
	"tls": []string{
		"tls",
	},
	"trn": []string{
		"trn",
	},
	"trousers": []string{
		"grub-0.97-13-ima",
		"grub-0.97-fc5-tcg",
		"openssl_tpm_engine",
		"testsuite",
		"tpm-tools",
		"tpm_keyring2",
		"trousers",
	},
	"tuxracer": []string{
		"tuxracer",
		"tuxracer-data",
		"tuxracer-win32",
	},
	"unimediaserver": []string{
		"UMS",
		"UMSBuilder",
		"VideoTestingSuite",
	},
	"vamp": []string{
		"match-vamp-plugin",
		"ofa-vamp-plugin",
		"vamp-aubio-plugins",
		"vamp-example-plugins",
		"vamp-libxtract-plugins",
		"vamp-onsetsds-plugin",
		"vamp-plugin-sdk",
		"vamp-plugin-tester",
		"vampy",
	},
	"virtualgl": []string{
		"VirtualGL",
	},
	"vnc-tight": []string{
		"tightvnc",
	},
	"warzone2100": []string{
		"QT",
		"SDL",
		"WMIT_dSYM",
		"WMIT_dSYMef",
		"WMIT_dSYMfa",
		"WarzoneHelp",
		"WarzoneHelp-a",
		"WarzoneHelp-b",
		"WarzoneHelp-c",
		"WarzoneHelp-d",
		"WarzoneHelp-e",
		"WarzoneHelp-fdf",
		"WarzoneModelImporterTool",
		"WarzoneModelImporterToolef",
		"WarzoneModelImporterToolfa",
		"gettext",
		"quesoglc",
		"warzone",
		"warzone2100",
		"warzone2100-master",
		"warzone2100-pre",
		"warzone2100-tags_master",
		"warzone2100-v",
		"wmit-win",
	},
	"webcamstudio": []string{
		"WebcamStudio",
		"WebcamStudioFX",
		"dist_WebcamStudio",
		"vloopback",
		"webcamstudio",
		"webcamstudio-module",
	},
	"wgois": []string{
		"OISv",
		"ois",
		"ois-v",
		"ois_v",
	},
	"wine": []string{
		"wine",
		"wine-dlls",
		"wine-mono",
		"wine-mozilla",
		"wine-prgs",
		"wine-w32api",
		"wine_gecko",
	},
	"winetools": []string{
		"winetools",
		"winetools-freebsd",
		"winetools-linux",
	},
	"worldforge": []string{
		"Atlas-C++",
		"WFUT",
		"cyphesis",
		"ember",
		"ember-dependencies-mingw",
		"ember-media",
		"entityforge",
		"equator",
		"eris",
		"libmodelfile",
		"libwfut",
		"mercator",
		"metaserver",
		"plunger",
		"sage",
		"sear",
		"sear-media",
		"skstream",
		"varconf",
		"wfmath",
	},
	"wvware": []string{
		"UTF8-UCS4-String",
		"libwmf",
		"wv",
		"wv2",
	},
	"wxpython": []string{
		"wxPython-demo",
		"wxPython-docs",
		"wxPython-newdocs",
		"wxPython-src",
		"wxPython2.5-win32-devel",
		"wxPython2.6-win32-devel",
		"wxPython2.7-win32-devel",
		"wxPython2.8-win32-devel",
		"wxPython2.8-win64-devel",
		"wxPython2.9-win32-devel",
		"wxPython2.9-win64-devel",
		"wxPythonDemo",
		"wxPythonDocs",
		"wxPythonNewDocs",
		"wxPythonSrc",
		"wxPythonWIN32-devel",
		"wxwidgets2.8",
		"wxwidgets2.9",
		"wxwidgets3.0",
	},
	"wxwindows": []string{
		"bb-gtk-linux-amd",
		"bb-gtk-linux-x",
		"contrib",
		"dialoged",
		"dialoged-win",
		"dialoged-win32",
		"jpeg",
		"mmedia",
		"ogl",
		"ogl3",
		"patch",
		"stc",
		"tex",
		"tex2rtf-win32",
		"tex2rtf2",
		"tiff",
		"utils",
		"wx",
		"wx-docs",
		"wxAll",
		"wxBASE",
		"wxBase",
		"wxDFB",
		"wxDemos",
		"wxGTK",
		"wxGTK-demos",
		"wxGTK-samples",
		"wxMAC",
		"wxMGL",
		"wxMGL-demos",
		"wxMGL-samples",
		"wxMOTIF",
		"wxMSW",
		"wxMac",
		"wxMotif",
		"wxOS2",
		"wxSamples",
		"wxUniv",
		"wxWidgets",
		"wxWidgets-docs",
		"wxWidgets-docs-chm",
		"wxWidgets-docs-html",
		"wxWidgets-snapshot",
		"wxWindows",
		"wxX11",
		"wxxrc",
	},
	"x3270": []string{
		"c3270",
		"pr3287",
		"s3270",
		"suite3270",
		"tcl3270",
		"wc3270",
		"wpr3287",
		"x026",
		"x3270",
	},
	"xdvi": []string{
		"xdvi",
		"xdvik",
	},
	"xfce": []string{
		"gtk+",
		"gtk-xfce-engine",
		"xfce",
	},
	"xindy": []string{
		"xindy",
		"xindy-kernel",
		"xindy-make-rules",
		"xindy-modules",
		"xindy-rte",
	},
	"xine": []string{
		"gxine",
		"libcdio",
		"vcdimager",
		"xine-lib",
		"xine-plugin",
		"xine-ui",
		"xine-vcdx",
	},
	"xmpppy": []string{
		"irc-transport",
		"xmppd",
		"xmpppy",
		"yahoo-transport",
	},
	"xvidcap": []string{
		"xvidcap",
		"xvidcap-orig",
	},
	"zbar": []string{
		"ZBarAndroidSDK",
		"zbar",
		"zebra",
	},
	"zint": []string{
		"zint",
		"zint_win_snapshot",
	},
	"zziplib": []string{
		"libgz",
		"zziplib",
		"zziplib-dll",
	},
}