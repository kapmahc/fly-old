package forum

import (
	"net/http"

	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexArticles(c *gin.Context) {
	var total int64
	var pag *web.Pagination
	err := p.Db.Model(&Article{}).Count(&total).Error
	if err == nil {
		pag = web.NewPagination(c.Request, total)
		var articles []Article
		err = p.Db.Select([]string{"id", "title", "summary", "user_id", "updated_at"}).
			Limit(pag.Limit()).Offset(pag.Offset()).
			Find(&articles).Error

		for _, it := range articles {
			pag.Items = append(pag.Items, it)
		}
	}

	web.JSON(c, pag, err)
}

type fmArticle struct {
	Title   string   `form:"title" validate:"required,max=255"`
	Summary string   `form:"summary" validate:"required,max=500"`
	Type    string   `form:"type" validate:"required,max=8"`
	Body    string   `form:"body" validate:"required,max=2000"`
	Tags    []string `form:"tags"`
}

func (p *Engine) createArticle(c *gin.Context) {
	user := c.MustGet(auth.CurrentUser).(*auth.User)
	var fm fmArticle
	err := c.Bind(&fm)
	if err == nil {
		var tags []Tag
		for _, it := range fm.Tags {
			var t Tag
			if err = p.Db.Select([]string{"id"}).Where("id = ?", it).First(&t).Error; err == nil {
				tags = append(tags, t)
			} else {
				break
			}
		}
		a := Article{
			Title:   fm.Title,
			Summary: fm.Summary,
			Body:    fm.Body,
			Type:    fm.Type,
			UserID:  user.ID,
		}
		if err == nil {
			err = p.Db.Create(&a).Error
		}
		if err == nil {
			err = p.Db.Model(a).Association("Tags").Append(tags).Error
		}
	}
	web.JSON(c, nil, err)
}

func (p *Engine) showArticle(c *gin.Context) {
	var a Article
	err := p.Db.Select("id = ?", c.Param("id")).First(&a).Error
	if err == nil {
		err = p.Db.Model(&a).Association("Tags").Find(&a.Tags).Error
	}
	web.JSON(c, a, err)
}

func (p *Engine) updateArticle(c *gin.Context) {
	a := c.MustGet("article").(*Article)
	var fm fmArticle
	err := c.Bind(&fm)
	if err == nil {
		var tags []Tag
		for _, it := range fm.Tags {
			var t Tag
			if err = p.Db.Select([]string{"id"}).Where("id = ?", it).First(&t).Error; err == nil {
				tags = append(tags, t)
			} else {
				break
			}
		}
		if err == nil {
			err = p.Db.Where("id = ?", a.ID).Updates(map[string]interface{}{
				"title":   fm.Title,
				"summary": fm.Summary,
				"body":    fm.Body,
				"type":    fm.Type,
			}).Error
		}
		if err == nil {
			err = p.Db.Model(a).Association("Tags").Replace(tags).Error
		}
	}
	web.JSON(c, nil, err)
}

func (p *Engine) destroyArticle(c *gin.Context) {
	a := c.MustGet("article").(*Article)
	err := p.Db.Model(a).Association("Tags").Clear().Error
	if err == nil {
		err = p.Db.Delete(a).Error
	}
	web.JSON(c, nil, err)
}

func (p *Engine) canEditArticle(c *gin.Context) {
	user := c.MustGet(auth.CurrentUser).(*auth.User)

	var a Article
	err := p.Db.Where("id = ?", c.Param("id")).First(&a).Error
	if err == nil {
		if user.ID == a.UserID || c.MustGet(auth.IsAdmin).(bool) {
			c.Set("article", &a)
			c.Next()
			return
		}
		c.AbortWithStatus(http.StatusForbidden)
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}

}
