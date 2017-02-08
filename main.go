package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kapmahc/fly/engines/site"
	_ "github.com/kapmahc/fly/routers"
	_ "github.com/lib/pq"
)

func main() {
	// database
	orm.Debug = true
	drv := beego.AppConfig.String("databasedriver")
	dsn := beego.AppConfig.String("databaseurl")
	orm.RegisterDataBase(
		"default", drv, dsn,
		beego.AppConfig.DefaultInt("databasemaxidle", 2),
		beego.AppConfig.DefaultInt("databasemaxconn", 18),
	)

	if err := site.OnMigrate(drv, dsn); err != nil {
		beego.Error(err)
		return
	}

	// worker
	// go base.Worker()

	// web
	beego.Run()
}
