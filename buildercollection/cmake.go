package buildercollection

import (
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["cmake"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_cmake(bs)
	}
}

type Builder_cmake struct {
	*Builder_std_cmake
}

func NewBuilder_cmake(bs basictypes.BuildingSiteCtlI) (*Builder_cmake, error) {
	self := new(Builder_cmake)
	if t, err := NewBuilder_std_cmake(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_cmake = t
	}

	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.AfterDistributeCB = self.AfterDistribute

	return self, nil
}

func (self *Builder_cmake) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ncurses_include_dir := path.Join(prefix, "include", "ncursesw")

	ncursrs_opt := "-DCURSES_INCLUDE_PATH=" + ncurses_include_dir

	log.Info("Calculated ncurses include dir is: " + ncurses_include_dir)

	ret = append(ret, ncursrs_opt)

	return ret, nil
}

func (self *Builder_cmake) AfterDistribute(log *logger.Logger, err error) error {
	if err != nil {
		return err
	}

	calc := self.bs.GetBuildingSiteValuesCalculator()

	prefix, err := calc.CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	prefix_doc := path.Join(prefix, "doc")
	prefix_share := path.Join(prefix, "share")
	prefix_share_doc := path.Join(prefix_share, "doc")

	if _, err := os.Stat(prefix_doc); err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	}

	os.MkdirAll(prefix_share, 0700)

	if _, err := os.Stat(prefix_share_doc); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	err = os.Rename(prefix_doc, prefix_share_doc)
	if err != nil {
		return err
	}

	return nil
}
