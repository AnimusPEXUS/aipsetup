all:
	go generate -v -x
	cd cmd/aipinfoeditor && go build
	cd cmd/aipsetup5 && go build
