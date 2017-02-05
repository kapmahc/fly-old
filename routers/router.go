package routers

import (
	"github.com/astaxie/beego"
	"github.com/kapmahc/fly/engines/site"
)

func init() {
	beego.Include(
		// &auth.Controller{},
		// &forum.Controller{},
		// &reading.Controller{},
		// &shop.Controller{},
		&site.Controller{},
	)
}
