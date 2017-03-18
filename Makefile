dst=dist
theme=bootstrap

build:
	go build -ldflags "-s -w -X github.com/kapmahc/fly/web.Version=`git rev-parse --short HEAD` -X github.com/kapmahc/fly/web.BuildTime=`date +%FT%T%z`" -o ${dst}/fly main.go
	-cp -rv locales db $(dst)/
	cd themes/$(theme) && npm run build
	-cp -rv themes/$(theme)/public $(dst)/public

init:
	go get -u github.com/kardianos/govendor
	govendor sync
	cd themes/$(theme) && npm install

clean:
	-rm -rv $(dst) themes/$(theme)/public
