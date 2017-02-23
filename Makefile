dst=dist
theme=bootstrap

all:
	go build -ldflags "-s -w -X github.com/kapmahc/fly/web.Version=`git rev-parse --short HEAD` -X github.com/kapmahc/fly/web.BuildTime=`date +%FT%T%z`" -o $(dst)/fly main.go
	cd themes/$(theme) && npm run build
	-cp -rv locales db $(dst)/
	mkdir -pv $(dst)/themes/$(theme)
	cp -rv themes/$(theme)/{assets,views} $(dst)/themes/$(theme)

init:
	go get -u github.com/kardianos/govendor
	govendor sync
	cd themes/$(theme) && npm install

clean:
	-rm -rv $(dst)
