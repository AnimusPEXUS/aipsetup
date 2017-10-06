package buildercollection

import (
	"os"

	"github.com/AnimusPEXUS/aipsetup"
)

type BuilderAutotoolsStd struct {
	site *aipsetup.BuildingSiteCtl
}

// builders are independent of anything so have no moto to return errors
func NewBuilderAutotoolsStd(buildingsite *aipsetup.BuildingSiteCtl) *BuilderAutotoolsStd {
	ret := new(BuilderAutotoolsStd)

	ret.site = buildingsite

	return ret
}

func (self *BuilderAutotoolsStd) DefineActions() []string {
	return []string{"dst_cleanup", "src_cleanup", "bld_cleanup",
		"primary_extract", "patch", "autogen", "configure", "build", "distribute"}
}

func (self *BuilderAutotoolsStd) BuildingSctipt(
	build_targets []string,
) error {

	//a_tools := new(buildingtools.Autotools)

	if IsStringIn("dst_cleanup", build_targets) {
		dst_dir := self.site.GetDIR_DESTDIR()
		os.RemoveAll(dst_dir)
		os.MkdirAll(dst_dir, 0700)
	}

	if IsStringIn("src_cleanup", build_targets) {
		src_dir := self.site.GetDIR_SOURCE()
		os.RemoveAll(src_dir)
		os.MkdirAll(src_dir, 0700)
	}

	if IsStringIn("bld_cleanup", build_targets) {
		bld_dir := self.site.GetDIR_BUILDING()
		os.RemoveAll(bld_dir)
		os.MkdirAll(bld_dir, 0700)
	}

	if IsStringIn("primary_extract", build_targets) {
		//self.site
	}

	return nil

}
