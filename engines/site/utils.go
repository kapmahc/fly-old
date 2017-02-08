package site

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/fly/engines/auth"
)

func (p *Controller) dbEmpty() bool {
	o := orm.NewOrm()
	count, err := o.QueryTable(&auth.User{}).Count()
	if err != nil {
		beego.Error(err)
		p.Abort(http.StatusInternalServerError)
	}
	return count == 0
}
