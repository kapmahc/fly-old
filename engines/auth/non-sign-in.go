package auth

import (
	"encoding/base64"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/fly/engines/base"
)

// GetSignIn sign in
// @router /sign-in [get]
func (p *Controller) GetSignIn() {
	p.HTML(p.T("auth.sign-in.title"), "auth/sign-in.html")
}

type fmSignIn struct {
	Email    string `form:"email" valid:"Email"`
	Password string `form:"password" valid:"Required"`
}

// PostSignIn sign in
// @router /sign-in [post]
func (p *Controller) PostSignIn() {
	var fm fmSignIn
	ok, fsh := p.ParseForm(&fm)
	if ok {
		user, err := p.signIn(fm.Email, fm.Password)

		if err == nil {
			p.SetSession(uidKey, user.UID)
			p.Redirect(p.URLFor("site.Controller.GetHome"), http.StatusFound)
			return
		}
		fsh.Error(err.Error())
	}
	fsh.Store(&p.Controller.Controller)
	p.Redirect(p.URLFor("auth.Controller.GetSignIn"), http.StatusFound)
}

func (p *Controller) signIn(email, password string) (*User, error) {
	var user User
	o := orm.NewOrm()
	if err := o.QueryTable(&user).
		Filter("provider_id", email).
		Filter("provider_type", UserTypeEmail).
		One(&user); err != nil {
		return nil, err
	}

	pwd, err := base64.StdEncoding.DecodeString(user.Password)

	if err != nil {
		return nil, err
	}

	if !base.HmacChk([]byte(password), pwd) {
		return nil, p.E("auth.email-password-not-match")
	}

	if !user.IsConfirm() {
		return nil, p.E("auth.user-not-confirm")
	}
	if user.IsLock() {
		return nil, p.E("auth.user-is-lock")
	}
	if err := SetSignIn(&user, p.Ctx.Input.IP()); err != nil {
		return nil, err
	}
	return &user, nil
}
