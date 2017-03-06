package site

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kapmahc/fly/web"
)

func (p *Engine) indexLeavewords(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckAdmin(w, r, true) {
		return
	}
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)

	var items []LeaveWord
	if err := p.Db.Order("created_at DESC").Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["items"] = items
	data["title"] = p.I18n.T(lang, "site.leave-words.index.title")
	p.Ctx.HTML(w, "site/leave-words/index", data)
}

type fmLeaveWord struct {
	Body string `form:"body" validate:"required,max=800"`
	Type string `form:"type" validate:"required,max=8"`
}

func (p *Engine) newLeaveword(w http.ResponseWriter, r *http.Request) {

	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)

	switch r.Method {
	case http.MethodPost:
		var fm fmLeaveWord
		err := p.Ctx.Bind(&fm, r)
		if err == nil {
			err = p.Db.Create(&LeaveWord{Type: fm.Type, Body: fm.Body}).Error
		}
		if err == nil {
			data[web.INFO] = p.I18n.T(lang, "success")
		} else {
			data[web.ERROR] = err.Error()
		}

	}

	data["title"] = p.I18n.T(lang, "site.leave-words.new.title")
	p.Ctx.HTML(w, "site/leave-words/new", data)
}

func (p *Engine) destroyLeaveword(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckAdmin(w, r, true) {
		return
	}

	if err := p.Db.
		Where("id = ?", mux.Vars(r)["id"]).
		Delete(LeaveWord{}).Error; err != nil {
		log.Error(err)
	}
	p.Ctx.JSON(w, web.H{})
}
