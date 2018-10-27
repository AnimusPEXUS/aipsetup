package buildercollection

import (
	"errors"
	"fmt"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["docbook_sgml3"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_docbook_sgml3(bs)
	}
}

type Builder_docbook_sgml3 struct {
	*Builder_std

	sgml_or_xml string
}

func NewBuilder_docbook_sgml3(bs basictypes.BuildingSiteCtlI) (*Builder_docbook_sgml3, error) {
	self := new(Builder_docbook_sgml3)

	self.sgml_or_xml = "sgml"

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditExtractMoreThanOneExtractedOkCB = func(log *logger.Logger, ret bool) (bool, error) {
		return true, nil
	}

	self.EditExtractUnwrapCB = func(log *logger.Logger, ret bool) (bool, error) {
		return false, nil
	}

	return self, nil
}

func (self *Builder_docbook_sgml3) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	for _, i := range []string{"sgml", "xml", "xsl"} {
		if self.sgml_or_xml == i {
			goto ok
		}
	}

	return nil, errors.New(
		"Builder_docbook_sgml3.sgml_or_xml must be sgml, xml or xsl",
	)

ok:

	for _, i := range []string{"autogen", "configure", "build"} {
		ret = ret.Remove(i)
	}

	err := ret.ReplaceShort("distribute", self.BuilderActionDistribute)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_docbook_sgml3) BuilderActionDistribute(log *logger.Logger) error {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	dst_install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	xml_dir := path.Join(dst_install_prefix, "share", self.sgml_or_xml, "docbook")

	xml_dir_plus_name := ""

	if self.sgml_or_xml == "sgml" || self.sgml_or_xml == "xml" {

		dirname := "dockbook"
		if self.sgml_or_xml == "xml" {
			dirname += "-xml"
		}

		xml_dir_plus_name = path.Join(
			xml_dir,
			fmt.Sprintf(
				"%s-%s",
				dirname,
				info.PackageVersion,
			),
		)

	} else if self.sgml_or_xml == "xsl" {

		xml_dir_plus_name = xml_dir

	}

	err = filetools.CopyTree(
		self.GetBuildingSiteCtl().GetDIR_SOURCE(),
		xml_dir_plus_name,
		false,
		false,
		false,
		true,
		log,
		true,
		true,
		filetools.CopyWithInfo,
	)
	if err != nil {
		return err
	}

	return nil
}
