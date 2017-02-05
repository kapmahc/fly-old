package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/kapmahc/fly/routers"
	_ "github.com/lib/pq"
)

func main() {
	beego.Run()
}
