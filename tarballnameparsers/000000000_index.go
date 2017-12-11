package tarballnameparsers

var Index = map[string](func() TarballNameParserI){
	"std": func() TarballNameParserI {
		return new(TarballNameParser_Std)
	},
}
