package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/kapmahc/fly/engines/blog"
	_ "github.com/kapmahc/fly/engines/erp"
	_ "github.com/kapmahc/fly/engines/forum"
	_ "github.com/kapmahc/fly/engines/mall"
	_ "github.com/kapmahc/fly/engines/ops/mail"
	_ "github.com/kapmahc/fly/engines/ops/vpn"
	_ "github.com/kapmahc/fly/engines/reading"
	_ "github.com/kapmahc/fly/engines/site"
	"github.com/kapmahc/fly/web"
)

func main() {
	if err := web.Main(); err != nil {
		log.Fatal(err)
	}
}
