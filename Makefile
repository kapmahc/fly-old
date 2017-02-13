dst=dist

all:
	go build -ldflags "-s -w -X github.com/kapmahc/fly/web.Version=`git rev-parse --short HEAD` -X github.com/kapmahc/fly/web.BuildTime=`date +%FT%T%z`" -o $(dst)/fly main.go
	-cp -rv locales db themes $(dst)/

clean:
	-rm -rv $(dst)
