package routers

import (
	"github.com/astaxie/beego"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/engines/forum"
	"github.com/kapmahc/fly/engines/reading"
	"github.com/kapmahc/fly/engines/shop"
	"github.com/kapmahc/fly/engines/site"
)

func init() {
	beego.Include(
		&site.Controller{},
	)
	beego.AddNamespace(
		beego.NewNamespace("/users", beego.NSInclude(&auth.Controller{})),
		beego.NewNamespace("/forum", beego.NSInclude(&forum.Controller{})),
		beego.NewNamespace("/reading", beego.NSInclude(&reading.Controller{})),
		beego.NewNamespace("/shop", beego.NSInclude(&shop.Controller{})),
	)
}
