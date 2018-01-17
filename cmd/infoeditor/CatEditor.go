package main

import (
	"fmt"
	"strings"

	"github.com/AnimusPEXUS/gotk3collection/explorer"
	"github.com/AnimusPEXUS/gotk3collection/treemodel"
	"github.com/AnimusPEXUS/utils/directory"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type CatEditorFileOperation uint

const (
	OperationNone CatEditorFileOperation = iota
	OperationCut
	OperationCopy
)

type CatEditor struct {
	mw             *UIMainWindow
	cated1, cated2 *explorer.Explorer
	cated_store    *gtk.TreeStore
	pan1           *gtk.Paned
	tb_reload_cats *gtk.ToolButton

	_FolderIconPixbuf, _FileIconPixbuf     *gdk.Pixbuf
	_FolderIconPixbufGV, _FileIconPixbufGV *glib.Value

	current_operation CatEditorFileOperation
	operation_subject []*explorer.SelectedForOperationItem

	d *directory.File
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
	}

	self.tb_reload_cats.Connect(
		"clicked",
		func() {
			self.ReloadInfo()
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

	fmt.Println("move/copy..")
	for _, i := range paths {
		fmt.Print("  ")
		if i.IsDir {
			fmt.Print("<dir> ")
		}
		fmt.Println(i.Path)
	}
	fmt.Println("    to", path.Path)

	m := self.cated_store

	t_path, err := treemodel.GetTreePathByValueStringPath(
		strings.Split(path.Path, "/"),
		m,
		1,
	)
	if err != nil {
		return err
	}

	for _, i := range paths {
		i_sep := strings.Split(i.Path, "/")

		tp, err := treemodel.GetTreePathByValueStringPath(i_sep, m, 1)
		if err != nil {
			return err
		}

		fmt.Println("before copy")

		c, err := treemodel.CopyTreeStore(
			m,
			tp,
		)
		if err != nil {
			return err
		}

		fmt.Println("printing c")
		fmt.Println(treemodel.TreeStoreString(c, nil))
		fmt.Println("before walk")

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
	fmt.Println("_Paste result", ret)
	return ret
}

func (self *CatEditor) _DeleteFunction(
	paths []*explorer.SelectedForOperationItem,
) error {
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

	self.d = directory.NewFile(nil, "", true, nil)

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

			gv, err = s.GetValue(it, 13)
			if err != nil {
				return err
			}

			cat_string, err := gv.GetString()
			if err != nil {
				return err
			}

			d := self.d.GetByPath(
				strings.Split(cat_string, "/"),
				true,
				false,
				nil,
			)
			d.MkFile(record_name, nil)
		}
	}

	self.d.Walk(
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
					// fmt.Println("folder setting pixbuf", self._FolderIconPixbuf, i.Name())
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
				// fmt.Println("file setting pixbuf", self._FolderIconPixbuf.GetPixbuf(), i.Name())
				s.SetValue(it, 0, self._FileIconPixbuf)
				s.SetValue(it, 1, i.Name())
				s.SetValue(it, 2, false)
			}

			return nil
		},
	)

	return nil
}
