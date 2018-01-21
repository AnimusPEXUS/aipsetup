package main

import (
	"bytes"
	"encoding/json"
	"go/build"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/cmd/infoeditor/ui"
	"github.com/AnimusPEXUS/utils/tags"
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
	[3]interface{}{6, glib.TYPE_BOOLEAN, "NonInstallable"},
	[3]interface{}{7, glib.TYPE_BOOLEAN, "Deprecated"},
	[3]interface{}{8, glib.TYPE_BOOLEAN, "PrimaryOnly"},
	[3]interface{}{9, glib.TYPE_STRING, "Building Dependencies"},
	[3]interface{}{10, glib.TYPE_STRING, "Shared Object Dependencies"},
	[3]interface{}{11, glib.TYPE_STRING, "Runtime Dependencies"},
	[3]interface{}{12, glib.TYPE_STRING, "Tags"},
	[3]interface{}{13, glib.TYPE_STRING, "Category"},
	[3]interface{}{14, glib.TYPE_STRING, "Groups"},
	[3]interface{}{15, glib.TYPE_STRING, "Tarball Name"},
	[3]interface{}{16, glib.TYPE_STRING, "Tarball Parser"},
	[3]interface{}{17, glib.TYPE_STRING, "Tarball Filters"},
	[3]interface{}{18, glib.TYPE_STRING, "Tarball Provider"},
	[3]interface{}{19, glib.TYPE_STRING, "Provider Arguments"},
	[3]interface{}{20, glib.TYPE_BOOLEAN, "Use Cache"},
	[3]interface{}{21, glib.TYPE_STRING, "Cache Preset Name"},
	[3]interface{}{22, glib.TYPE_INT, "Tarball Sync Depth"},
	[3]interface{}{23, glib.TYPE_STRING, "Version Tool"},
}

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
						c.SetFixedWidth(400)
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
								[]interface{}{new_text},
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
					c.SetFixedWidth(50)
					c.SetMaxWidth(50)
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
				self.LoadTable()
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
			go func() {
				self.LoadTable()
			}()
		},
	)

	return self, nil
}

func (self *UIMainWindow) Show() {
	self.window.ShowAll()
}

func (self *UIMainWindow) LoadTable() error {
	glib.IdleAdd(
		func() {
			self.mb.SetSensitive(false)
			self.nb.SetSensitive(false)
			// self.nb.SetVisible(false)
			self.pb.Show()
		},
	)
	defer glib.IdleAdd(
		func() {
			self.pb.Hide()
			self.mb.SetSensitive(true)
			// self.nb.SetVisible(true)
			self.nb.SetSensitive(true)
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

	pi := make(map[string]*basictypes.PackageInfo)
	for ii, i := range stats {
		if strings.HasSuffix(i.Name(), ".json") && !i.IsDir() {

			tb, err := ioutil.ReadFile(path.Join(d, i.Name()))
			if err != nil {
				return err
			}

			t := new(basictypes.PackageInfo)
			err = json.Unmarshal(tb, t)
			if err != nil {
				return err
			}

			name := i.Name()[0 : len(i.Name())-5]
			pi[name] = t

		}
		c := make(chan bool)
		glib.IdleAdd(
			func() {
				f := 0.5 / float64(len(stats)) * float64(ii+1)
				self.pb.SetFraction(f)
				c <- true
			},
		)
		<-c
	}

	{
		c := make(chan bool)
		glib.IdleAdd(
			func() {
				self.info_table_store.Clear()
				c <- true
			},
		)
		<-c

	}
	{
		c := make(chan bool)
		glib.IdleAdd(
			func() {

				i := 0
				for k, v := range pi {
					iter := self.info_table_store.Append()
					err := self.info_table_store.Set(
						iter,
						[]int{
							0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
							10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
							20, 21, 22, 23,
						},
						[]interface{}{
							k,
							v.Description,
							v.HomePage,

							v.BuilderName,

							v.Removable,
							v.Reducible,
							v.NonInstallable,
							v.Deprecated,
							v.PrimaryInstallOnly,

							strings.Join(v.BuildDeps, "\n"),
							strings.Join(v.SODeps, "\n"),
							strings.Join(v.RunTimeDeps, "\n"),

							tags.New(v.Tags).String(),
							v.Category,
							strings.Join(v.Groups, "\n"),

							v.TarballName,
							v.TarballFileNameParser,
							strings.Join(v.TarballFilters, "\n"),
							v.TarballProvider,
							strings.Join(v.TarballProviderArguments, "\n"),
							v.TarballProviderUseCache,
							v.TarballProviderCachePresetName,
							v.TarballProviderVersionSyncDepth,
							v.TarballVersionTool,
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
			},
		)
		<-c
	}
	return nil
}

func (self *UIMainWindow) SaveTable() error {

	glib.IdleAdd(
		func() {
			self.mb.SetSensitive(false)
			self.nb.SetSensitive(false)
			self.pb.Show()
		},
	)
	defer glib.IdleAdd(
		func() {
			self.pb.Hide()
			self.mb.SetSensitive(true)
			self.nb.SetSensitive(true)
		},
	)

	out_dir, err := InfoJSONDir()
	if err != nil {
		return err
	}

	// pi := pkginfodb.Index

	pi := make(map[string]*basictypes.PackageInfo)

	ok := true
	{
		var i *gtk.TreeIter

		for {
			glib.IdleAdd(
				func() {
					self.pb.Pulse()
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
			func() {
				f := 1.0 / float64(len(pi)) * float64(i+1)
				self.pb.SetFraction(f)
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
			ret.NonInstallable = vv.(bool)
		case 7:
			ret.Deprecated = vv.(bool)
		case 8:
			ret.PrimaryInstallOnly = vv.(bool)
		case 9:
			ret.BuildDeps = strings.Split(vv.(string), "\n")
		case 10:
			ret.SODeps = strings.Split(vv.(string), "\n")
		case 11:
			ret.RunTimeDeps = strings.Split(vv.(string), "\n")
		case 12:
			ret.Tags = tags.NewFromString(vv.(string)).Values()
		case 13:
			ret.Category = vv.(string)
		case 14:
			ret.Groups = strings.Split(vv.(string), "\n")
		case 15:
			ret.TarballName = vv.(string)
		case 16:
			ret.TarballFileNameParser = vv.(string)
		case 17:
			ret.TarballFilters = strings.Split(vv.(string), "\n")
		case 18:
			ret.TarballProvider = vv.(string)
		case 19:
			ret.TarballProviderArguments = strings.Split(vv.(string), "\n")
		case 20:
			ret.TarballProviderUseCache = vv.(bool)
		case 21:
			ret.TarballProviderCachePresetName = vv.(string)
		case 22:
			ret.TarballProviderVersionSyncDepth = vv.(int)
		case 23:
			ret.TarballVersionTool = vv.(string)
		}
	}

	return name, ret, nil
}

func _EditorDir() (string, error) {

	d, err := build.Import(
		"github.com/AnimusPEXUS/aipsetup/cmd/infoeditor",
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
