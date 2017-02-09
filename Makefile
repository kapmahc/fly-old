dst=dist

all:
	go build -ldflags "-s -w -X github.com/kapmahc/fly/engines/base.Version=`git rev-parse --short HEAD` -X github.com/kapmahc/fly/engines/base.BuildTime=`date +%FT%T%z`" -o $(dst)/fly main.go
	-cp -rv conf Makefile themes $(dst)/

clean:
	-rm -rv $(dst)
