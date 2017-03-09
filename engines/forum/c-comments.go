package forum

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/kapmahc/fly/web"
)

func (p *Engine) indexComments(w http.ResponseWriter, r *http.Request) {
	// TODO pager
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	var comments []Comment
	if err := p.Db.
		Select([]string{"type", "body"}).
		Order("updated_at DESC").
		Find(&comments).Error; err != nil {
		log.Error(err)
	}
	data["comments"] = comments
	data["title"] = p.I18n.T(lang, "forum.comments.index.title")
	p.Ctx.HTML(w, "forum/comments/index", data)
}

type fmCommentAdd struct {
	Body      string `form:"body" validate:"required,max=800"`
	Type      string `form:"type" validate:"required,max=8"`
	ArticleID uint   `form:"article_id" validate:"required"`
}

func (p *Engine) newComment(w http.ResponseWriter, r *http.Request) {
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
			data[web.INFO] = p.I18n.T(lang, "success")
		} else {
			data[web.ERROR] = err.Error()
		}

	}

	data["title"] = p.I18n.T(lang, "buttons.new")
	p.Ctx.HTML(w, "forum/comments/new", data)
}
func (p *Engine) editComment(w http.ResponseWriter, r *http.Request) {

}
func (p *Engine) destroyComment(w http.ResponseWriter, r *http.Request) {

}
