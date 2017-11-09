package buildercollection

import (
	"errors"
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/gologger"
)

type CrossBuildEnum uint

const (
	NoAction CrossBuildEnum = iota
	Force
	Forbid
)

type BuilderAutotoolsStd struct {

	// # this is for builder_action_autogen() method
	ForcedAutogen                bool
	SeparateBuildDir             bool
	SourceConfigureRelPath       string
	ForcedTarget                 bool
	ApplyHostSpecCompilerOptions bool

	// # None - not used, bool - force value
	ForceCrossbuilder CrossBuildEnum
	ForceCrossbuild   CrossBuildEnum

	site basictypes.BuildingSiteCtlI
}

// builders are independent of anything so have no moto to return errors
func NewBuilderAutotoolsStd(buildingsite basictypes.BuildingSiteCtlI) *BuilderAutotoolsStd {
	ret := new(BuilderAutotoolsStd)

	ret.site = buildingsite

	ret.ForcedAutogen = false

	ret.SeparateBuildDir = false
	ret.SourceConfigureRelPath = "."
	ret.ForcedTarget = false
	ret.ApplyHostSpecCompilerOptions = true

	self.force_crossbuilder = NoAction
	self.force_crossbuild = NoAction

	return ret
}

// func (self *BuilderAutotoolsStd) SetBuildingSite(bs basictypes.BuildingSiteCtlI) {
// 	self.site = bs
// }

func (self *BuilderAutotoolsStd) DefineActions() (
	[]string,
	basictypes.BuilderActions,
) {
	ret0 := []string{
		"dst_cleanup",
		"src_cleanup",
		"bld_cleanup",
		"primary_extract",
		//		"patch",
		//	"autogen",
		"configure",
		"build",
		"distribute",
	}
	ret := make(basictypes.BuilderActions)

	ret["dst_cleanup"] = self.BuilderActionDstCleanup
	ret["src_cleanup"] = self.BuilderActionSrcCleanup
	ret["bld_cleanup"] = self.BuilderActionBldCleanup
	ret["primary_extract"] = self.BuilderActionPrimaryExtract
	//ret["patch"] = self.BuilderActionPatch
	//ret["autogen"] = self.BuilderActionAutogen
	ret["configure"] = self.BuilderActionConfigure
	ret["build"] = self.BuilderActionBuild
	ret["distribute"] = self.BuilderActionDistribute

	return ret0, ret
}

func (self *BuilderAutotoolsStd) BuilderActionDstCleanup(
	log *gologger.Logger,
) error {
	dst_dir := self.site.GetDIR_DESTDIR()
	os.RemoveAll(dst_dir)
	os.MkdirAll(dst_dir, 0700)
	return nil
}

func (self *BuilderAutotoolsStd) BuilderActionSrcCleanup(
	log *gologger.Logger,
) error {
	src_dir := self.site.GetDIR_SOURCE()
	os.RemoveAll(src_dir)
	os.MkdirAll(src_dir, 0700)
	return nil
}
func (self *BuilderAutotoolsStd) BuilderActionBldCleanup(
	log *gologger.Logger,
) error {
	bld_dir := self.site.GetDIR_BUILDING()
	os.RemoveAll(bld_dir)
	os.MkdirAll(bld_dir, 0700)
	return nil
}

func (self *BuilderAutotoolsStd) BuilderActionPrimaryExtract(
	log *gologger.Logger,
) error {
	a_tools := new(buildingtools.Autotools)

	info, err := self.site.ReadInfo()
	if err != nil {
		return err
	}

	if len(info.Sources) == 0 {
		return errors.New("no tarballs supplied. primary tarball is required")
	}
	tarball := info.Sources[0]
	tarball = path.Join(self.site.GetDIR_TARBALL(), tarball)
	err = a_tools.Extract(
		tarball,
		self.site.GetDIR_SOURCE(),
		path.Join(self.site.GetDIR_TEMP(), "primary_tarball"),
		true,
		false,
		"",
		false,
		false,
		log,
	)
	if err != nil {
		return err
	}
	return nil
}

func (self *BuilderAutotoolsStd) BuilderActionPatch(
	log *gologger.Logger,
) error {
	return errors.New("not impl")
}

func (self *BuilderAutotoolsStd) BuilderActionAutogen(
	log *gologger.Logger,
) error {
	return errors.New("not impl")
}

func (self *BuilderAutotoolsStd) BuilderActionConfigure(
	log *gologger.Logger,
) error {
	a_tools := new(buildingtools.Autotools)
	a_tools.Configure()
	return errors.New("not impl")
}

func (self *BuilderAutotoolsStd) BuilderActionBuild(
	log *gologger.Logger,
) error {
	return errors.New("not impl")
}

func (self *BuilderAutotoolsStd) BuilderActionDistribute(
	log *gologger.Logger,
) error {
	return errors.New("not impl")
}
