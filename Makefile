clean:
	rm -r bin/*

static:
	go-bindata -o src/data/static.go -pkg data static/...

darwin:
	go build -o bin/erpass-darwin ./src 

test:
	cd bin && ./erpass-darwin

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/erpass-linux-amd64 ./src 

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/erpass-windows-amd64.exe ./src 

build:
	make darwin && make test

all:
	make darwin && make linux && make windows