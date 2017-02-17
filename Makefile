dst=dist
theme=bootstrap

all:
	go build -ldflags "-s -w -X github.com/kapmahc/fly/web.Version=`git rev-parse --short HEAD` -X github.com/kapmahc/fly/web.BuildTime=`date +%FT%T%z`" -o $(dst)/fly main.go
	-cp -rv locales db themes $(dst)/
	cd themes/bootstrap/$(theme) && npm run build

clean:
	-rm -rv $(dst)
