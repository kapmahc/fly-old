package forum

import (
	"fmt"
	"net/http"

	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) myComments(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "forum.comments.my.title")
	tpl := "forum-comments-my"

	user := c.MustGet(auth.CurrentUser).(*auth.User)
	isa := c.MustGet(auth.IsAdmin).(bool)
	var comments []Comment
	qry := p.Db.Select([]string{"body", "article_id", "updated_at", "id"})
	if !isa {
		qry = qry.Where("user_id = ?", user.ID)
	}
	if err := qry.Order("updated_at DESC").Find(&comments).Error; err != nil {
		return tpl, err
	}
	data["items"] = comments
	return tpl, nil
}

func (p *Engine) indexComments(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "forum.comments.index.title")
	tpl := "forum-comments-index"
	var total int64
	if err := p.Db.Model(&Comment{}).Count(&total).Error; err != nil {
		return tpl, err
	}
	var pag *web.Pagination

	pag = web.NewPagination(c.Request, total)

	var comments []Comment
	if err := p.Db.Select([]string{"id", "type", "body", "article_id", "updated_at"}).
		Limit(pag.Limit()).Offset(pag.Offset()).
		Find(&comments).Error; err != nil {
		return tpl, err
	}
	for _, it := range comments {
		pag.Items = append(pag.Items, it)
	}
	data["pager"] = pag
	return tpl, nil
}

type fmCommentAdd struct {
	Body      string `form:"body" binding:"required,max=800"`
	Type      string `form:"type" binding:"required,max=8"`
	ArticleID uint   `form:"articleId" binding:"required"`
}

func (p *Engine) createComment(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.new")
	tpl := "forum-comments-new"
	user := c.MustGet(auth.CurrentUser).(*auth.User)
	if c.Request.Method == http.MethodPost {
		var fm fmCommentAdd
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		cm := Comment{
			Body:      fm.Body,
			Type:      fm.Type,
			ArticleID: fm.ArticleID,
			UserID:    user.ID,
		}

		if err := p.Db.Create(&cm).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, fmt.Sprintf("/forum/articles/show/%d", cm.ArticleID))
		return "", nil
	}
	return tpl, nil
}

type fmCommentEdit struct {
	Body string `form:"body" binding:"required,max=800"`
	Type string `form:"type" binding:"required,max=8"`
}

func (p *Engine) updateComment(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "forum.comments.edit.title")
	tpl := "forum-comments-edit"
	cm := c.MustGet("comment").(*Comment)

	switch c.Request.Method {
	case http.MethodPost:
		var fm fmCommentEdit
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		if err := p.Db.Model(cm).Updates(map[string]interface{}{
			"body": fm.Body,
			"type": fm.Type,
		}).Error; err != nil {
			return tpl, err
		}

		c.Redirect(http.StatusFound, fmt.Sprintf("/forum/articles/show/%d", cm.ArticleID))
		return "", nil
	}
	data["item"] = cm
	return tpl, nil
}

func (p *Engine) destroyComment(c *gin.Context) (interface{}, error) {
	comment := c.MustGet("comment").(*Comment)
	err := p.Db.Delete(comment).Error
	return gin.H{}, err
}

func (p *Engine) canEditComment(c *gin.Context) {
	user := c.MustGet(auth.CurrentUser).(*auth.User)

	var o Comment
	err := p.Db.Where("id = ?", c.Param("id")).First(&o).Error
	if err == nil {
		if user.ID == o.UserID || c.MustGet(auth.IsAdmin).(bool) {
			c.Set("comment", &o)
			c.Next()
			return
		}
		c.AbortWithStatus(http.StatusForbidden)
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}
