package site

import (
	"net/http"
	"time"

	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/sky"
	"github.com/kapmahc/sky/widgets"
)

func (p *Engine) getInstall(c *sky.Context) error {
	data := c.Get(sky.DATA).(sky.H)
	lang := c.Get(sky.LOCALE).(string)

	data["form"] = widgets.NewForm(
		c.Request,
		lang,
		p.Layout.URLFor("site.install"),
		p.Layout.URLFor("auth.users.sign-in"),
		p.I18n.T(lang, "site.install.title"),
		widgets.NewTextField("title", p.I18n.T(lang, "site.attributes.title"), ""),
		widgets.NewTextField("subTitle", p.I18n.T(lang, "site.attributes.subTitle"), ""),
		widgets.NewEmailField("email", p.I18n.T(lang, "attributes.email"), ""),
		widgets.NewPasswordField("password", p.I18n.T(lang, "attributes.password"), p.I18n.T(lang, "helpers.password")),
		widgets.NewPasswordField("passwordConfirmation", p.I18n.T(lang, "attributes.passwordConfirmation"), p.I18n.T(lang, "helpers.passwordConfirmation")),
	)
	c.HTML(http.StatusOK, "form", data)
	return nil
}

type fmInstall struct {
	Title                string `form:"title" validate:"required"`
	SubTitle             string `form:"subTitle" validate:"required"`
	Email                string `form:"email" validate:"required,email"`
	Password             string `form:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

func (p *Engine) postInstall(c *sky.Context) error {
	var fm fmInstall
	if err := c.Bind(&fm); err != nil {
		return err
	}
	lang := c.Get(sky.LOCALE).(string)
	// ------------
	p.I18nStore.Set(lang, "site.title", fm.Title, true)
	p.I18nStore.Set(lang, "site.subTitle", fm.SubTitle, true)
	// -------------
	user, err := p.Dao.AddEmailUser("root", fm.Email, fm.Password)
	if err != nil {
		return err
	}
	for _, r := range []string{auth.RoleAdmin, auth.RoleRoot} {
		role, er := p.Dao.Role(r, auth.DefaultResourceType, auth.DefaultResourceID)
		if er != nil {
			return er
		}
		er = p.Dao.Allow(role.ID, user.ID, 50, 0, 0)
		if er != nil {
			return er
		}
	}
	if err = p.Db.Model(user).UpdateColumn("confirmed_at", time.Now()).Error; err != nil {
		return err
	}
	// ---------------
	c.JSON(http.StatusOK, sky.H{})
	return nil
}

func (p *Engine) mustDatabaseEmpty(c *sky.Context) error {
	lang := c.Get(sky.LOCALE).(string)
	var count int
	if err := p.Db.Model(&auth.User{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return p.I18n.E(lang, "errors.forbidden")
	}
	return c.Next()
}
