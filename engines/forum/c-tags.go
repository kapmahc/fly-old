package forum

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kapmahc/fly/web"
)

func (p *Engine) indexTags(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	var tags []Tag
	if err := p.Db.Select([]string{"name", "id"}).
		Find(&tags).Error; err != nil {
		log.Error(err)
	}
	data["tags"] = tags
	data["title"] = p.I18n.T(lang, "forum.tags.index.title")
	p.Ctx.HTML(w, "forum/tags/index", data)
}

type fmTag struct {
	Name string `form:"name" validate:"required,max=255"`
}

func (p *Engine) newTag(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckAdmin(w, r, true) {
		return
	}
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)

	switch r.Method {
	case http.MethodPost:
		var fm fmTag
		err := p.Ctx.Bind(&fm, r)
		if err == nil {
			err = p.Db.Create(&Tag{Name: fm.Name}).Error
		}
		if err == nil {
			p.Ctx.Redirect(w, r, p.Ctx.URLFor("forum.dashboard.tags"))
			return
		}
		data[web.ERROR] = err.Error()
	}

	data["title"] = p.I18n.T(lang, "buttons.new")
	p.Ctx.HTML(w, "forum/tags/new", data)
}

func (p *Engine) showTag(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value(web.DATA).(web.H)
	var tag Tag
	err := p.Db.Where("id = ?", mux.Vars(r)["id"]).First(&tag).Error
	if !p.Ctx.Check(w, err) {
		return
	}
	err = p.Db.Model(&tag).Association("Articles").Find(&tag.Articles).Error
	if !p.Ctx.Check(w, err) {
		return
	}
	data["tag"] = tag
	data["title"] = tag.Name
	p.Ctx.HTML(w, "forum/tags/show", data)
}

func (p *Engine) editTag(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckAdmin(w, r, true) {
		return
	}
	id := mux.Vars(r)["id"]
	var tag Tag
	err := p.Db.Where("id = ?", id).First(&tag).Error
	if !p.Ctx.Check(w, err) {
		return
	}

	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)

	switch r.Method {
	case http.MethodPost:
		var fm fmTag
		err := p.Ctx.Bind(&fm, r)
		if err == nil {
			err = p.Db.Model(&tag).Update("name", fm.Name).Error
		}
		if err == nil {
			p.Ctx.Redirect(w, r, p.Ctx.URLFor("forum.dashboard.tags"))
			return
		}
		data[web.ERROR] = err.Error()
	}

	data["title"] = p.I18n.T(lang, "buttons.edit")
	data["tag"] = tag
	p.Ctx.HTML(w, "forum/tags/edit", data)
}

func (p *Engine) destroyTag(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckAdmin(w, r, true) {
		return
	}
	id := mux.Vars(r)["id"]
	var tag Tag
	err := p.Db.Where("id = ?", id).First(&tag).Error
	if !p.Ctx.Check(w, err) {
		return
	}
	err = p.Db.Model(&tag).Association("Articles").Clear().Error
	if !p.Ctx.Check(w, err) {
		return
	}
	err = p.Db.Delete(&tag).Error
	if !p.Ctx.Check(w, err) {
		return
	}
	p.Ctx.JSON(w, web.H{})
}
