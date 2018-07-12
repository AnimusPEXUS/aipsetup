package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["xfs"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_xfs(bs), nil
	}
}

type Builder_xfs struct {
	*Builder_std
}

func NewBuilder_xfs(bs basictypes.BuildingSiteCtlI) *Builder_xfs {
	self := new(Builder_xfs)

	self.Builder_std = NewBuilder_std(bs)

	self.EditDistributeArgsCB = self.EditDistributeArgs

	self.ForcedAutogen = true

	return self
}

func (self *Builder_xfs) EditDistributeArgs(
	log *logger.Logger,
	ret []string,
) (
	[]string,
	error,
) {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	ret = make([]string, 0)

	ret = append(ret, "install")

	// NOTE: looks like install-lib and install-dev has been deprecated

	{
		add_install_dev := true
		for _, i := range []string{"acl", "attr"} {
			if info.PackageName == i {
				add_install_dev = false
				break
			}
		}
		if add_install_dev {
			ret = append(ret, "install-dev")
		}
	}

	//	{
	//		add_install_lib := true
	//		for _, i := range []string{"acl", "attr", "xfsprogs", "xfsdump", "dmapi"} {
	//			if info.PackageName == i {
	//				add_install_lib = false
	//				break
	//			}
	//		}
	//		if add_install_lib {
	//			ret = append(ret, "install-lib")
	//		}
	//	}

	ret = append(ret, "DESTDIR="+self.bs.GetDIR_DESTDIR())

	return ret, nil
}
