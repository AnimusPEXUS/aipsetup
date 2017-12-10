# This is a -*-perl-*- script
#
# Set variables that were defined by configure, in case we need them
# during the tests.

%CONFIG_FLAGS = (
    AM_LDFLAGS   => '-Wl,--export-dynamic',
    AR           => 'ar',
    CC           => 'x86_64-pc-linux-gnu-gcc -m64',
    CFLAGS       => '-g -O2',
    CPP          => 'x86_64-pc-linux-gnu-gcc -m64 -E',
    CPPFLAGS     => '',
    GUILE_CFLAGS => '-pthread -I/multihost/x86_64-pc-linux-gnu/include/guile/2.0',
    GUILE_LIBS   => '-L/multihost/x86_64-pc-linux-gnu/lib64 -lguile-2.0 -lgc',
    LDFLAGS      => '',
    LIBS         => '-ldl '
);

1;
