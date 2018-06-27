#!/usr/bin/python3

import glob
import sys

LD_LIBRARY_PATH = []
PATH=[]

# calc LD_LIBRARY_PATH

LD_LIBRARY_PATH += glob.glob('/multihost/*/lib')
LD_LIBRARY_PATH += glob.glob('/multihost/*/lib64')

LD_LIBRARY_PATH += glob.glob('/multihost/*/multiarch/*/lib')
LD_LIBRARY_PATH += glob.glob('/multihost/*/multiarch/*/lib64')

LD_LIBRARY_PATH += ['/lib', '/lib64']


for i in range(len(LD_LIBRARY_PATH)-1,-1,-1):
    for j in ['_primary']:
        if j in LD_LIBRARY_PATH[i]:
            del LD_LIBRARY_PATH[i]
            break

# calc PATH

PATH += glob.glob('/multihost/*/bin')
PATH += glob.glob('/multihost/*/sbin')
PATH += ['/usr/bin', '/usr/sbin']
PATH += ['/bin', '/sbin']
PATH += glob.glob('/multihost/*/multiarch/*/bin')
PATH += glob.glob('/multihost/*/multiarch/*/sbin')

for i in range(len(PATH)-1,-1,-1):
    for j in ['_primary']:
        if j in PATH[i]:
            del PATH[i]
            break

# print results

#print("export LD_LIBRARY_PATH={}".format(':'.join(LD_LIBRARY_PATH)))
print("export PATH={}".format(':'.join(PATH)))

# print table for ld.so.conf

if len(sys.argv) > 1 and sys.argv[1] == '-p':
    print('\n'.join(LD_LIBRARY_PATH))
exit(0)
