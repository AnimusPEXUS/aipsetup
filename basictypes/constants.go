package basictypes

const (
	AIPSETUP_SYSTEM_CONFIG_FILENAME = "aipsetup5.system.ini"

	LILITH_ROOT_MULTIHOST_DIRNAME         = "multihost"
	LILITH_MULTIHOST_MULTIARCH_DIRNAME    = "multiarch"
	LILITH_MULTIHOST_CROSSBULDERS_DIRNAME = "crossbuilders"

	DIR_TARBALL    = "00.TARBALL"
	DIR_SOURCE     = "01.SOURCE"
	DIR_PATCHES    = "02.PATCHES"
	DIR_BUILDING   = "03.BUILDING"
	DIR_DESTDIR    = "04.DESTDIR"
	DIR_BUILD_LOGS = "05.BUILD_LOGS"
	DIR_LISTS      = "06.LISTS"
	DIR_TEMP       = "07.TEMP"

	MASSBUILDER_INFO_FILENAME = "massbuilder.info"
	MASSBUILDER_TARBALLS_DIR  = "00.tarballs"
	MASSBUILDER_ASPS_DIR      = "01.asps"

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
)

var (
	POSSIBLE_LIBDIR_NAMES = []string{DIRNAME_LIB, DIRNAME_LIB64}

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
