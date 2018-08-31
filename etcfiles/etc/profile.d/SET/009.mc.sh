#!/bin/bash

if [ -x /multihost/_primary/share/mc/bin/mc-wrapper.sh ]
then
    alias mc=". /multihost/_primary/share/mc/bin/mc-wrapper.sh"
else
    if [ -x /multihost/_primary/libexec/mc/mc-wrapper.sh ]
    then
        alias mc=". /multihost/_primary/libexec/mc/mc-wrapper.sh"
    fi
fi
