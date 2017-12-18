package infoeditor

// List tarball names which have "Gnome" version scheme

var PACKAGES_WITH_GNOME_VERSION_SCHEME []string

func init() {
	PACKAGES_WITH_GNOME_VERSION_SCHEME = append(
		PACKAGES_WITH_GNOME_VERSION_SCHEME,
		GNOME_PROJECTS...,
	)

	PACKAGES_WITH_GNOME_VERSION_SCHEME = append(
		PACKAGES_WITH_GNOME_VERSION_SCHEME,
		GTK_PROJECTS...,
	)

}
