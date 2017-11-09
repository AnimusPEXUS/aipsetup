#!/usr/bin/python3

import re
import os

reg = re.compile(
'TarballName\:\s*map\[string\]string\{\s*"prefix"\:\s*"(?P<value>.*?)",?\s*\},'
)


lst = os.listdir()

found_count = 0
not_found_count = 0

for i in lst:

	f=open(i)
	ft=f.read()
	f.close()

	search_res = reg.search(ft)

	if search_res is not None:
		print("Name {}".format(search_res.group("value")))


		ftn = (ft[:search_res.start()] +
		 "TarballName: \"{}\",".format(search_res.group("value")) +
		 ft[search_res.end():]
		 )

		print("---")
		print(ftn)
		print("---")

		f=open(i,"w")
		f.write(ftn)
		f.close()

		found_count+=1
	else:
		not_found_count+=1

print("found", found_count, "not found", not_found_count)
