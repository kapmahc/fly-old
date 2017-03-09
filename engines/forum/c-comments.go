package forum

import (
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
)

func (p *Engine) indexComments(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))

	var total int64
	if err := p.Db.Model(&Comment{}).Count(&total).Error; err != nil {
		log.Error(err)
	}
	pag := web.NewPagination(
		p.Ctx.URLFor("reading.books.index"),
		int64(page), int64(size), total,
	)

	var comments []Comment
	if err := p.Db.Select([]string{"id", "type", "body", "article_id", "updated_at"}).
		Limit(pag.Limit()).Offset(pag.Offset()).
		Find(&comments).Error; err != nil {
		log.Error(err)
	}
	for _, it := range comments {
		pag.Items = append(pag.Items, it)
	}

	data["comments"] = pag
	data["title"] = p.I18n.T(lang, "forum.comments.index.title")
	p.Ctx.HTML(w, "forum/comments/index", data)
}

type fmComment struct {
	Body      string `form:"body" validate:"required,max=800"`
	Type      string `form:"type" validate:"required,max=8"`
	ArticleID uint   `form:"article_id" validate:"required"`
}

func (p *Engine) newComment(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckSignIn(w, r, true) {
		return
	}
	user := p.Session.CurrentUser(r)
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)

	switch r.Method {
	case http.MethodGet:
		data["article_id"] = r.URL.Query().Get("article_id")
	case http.MethodPost:
		var fm fmComment
		err := p.Ctx.Bind(&fm, r)
		if err == nil {
			err = p.Db.Create(&Comment{
				Body:      fm.Body,
				Type:      fm.Type,
				ArticleID: fm.ArticleID,
				UserID:    user.ID,
			}).Error
		}
		if err == nil {
			p.Ctx.Redirect(w, r, p.Ctx.URLFor("forum.articles.show", "id", fm.ArticleID))
			return
		}
		data["article_id"] = fm.ArticleID
		data[web.ERROR] = err.Error()
	}

	data["title"] = p.I18n.T(lang, "buttons.new")
	p.Ctx.HTML(w, "forum/comments/new", data)
}

func (p *Engine) editComment(w http.ResponseWriter, r *http.Request) {
	comment := p.canEditComment(w, r)
	if comment == nil {
		return
	}
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)

	switch r.Method {
	case http.MethodPost:
		var fm fmComment
		err := p.Ctx.Bind(&fm, r)
		if err == nil {
			err = p.Db.Where("id = ?", mux.Vars(r)["id"]).Updates(map[string]interface{}{
				"body": fm.Body,
				"type": fm.Type,
			}).Error
		}
		if err == nil {
			p.Ctx.Redirect(w, r, p.Ctx.URLFor("forum.articles.show", "id", fm.ArticleID))
			return
		}
		data[web.ERROR] = err.Error()
	}

	data["comment"] = comment
	data["title"] = p.I18n.T(lang, "buttons.edit")
	p.Ctx.HTML(w, "forum/comments/edit", data)
}

func (p *Engine) destroyComment(w http.ResponseWriter, r *http.Request) {
	comment := p.canEditComment(w, r)
	if comment == nil {
		return
	}
	err := p.Db.Delete(comment).Error
	if !p.Ctx.Check(w, err) {
		return
	}
	p.Ctx.JSON(w, web.H{})
}

func (p *Engine) canEditComment(w http.ResponseWriter, r *http.Request) *Comment {
	if !p.Session.CheckSignIn(w, r, true) {
		return nil
	}
	var comment Comment
	err := p.Db.Where("id = ?", mux.Vars(r)["id"]).First(&comment).Error
	if !p.Ctx.Check(w, err) {
		return nil
	}
	user := p.Session.CurrentUser(r)
	if user.ID != comment.UserID && !p.Dao.Is(user.ID, auth.RoleAdmin) {
		return nil
	}
	return &comment
}
