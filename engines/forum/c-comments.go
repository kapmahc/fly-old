package forum

import (
	"net/http"

	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexComments(c *gin.Context) {
	var total int64
	err := p.Db.Model(&Comment{}).Count(&total).Error
	var pag *web.Pagination
	if err == nil {
		pag = web.NewPagination(c.Request, total)

		var comments []Comment
		err = p.Db.Select([]string{"id", "type", "body", "article_id", "updated_at"}).
			Limit(pag.Limit()).Offset(pag.Offset()).
			Find(&comments).Error
		for _, it := range comments {
			pag.Items = append(pag.Items, it)
		}
	}

	web.JSON(c, pag, err)
}

type fmComment struct {
	Body      string `form:"body" binding:"required,max=800"`
	Type      string `form:"type" binding:"required,max=8"`
	ArticleID uint   `form:"article_id" binding:"required"`
}

func (p *Engine) createComment(c *gin.Context) {
	user := c.MustGet(auth.CurrentUser).(*auth.User)

	var fm fmComment
	err := c.Bind(&fm)
	if err == nil {
		err = p.Db.Create(&Comment{
			Body:      fm.Body,
			Type:      fm.Type,
			ArticleID: fm.ArticleID,
			UserID:    user.ID,
		}).Error
	}

	web.JSON(c, nil, err)
}

func (p *Engine) updateComment(c *gin.Context) {
	comment := c.MustGet("comment").(*Comment)

	var fm fmComment
	err := c.Bind(&fm)
	if err == nil {
		err = p.Db.Where("id = ?", comment.ID).Updates(map[string]interface{}{
			"body": fm.Body,
			"type": fm.Type,
		}).Error
	}
	web.JSON(c, nil, err)
}

func (p *Engine) destroyComment(c *gin.Context) {
	comment := c.MustGet("comment").(*Comment)
	err := p.Db.Delete(comment).Error
	web.JSON(c, nil, err)
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
