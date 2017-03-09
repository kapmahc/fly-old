package forum

import (
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
)

func (p *Engine) indexArticles(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))

	var total int64
	if err := p.Db.Model(&Article{}).Count(&total).Error; err != nil {
		log.Error(err)
	}
	pag := web.NewPagination(
		p.Ctx.URLFor("reading.books.index"),
		int64(page), int64(size), total,
	)

	var articles []Article
	if err := p.Db.Select([]string{"id", "title", "summary", "user_id", "updated_at"}).
		Limit(pag.Limit()).Offset(pag.Offset()).
		Find(&articles).Error; err != nil {
		log.Error(err)
	}
	for _, it := range articles {
		pag.Items = append(pag.Items, it)
	}

	data["articles"] = pag
	data["title"] = p.I18n.T(lang, "forum.articles.index.title")
	p.Ctx.HTML(w, "forum/articles/index", data)
}

type fmArticle struct {
	Title   string   `form:"type" validate:"required,max=255"`
	Summary string   `form:"type" validate:"required,max=500"`
	Type    string   `form:"type" validate:"required,max=8"`
	Body    string   `form:"body" validate:"required,max=2000"`
	Tags    []string `form:"tags"`
}

func (p *Engine) newArticle(w http.ResponseWriter, r *http.Request) {
	if !p.Session.CheckSignIn(w, r, true) {
		return
	}
	user := p.Session.CurrentUser(r)
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)

	switch r.Method {
	case http.MethodPost:
		var fm fmArticle
		err := p.Ctx.Bind(&fm, r)
		if err == nil {
			var tags []Tag
			for _, it := range fm.Tags {
				var t Tag
				if err = p.Db.Select([]string{"id"}).Where("id = ?", it).First(&t).Error; err == nil {
					tags = append(tags, t)
				} else {
					log.Error(err)
				}
			}
			a := Article{
				Title:   fm.Title,
				Summary: fm.Summary,
				Body:    fm.Body,
				Type:    fm.Type,
				UserID:  user.ID,
				Tags:    tags,
			}
			if err = p.Db.Create(&a).Error; err != nil {
				log.Error(err)
			}
		}
		if err == nil {
			data[web.INFO] = p.I18n.T(lang, "success")
		} else {
			data[web.ERROR] = err.Error()
		}
	}

	data["title"] = p.I18n.T(lang, "forum.articles.new")
	var tags []Tag
	if err := p.Db.Select([]string{"id", "name"}).Find(&tags).Error; err != nil {
		log.Error(err)
	}
	data["tags"] = tags
	p.Ctx.HTML(w, "forum/articles/new", data)
}

func (p *Engine) showArticle(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value(web.DATA).(web.H)

	var a Article
	err := p.Db.Select("id = ?", mux.Vars(r)["id"]).First(&a).Error
	if !p.Ctx.Check(w, err) {
		return
	}
	data["title"] = a.Title
	data["article"] = a
	p.Ctx.HTML(w, "forum/articles/show", data)
}

func (p *Engine) editArticle(w http.ResponseWriter, r *http.Request) {
	a := p.canEditArticle(w, r)
	if a == nil {
		return
	}
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)

	switch r.Method {
	case http.MethodPost:
		var fm fmArticle
		err := p.Ctx.Bind(&fm, r)
		if err == nil {
			var tags []Tag
			for _, it := range fm.Tags {
				var t Tag
				if err = p.Db.Select([]string{"id"}).Where("id = ?", it).First(&t).Error; err == nil {
					tags = append(tags, t)
				} else {
					log.Error(err)
				}
			}
			if err = p.Db.Where("id = ?", a.ID).Updates(map[string]interface{}{
				"title":   fm.Title,
				"summary": fm.Summary,
				"body":    fm.Body,
				"type":    fm.Type,
			}).Error; err != nil {
				log.Error(err)
			}
			if err = p.Db.Model(a).Association("Tags").Replace(tags).Error; err != nil {
				log.Error(err)
			}
		}
		if err == nil {
			p.Ctx.Redirect(w, r, p.Ctx.URLFor("forum.articles.show", "id", a.ID))
			return
		}
		a.Title = fm.Title
		a.Type = fm.Type
		a.Body = fm.Body
		a.Summary = fm.Summary
		data[web.ERROR] = err.Error()
	}

	data["title"] = p.I18n.T(lang, "buttons.edit")
	data["article"] = a
	var tags []Tag
	if err := p.Db.Select([]string{"id", "name"}).Find(&tags).Error; err != nil {
		log.Error(err)
	}
	data["tags"] = tags
	p.Ctx.HTML(w, "forum/articles/edit", data)
}

func (p *Engine) destroyArticle(w http.ResponseWriter, r *http.Request) {
	a := p.canEditArticle(w, r)
	if a == nil {
		return
	}
	err := p.Db.Delete(a).Error
	if !p.Ctx.Check(w, err) {
		return
	}
	p.Ctx.JSON(w, web.H{})
}

func (p *Engine) canEditArticle(w http.ResponseWriter, r *http.Request) *Article {
	if !p.Session.CheckSignIn(w, r, true) {
		return nil
	}
	var a Article
	err := p.Db.Where("id = ?", mux.Vars(r)["id"]).First(&a).Error
	if !p.Ctx.Check(w, err) {
		return nil
	}
	user := p.Session.CurrentUser(r)
	if user.ID != a.UserID && !p.Dao.Is(user.ID, auth.RoleAdmin) {
		return nil
	}
	return &a
}
