package main

import (
	"fmt"
	"log"

	"github.com/kapmahc/fly/web"
)

var (
	version   string
	buildTime string
)

func main() {
	app := web.New()
	if err := app.Main(
		"FLY - A complete open source e-commerce solution by the Go language.",
		fmt.Sprintf("%s(%s)", version, buildTime),
	); err != nil {
		log.Fatal(err)
	}
}
