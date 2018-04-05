all:
	go generate -v -x
	cd cmd/aipinfoeditor && go build -v
	cd cmd/aipsetup5 && go build -v
