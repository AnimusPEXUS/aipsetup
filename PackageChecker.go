package aipsetup

import (
	"archive/tar"
	"bytes"
	"errors"
	"io"
	"os"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

type PackageChecker struct {
}

func NewPackageChecker() *PackageChecker {
	self := new(PackageChecker)
	return self
}

func (self *PackageChecker) _MakeReader(
	filename string,
	name string,
) (*tar.Header, *tar.Reader, error) {

	tar_file_obj, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	tar_obj := tar.NewReader(tar_file_obj)

	var head *tar.Header

	for {

		head, err = tar_obj.Next()
		if err != nil {
			break
		}

		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
		}

		if head.Name == name {
			return head, tar_obj, nil
		}
	}

	return nil, nil, errors.New("not found")
}

// return bool, error: bool - ok or no
func (self *PackageChecker) CheckByFilename(filename string) (bool, error) {

	_, err := basictypes.NewASPNameFromString(filename)
	if err != nil {
		return false, err
	}

	tar_file_obj_info, err := os.Stat(filename)
	if err != nil {
		return false, err
	}

	if tar_file_obj_info.IsDir() {
		return false, errors.New("not a file")
	}

	if tar_file_obj_info.Size() == 0 {
		return false, errors.New("wrong contents")
	}

	{
		_, tar_obj, err := self._MakeReader(
			filename,
			basictypes.PACKAGE_INFO_FILENAME_V5,
		)
		if err != nil {
			return false, err
		}

		{
			_ts := new(bytes.Buffer)
			_, err = io.Copy(_ts, tar_obj)
			if err != nil {
				return false, err
			}

			_, err = basictypes.NewBuildingSiteInfoFromByteSlice(_ts.Bytes())
			if err != nil {
				return false, err
			}
		}

	}

	return true, nil
}

func CheckAspPackageByFilename(filename string) (bool, error) {
	c := NewPackageChecker()
	return c.CheckByFilename(filename)
}
