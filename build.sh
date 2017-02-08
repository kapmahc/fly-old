#!/bin/sh

#bee pack -ba "-ldflags '-s'"
go build -ldflags "-s -X github.com/kapmahc/fly/engines/base.Version=`git rev-parse --short HEAD` -X github.com/kapmahc/fly/engines/base.BuildTime=`date +%FT%T%z`" -o fly main.go
