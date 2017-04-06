dst=dist
theme=bootstrap

build:
	mkdir -pv $(dst)/themes/$(theme)/public
	go build -ldflags "-s -w -X github.com/kapmahc/fly/web.Version=`git rev-parse --short HEAD` -X github.com/kapmahc/fly/web.BuildTime=`date +%FT%T%z`" -o ${dst}/fly main.go
	-cp -rv locales db $(dst)/
	-cp -rv themes/$(theme)/views themes/$(theme)/assets $(dst)/themes/$(theme)/	
	tar jcvf $(dst).tar.bz2 $(dst)

init:
	go get -u github.com/kardianos/govendor
	govendor sync
	cd themes/$(theme) && npm install

clean:
	-rm -rv $(dst) $(dst).tar.bz2
