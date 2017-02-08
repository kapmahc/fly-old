package site

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/kapmahc/fly/engines/auth"
)

// GetInstall install
// @router /install [get]
func (p *Controller) GetInstall() {
	if !p.dbEmpty() {
		p.Abort(http.StatusNotFound)
	}
	p.HTML(p.T("site.install.title"), "site/install.html")
}

type fmInstall struct {
	Name                 string `form:"name" valid:"Required; MaxSize(32)"`
	Email                string `form:"email" valid:"Email; MaxSize(100)"`
	Password             string `form:"password" valid:"MaxSize(32); MinSize(6)"`
	PasswordConfirmation string `form:"passwordConfirmation"`
}

func (p *fmInstall) Valid(v *validation.Validation) {
	if p.Password != p.PasswordConfirmation {
		v.SetError("PasswordConfirmation", "passwords must match")
	}
}

// PostInstall install
// @router /install [post]
func (p *Controller) PostInstall() {
	if !p.dbEmpty() {
		p.Abort(http.StatusNotFound)
	}

	var fm fmInstall
	ok, fsh := p.ParseForm(&fm)
	if ok {
		user, err := auth.AddEmailUser(fm.Name, fm.Email, fm.Password)
		if err == nil {
			err = auth.ConfirmUser(user)
		}
		if err == nil {
			for _, r := range []string{auth.RoleAdmin, auth.RoleRoot} {
				if err = auth.Allow(user, r, 100, 0, 0); err != nil {
					break
				}
			}
		}
		if err == nil {
			p.Redirect(p.URLFor("site.Controller.GetHome"), http.StatusFound)
			return
		}
		fsh.Error(err.Error())
	}
	fsh.Store(&p.Controller.Controller)
	p.Redirect(p.URLFor("site.Controller.GetInstall"), http.StatusFound)
}
