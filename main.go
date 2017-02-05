package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/kapmahc/fly/engines/auth"
	_ "github.com/kapmahc/fly/engines/forum"
	_ "github.com/kapmahc/fly/engines/reading"
	_ "github.com/kapmahc/fly/engines/shop"
	"github.com/kapmahc/fly/web"
)

func main() {
	if err := web.Main(); err != nil {
		log.Fatal(err)
	}
}
