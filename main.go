package main

import (
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/kapmahc/fly/engines/admin"
	// _ "github.com/kapmahc/fly/engines/forum"
	// _ "github.com/kapmahc/fly/engines/ops/mail"
	// _ "github.com/kapmahc/fly/engines/ops/vpn"
	// _ "github.com/kapmahc/fly/engines/reading"
	// _ "github.com/kapmahc/fly/engines/shop"
	"github.com/kapmahc/sky"
)

func main() {
	app := sky.New()
	app.Usage = "FLY - A complete open source e-commerce solution by the Go language."
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
