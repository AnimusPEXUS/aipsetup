package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

func init() {
	// NOTE: not an error sgml3, sgml4 and xml4 are monstly the same
	Index["docbook_xml4"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		ret, err := NewBuilder_docbook_sgml3(bs)
		if err != nil {
			return nil, err
		}
		ret.sgml_or_xml = "xml"
		return ret, nil
	}
}
