package forum

import (
	"net/http"

	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) myArticles(c *gin.Context) (interface{}, error) {
	user := c.MustGet(auth.CurrentUser).(*auth.User)
	isa := c.MustGet(auth.IsAdmin).(bool)
	var articles []Article
	qry := p.Db.Select([]string{"title", "updated_at", "id"})
	if !isa {
		qry = qry.Where("user_id = ?", user.ID)
	}
	err := qry.Order("updated_at DESC").Find(&articles).Error

	return articles, err
}

func (p *Engine) indexArticles(c *gin.Context) (interface{}, error) {
	var total int64
	var pag *web.Pagination
	if err := p.Db.Model(&Article{}).Count(&total).Error; err != nil {
		return nil, err
	}

	pag = web.NewPagination(c.Request, total)
	var articles []Article
	if err := p.Db.Select([]string{"id", "title", "summary", "user_id", "updated_at"}).
		Limit(pag.Limit()).Offset(pag.Offset()).
		Find(&articles).Error; err != nil {
		return nil, err
	}

	for _, it := range articles {
		pag.Items = append(pag.Items, it)
	}

	return pag, nil
}

type fmArticle struct {
	Title   string   `form:"title" binding:"required,max=255"`
	Summary string   `form:"summary" binding:"required,max=500"`
	Type    string   `form:"type" binding:"required,max=8"`
	Body    string   `form:"body" binding:"required,max=2000"`
	Tags    []string `form:"tags"`
}

func (p *Engine) createArticle(c *gin.Context) (interface{}, error) {
	user := c.MustGet(auth.CurrentUser).(*auth.User)
	var fm fmArticle
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var tags []Tag
	for _, it := range fm.Tags {
		var t Tag
		if err := p.Db.Select([]string{"id"}).Where("id = ?", it).First(&t).Error; err == nil {
			tags = append(tags, t)
		} else {
			return nil, err
		}
	}
	a := Article{
		Title:   fm.Title,
		Summary: fm.Summary,
		Body:    fm.Body,
		Type:    fm.Type,
		UserID:  user.ID,
	}

	if err := p.Db.Create(&a).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Model(&a).Association("Tags").Append(tags).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func (p *Engine) showArticle(c *gin.Context) (interface{}, error) {
	var a Article
	if err := p.Db.Where("id = ?", c.Param("id")).First(&a).Error; err != nil {
		return nil, err
	}

	if err := p.Db.Model(&a).Related(&a.Comments).Error; err != nil {
		return nil, err
	}

	if err := p.Db.Model(&a).Association("Tags").Find(&a.Tags).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func (p *Engine) updateArticle(c *gin.Context) (interface{}, error) {
	a := c.MustGet("article").(*Article)
	var fm fmArticle
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	var tags []Tag
	for _, it := range fm.Tags {
		var t Tag
		if err := p.Db.Select([]string{"id"}).Where("id = ?", it).First(&t).Error; err == nil {
			tags = append(tags, t)
		} else {
			return nil, err
		}
	}

	if err := p.Db.Model(a).Updates(map[string]interface{}{
		"title":   fm.Title,
		"summary": fm.Summary,
		"body":    fm.Body,
		"type":    fm.Type,
	}).Error; err != nil {
		return nil, err
	}

	err := p.Db.Model(a).Association("Tags").Replace(tags).Error
	return a, err
}

func (p *Engine) destroyArticle(c *gin.Context) (interface{}, error) {
	a := c.MustGet("article").(*Article)
	if err := p.Db.Model(a).Association("Tags").Clear().Error; err != nil {
		return nil, err
	}
	err := p.Db.Delete(a).Error
	return gin.H{}, err
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
