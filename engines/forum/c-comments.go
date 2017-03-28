package forum

import (
	"net/http"

	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) myComments(c *gin.Context) (interface{}, error) {

	user := c.MustGet(auth.CurrentUser).(*auth.User)
	isa := c.MustGet(auth.IsAdmin).(bool)
	var comments []Comment
	qry := p.Db.Select([]string{"body", "article_id", "updated_at", "id"})
	if !isa {
		qry = qry.Where("user_id = ?", user.ID)
	}
	if err := qry.Order("updated_at DESC").Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (p *Engine) indexComments(c *gin.Context) (interface{}, error) {

	var total int64
	if err := p.Db.Model(&Comment{}).Count(&total).Error; err != nil {
		return nil, err
	}
	var pag *web.Pagination

	pag = web.NewPagination(c.Request, total)

	var comments []Comment
	if err := p.Db.Select([]string{"id", "type", "body", "article_id", "updated_at"}).
		Limit(pag.Limit()).Offset(pag.Offset()).
		Find(&comments).Error; err != nil {
		return nil, err
	}
	for _, it := range comments {
		pag.Items = append(pag.Items, it)
	}

	return pag, nil
}

type fmCommentAdd struct {
	Body      string `form:"body" binding:"required,max=800"`
	Type      string `form:"type" binding:"required,max=8"`
	ArticleID uint   `form:"articleId" binding:"required"`
}

func (p *Engine) createComment(c *gin.Context) (interface{}, error) {

	user := c.MustGet(auth.CurrentUser).(*auth.User)

	var fm fmCommentAdd
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	cm := Comment{
		Body:      fm.Body,
		Type:      fm.Type,
		ArticleID: fm.ArticleID,
		UserID:    user.ID,
	}

	if err := p.Db.Create(&cm).Error; err != nil {
		return nil, err
	}

	return cm, nil
}

type fmCommentEdit struct {
	Body string `form:"body" binding:"required,max=800"`
	Type string `form:"type" binding:"required,max=8"`
}

func (p *Engine) updateComment(c *gin.Context) (interface{}, error) {

	cm := c.MustGet("comment").(*Comment)

	var fm fmCommentEdit
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	if err := p.Db.Model(cm).Updates(map[string]interface{}{
		"body": fm.Body,
		"type": fm.Type,
	}).Error; err != nil {
		return nil, err
	}

	return cm, nil
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
