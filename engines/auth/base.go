package auth

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/fly/engines/base"
)

const (
	uidKey         = "uid"
	currentUserKey = "currentUser"
)

// UserController currnet-user
type UserController struct {
	base.Controller
}

// Prepare prepare
func (p *UserController) Prepare() {
	p.Controller.Prepare()
	p.setCurrentUser()
}

func (p *UserController) setCurrentUser() {
	uid := p.GetSession(uidKey)
	if uid == nil {
		return
	}

	var user User
	o := orm.NewOrm()
	if err := o.QueryTable(&user).Filter("uid", uid).One(&user); err != nil {
		return
	}
	p.Data[currentUserKey] = &user
}

// CurrentUser current-user
func (p *UserController) CurrentUser() *User {
	user, ok := p.Data[currentUserKey]
	if ok {
		return user.(*User)
	}
	return nil
}
