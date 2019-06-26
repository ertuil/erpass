build:
	make clean & make static && make darwin && make test

.PHONY:clean static darwin test linux windows all 

static:
	go-bindata -o src/data/static.go -pkg data static/... templates/... ReadMe.md

darwin:
	go build -o bin/erpass-darwin ./src 

test:
	cd bin && ./erpass-darwin --host 127.0.0.1 --port 8080

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/erpass-linux-amd64 ./src 

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/erpass-windows-amd64.exe ./src 

all:
	make darwin && make linux && make windows

clean:
	rm -r bin/static bin/templates bin/erpass-* bin/erpass.log