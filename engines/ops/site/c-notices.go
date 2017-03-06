package site

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kapmahc/fly/web"
)

func (p *Engine) indexNotices(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)

	var items []Notice
	if err := p.Db.Order("updated_at DESC").Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["items"] = items
	data["title"] = p.I18n.T(lang, "site.notices.index.title")
	p.Ctx.HTML(w, "site/notices/index", data)
}

type fmNotice struct {
	Body string `form:"body" validate:"required,max=800"`
	Type string `form:"type" validate:"required,max=8"`
}

func (p *Engine) newNotice(w http.ResponseWriter, r *http.Request) {

	if !p.Session.CheckAdmin(w, r, true) {
		return
	}
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)

	switch r.Method {
	case http.MethodPost:
		var fm fmNotice
		err := p.Ctx.Bind(&fm, r)
		if err == nil {
			err = p.Db.Create(&Notice{Type: fm.Type, Body: fm.Body}).Error
		}
		if err == nil {
			data[web.INFO] = p.I18n.T(lang, "success")
		} else {
			data[web.ERROR] = err.Error()
		}

	}

	data["title"] = p.I18n.T(lang, "buttons.new")
	p.Ctx.HTML(w, "site/notices/new", data)
}

func (p *Engine) editNotice(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckAdmin(w, r, true) {
		return
	}
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	id := mux.Vars(r)["id"]

	switch r.Method {
	case http.MethodPost:
		var fm fmNotice
		err := p.Ctx.Bind(&fm, r)
		if err == nil {
			err = p.Db.Model(&Notice{}).Where("id = ?", id).Updates(map[string]interface{}{
				"body": fm.Body,
				"type": fm.Type,
			}).Error
		}
		if err == nil {
			data[web.INFO] = p.I18n.T(lang, "success")
		} else {
			data[web.ERROR] = err.Error()
		}

	}

	data["title"] = p.I18n.T(lang, "buttons.edit")

	var item Notice
	if err := p.Db.Where("id = ?", id).First(&item).Error; err != nil {
		log.Error(err)
	}
	data["item"] = item

	p.Ctx.HTML(w, "site/notices/edit", data)
}

func (p *Engine) destroyNotice(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckAdmin(w, r, true) {
		return
	}

	if err := p.Db.
		Where("id = ?", mux.Vars(r)["id"]).
		Delete(Notice{}).Error; err != nil {
		log.Error(err)
	}
	p.Ctx.JSON(w, web.H{})
}
