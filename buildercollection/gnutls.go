package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["gnutls"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_gnutls(bs)
	}
}

type Builder_gnutls struct {
	*Builder_std
}

func NewBuilder_gnutls(bs basictypes.BuildingSiteCtlI) (*Builder_gnutls, error) {

	self := new(Builder_gnutls)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_gnutls) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	// 3.6.3 can't configure with guile 2.2

	suffix := "-2.0"

	guile_config, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().
		CalculateInstallPrefixExecutable("guile-config" + suffix)
	if err != nil {
		return nil, err
	}

	guile_snarf, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().
		CalculateInstallPrefixExecutable("guile-snarf" + suffix)
	if err != nil {
		return nil, err
	}

	guile_tools, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().
		CalculateInstallPrefixExecutable("guile-tools" + suffix)
	if err != nil {
		return nil, err
	}

	guile, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().
		CalculateInstallPrefixExecutable("guile" + suffix)
	if err != nil {
		return nil, err
	}

	//	pkg_config, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().
	//		GetPrefixPkgConfig()
	//	if err != nil {
	//		return nil, err
	//	}

	//	guile_cflags, err := pkg_config.CommandOutput("--cflags", "guile-2.2")
	//	if err != nil {
	//		return nil, err
	//	}

	//	guile_libs, err := pkg_config.CommandOutput("--libs", "guile-2.2")
	//	if err != nil {
	//		return nil, err
	//	}

	ret = append(
		ret,
		[]string{
			"GUILE=" + guile,
			"GUILE_CONFIG=" + guile_config,
			"GUILE_TOOLS=" + guile_tools,
			"GUILE_SNARF=" + guile_snarf,
		}...,
	)

	return ret, nil
}
