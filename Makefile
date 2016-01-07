all: bin/eecsproject

clean:
	rm -rf bin/*
	rm -rf pkg/*
	rm -rf vendor/*
	rm src/templates/data.go

bin/go-bindata:
	gb vendor fetch github.com/jteeuwen/go-bindata
	gb build github.com/jteeuwen/go-bindata/go-bindata

src/templates/data.go: bin/go-bindata
	bin/go-bindata -o temp.go assets/...
	sed 's/package main/package templates/' <temp.go >src/templates/data.go
	rm temp.go

bin/eecsproject: src/templates/data.go src/cmd/eecsproject/main.go
	gb build all
