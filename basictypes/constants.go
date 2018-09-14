package basictypes

import (
	"errors"
	"sort"
)

const (
	AIPSETUP_SYSTEM_CONFIG_FILENAME = "aipsetup5.system.ini"

	HORIZON_ROOT_MULTIHOST_DIRNAME         = "multihost"
	HORIZON_MULTIHOST_MULTIARCH_DIRNAME    = "multiarch"
	HORIZON_MULTIHOST_CROSSBULDERS_DIRNAME = "crossbuilders"

	DIR_TARBALL    = "00.TARBALL"
	DIR_SOURCE     = "01.SOURCE"
	DIR_PATCHES    = "02.PATCHES"
	DIR_BUILDING   = "03.BUILDING"
	DIR_DESTDIR    = "04.DESTDIR"
	DIR_BUILD_LOGS = "05.BUILD_LOGS"
	DIR_LISTS      = "06.LISTS"
	DIR_TEMP       = "07.TEMP"

	MASSBUILDER_INFO_FILENAME = "00.massbuilder.info"
	MASSBUILDER_ASPS_DIR      = "01.asps"
	MASSBUILDER_DONE_TARBALLS = "03.done_tarballs"

	PACKAGE_INFO_FILENAME     = "package_info.json"
	PACKAGE_INFO_FILENAME_V5  = "package_info_v5.json"
	PACKAGE_CHECKSUM_FILENAME = "package.sha512"

	I686_PC_LINUX_GNU   = "i686-pc-linux-gnu"
	X86_64_PC_LINUX_GNU = "x86_64-pc-linux-gnu"

	DIRNAME_BIN     = "bin"
	DIRNAME_SBIN    = "sbin"
	DIRNAME_INCLUDE = "include"
	DIRNAME_LIB     = "lib"
	DIRNAME_LIB32   = "lib32"
	DIRNAME_LIBX32  = "libx32"
	DIRNAME_LIB64   = "lib64"
	DIRNAME_SHARE   = "share"
	DIRNAME_MAN     = "man"
	DIRNAME_DOC     = "doc"
	DIRNAME_DOCS    = "docs"
	DIRNAME_VAR     = "var"
	DIRNAME_OPT     = "opt"
	DIRNAME_PROC    = "proc"

	DIRNAME_DAEMONS = "daemons"

	SYS_UID_MAX = 999

	BOOT_IMAGE_BOOT_PARTITION_UUID    = "8c19d732-192f-2e46-bce3-53e8dea186ff"
	BOOT_IMAGE_BOOT_PARTITION_FS_UUID = "2417c10c-1ea5-4107-bd97-86384692c09b"
)

var (
	POSSIBLE_LIBDIR_NAMES = []string{
		DIRNAME_LIB,
		DIRNAME_LIB32,
		DIRNAME_LIBX32,
		DIRNAME_LIB64,
	}

	PATH_CALCULATOR_BIN_DIR_NAMES = []string{DIRNAME_BIN, DIRNAME_SBIN}

	DIR_ALL = []string{
		DIR_TARBALL,
		DIR_SOURCE,
		DIR_PATCHES,
		DIR_BUILDING,
		DIR_DESTDIR,
		DIR_BUILD_LOGS,
		DIR_LISTS,
		DIR_TEMP,
	}

	AIPSETUP_SUPPORTED_HOST_ARCHS = map[string]([]string){
		"x86_64-pc-linux-gnu": []string{
			"i686-pc-linux-gnu",
			"i586-pc-linux-gnu",
			"i486-pc-linux-gnu",
		},
	}

	AIPSETUP_SUPPORTED_HOST_TARGETS = map[string]([]string){
		"x86_64-pc-linux-gnu": []string{
			"i686-pc-linux-gnu",
		},
		"i686-pc-linux-gnu": []string{
			"x86_64-pc-linux-gnu",
		},
	}

	USERS = map[int]string{

		// TODO: this list requires serious cleanup

		// # users for groups

		// # lspecial users 1-9
		1: "nobody",
		2: "nogroup",
		3: "bin",
		4: "ftp",
		5: "mail",
		6: "adm",
		7: "gdm",
		8: "wheel",

		// # terminals 10-19
		10: "pts",
		11: "tty",

		// # devices 20-38
		20: "disk",
		21: "usb",
		22: "flash",
		23: "mouse",
		24: "lp",
		25: "floppy",
		26: "video",
		27: "audio",
		28: "cdrom",
		29: "tape",
		30: "pulse",
		31: "pulse-access",
		32: "usbfs",
		33: "usbdev",
		34: "usbbus",
		35: "usblist",
		36: "alsa",
		37: "render",

		// # daemons 40-99
		39: "colord",

		40: "messagebus",
		41: "sshd",
		42: "haldaemon",
		//	43: "clamav",
		44: "mysql",
		45: "exim",
		46: "postgres",
		47: "httpd",
		48: "cron",
		//	49: "mrim",
		//	50: "icq",
		//	51: "pyvkt",
		//	52: "j2j",
		//	53: "gnunet",
		//	54: "ejabberd",
		55: "cupsd",
		//	56: "bandersnatch",
		//	57: "torrent",
		58: "ssl",
		//	59: "dovecot",
		//	60: "dovenull",
		//	61: "spamassassin",
		//	62: "yacy",
		//	63: "irc",
		//	64: "hub",
		//	65: "cynin",
		//	66: "mailman",
		//	67: "asterisk",
		//	68: "bitcoin",
		//	69: "adch",

		//	70: "dialout",
		71: "kmem",
		72: "polkituser",
		//	73: "nexuiz",
		//	74: "couchdb",
		75: "polkitd",
		76: "kvm",

		90: "mine", // TODO: remember what it is. minetest?

		91: "utmp",
		92: "lock",
		93: "avahi",
		94: "avahi-autoipd",
		95: "netdev",
		//	96: "freenet",
		//	97: "jabberd2",
		//	98: "mongodb",
		99: "aipsetupserv",

		100: "systemd-bus-proxy",
		101: "systemd-network",
		102: "systemd-resolve",
		103: "systemd-timesync",
		104: "systemd-journal",
		105: "systemd-journal-gateway",
		106: "systemd-journal-remote",
		107: "systemd-journal-upload",
		108: "systemd-coredump",

		200: "tor",
		//	201: "shinken",
	}

	HORIZON_GROUPS = []string{
		"cross", "core0", "core1", "sdl", "perlmod", "gl",
		"gtk", "crypt", "llvm", "media", "netfilter", "qt",
		"rust", "sec", "wayland", "web", "xml", "lang",
	}

	HORIZON_CATEGORIES = []string{
		"x", "freedesktop",
		"media_alsa",
	}
)

func IsAipsetuHostSupported(name string) bool {
	for k, _ := range AIPSETUP_SUPPORTED_HOST_ARCHS {
		if k == name {
			return true
		}
	}
	return false
}

func IsAipsetupHostArchSupported(host, hostarch string) bool {

	host_supported := IsAipsetuHostSupported(host)

	if !host_supported {
		return false
	}

	if host == hostarch {
		return true
	}

	for _, i := range AIPSETUP_SUPPORTED_HOST_ARCHS[host] {
		if i == hostarch {
			return true
		}
	}

	return false
}

func IsAipsetupHostTargetSupported(host, target string) bool {

	host_supported := IsAipsetuHostSupported(host)

	if !host_supported {
		return false
	}

	if host == target {
		return true
	}

	for _, i := range AIPSETUP_SUPPORTED_HOST_TARGETS[host] {
		if i == target {
			return true
		}
	}

	return false
}

func UserKeysSortedSlice() []int {
	ret := make([]int, 0)

	for k, _ := range USERS {
		ret = append(ret, k)
	}
	sort.Ints(ret)
	return ret
}

func UserIdByName(name string) (int, error) {
	for k, v := range USERS {
		if v == name {
			return k, nil
		}
	}
	return -1, errors.New("not found")
}

func UserNameById(id int) (string, error) {
	name, ok := USERS[id]
	if !ok {
		return "", errors.New("not found")
	}
	return name, nil
}
