#!/bin/bash

export XDG_DATA_DIRS=

for i in /multihost/*/share/
do
    if [ -z $XDG_DATA_DIRS ]
    then
        XDG_DATA_DIRS=$i
    else
        XDG_DATA_DIRS=$XDG_DATA_DIRS:$i
    fi
done

unset i
