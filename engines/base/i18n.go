package base

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
)

// Tr translate content to target language
func Tr(lang, format string, args ...interface{}) string {
	o := orm.NewOrm()
	var l Locale
	err := o.QueryTable(&l).Filter("lang", lang).Filter("code", format).One(&l, "message")
	if err == nil {
		return fmt.Sprintf(l.Message, args...)
	}
	if err != orm.ErrNoRows {
		beego.Error(err)
	}
	return i18n.Tr(lang, format, args...)
}

func init() {
	beego.AddFuncMap("t", Tr)
	beego.AddFuncMap("fmt", fmt.Sprintf)
	beego.AddFuncMap("eq", func(arg1, arg2 interface{}) bool {
		return arg1 == arg2
	})
}
