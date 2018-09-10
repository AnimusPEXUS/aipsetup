package buildercollection

// NOTE: libxml2 is the same. butI won't delete this file for now

//import (
//	"path"

//	"github.com/AnimusPEXUS/aipsetup/basictypes"
//	"github.com/AnimusPEXUS/utils/logger"
//)

//func init() {
//	Index["libxslt"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
//		return NewBuilder_libxslt(bs), nil
//	}
//}

//type Builder_libxslt struct {
//	*Builder_std

//	python_name string
//}

//func NewBuilder_libxslt(bs basictypes.BuildingSiteCtlI) *Builder_libxslt {

//	self := new(Builder_libxslt)

//	self.Builder_std = NewBuilder_std(bs)
//	self.EditConfigureArgsCB = self.EditConfigureArgs

//	self.python_name = "python2.7"

//	return self
//}

//func (self *Builder_libxslt) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

//	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

//	python2, err := calc.CalculateInstallPrefixExecutable(self.python_name)
//	if err != nil {
//		return nil, err
//	}

//	install_prefix, err := calc.CalculateInstallPrefix()
//	if err != nil {
//		return nil, err
//	}

//	ret = append(
//		ret,
//		[]string{
//			"--with-python=" + python2,
//			"--with-python-install-dir=" + path.Join(
//				install_prefix,
//				"lib",
//				self.python_name,
//				"site-packages",
//			),
//			"--with-python=" + install_prefix,
//			"PYTHON=" + python2,
//		}...,
//	)

//	return ret, nil
//}
