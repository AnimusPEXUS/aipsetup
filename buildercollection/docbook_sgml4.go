package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

func init() {
	// NOTE: not an error sgml3, sgml4 and xml4 are monstly the same
	Index["docbook_sgml4"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_docbook_sgml3(bs)
	}
}
