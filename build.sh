go-bindata -o src/data/static.go -pkg data static/...
go build -o bin/erpass ./src && cd bin && ./erpass