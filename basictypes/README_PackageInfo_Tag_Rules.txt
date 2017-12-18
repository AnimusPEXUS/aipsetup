Some tags have special meaning

Calling 'info-db code' will save new distropkginfodb with according changes.

'group\:(.*)' - used for grouping packages

'version_scheme\:(.*)' - versioning scheme used to determine stability of
          the source in tarball.

          only 2 values of $1 have meaning to aipsetup:

            'gnome' - then odd numbers means stable versions, and
                  even and 9xx versions means testing and unstable version

            'gcc' - then x.0.0 versions means development versions

'sf_project\:(.*)' - if this tag is present, then this packages tarball
          considered to be found on sf.net site under project named $1

          TarballProvider field will be changed to 'sf'
          TarballProviderArguments will be changed to $1

          each package with this tag will get separate cache

'kernel_project' - edits info to target kernel.org. also all packages with
          this tag will get same listdir cache

'cairo_project' - simmilar to kernel_project

'gtk_project' - simmilar to kernel_project, but with tag 'version_scheme:gnome'
        will be added

'gnome_project' - simmilar to gtk_project

'perl_module' - treats package info as Perl's module
