package reading

import (
	"net/http"

	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) myNotes(c *gin.Context) (interface{}, error) {

	user := c.MustGet(auth.CurrentUser).(*auth.User)
	isa := c.MustGet(auth.IsAdmin).(bool)
	var notes []Note
	qry := p.Db
	if !isa {
		qry = qry.Where("user_id = ?", user.ID)
	}
	if err := qry.Order("updated_at DESC").Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

func (p *Engine) indexNotes(c *gin.Context) (interface{}, error) {

	var total int64
	var pag *web.Pagination
	if err := p.Db.Model(&Note{}).Count(&total).Error; err != nil {
		return nil, err
	}

	pag = web.NewPagination(c.Request, total)
	var notes []Note
	if err := p.Db.
		Limit(pag.Limit()).Offset(pag.Offset()).
		Find(&notes).Error; err != nil {
		return nil, err
	}

	for _, it := range notes {
		pag.Items = append(pag.Items, it)
	}

	return pag, nil
}

type fmNoteNew struct {
	Type   string `form:"type" binding:"required,max=8"`
	Body   string `form:"body" binding:"required,max=2000"`
	BookID uint   `form:"bookId"`
}

func (p *Engine) createNote(c *gin.Context) (interface{}, error) {

	user := c.MustGet(auth.CurrentUser).(*auth.User)

	var fm fmNoteNew
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	item := Note{
		Type:   fm.Type,
		Body:   fm.Body,
		BookID: fm.BookID,
		UserID: user.ID,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

type fmNoteEdit struct {
	Type string `form:"type" binding:"required,max=8"`
	Body string `form:"body" binding:"required,max=2000"`
}

func (p *Engine) updateNote(c *gin.Context) (interface{}, error) {

	note := c.MustGet("note").(*Note)

	var fm fmNoteEdit
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}

	if err := p.Db.Model(note).
		Updates(map[string]interface{}{
			"body": fm.Body,
			"type": fm.Type,
		}).Error; err != nil {
		return nil, err
	}

	return gin.H{}, nil
}

func (p *Engine) destroyNote(c *gin.Context) (interface{}, error) {
	n := c.MustGet("note").(*Note)
	err := p.Db.Delete(n).Error
	return gin.H{}, err
}

func (p *Engine) canEditNote(c *gin.Context) {
	user := c.MustGet(auth.CurrentUser).(*auth.User)

	var n Note
	err := p.Db.Where("id = ?", c.Param("id")).First(&n).Error
	if err == nil {
		if user.ID == n.UserID || c.MustGet(auth.IsAdmin).(bool) {
			c.Set("note", &n)
			c.Next()
			return
		}
		c.AbortWithStatus(http.StatusForbidden)
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}
