package buildercollection

import (
	"errors"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["nspr"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_nspr(bs)
	}
}

type Builder_nspr struct {
	*Builder_std
}

func NewBuilder_nspr(bs basictypes.BuildingSiteCtlI) (*Builder_nspr, error) {

	self := new(Builder_nspr)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.EditConfigureDirCB = self.EditConfigureDir
	self.EditConfigureWorkingDirCB = self.EditConfigureWorkingDir

	return self, nil
}

func (self *Builder_nspr) EditConfigureDir(log *logger.Logger, ret string) (string, error) {
	return path.Join(self.bs.GetDIR_SOURCE(), "nspr"), nil
}

func (self *Builder_nspr) EditConfigureWorkingDir(log *logger.Logger, ret string) (string, error) {
	return self.EditConfigureDir(log, ret)
}

func (self *Builder_nspr) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret = append(
		ret,
		[]string{
			"--with-mozilla",
			"--with-pthreads",
			"--enable-ipv6",
		}...,
	)

	if variant, err := self.bs.GetBuildingSiteValuesCalculator().
		CalculateMultilibVariant(); err != nil {
		return nil, err
	} else {
		switch variant {
		case "32":
			ret = append(ret, "--disable-64bit")
		case "64":
			ret = append(ret, "--enable-64bit")
		default:
			return nil, errors.New("requested multilib variant is not supported")
		}
	}

	return ret, nil
}
