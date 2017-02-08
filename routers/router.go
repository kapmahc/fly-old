package routers

import (
	"os"
	"path/filepath"

	"golang.org/x/text/language"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/engines/forum"
	"github.com/kapmahc/fly/engines/posts"
	"github.com/kapmahc/fly/engines/reading"
	"github.com/kapmahc/fly/engines/shop"
	"github.com/kapmahc/fly/engines/site"
)

func init() {
	// i18n
	if err := filepath.Walk(filepath.Join("conf", "locales"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		const ext = ".ini"
		if info.IsDir() {
			return nil
		}
		name := info.Name()
		if filepath.Ext(name) != ext {
			beego.Warn("ingnore file", name)
			return nil
		}
		lang, err := language.Parse(name[:len(name)-len(ext)])
		if err != nil {
			return err
		}
		beego.Info("find locale", lang)
		return i18n.SetMessage(lang.String(), path)
	}); err != nil {
		beego.Error(err)
	}

	// controllers
	beego.Include(
		&site.Controller{},
	)
	beego.AddNamespace(
		beego.NewNamespace("/users", beego.NSInclude(&auth.Controller{})),
		beego.NewNamespace("/forum", beego.NSInclude(&forum.Controller{})),
		beego.NewNamespace("/reading", beego.NSInclude(&reading.Controller{})),
		beego.NewNamespace("/shop", beego.NSInclude(&shop.Controller{})),
		beego.NewNamespace("/posts", beego.NSInclude(&posts.Controller{})),
	)
}
