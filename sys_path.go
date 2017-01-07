package aipsetup

import (
	"path/filepath"
	"strings"
)

func IsRootPath(path string) bool {
	abs := filepath.Abs(path)
	res:=augfilepath.Split(path)
	return len(res) == 2 && res[0] == "" && res[1] == ""
}

func IsUsrPath(path string) bool {
}

func IsPathAHostRoot(path string) bool {

	splitted := filepath.SplitList(dir_i)

	if len(splitted) != 2 {
		return false
	}

	if splitted[0] != MULTIHOST_DIR_NAME {
		return false
	}

	return true

}

func IsPathAnArchDir(path string) bool {

	splitted := filepath.SplitList(dir_i)

	if len(splitted) != 4 {
		return false
	}

	if splitted[0] != MULTIHOST_DIR_NAME {
		return false
	}

	if splitted[2] != MULTIARCH_DIR_NAME {
		return false
	}

	return true

}

func NameIsLibDirName(base_dir_name string) bool {

	if !strings.HasPrefix(base_dir_name, "lib") {
		return false
	}

	return true

}

func IsALibDirPath(path string) bool {
	ret := ((IsPathAHostRoot(path) || IsPathAnArchDir(path)) &&
		NameIsLibDirName(filepath.Base(path)))
	return ret
}
