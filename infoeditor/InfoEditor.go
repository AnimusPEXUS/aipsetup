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
	"github.com/AnimusPEXUS/utils/tags"
)

type InfoEditor struct {
}

func (self *InfoEditor) Run() error {

	index_copy := make(map[string]*basictypes.PackageInfo)

	for k, v := range distropkginfodb.Index {
		index_copy[k] = v
	}

	err := self.Edit(index_copy)
	if err != nil {
		return err
	}

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

	for k, _ := range index_copy {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for ii, i := range keys {
		info := index_copy[i]

		// correct_base_name := self.MakeInfoName(i, ii)
		correct_base_filename := self.MakeFileName(i, ii)

		index_t += fmt.Sprintf("\"%s\": %s,\n", i, self.MakeVariableName(i))

		file_path := path.Join(outdir, correct_base_filename)

		// _, err := os.Stat(file_path)
		// if err == nil {
		// 	return fmt.Errorf("file %s already exists, but shouldn't", correct_base_filename)
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

func (self *InfoEditor) ReplaceInvalidFileNameSymbols(name string) string {
	name = strings.Replace(name, "+", "plus", -1)
	name = strings.Replace(name, ".", "_", -1)
	name = strings.Replace(name, ".", "_", -1)
	name = strings.Replace(name, "@", "_", -1)
	return name
}

func (self *InfoEditor) ReplaceInvalidVariableSymbols(name string) string {
	name = self.ReplaceInvalidFileNameSymbols(name)
	name = strings.Replace(name, "-", "_", -1)
	return name
}

func (self *InfoEditor) MakeVariableName(name string) string {
	return fmt.Sprintf(
		"DistroPackageInfo_%s",
		self.ReplaceInvalidVariableSymbols(name),
	)
}

func (self *InfoEditor) MakeFileName(name string, id int) string {
	name = self.ReplaceInvalidFileNameSymbols(name)
	ret := fmt.Sprintf("%s_id%d.go", name, id)
	return ret
}

// func (self *InfoEditor) MakeVariableName(name string) string {
// 	n := self.MakeInfoName(name)
// 	return fmt.Sprintf("DistroPackageInfo_%s", n)
// }
//
// func (self *InfoEditor) MakeInfoName(name string) string {
//
// 	name = strings.Replace(name, "-", "_", -1)
// 	name = strings.Replace(name, "+", "plus", -1)
// 	name = strings.Replace(name, ".", "_", -1)
// 	name = strings.Replace(name, ".", "_", -1)
// 	name = strings.Replace(name, "@", "_", -1)
//
// 	return name
// }
//
// func (self *InfoEditor) MakeFileName(name string, id int) string {
// 	name = self.MakeInfoName(name)
// 	ret := fmt.Sprintf("%s_id%d.go", name, id)
// 	return ret
// }

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

func (self *InfoEditor) Edit(index map[string]*basictypes.PackageInfo) error {

	for _, i := range [](func(map[string]*basictypes.PackageInfo) error){
		//		self.ApplyBuilders,

		self.ApplyGroups,

		self.ApplyGnome,
		self.ApplySFNet,
	} {
		err := i(index)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *InfoEditor) ApplyGroups(index map[string]*basictypes.PackageInfo) error {
	for k1, v1 := range index {
		t := tags.New(v1.Tags)
		t.DeleteGroup("group")
		for k2, v2 := range GROUPS {
			for _, v3 := range v2 {
				if v3 == k1 {
					t.Add("group", k2)
				}
			}
		}
		v1.Tags = t.Values()
	}
	return nil
}

func (self *InfoEditor) ApplyGnome(index map[string]*basictypes.PackageInfo) error {
	for pkgname, info := range index {
		t := tags.New(info.Tags)
		gnome_p_found := false

		for _, i := range GNOME_PROJECTS {
			if i == info.TarballName {
				gnome_p_found = true
				break
			}
		}

		if t.Have("", "gnome_project") || gnome_p_found {
			info.TarballName = pkgname
			info.TarballProvider = "gnome"
			info.TarballProviderArguments = []string{pkgname}
			info.TarballVersionTool = "gnome"
			info.TarballProviderCachePresetName = "gnome"
			info.HomePage = "https://gnome.org/"
			// info.BuilderName = "std"

			// t.Add("group", "gnome")
			t.Add("", "gnome_project")

			info.Tags = t.Values()
		}
	}
	return nil
}

func (self *InfoEditor) ApplySFNet(index map[string]*basictypes.PackageInfo) error {
	for _, v1 := range index {

		for k2, v2 := range SOURCEFORGE_PROJECTS {
			for _, v3 := range v2 {

				if v3 == v1.TarballName {

					t := tags.New(v1.Tags)

					v1.HomePage = "https://sourceforge.net/projects/" + k2
					v1.TarballProvider = "sf"
					v1.TarballProviderArguments = []string{k2}

					t.Add("sf_project", k2)

					v1.Tags = t.Values()
				}
			}
		}
	}
	return nil
}
