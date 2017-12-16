package infoeditor

import (
	"fmt"
	"go/format"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/distropkginfodb"
)

type InfoEditor struct {
}

func (self *InfoEditor) Run() error {
	outdir, err := filepath.Abs("distropkginfodb_new")
	if err != nil {
		return err
	}

	os.MkdirAll(outdir, 0700)

	index_t := `
		package distropkginfodb

		// WARNING: Generated using infoeditor.
		//          Edit items, compile and use "info-db code" cmd for regenerating.


		import (
			"github.com/AnimusPEXUS/aipsetup/basictypes"
		)

		var Index = map[string]*basictypes.PackageInfo{
`

	keys := make([]string, 0)

	for k, _ := range distropkginfodb.Index {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for ii, i := range keys {
		info := distropkginfodb.Index[i]

		// correct_base_name := self.MakeInfoName(i, ii)
		correct_base_filename := self.MakeFileName(i, ii)

		index_t += fmt.Sprintf("\"%s\": %s,\n", i, self.MakeVariableName(i))

		file_path := path.Join(outdir, correct_base_filename)

		// _, err := os.Stat(file_path)
		// if err == nil {
		// 	return fmt.Errorf("file %s already exists, but shouldn't", correct_base_filename)
		// }

		// err=self.Edit(i, info)
		// if err != nil {
		// 	return err
		// }

		txt := self.RenderInfoFileText(i, info)

		f, err := os.Create(file_path)
		if err != nil {
			return err
		}

		b, err := format.Source([]byte(txt))
		if err != nil {
			fmt.Println(txt)
			return err
		}

		f.Write(b)

		f.Close()

	}
	index_t += "}\n"

	index_f, err := os.Create(path.Join(outdir, "00001_Index.go"))
	if err != nil {
		return err
	}
	b, err := format.Source([]byte(index_t))
	if err != nil {
		return err
	}
	index_f.Write(b)
	index_f.Close()
	return nil
}

func (self *InfoEditor) MakeVariableName(name string) string {
	n := self.MakeInfoName(name)
	return fmt.Sprintf("DistroPackageInfo_%s", n)
}

func (self *InfoEditor) MakeInfoName(name string) string {

	name = strings.Replace(name, "-", "_", -1)
	name = strings.Replace(name, "+", "plus", -1)
	name = strings.Replace(name, ".", "_", -1)
	name = strings.Replace(name, ".", "_", -1)
	name = strings.Replace(name, "@", "_", -1)

	return name
}

func (self *InfoEditor) MakeFileName(name string, id int) string {
	name = self.MakeInfoName(name)
	ret := fmt.Sprintf("%s_id%d.go", name, id)
	return ret
}

func (self *InfoEditor) Edit(pkgname string, info *basictypes.PackageInfo) error {
	self.ApplyGnome(pkgname, info)
	return nil
}

func (self *InfoEditor) RenderInfoFileText(
	pkgname string,
	info *basictypes.PackageInfo,
) string {
	ret := fmt.Sprintf(
		InfoFileTemplate,
		self.MakeVariableName(pkgname),
		self.AsMultiline(info.Description),
		self.AsSingleline(info.HomePage),
		self.AsSingleline(info.TarballFileNameParser),
		self.AsSingleline(info.TarballName),
		self.AsStringSlice(info.Filters),

		self.AsSingleline(info.BuilderName),

		info.Removable,
		info.Reducible,
		info.NonInstallable,
		info.Deprecated,
		info.PrimaryInstallOnly,

		self.AsStringSlice(info.BuildDeps),
		self.AsStringSlice(info.SODeps),
		self.AsStringSlice(info.RunTimeDeps),

		self.AsStringSlice(info.Tags),

		self.AsSingleline(info.TarballVersionTool),

		self.AsSingleline(info.TarballProvider),
		self.AsStringSlice(info.TarballProviderArguments),
		info.TarballProviderUseCache,
		self.AsSingleline(info.TarballProviderCachePresetName),
		info.TarballProviderVersionSyncDepth,
	)
	return ret
}

func (self *InfoEditor) AsMultiline(in string) string {
	return fmt.Sprintf("`%s`", in)
}

func (self *InfoEditor) AsSingleline(in string) string {
	return fmt.Sprintf("\"%s\"", in)
}

func (self *InfoEditor) AsStringSlice(in []string) string {
	ret := "[]string{\n"
	for _, i := range in {
		ret += fmt.Sprintf("  %s,", self.AsSingleline(i))
	}
	ret += "}"
	return ret
}

func (self *InfoEditor) ApplyGnome(pkgname string, info *basictypes.PackageInfo) {
	if info.Tags.HaveTag("project", "gnome") {
		info.TarballName = pkgname
		info.TarballProvider = "gnome"
		info.TarballProviderArguments = []string{pkgname}
		info.TarballVersionTool = "gnome"
		info.TarballProviderCachePresetName = "gnome"
		info.HomePage = "https://gnome.org/"
		info.BuilderName = "std"
	}
}
