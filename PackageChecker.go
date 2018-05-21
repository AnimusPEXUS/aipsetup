package aipsetup

import (
	"errors"
	"os"
)

type PackageChecker struct {
}

func NewPackageChecker() *PackageChecker {
	self := new(PackageChecker)
	return self
}

// return bool, error: bool - ok or no
func (self *PackageChecker) CheckByFilename(filename string) (bool, error) {

	tar_file_obj, err := os.Open(filename)
	if err != nil {
		return false, err
	}

	defer tar_file_obj.Close()

	tar_file_obj_info, err := tar_file_obj.Stat()
	if err != nil {
		return false, err
	}

	if tar_file_obj_info.IsDir() {
		return false, errors.New("not a file")
	}

	if tar_file_obj_info.Size() == 0 {
		return false, errors.New("wrong contents")
	}

	//	tar_obj := tar.NewReader(tar_file_obj)

	//	package_name := parsed.String()

	//	var head *tar.Header

	//	for {
	//		var err error

	//		head, err = tar_obj.Next()
	//		if err != nil {
	//			break
	//		}
	//	}

	return true, nil
}

func CheckAspPackageByFilename(filename string) (bool, error) {
	c := NewPackageChecker()
	return c.CheckByFilename(filename)
}
