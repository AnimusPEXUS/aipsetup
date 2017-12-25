package infoeditor

type RecordK struct {
	TarballName     string
	WholeTagRegExp  string
	TagPrefixRegExp string
	TagSuffixRegExp string
	TarballFormat   string
}

const GithubDefaultTarballName = "v"
const GithubDefaultWholeTagRegExp = STANDARD_GITHUB_TAG_REGEXP
const GithubDefaultTagPrefixRegExp = "v"
const GithubDefaultTagSuffixRegExp = `^$`
const GithubDefaultTarballFormat = "tar.xz"

const STANDARD_GITHUB_TAG_REGEXP = `` +
	`^` +
	`((?P<prefix>.*?)[\-\_]?)?` +
	`(?P<version>\d+(?P<delim>[\_\-\.])?(\d+(?P=delim)?)*)` +
	`([\-\_]??(?P<suffix>.*?)??)??` +
	`$`

var GITHUB_HOSTED = map[string](map[string]([]*RecordK)){

	"FFTW": map[string]([]*RecordK){

		"fftw3": []*RecordK{

			&RecordK{
				TarballName:     "fftw",
				TagPrefixRegExp: `fftw`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"FFmpeg": map[string]([]*RecordK){

		"FFmpeg": []*RecordK{

			&RecordK{
				TarballName:     "ffmpeg",
				TagPrefixRegExp: `n`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"LMMS": map[string]([]*RecordK){

		"lmms": []*RecordK{

			&RecordK{
				TarballName:     "lmms",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"MidnightCommander": map[string]([]*RecordK){

		"mc": []*RecordK{

			&RecordK{
				TarballName:     "mc",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"OpenImageIO": map[string]([]*RecordK){

		"oiio": []*RecordK{

			&RecordK{
				TarballName:     "openimageio",
				TagPrefixRegExp: `Release`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"SELinuxProject": map[string]([]*RecordK){

		"selinux": []*RecordK{

			&RecordK{
				TarballName:     "checkpolicy",
				TagPrefixRegExp: `checkpolicy`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},

			&RecordK{
				TarballName:     "libselinux",
				TagPrefixRegExp: `libselinux`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},

			&RecordK{
				TarballName:     "libsemanage",
				TagPrefixRegExp: `libsemanage`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},

			&RecordK{
				TarballName:     "libsepol",
				TagPrefixRegExp: `libsepol`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},

			&RecordK{
				TarballName:     "policycoreutils",
				TagPrefixRegExp: `policycoreutils`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},

			&RecordK{
				TarballName:     "sepolgen",
				TagPrefixRegExp: `sepolgen`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"SRombauts": map[string]([]*RecordK){

		"SQLiteCpp": []*RecordK{

			&RecordK{
				TarballName:     "sqlitecpp",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"anholt": map[string]([]*RecordK){

		"libepoxy": []*RecordK{

			&RecordK{
				TarballName:     "libepoxy",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"antirez": map[string]([]*RecordK){

		"hping": []*RecordK{},
	},

	"apple": map[string]([]*RecordK){

		"cups": []*RecordK{

			&RecordK{
				TarballName:     "cups",
				TagPrefixRegExp: `release`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"arvidn": map[string]([]*RecordK){

		"libtorrent": []*RecordK{

			&RecordK{
				TarballName:     "libtorrent-rasterbar",
				TagPrefixRegExp: `^libtorrent$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"audacity": map[string]([]*RecordK){

		"audacity": []*RecordK{

			&RecordK{
				TarballName:     "audacity-minsrc",
				TagPrefixRegExp: `Audacity`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"dosfstools": map[string]([]*RecordK){

		"dosfstools": []*RecordK{

			&RecordK{
				TarballName:     "dosfstools",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"ethereum": map[string]([]*RecordK){

		"mist": []*RecordK{

			&RecordK{
				TarballName:     "etherium",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"google": map[string]([]*RecordK){

		"googletest": []*RecordK{

			&RecordK{
				TarballName:     "googletest",
				TagPrefixRegExp: `release`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},

		"protobuf": []*RecordK{

			&RecordK{
				TarballName:     "protobuf",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"hamcrest": map[string]([]*RecordK){

		"JavaHamcrest": []*RecordK{

			&RecordK{
				TarballName:     "hamcrest-java",
				TagPrefixRegExp: `hamcrest-java`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"ibus": map[string]([]*RecordK){

		"ibus": []*RecordK{

			&RecordK{
				TarballName:     "ibus",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},

		"ibus-anthy": []*RecordK{

			&RecordK{
				TarballName:     "ibus-anthy",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},

		"ibus-cros": []*RecordK{

			&RecordK{
				TarballName:     "ibus-cros",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},

		"ibus-m17n": []*RecordK{

			&RecordK{
				TarballName:     "ibus-m17n",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},

		"ibus-pinyin": []*RecordK{

			&RecordK{
				TarballName:     "ibus-pinyin",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},

		"ibus-qt": []*RecordK{

			&RecordK{
				TarballName:     "ibus-qt",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},

		"ibus-xkb": []*RecordK{

			&RecordK{
				TarballName:     "ibus-xkb",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"jackaudio": map[string]([]*RecordK){

		"jack1": []*RecordK{

			&RecordK{
				TarballName:     "jack",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},

		"jack2": []*RecordK{

			&RecordK{
				TarballName:     "jack",
				TagPrefixRegExp: `^v$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"jpirko": map[string]([]*RecordK){

		"libndp": []*RecordK{

			&RecordK{
				TarballName:     "libndp",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"json-c": map[string]([]*RecordK){

		"json-c": []*RecordK{

			&RecordK{
				TarballName:     "json-c",
				TagPrefixRegExp: `json-c`,
				TagSuffixRegExp: `^\d{8}$`,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"junit-team": map[string]([]*RecordK){

		"junit4": []*RecordK{

			&RecordK{
				TarballName:     "junit",
				TagPrefixRegExp: `r`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},

		"junit5": []*RecordK{

			&RecordK{
				TarballName:     "junit",
				TagPrefixRegExp: `r`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"kripken": map[string]([]*RecordK){

		"emscripten": []*RecordK{

			&RecordK{
				TarballName:     "emscripten",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"libevent": map[string]([]*RecordK){

		"libevent": []*RecordK{

			&RecordK{
				TarballName:     "libevent",
				TagPrefixRegExp: `release`,
				TagSuffixRegExp: `(stable)?`,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"libfuse": map[string]([]*RecordK){

		"libfuse": []*RecordK{

			&RecordK{
				TarballName:     "fuse",
				TagPrefixRegExp: `fuse`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},

		"sshfs": []*RecordK{

			&RecordK{
				TarballName:     "sshfs-fuse",
				TagPrefixRegExp: `sshfs`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"libgit2": map[string]([]*RecordK){

		"libgit2": []*RecordK{

			&RecordK{
				TarballName:     "libgit2",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"libproxy": map[string]([]*RecordK){

		"libproxy": []*RecordK{

			&RecordK{
				TarballName:     "libproxy",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"lloyd": map[string]([]*RecordK){

		"yajl": []*RecordK{

			&RecordK{
				TarballName:     "yajl",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"miniupnp": map[string]([]*RecordK){

		"miniupnp": []*RecordK{

			&RecordK{
				TarballName:     "miniupnpc",
				TagPrefixRegExp: `miniupnpc`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},

			&RecordK{
				TarballName:     "miniupnpd",
				TagPrefixRegExp: `miniupnpd`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"mongodb": map[string]([]*RecordK){

		"mongo": []*RecordK{

			&RecordK{
				TarballName:     "mongodb-src",
				TagPrefixRegExp: `r`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"nemequ": map[string]([]*RecordK){

		"vala-extra-vapis": []*RecordK{},
	},

	"networkupstools": map[string]([]*RecordK){

		"nut": []*RecordK{

			&RecordK{
				TarballName:     "nut",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"open-mpi": map[string]([]*RecordK){

		"ompi-release": []*RecordK{

			&RecordK{
				TarballName:     "openmpi",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"qbittorrent": map[string]([]*RecordK){

		"qBittorrent": []*RecordK{

			&RecordK{
				TarballName:     "qbittorrent",
				TagPrefixRegExp: `^release$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"rust-lang": map[string]([]*RecordK){

		"cargo": []*RecordK{

			&RecordK{
				TarballName:     "cargo",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},

		"rust": []*RecordK{

			&RecordK{
				TarballName:     "rustc",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"rust-lang-nursery": map[string]([]*RecordK){

		"rustfmt": []*RecordK{

			&RecordK{
				TarballName:     "rustfmt",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"sphair": map[string]([]*RecordK){

		"ClanLib": []*RecordK{

			&RecordK{
				TarballName:     "ClanLib",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"sqlcipher": map[string]([]*RecordK){

		"sqlcipher": []*RecordK{

			&RecordK{
				TarballName:     "sqlcipher",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"systemd": map[string]([]*RecordK){

		"systemd": []*RecordK{

			&RecordK{
				TarballName:     "systemd",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"tgraf": map[string]([]*RecordK){

		"libnl": []*RecordK{

			&RecordK{
				TarballName:     "libnl",
				TagPrefixRegExp: `libnl`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"transmission": map[string]([]*RecordK){

		"transmission": []*RecordK{

			&RecordK{
				TarballName:     "transmission",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"traviscross": map[string]([]*RecordK){

		"mtr": []*RecordK{

			&RecordK{
				TarballName:     "mtr",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"uclouvain": map[string]([]*RecordK){

		"openjpeg": []*RecordK{

			&RecordK{
				TarballName:     "openjpeg",
				TagPrefixRegExp: `version.`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"ukoethe": map[string]([]*RecordK){

		"vigra": []*RecordK{

			&RecordK{
				TarballName:     "vigra",
				TagPrefixRegExp: `Version`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"unittest-cpp": map[string]([]*RecordK){

		"unittest-cpp": []*RecordK{

			&RecordK{
				TarballName:     "unittest-cpp",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"vancegroup": map[string]([]*RecordK){

		"freealut": []*RecordK{

			&RecordK{
				TarballName:     "freealut",
				TagPrefixRegExp: `freealut`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"waf-project": map[string]([]*RecordK){

		"waf": []*RecordK{

			&RecordK{
				TarballName:     "waf",
				TagPrefixRegExp: `waf`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"webmproject": map[string]([]*RecordK){

		"libvpx": []*RecordK{

			&RecordK{
				TarballName:     "libvpx",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},

		"libwebm": []*RecordK{

			&RecordK{
				TarballName:     "libwebm",
				TagPrefixRegExp: `libwebm`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},

		"libwebp": []*RecordK{

			&RecordK{
				TarballName:     "libwebp",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"xkbcommon": map[string]([]*RecordK){

		"libxkbcommon": []*RecordK{

			&RecordK{
				TarballName:     "libxkbcommon",
				TagPrefixRegExp: `xkbcommon`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"yaml": map[string]([]*RecordK){

		"libyaml": []*RecordK{

			&RecordK{
				TarballName:     "libyaml",
				TagPrefixRegExp: `^$`,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},

	"zeromq": map[string]([]*RecordK){

		"cppzmq": []*RecordK{

			&RecordK{
				TarballName:     "cppzmq",
				TagPrefixRegExp: GithubDefaultTagPrefixRegExp,
				TagSuffixRegExp: GithubDefaultTagSuffixRegExp,
				WholeTagRegExp:  STANDARD_GITHUB_TAG_REGEXP,
				TarballFormat:   GithubDefaultTarballFormat,
			},
		},
	},
}
