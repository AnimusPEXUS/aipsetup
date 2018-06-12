package main

import (
	"errors"
	"strings"

	"github.com/AnimusPEXUS/gotk3collection/explorer"
	"github.com/AnimusPEXUS/gotk3collection/treemodel"
	"github.com/AnimusPEXUS/utils/directory"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type CatEditor struct {
	mw             *UIMainWindow
	cated1, cated2 *explorer.Explorer
	cated_store    *gtk.TreeStore
	pan1           *gtk.Paned
	tb_reload_cats *gtk.ToolButton
	tb_apply_cats  *gtk.ToolButton

	_FolderIconPixbuf, _FileIconPixbuf     *gdk.Pixbuf
	_FolderIconPixbufGV, _FileIconPixbufGV *glib.Value

	current_operation FileOperation
	operation_subject []*explorer.SelectedForOperationItem
}

func CatEditorNew(mw *UIMainWindow, builder *gtk.Builder) (*CatEditor, error) {
	self := &CatEditor{
		mw: mw,
	}

	err := self._UpdateIconPixbufs()
	if err != nil {
		return nil, err
	}

	if t, err := gtk.TreeStoreNew(
		gdk.PixbufGetType(),
		glib.TYPE_STRING,
		glib.TYPE_BOOLEAN,
	); err != nil {
		return nil, err
	} else {
		self.cated_store = t
	}

	if t, err := builder.GetObject("pan1"); err != nil {
		return nil, err
	} else {
		self.pan1 = t.(*gtk.Paned)
	}

	if t, err := builder.GetObject("tb_reload_cats"); err != nil {
		return nil, err
	} else {
		self.tb_reload_cats = t.(*gtk.ToolButton)
	}

	if t, err := builder.GetObject("tb_apply_cats"); err != nil {
		return nil, err
	} else {
		self.tb_apply_cats = t.(*gtk.ToolButton)
	}

	if t, err := explorer.ExplorerNew(); err != nil {
		return nil, err
	} else {
		self.cated1 = t
	}

	if t, err := explorer.ExplorerNew(); err != nil {
		return nil, err
	} else {
		self.cated2 = t
	}

	for _, i := range []*explorer.Explorer{
		self.cated1, self.cated2,
	} {
		self.pan1.Add(i.GetWidget())

		i.SetColumns(0, 1, 2)
		i.SetControlls(
			false, true,
			true, false, true,
		)
		i.SetModel(self.cated_store)

		i.SetCutFunction(self._CutFunction)
		i.SetPasteFunction(self._PasteFunction)
		i.SetDeleteFunction(self._DeleteFunction)
		i.SetMkDirFunction(self._MkDirFunction)
	}

	self.tb_reload_cats.Connect(
		"clicked",
		func() {
			self.ReloadInfo()
		},
	)

	self.tb_apply_cats.Connect(
		"clicked",
		func() {
			self.ApplyInfo()
		},
	)

	self.pan1.Connect(
		"style-set",
		func() {
			self._UpdateIconPixbufs()
		},
	)

	return self, nil
}

func (self *CatEditor) _MoveCopyOperation(
	paths []*explorer.SelectedForOperationItem,
	path *explorer.SelectedForOperationItem,
	move bool,
) error {

	m := self.cated_store

	var t_path *gtk.TreePath
	var err error

	if path != nil {
		t_path, err = treemodel.GetTreePathByValueStringPath(
			strings.Split(path.Path, "/"),
			m,
			1,
		)
		if err != nil {
			return err
		}
	}

	for _, i := range paths {
		i_sep := strings.Split(i.Path, "/")

		tp, err := treemodel.GetTreePathByValueStringPath(i_sep, m, 1)
		if err != nil {
			return err
		}

		c, err := treemodel.CopyTreeStore(
			m,
			tp,
		)
		if err != nil {
			return err
		}

		err = treemodel.PasteTreeStore(c, m, t_path)
		if err != nil {
			return err
		}

		if move {
			tp_it, err := m.GetIter(tp)
			if err != nil {
				return err
			}
			m.Remove(tp_it)
		}
	}

	return nil
}

func (self *CatEditor) _CutFunction(
	paths []*explorer.SelectedForOperationItem,
) error {
	self.current_operation = OperationCut
	self.operation_subject = paths
	return nil
}

func (self *CatEditor) _CopyFunction(
	paths []*explorer.SelectedForOperationItem,
) error {
	self.current_operation = OperationCopy
	self.operation_subject = paths
	return nil
}

func (self *CatEditor) _PasteFunction(
	path *explorer.SelectedForOperationItem,
) error {
	var ret error
	switch self.current_operation {
	case OperationCut:
		ret = self._MoveCopyOperation(
			self.operation_subject,
			path,
			true,
		)
	case OperationCopy:
		ret = self._MoveCopyOperation(
			self.operation_subject,
			path,
			false,
		)
	default:
		// nothing
	}
	self.current_operation = OperationNone
	return ret
}

func (self *CatEditor) _DeleteFunction(
	paths []*explorer.SelectedForOperationItem,
) error {
	err := self._MoveCopyOperation(paths, nil, true)
	return err
}

func (self *CatEditor) _MkDirFunction(target *explorer.SelectedForOperationItem) error {

	return nil
}

func (self *CatEditor) _UpdateIconPixbufs() error {
	it, err := gtk.IconThemeGetDefault()
	if err != nil {
		return err
	}

	if t, err := it.LoadIcon("folder", 16, 0); err != nil {
		return err
	} else {
		self._FolderIconPixbuf = t
	}

	if t, err := it.LoadIcon("text-x-generic", 16, 0); err != nil {
		return err
	} else {
		self._FileIconPixbuf = t
	}

	return nil
}

func (self *CatEditor) ReloadInfo() error {
	s := self.cated_store

	s.Clear()

	d := directory.NewFile(nil, "", true, nil)

	{
		s := self.mw.info_table_store

		var it *gtk.TreeIter
		ok := true

		for {
			if it == nil {
				it, ok = self.mw.info_table_store.GetIterFirst()
			} else {
				ok = self.mw.info_table_store.IterNext(it)
			}
			if !ok {
				break
			}

			gv, err := s.GetValue(it, 0)
			if err != nil {
				return err
			}

			record_name, err := gv.GetString()
			if err != nil {
				return err
			}

			gv, err = s.GetValue(it, 15)
			if err != nil {
				return err
			}

			cat_string, err := gv.GetString()
			if err != nil {
				return err
			}

			cat_string = strings.Trim(cat_string, "/ ")

			gbpl := []string{}
			if len(cat_string) != 0 {
				gbpl = strings.Split(cat_string, "/")
			}

			dd, err := d.GetByPath(
				gbpl,
				true,
				false,
				nil,
			)
			if err != nil {
				return err
			}
			dd.MkFile(record_name, nil)
		}
	}

	d.Walk(
		func(path []*directory.File, dirs, files []*directory.File) error {

			var cur_lvl_path *gtk.TreePath
			var err error
			var cur_lvl_path_iter *gtk.TreeIter

			if cur_lvl_path != nil {
				cur_lvl_path_iter, err = s.GetIter(cur_lvl_path)
				if err != nil {
					return err
				}
			}

			for _, i := range path {

				pa, ok, err := treemodel.FindTreePathByStringAndColOnSameLevel(
					s,
					cur_lvl_path,
					i.Name(),
					1,
				)
				if err != nil {
					return err
				}

				var t_it *gtk.TreeIter

				if ok {
					t_it, err = s.GetIter(pa)
					if err != nil {
						return err
					}
				} else {
					t_it = s.Append(cur_lvl_path_iter)
					s.SetValue(t_it, 0, self._FolderIconPixbuf)
					s.SetValue(t_it, 1, i.Name())
					s.SetValue(t_it, 2, true)
				}

				pa, err = s.GetPath(t_it)
				if err != nil {
					return err
				}

				cur_lvl_path = pa
				cur_lvl_path_iter = t_it
			}

			for _, i := range files {
				it := s.Append(cur_lvl_path_iter)
				s.SetValue(it, 0, self._FileIconPixbuf)
				s.SetValue(it, 1, i.Name())
				s.SetValue(it, 2, false)
			}

			return nil
		},
	)

	return nil
}

func (self *CatEditor) ApplyInfo() error {

	new_tree := directory.NewFile(nil, "", true, nil)

	{
		m := self.cated_store
		err := treemodel.WalkTreeStore(
			m, nil,
			func(
				m *gtk.TreeStore,
				path *gtk.TreePath,
				values [][]*treemodel.Value,
			) error {

				path_str_lst, err := treemodel.RenderTreePathString(
					m,
					path,
					1,
				)
				if err != nil {
					return err
				}

				dir, err := new_tree.GetByPath(path_str_lst, true, false, nil)
				if err != nil {
					return err
				}

				for _, i := range values {
					isdir_i, err := i[2].Interface()
					if err != nil {
						return err
					}
					isdir_ib, ok := isdir_i.(bool)
					if !ok {
						return errors.New("can't convert to bool")
					}
					if !isdir_ib {
						n, err := i[1].GValue.GetString()
						if err != nil {
							return err
						}
						_, err = dir.MkFile(n, nil)
						if err != nil {
							return err
						}
					}
				}

				return nil
			},
		)
		if err != nil {
			return err
		}
	}

	{
		s := self.mw.info_table_store

		ok := true

		var iter *gtk.TreeIter

		for {
			if iter == nil {
				iter, ok = s.GetIterFirst()
				if !ok {
					return nil
				}
			} else {
				ok = s.IterNext(iter)
			}
			if !ok {
				break
			}

			name_v, err := s.GetValue(iter, 0)
			if err != nil {
				return err
			}
			name, err := name_v.GetString()
			if err != nil {
				return err
			}
			fres, err := new_tree.FindFile(name)
			if err != nil {
				return err
			}

			var cat string

			switch len(fres) {
			default:
				return errors.New("duplicated names found in tree")
			case 0:
				cat = ""
				// ignore. but potentially this is programming error
			case 1:

				if len(fres[0]) == 0 {
					cat = ""
				} else {
					spl_fres0 := strings.Split(fres[0], "/")
					cat = strings.Join(
						spl_fres0[:len(spl_fres0)-1],
						"/",
					)
				}

			}
			err = s.SetValue(iter, 15, cat)
			if err != nil {
				return err
			}

		}

	}

	return nil
}
