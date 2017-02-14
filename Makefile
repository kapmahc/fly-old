dst=dist

all:
	go build -ldflags "-s -w -X main.version=`git rev-parse --short HEAD` -X main.buildTime=`date +%FT%T%z`" -o $(dst)/fly main.go
	# -cp -rv locales db themes $(dst)/

clean:
	-rm -rv $(dst)
