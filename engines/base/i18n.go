package base

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
)

// Tr translate content to target language
func Tr(lang, format string, args ...interface{}) string {
	key := fmt.Sprintf("locales/%s/%s", lang, format)
	msg := cache.Get(key)
	if msg == nil {
		o := orm.NewOrm()
		var l Locale
		err := o.QueryTable(&l).Filter("lang", lang).Filter("code", format).One(&l, "message")
		if err == nil {
			CachePut(key, l.Message, time.Hour*24)
			msg = l.Message
		}

		if err == orm.ErrNoRows {
			CachePut(key, "", time.Hour*24)
		} else {
			beego.Error(err)
		}
	}

	if msg == nil || len(msg.([]byte)) == 0 {
		return i18n.Tr(lang, format, args...)
	}
	return fmt.Sprintf(msg.(string), args...)
}

func init() {
	beego.AddFuncMap("t", Tr)
	beego.AddFuncMap("fmt", fmt.Sprintf)
	// beego.AddFuncMap("md2ht", func(md string) template.HTML {
	// 	return template.HTML(blackfriday.MarkdownBasic([]byte(md)))
	// })
}
