package pkginfodb

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb/db"
)

var Index map[string]*basictypes.PackageInfo

func init() {
	Index = make(map[string]*basictypes.PackageInfo)
	for _, i := range db.AssetNames() {
		s := strings.Split(i, "/")
		name := s[len(s)-1]
		name = name[:len(name)-5]
		t := new(basictypes.PackageInfo)
		err := json.Unmarshal(
			db.MustAsset(
				fmt.Sprintf("../cmd/infoeditor/infojson/%s.json", name),
			),
			t,
		)
		if err != nil {
			panic(err)
		}
		Index[name] = t
	}
}