package aipsetup

type (
	ASPackageConstitution struct {
		Cc               string   "yaml:'CC' json:'CC'"
		Cxx              string   "yaml:'CXX' json:'CXX'"
		Host             string   "yaml:'host' json:'host'"
		Arch             string   "yaml:'arch' json:'arch'"
		Build            string   "yaml:'build' json:'build'"
		Target           string   "yaml:'target' json:'target'"
		MultilibVariants []string "yaml:'multilib_variants' json:'multilib_variants'"
		SystemTitle      string   "yaml:'system_title' json:'system_title'"
		SystemVersion    string   "yaml:'system_version' json:'system_version'"
	}

	ASPackageInfo struct {
		Constitution ASPackageConstitution "yaml:'constitution' json:'constitution'"
		PkgInfo      OveralPackageInfo     "yaml:'pkg_info' json:'pkg_info'"
	}
)
