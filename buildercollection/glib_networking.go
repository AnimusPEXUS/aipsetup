package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["glib_networking"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_glib_networking(bs)
	}
}

type Builder_glib_networking struct {
	*Builder_std_meson
}

func NewBuilder_glib_networking(bs basictypes.BuildingSiteCtlI) (*Builder_glib_networking, error) {

	self := new(Builder_glib_networking)

	Builder_std_meson, err := NewBuilder_std_meson(bs)
	if err != nil {
		return nil, err
	}

	self.Builder_std_meson = Builder_std_meson

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_glib_networking) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			//			"--with-ca-certificates=/etc/ssl/cert.pem",
		}...,
	)

	return ret, nil
}
