package tarballnameparsers

// TODO: move interface to basictypes if possible
// NOTE: on other tout, possibly this package will be separated from aipsetup
//       someday
type TarballNameParserI interface {
	ParseName(value string) (*ParseResult, error)
}

var Index = map[string](func() TarballNameParserI){
	"std": func() TarballNameParserI {
		return new(TarballNameParser_Std)
	},
}
