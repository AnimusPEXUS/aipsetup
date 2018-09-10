all:
	go generate -v -x
#	cd cmd/aipinfoeditor && go build -v
	cd cmd/aipsetup5 && go build -v

chmod:
	find -type d -exec chmod 700 "{}" ";" -print
	find -type f -exec chmod 600 "{}" ";" -print
