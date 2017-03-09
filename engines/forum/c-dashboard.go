package forum

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
)

func (p *Engine) dashboardComments(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckSignIn(w, r, true) {
		return
	}
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)

	var comments []Comment
	qry := p.Db.
		Select([]string{"id", "article_id", "body"}).
		Order("updated_at DESC")
	user := p.Session.CurrentUser(r)
	if !p.Dao.Is(user.ID, auth.RoleAdmin) {
		qry = qry.Where("user_id = ?", user.ID)
	}
	if err := qry.Find(&comments).Error; err != nil {
		log.Error(err)
	}
	data["comments"] = comments
	data["title"] = p.I18n.T(lang, "forum.dashboard.comments.title")
	p.Ctx.HTML(w, "forum.dashboard.comments", data)
}

func (p *Engine) dashboardArticles(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckSignIn(w, r, true) {
		return
	}
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)

	var articles []Article
	qry := p.Db.
		Select([]string{"id", "title"}).
		Order("updated_at DESC")
	user := p.Session.CurrentUser(r)
	if !p.Dao.Is(user.ID, auth.RoleAdmin) {
		qry = qry.Where("user_id = ?", user.ID)
	}

	if err := qry.Find(&articles).Error; err != nil {
		log.Error(err)
	}
	data["articles"] = articles
	data["title"] = p.I18n.T(lang, "forum.dashboard.articles.title")
	p.Ctx.HTML(w, "forum/dashboard/articles", data)
}

func (p *Engine) dashboardTags(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckAdmin(w, r, true) {
		return
	}
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)

	var tags []Tag
	if err := p.Db.
		Select([]string{"id", "name"}).
		Order("updated_at DESC").Find(&tags).Error; err != nil {
		log.Error(err)
	}
	data["tags"] = tags
	data["title"] = p.I18n.T(lang, "forum.dashboard.tags.title")
	p.Ctx.HTML(w, "forum/dashboard/tags", data)
}
