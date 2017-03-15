package auth

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) showAttachment(c *gin.Context) (interface{}, error) {
	var a Attachment
	err := p.Db.Where("id = ?", c.Param("id")).First(&a).Error
	return a, err
}

func (p *Engine) createAttachment(c *gin.Context) (interface{}, error) {
	user := c.MustGet(CurrentUser).(*User)
	if err := c.Request.ParseMultipartForm(10 * 1024); err != nil {
		return nil, err
	}

	var items []Attachment

	for _, f := range c.Request.MultipartForm.File["files"] {
		url, size, err := p.Uploader.Save(f)
		if err != nil {
			return nil, err
		}
		fd, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer fd.Close()

		// http://golang.org/pkg/net/http/#DetectContentType
		buf := make([]byte, 512)
		if _, err = fd.Read(buf); err != nil {
			return nil, err
		}

		a := Attachment{
			Title:     f.Filename,
			URL:       url,
			UserID:    user.ID,
			MediaType: http.DetectContentType(buf),
			Length:    size / 1024,
		}
		if err := p.Db.Create(&a).Error; err != nil {
			return nil, err
		}
		items = append(items, a)
	}
	return items, nil
}

type fmAttachmentEdit struct {
	Title string `form:"title" binding:"required,max=255"`
}

func (p *Engine) updateAttachment(c *gin.Context) (interface{}, error) {
	a := c.MustGet("attachment").(*Attachment)
	var fm fmAttachmentEdit

	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	err := p.Db.Model(a).Update("title", fm.Title).Error
	return a, err
}

func (p *Engine) destroyAttachment(c *gin.Context) (interface{}, error) {
	a := c.MustGet("attachment").(*Attachment)
	err := p.Db.Delete(a).Error
	return a, err
}
func (p *Engine) indexAttachments(c *gin.Context) (interface{}, error) {
	user := c.MustGet(CurrentUser).(*User)
	isa := c.MustGet(IsAdmin).(bool)
	var items []Attachment
	qry := p.Db
	if !isa {
		qry = qry.Where("user_id = ?", user.ID)
	}
	err := qry.Order("updated_at DESC").Find(&items).Error
	return items, err
}

func (p *Engine) canEditAttachment(c *gin.Context) {
	user := c.MustGet(CurrentUser).(*User)

	var a Attachment
	err := p.Db.Where("id = ?", c.Param("id")).First(&a).Error
	if err == nil {
		if user.ID == a.UserID || c.MustGet(IsAdmin).(bool) {
			c.Set("attachment", &a)
			c.Next()
			return
		}
		c.AbortWithStatus(http.StatusForbidden)
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}

}
