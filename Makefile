dst=dist
theme=bootstrap

all:
	mkdir -pv $(dst)/themes/$(theme)
	go build -ldflags "-s -w -X github.com/kapmahc/fly/web.Version=`git rev-parse --short HEAD` -X github.com/kapmahc/fly/web.BuildTime=`date +%FT%T%z`" -o $(dst)/fly main.go
	cd themes/$(theme) && npm run build
	-cp -rv themes/$(theme)/{assets,views} $(dst)/themes/$(theme)
	-cp -rv locales db $(dst)/

clean:
	-rm -rv $(dst)
