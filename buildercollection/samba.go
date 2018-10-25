package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["samba"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_samba(bs)
	}
}

type Builder_samba struct {
	*Builder_std
}

func NewBuilder_samba(bs basictypes.BuildingSiteCtlI) (*Builder_samba, error) {

	self := new(Builder_samba)

	//	Builder_std_waf, err := NewBuilder_std_waf(bs)
	//	if err != nil {
	//		return nil, err
	//	}

	//	self.Builder_std_waf = Builder_std_waf

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_samba) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	//	log.Info(fmt.Sprintf("%#v\n", pkginfodb.Index["samba"]))

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return nil, err
	}

	variant, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateMultilibVariant()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"--with-pam",
			//			"--with-pam_smbpass",
			"--enable-fhs",
			"--with-systemd",
			"--sysconfdir=/etc/samba",
			"--with-configdir=/etc/samba",
			"--with-privatedir=/etc/samba/private",
			"--hostcc=" + info.Host + "-gcc -m" + variant,
		}...,
	)

	ret, err = buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{"--enable-shared"},
		[]string{"CC=", "CXX=", "GCC=", "--host", "--build"},
		//, "--docdir"
	)
	if err != nil {
		return nil, err
	}

	//	ret, err := buildingtools.FilterAutotoolsConfigOptions(
	//		ret,
	//		[]string{"--enable-shared"},
	//		[]string{"CC=", "CXX=", "GCC="},
	//		//"--host", "--build", "--docdir"
	//	)
	//	if err != nil {
	//		return nil, err
	//	}

	return ret, nil
}
