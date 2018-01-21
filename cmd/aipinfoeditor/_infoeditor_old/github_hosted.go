package infoeditor

type RecordK struct {
	TarballName string
	TagParser   string
	TagName     string
	TagStatus   string
}

const GithubDefaultTagParser = "std"
const GithubDefaultTagName = "v"
const GithubDefaultTagStatus = `^$`

var GITHUB_HOSTED = map[string](map[string]([]*RecordK)){

	"FFTW": map[string]([]*RecordK){

		"fftw3": []*RecordK{

			&RecordK{
				TarballName: "fftw",
				TagName:     `fftw`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"FFmpeg": map[string]([]*RecordK){

		"FFmpeg": []*RecordK{

			&RecordK{
				TarballName: "ffmpeg",
				TagName:     `n`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"LMMS": map[string]([]*RecordK){

		"lmms": []*RecordK{

			&RecordK{
				TarballName: "lmms",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"MidnightCommander": map[string]([]*RecordK){

		"mc": []*RecordK{

			&RecordK{
				TarballName: "mc",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"OpenImageIO": map[string]([]*RecordK){

		"oiio": []*RecordK{

			&RecordK{
				TarballName: "openimageio",
				TagName:     `Release`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"SELinuxProject": map[string]([]*RecordK){

		"selinux": []*RecordK{

			&RecordK{
				TarballName: "checkpolicy",
				TagName:     `checkpolicy`,

				TagParser: GithubDefaultTagParser,
			},

			&RecordK{
				TarballName: "libselinux",
				TagName:     `libselinux`,
				TagStatus:   GithubDefaultTagStatus,
				TagParser:   GithubDefaultTagParser,
			},

			&RecordK{
				TarballName: "libsemanage",
				TagName:     `libsemanage`,

				TagParser: GithubDefaultTagParser,
			},

			&RecordK{
				TarballName: "libsepol",
				TagName:     `libsepol`,

				TagParser: GithubDefaultTagParser,
			},

			&RecordK{
				TarballName: "policycoreutils",
				TagName:     `policycoreutils`,

				TagParser: GithubDefaultTagParser,
			},

			&RecordK{
				TarballName: "sepolgen",
				TagName:     `sepolgen`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"SRombauts": map[string]([]*RecordK){

		"SQLiteCpp": []*RecordK{

			&RecordK{
				TarballName: "sqlitecpp",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"anholt": map[string]([]*RecordK){

		"libepoxy": []*RecordK{

			&RecordK{
				TarballName: "libepoxy",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"antirez": map[string]([]*RecordK){

		"hping": []*RecordK{},
	},

	"apple": map[string]([]*RecordK){

		"cups": []*RecordK{

			&RecordK{
				TarballName: "cups",
				TagName:     `release`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"arvidn": map[string]([]*RecordK){

		"libtorrent": []*RecordK{

			&RecordK{
				TarballName: "libtorrent-rasterbar",
				TagName:     `^libtorrent$`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"audacity": map[string]([]*RecordK){

		"audacity": []*RecordK{

			&RecordK{
				TarballName: "audacity-minsrc",
				TagName:     `Audacity`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"dosfstools": map[string]([]*RecordK){

		"dosfstools": []*RecordK{

			&RecordK{
				TarballName: "dosfstools",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"ethereum": map[string]([]*RecordK){

		"mist": []*RecordK{

			&RecordK{
				TarballName: "etherium",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"google": map[string]([]*RecordK){

		"googletest": []*RecordK{

			&RecordK{
				TarballName: "googletest",
				TagName:     `release`,

				TagParser: GithubDefaultTagParser,
			},
		},

		"protobuf": []*RecordK{

			&RecordK{
				TarballName: "protobuf",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"hamcrest": map[string]([]*RecordK){

		"JavaHamcrest": []*RecordK{

			&RecordK{
				TarballName: "hamcrest-java",
				TagName:     `hamcrest-java`,
				TagParser:   GithubDefaultTagParser,
			},
		},
	},

	"ibus": map[string]([]*RecordK){

		"ibus": []*RecordK{

			&RecordK{
				TarballName: "ibus",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},

		"ibus-anthy": []*RecordK{

			&RecordK{
				TarballName: "ibus-anthy",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},

		"ibus-cros": []*RecordK{

			&RecordK{
				TarballName: "ibus-cros",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},

		"ibus-m17n": []*RecordK{

			&RecordK{
				TarballName: "ibus-m17n",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},

		"ibus-pinyin": []*RecordK{

			&RecordK{
				TarballName: "ibus-pinyin",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},

		"ibus-qt": []*RecordK{

			&RecordK{
				TarballName: "ibus-qt",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},

		"ibus-xkb": []*RecordK{

			&RecordK{
				TarballName: "ibus-xkb",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"jackaudio": map[string]([]*RecordK){

		"jack1": []*RecordK{

			&RecordK{
				TarballName: "jack",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},

		"jack2": []*RecordK{

			&RecordK{
				TarballName: "jack",
				TagName:     `^v$`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"jpirko": map[string]([]*RecordK){

		"libndp": []*RecordK{

			&RecordK{
				TarballName: "libndp",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"json-c": map[string]([]*RecordK){

		"json-c": []*RecordK{

			&RecordK{
				TarballName: "json-c",
				TagName:     `json-c`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"junit-team": map[string]([]*RecordK){

		"junit4": []*RecordK{

			&RecordK{
				TarballName: "junit",
				TagName:     `r`,

				TagParser: GithubDefaultTagParser,
			},
		},

		"junit5": []*RecordK{

			&RecordK{
				TarballName: "junit",
				TagName:     `r`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"kripken": map[string]([]*RecordK){

		"emscripten": []*RecordK{

			&RecordK{
				TarballName: "emscripten",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"libevent": map[string]([]*RecordK){

		"libevent": []*RecordK{

			&RecordK{
				TarballName: "libevent",
				TagName:     `release`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"libfuse": map[string]([]*RecordK){

		"libfuse": []*RecordK{

			&RecordK{
				TarballName: "fuse",
				TagName:     `fuse`,

				TagParser: GithubDefaultTagParser,
			},
		},

		"sshfs": []*RecordK{

			&RecordK{
				TarballName: "sshfs-fuse",
				TagName:     `sshfs`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"libgit2": map[string]([]*RecordK){

		"libgit2": []*RecordK{

			&RecordK{
				TarballName: "libgit2",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"libproxy": map[string]([]*RecordK){

		"libproxy": []*RecordK{

			&RecordK{
				TarballName: "libproxy",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"lloyd": map[string]([]*RecordK){

		"yajl": []*RecordK{

			&RecordK{
				TarballName: "yajl",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"miniupnp": map[string]([]*RecordK){

		"miniupnp": []*RecordK{

			&RecordK{
				TarballName: "miniupnpc",
				TagName:     `miniupnpc`,

				TagParser: GithubDefaultTagParser,
			},

			&RecordK{
				TarballName: "miniupnpd",
				TagName:     `miniupnpd`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"mongodb": map[string]([]*RecordK){

		"mongo": []*RecordK{

			&RecordK{
				TarballName: "mongodb-src",
				TagName:     `r`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"nemequ": map[string]([]*RecordK){

		"vala-extra-vapis": []*RecordK{},
	},

	"networkupstools": map[string]([]*RecordK){

		"nut": []*RecordK{

			&RecordK{
				TarballName: "nut",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"open-mpi": map[string]([]*RecordK){

		"ompi-release": []*RecordK{

			&RecordK{
				TarballName: "openmpi",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"qbittorrent": map[string]([]*RecordK){

		"qBittorrent": []*RecordK{

			&RecordK{
				TarballName: "qbittorrent",
				TagName:     `^release$`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"rust-lang": map[string]([]*RecordK){

		"cargo": []*RecordK{

			&RecordK{
				TarballName: "cargo",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},

		"rust": []*RecordK{

			&RecordK{
				TarballName: "rustc",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"rust-lang-nursery": map[string]([]*RecordK){

		"rustfmt": []*RecordK{

			&RecordK{
				TarballName: "rustfmt",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"sphair": map[string]([]*RecordK){

		"ClanLib": []*RecordK{

			&RecordK{
				TarballName: "ClanLib",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"sqlcipher": map[string]([]*RecordK){

		"sqlcipher": []*RecordK{

			&RecordK{
				TarballName: "sqlcipher",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"systemd": map[string]([]*RecordK){

		"systemd": []*RecordK{

			&RecordK{
				TarballName: "systemd",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"tgraf": map[string]([]*RecordK){

		"libnl": []*RecordK{

			&RecordK{
				TarballName: "libnl",
				TagName:     `libnl`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"transmission": map[string]([]*RecordK){

		"transmission": []*RecordK{

			&RecordK{
				TarballName: "transmission",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"traviscross": map[string]([]*RecordK){

		"mtr": []*RecordK{

			&RecordK{
				TarballName: "mtr",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"uclouvain": map[string]([]*RecordK){

		"openjpeg": []*RecordK{

			&RecordK{
				TarballName: "openjpeg",
				TagName:     `version.`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"ukoethe": map[string]([]*RecordK){

		"vigra": []*RecordK{

			&RecordK{
				TarballName: "vigra",
				TagName:     `Version`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"unittest-cpp": map[string]([]*RecordK){

		"unittest-cpp": []*RecordK{

			&RecordK{
				TarballName: "unittest-cpp",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"vancegroup": map[string]([]*RecordK){

		"freealut": []*RecordK{

			&RecordK{
				TarballName: "freealut",
				TagName:     `freealut`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"waf-project": map[string]([]*RecordK){

		"waf": []*RecordK{

			&RecordK{
				TarballName: "waf",
				TagName:     `waf`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"webmproject": map[string]([]*RecordK){

		"libvpx": []*RecordK{

			&RecordK{
				TarballName: "libvpx",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},

		"libwebm": []*RecordK{

			&RecordK{
				TarballName: "libwebm",
				TagName:     `libwebm`,

				TagParser: GithubDefaultTagParser,
			},
		},

		"libwebp": []*RecordK{

			&RecordK{
				TarballName: "libwebp",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"xkbcommon": map[string]([]*RecordK){

		"libxkbcommon": []*RecordK{

			&RecordK{
				TarballName: "libxkbcommon",
				TagName:     `xkbcommon`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"yaml": map[string]([]*RecordK){

		"libyaml": []*RecordK{

			&RecordK{
				TarballName: "libyaml",
				TagName:     `^$`,

				TagParser: GithubDefaultTagParser,
			},
		},
	},

	"zeromq": map[string]([]*RecordK){

		"cppzmq": []*RecordK{

			&RecordK{
				TarballName: "cppzmq",
				TagName:     GithubDefaultTagName,

				TagParser: GithubDefaultTagParser,
			},
		},
	},
}
