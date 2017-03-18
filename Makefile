dst=dist
theme=bootstrap

all: backend frontend


backend:
	go build -ldflags "-s -w -X github.com/kapmahc/fly/web.Version=`git rev-parse --short HEAD` -X github.com/kapmahc/fly/web.BuildTime=`date +%FT%T%z`" -o ${dst}/fly demo/main.go
	-cp -rv demo/locales demo/db $(dst)/

frontend:
	cd themes/$(theme) && npm run build
	-cp -rv themes/$(theme)/assets $(dst)/public

init:
	go get -u github.com/kardianos/govendor
	govendor sync
	cd themes/$(theme) && npm install

clean:
	-rm -rv $(dst) themes/$(theme)/build
