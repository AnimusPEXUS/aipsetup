# This is a -*-perl-*- script
#
# Set variables that were defined by configure, in case we need them
# during the tests.

%CONFIG_FLAGS = (
    AM_LDFLAGS   => '-Wl,--export-dynamic',
    AR           => 'ar',
    CC           => 'gcc',
    CFLAGS       => '-g -O2',
    CPP          => 'gcc -E',
    CPPFLAGS     => '',
    GUILE_CFLAGS => '-pthread -I/multihost/x86_64-pc-linux-gnu/include/guile/2.0 -I/multihost/x86_64-pc-linux-gnu/include',
    GUILE_LIBS   => '-L/multihost/x86_64-pc-linux-gnu/lib64 -lguile-2.0 -lgc',
    LDFLAGS      => '',
    LIBS         => '-ldl '
);

1;
