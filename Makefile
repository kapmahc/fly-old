dst=dist
theme=bootstrap

build: back front
	tar jcvf dist.tar.bz2 dist

back:
	mkdir -pv $(dst)/themes/$(theme)/public
	go build -ldflags "-s -w -X github.com/kapmahc/fly/web.Version=`git rev-parse --short HEAD` -X github.com/kapmahc/fly/web.BuildTime=`date +%FT%T%z`" -o ${dst}/fly main.go
	-cp -rv locales db $(dst)/

front:
	cd frontend && npm run build
	-cp -rv frontend/dist $(dst)/public

init:
	go get -u github.com/kardianos/govendor
	govendor sync
	npm install

clean:
	-rm -rv $(dst) frontend/dist dist.tar.bz2
