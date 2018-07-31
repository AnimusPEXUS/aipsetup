package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/cmd/aipinfoeditor/ui"
	"github.com/AnimusPEXUS/utils/tags"
	"github.com/AnimusPEXUS/utils/textlist"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
)

var TableStructure = [][3]interface{}{
	[3]interface{}{0, glib.TYPE_STRING, "Name"},
	[3]interface{}{1, glib.TYPE_STRING, "Description"},
	[3]interface{}{2, glib.TYPE_STRING, "Homepage"},
	[3]interface{}{3, glib.TYPE_STRING, "Builder Name"},
	[3]interface{}{4, glib.TYPE_BOOLEAN, "Removable"},
	[3]interface{}{5, glib.TYPE_BOOLEAN, "Reducible"},
	[3]interface{}{6, glib.TYPE_BOOLEAN, "Auto Reduce"},
	[3]interface{}{7, glib.TYPE_BOOLEAN, "Dont Preserve Shared Objects"},
	[3]interface{}{8, glib.TYPE_BOOLEAN, "Non Buildable"},
	[3]interface{}{9, glib.TYPE_BOOLEAN, "Non Installable"},
	[3]interface{}{10, glib.TYPE_BOOLEAN, "Deprecated"},
	[3]interface{}{11, glib.TYPE_BOOLEAN, "PrimaryOnly"},
	[3]interface{}{12, glib.TYPE_STRING, "Package Building Dependencies"},
	[3]interface{}{13, glib.TYPE_STRING, "Building Dependencies"},
	[3]interface{}{14, glib.TYPE_STRING, "Shared Object Dependencies"},
	[3]interface{}{15, glib.TYPE_STRING, "Runtime Dependencies"},
	[3]interface{}{16, glib.TYPE_STRING, "Tags"},
	[3]interface{}{17, glib.TYPE_STRING, "Category"},
	[3]interface{}{18, glib.TYPE_STRING, "Groups"},
	[3]interface{}{19, glib.TYPE_STRING, "Tarball Parser"},
	[3]interface{}{20, glib.TYPE_STRING, "Tarball Name"},
	[3]interface{}{21, glib.TYPE_STRING, "Tarball Filters"},
	[3]interface{}{22, glib.TYPE_STRING, "Tarball Provider"},
	[3]interface{}{23, glib.TYPE_STRING, "Provider Arguments"},
	[3]interface{}{24, glib.TYPE_STRING, "Tarball Stability Classifier"},
	[3]interface{}{25, glib.TYPE_STRING, "Tarball Version Comparator"},
	[3]interface{}{26, glib.TYPE_INT, "Tarball Sync Depth"},
	[3]interface{}{27, glib.TYPE_BOOLEAN, "Download Patches"},
	[3]interface{}{28, glib.TYPE_STRING, "Patches Downloading Script Text"},
}

const CAT_EDITOR_CAT_COLUMN = 17
const GROUP_EDITOR_GROUPS_COLUMN = 18

type UIMainWindow struct {
	window     *gtk.Window
	info_table *gtk.TreeView

	info_table_store *gtk.ListStore

	mb        *gtk.MenuBar
	mi_reload *gtk.MenuItem
	mi_save   *gtk.MenuItem
	mi_quit   *gtk.MenuItem

	nb *gtk.Notebook
	pb *gtk.ProgressBar

	cat_editor   *CatEditor
	group_editor *GroupEditor
}

func UIMainWindowNew() (*UIMainWindow, error) {

	self := new(UIMainWindow)

	builder, err := gtk.BuilderNew()
	if err != nil {
		panic(err.Error())
	}

	data := ui.MustAsset("ui/main.glade")

	err = builder.AddFromString(string(data))
	if err != nil {
		panic(err.Error())
	}

	if t, err := CatEditorNew(self, builder); err != nil {
		return nil, err
	} else {
		self.cat_editor = t
	}

	if t, err := GroupEditorNew(self, builder); err != nil {
		return nil, err
	} else {
		self.group_editor = t
	}

	if t, err := builder.GetObject("root"); err != nil {
		return nil, err
	} else {
		self.window = t.(*gtk.Window)
	}

	self.window.Maximize()

	if t, err := builder.GetObject("mb"); err != nil {
		return nil, err
	} else {
		self.mb = t.(*gtk.MenuBar)
	}

	if t, err := builder.GetObject("nb"); err != nil {
		return nil, err
	} else {
		self.nb = t.(*gtk.Notebook)
	}

	table_structure_lst := make([]glib.Type, 0)
	for _, i := range TableStructure {
		table_structure_lst = append(table_structure_lst, i[1].(glib.Type))
	}

	if t, err := builder.GetObject("info_table"); err != nil {
		return nil, err
	} else {
		self.info_table = t.(*gtk.TreeView)

		for ii, i := range TableStructure {
			switch i[1] {
			case glib.TYPE_STRING:
				{
					r, err := gtk.CellRendererTextNew()
					if err != nil {
						return nil, err
					}

					if i[0].(int) != 0 {
						r.SetProperty("editable", true)
					}

					c, err := gtk.TreeViewColumnNewWithAttribute(
						i[2].(string), r, "text", ii,
					)
					if err != nil {
						return nil, err
					}

					c.SetResizable(true)
					c.SetClickable(true)
					c.SetSortColumnID(i[0].(int))
					c.SetFixedWidth(100)

					switch i[0].(int) {
					case 1:
						r.SetProperty("wrap-mode", pango.WRAP_WORD)
						r.SetProperty("wrap-width", 0)
						c.SetFixedWidth(200)
					case 2:
						c.SetFixedWidth(250)
					case 15:
						fallthrough
					case 16:
						fallthrough
					case 17:
						c.SetFixedWidth(150)
					case 19:
						c.SetFixedWidth(150)
					case 20:
						c.SetFixedWidth(200)
					case 22:
						c.SetFixedWidth(400)
					case 27:
						c.SetFixedWidth(500)
					}

					self.info_table.AppendColumn(c)

					r.Connect(
						"edited",
						func(
							cell_renderer *gtk.CellRendererText,
							path string,
							new_text string,
							udata int,
						) {

							ite, err := self.info_table_store.GetIterFromString(path)
							if err != nil {
								panic(err)
							}

							err = self.info_table_store.Set(
								ite,
								[]int{udata},
								[]interface{}{strings.Replace(new_text, "\\n", "\n", -1)},
							)
							if err != nil {
								panic(err)
							}

						},
						i[0].(int),
					)

				}
			case glib.TYPE_BOOLEAN:
				{
					r, err := gtk.CellRendererToggleNew()
					if err != nil {
						return nil, err
					}

					r.SetProperty("activatable", true)

					c, err := gtk.TreeViewColumnNewWithAttribute(
						i[2].(string), r, "active", ii,
					)
					if err != nil {
						return nil, err
					}

					c.SetResizable(true)
					c.SetClickable(true)
					c.SetFixedWidth(70)
					//					c.SetMaxWidth(70)
					c.SetSortColumnID(i[0].(int))

					self.info_table.AppendColumn(c)

					r.Connect(
						"toggled",
						func(
							cell_renderer *gtk.CellRendererToggle,
							path string,
							udata int,
						) {

							ite, err := self.info_table_store.GetIterFromString(path)
							if err != nil {
								panic(err)
							}

							val, err := self.info_table_store.GetValue(ite, udata)
							if err != nil {
								panic(err)
							}

							valbt, err := val.GoValue()
							if err != nil {
								panic(err)
							}

							valb := valbt.(bool)

							err = self.info_table_store.Set(
								ite,
								[]int{udata},
								[]interface{}{!valb},
							)
							if err != nil {
								panic(err)
							}

						},
						i[0].(int),
					)
				}
			case glib.TYPE_INT:
				{
					r, err := gtk.CellRendererTextNew()
					if err != nil {
						return nil, err
					}

					r.SetProperty("editable", true)

					c, err := gtk.TreeViewColumnNewWithAttribute(
						i[2].(string), r, "text", ii,
					)
					if err != nil {
						return nil, err
					}

					c.SetResizable(true)
					c.SetClickable(true)
					c.SetFixedWidth(50)
					c.SetSortColumnID(i[0].(int))

					self.info_table.AppendColumn(c)

					r.Connect(
						"edited",
						func(
							cell_renderer *gtk.CellRendererText,
							path string,
							new_text string,
							udata int,
						) {

							inty, err := strconv.Atoi(new_text)
							if err != nil {
								return
							}

							ite, err := self.info_table_store.GetIterFromString(path)
							if err != nil {
								panic(err)
							}

							err = self.info_table_store.Set(
								ite,
								[]int{udata},
								[]interface{}{inty},
							)
							if err != nil {
								panic(err)
							}

						},
						i[0].(int),
					)
				}
			}
		}
	}

	if t, err := builder.GetObject("mi_quit"); err != nil {
		return nil, err
	} else {
		self.mi_quit = t.(*gtk.MenuItem)
	}

	if t, err := builder.GetObject("mi_save"); err != nil {
		return nil, err
	} else {
		self.mi_save = t.(*gtk.MenuItem)
	}

	if t, err := builder.GetObject("mi_reload"); err != nil {
		return nil, err
	} else {
		self.mi_reload = t.(*gtk.MenuItem)
	}

	if t, err := builder.GetObject("pb"); err != nil {
		return nil, err
	} else {
		self.pb = t.(*gtk.ProgressBar)
	}

	if t, err := gtk.ListStoreNew(
		table_structure_lst...,
	); err != nil {
		return nil, err
	} else {
		self.info_table_store = t
	}

	self.info_table.SetModel(self.info_table_store)

	self.window.Connect(
		"destroy",
		func() {
			gtk.MainQuit()
		},
	)

	self.mi_quit.Connect(
		"activate",
		func() {
			gtk.MainQuit()
		},
	)

	self.mi_reload.Connect(
		"activate",
		func() {
			go func() {
				err := self.LoadTable()
				if err != nil {
					fmt.Println("error")
					fmt.Println(err)
				}
			}()
		},
	)

	self.mi_save.Connect(
		"activate",
		func() {
			go func() {
				self.SaveTable()
			}()
		},
	)

	glib.IdleAdd(
		func() {
			self.mi_reload.Activate()
		},
	)

	return self, nil
}

func (self *UIMainWindow) Show() {
	self.window.ShowAll()
}

func (self *UIMainWindow) LoadTable() error {
	glib.IdleAdd(
		func() bool {
			self.mb.SetSensitive(false)
			self.nb.SetSensitive(false)
			// self.nb.SetVisible(false)
			self.pb.Show()
			return false
		},
	)
	defer glib.IdleAdd(
		func() bool {
			self.pb.Hide()
			self.mb.SetSensitive(true)
			// self.nb.SetVisible(true)
			self.nb.SetSensitive(true)
			return false
		},
	)

	d, err := InfoJSONDir()
	if err != nil {
		return err
	}

	stats, err := ioutil.ReadDir(d)
	if err != nil {
		return err
	}

	// fmt.Println("len stats", len(stats))

	pi := make(map[string]*basictypes.PackageInfo)
	for ii, i := range stats {
		if strings.HasSuffix(i.Name(), ".json") && !i.IsDir() {

			tb, err := ioutil.ReadFile(path.Join(d, i.Name()))
			if err != nil {
				return errors.New("reading " + i.Name() + " " + err.Error())
			}

			t := new(basictypes.PackageInfo)
			err = json.Unmarshal(tb, t)
			if err != nil {
				return errors.New("unmarshal " + i.Name() + " " + err.Error())
			}

			name := i.Name()[0 : len(i.Name())-5]
			pi[name] = t

		}
		c := make(chan bool)
		glib.IdleAdd(
			func() bool {
				f := 0.5 / float64(len(stats)) * float64(ii+1)
				self.pb.SetFraction(f)
				c <- true
				return false
			},
		)
		<-c
	}

	// fmt.Println("len pi", len(pi))

	{
		c := make(chan bool)
		glib.IdleAdd(
			func() bool {
				self.info_table_store.Clear()
				c <- true
				return false
			},
		)
		<-c

	}
	{
		c := make(chan bool)
		glib.IdleAdd(
			func() bool {

				i := 0
				pi_keys := make([]string, 0)
				for k, _ := range pi {
					pi_keys = append(pi_keys, k)
				}
				sort.Strings(pi_keys)
				for _, k := range pi_keys {
					v := pi[k]
					iter := self.info_table_store.Append()
					err := self.info_table_store.Set(
						iter,
						[]int{
							0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
							10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
							20, 21, 22, 23, 24, 25, 26, 27, 28,
						},
						[]interface{}{
							k,
							v.Description,
							v.HomePage,

							v.BuilderName,

							v.Removable,
							v.Reducible,
							v.AutoReduce,
							v.DontPreserveSharedObjects,
							v.NonBuildable,
							v.NonInstallable,
							v.Deprecated,
							v.PrimaryInstallOnly,

							strings.Join(v.BuildPkgDeps, "\n"),

							strings.Join(v.BuildDeps, "\n"),
							strings.Join(v.SODeps, "\n"),
							strings.Join(v.RunTimeDeps, "\n"),

							tags.New(v.Tags).String(),
							v.Category,
							strings.Join(tags.New(v.Groups).Values(), "\n"),

							v.TarballFileNameParser,
							v.TarballName,
							strings.Join(v.TarballFilters, "\n"),
							v.TarballProvider,
							strings.Join(v.TarballProviderArguments, "\n"),
							v.TarballStabilityClassifier,
							v.TarballVersionComparator,
							v.TarballProviderVersionSyncDepth,

							v.DownloadPatches,
							v.PatchesDownloadingScriptText,
						},
					)
					if err != nil {
						panic(err)
					}

					i++
					f := (0.5 / float64(len(pi)) * float64(i+1)) + 0.5
					self.pb.SetFraction(f)

				}
				c <- true
				return false
			},
		)
		<-c
	}
	return nil
}

func (self *UIMainWindow) SaveTable() error {

	glib.IdleAdd(
		func() bool {
			self.mb.SetSensitive(false)
			self.nb.SetSensitive(false)
			self.pb.Show()
			return false
		},
	)
	defer glib.IdleAdd(
		func() bool {
			self.pb.Hide()
			self.mb.SetSensitive(true)
			self.nb.SetSensitive(true)
			return false
		},
	)

	out_dir, err := InfoJSONDir()
	if err != nil {
		return err
	}

	pi := make(map[string]*basictypes.PackageInfo)

	ok := true
	{
		var i *gtk.TreeIter

		for {
			glib.IdleAdd(
				func() bool {
					self.pb.Pulse()
					return false
				},
			)

			if i == nil {
				i, ok = self.info_table_store.GetIterFirst()
			} else {
				ok = self.info_table_store.IterNext(i)
			}
			if !ok {
				break
			}

			name, info, err := self._IterToPackageInfo(i)
			if err != nil {
				return err
			}

			pi[name] = info

		}
	}

	err = os.RemoveAll(out_dir)
	if err != nil {
		return err
	}

	err = os.MkdirAll(out_dir, 0700)
	if err != nil {
		return err
	}

	i := 0
	for k, v := range pi {
		outfile, err := os.Create(path.Join(out_dir, k+".json"))
		if err != nil {
			return err
		}

		v.Groups = tags.New(v.Groups).Values()

		by, err := json.Marshal(v)
		if err != nil {
			return err
		}

		b := &bytes.Buffer{}
		if err != nil {
			return err
		}

		err = json.Indent(b, by, "", "  ")
		if err != nil {
			return err
		}

		b.WriteTo(outfile)

		outfile.Write([]byte("\n"))
		outfile.Close()

		i++
		glib.IdleAdd(
			func() bool {
				f := 1.0 / float64(len(pi)) * float64(i+1)
				self.pb.SetFraction(f)
				return false
			},
		)
	}

	return nil
}

func (self *UIMainWindow) _IterToPackageInfo(iter *gtk.TreeIter) (
	string,
	*basictypes.PackageInfo,
	error,
) {

	var name string

	ret := new(basictypes.PackageInfo)

	for _, i := range TableStructure {

		i_0_int := i[0].(int)

		v, err := self.info_table_store.GetValue(iter, i_0_int)
		if err != nil {
			return "", nil, err
		}

		vv, err := v.GoValue()
		if err != nil {
			return "", nil, err
		}

		switch i_0_int {
		case 0:
			name = vv.(string)
		case 1:
			ret.Description = vv.(string)
		case 2:
			ret.HomePage = vv.(string)
		case 3:
			ret.BuilderName = vv.(string)
		case 4:
			ret.Removable = vv.(bool)
		case 5:
			ret.Reducible = vv.(bool)
		case 6:
			ret.AutoReduce = vv.(bool)
		case 7:
			ret.DontPreserveSharedObjects = vv.(bool)
		case 8:
			ret.NonBuildable = vv.(bool)
		case 9:
			ret.NonInstallable = vv.(bool)
		case 10:
			ret.Deprecated = vv.(bool)
		case 11:
			ret.PrimaryInstallOnly = vv.(bool)
		case 12:
			ret.BuildPkgDeps = textlist.RemoveZeroLengthItems(strings.Split(vv.(string), "\n"))
		case 13:
			ret.BuildDeps = textlist.RemoveZeroLengthItems(strings.Split(vv.(string), "\n"))
		case 14:
			ret.SODeps = textlist.RemoveZeroLengthItems(strings.Split(vv.(string), "\n"))
		case 15:
			ret.RunTimeDeps = textlist.RemoveZeroLengthItems(strings.Split(vv.(string), "\n"))
		case 16:
			ret.Tags = textlist.RemoveZeroLengthItems(tags.NewFromString(vv.(string)).Values())
		case 17:
			ret.Category = vv.(string)
		case 18:
			ret.Groups = textlist.RemoveZeroLengthItems(strings.Split(vv.(string), "\n"))
		case 19:
			ret.TarballFileNameParser = vv.(string)
		case 20:
			ret.TarballName = vv.(string)
		case 21:
			ret.TarballFilters = textlist.RemoveZeroLengthItems(strings.Split(vv.(string), "\n"))
		case 22:
			ret.TarballProvider = vv.(string)
		case 23:
			ret.TarballProviderArguments = textlist.RemoveZeroLengthItems(strings.Split(vv.(string), "\n"))
		case 24:
			ret.TarballStabilityClassifier = vv.(string)
		case 25:
			ret.TarballVersionComparator = vv.(string)
		case 26:
			ret.TarballProviderVersionSyncDepth = vv.(int)
		case 27:
			ret.DownloadPatches = vv.(bool)
		case 28:
			ret.PatchesDownloadingScriptText = vv.(string)
		}
	}

	return name, ret, nil
}

func _EditorDir() (string, error) {

	d, err := build.Import(
		"github.com/AnimusPEXUS/aipsetup/cmd/aipinfoeditor",
		"",
		build.FindOnly,
	)
	if err != nil {
		return "", err
	}

	return d.Dir, nil
}

func InfoJSONDir() (string, error) {
	d, err := _EditorDir()
	if err != nil {
		return "", err
	}

	return path.Join(d, "infojson"), nil
}

func InfoBindataDir() (string, error) {
	d, err := _EditorDir()
	if err != nil {
		return "", err
	}

	return path.Join(d, "infobindata"), nil
}
