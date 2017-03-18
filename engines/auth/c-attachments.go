package auth

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) newAttachment(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.upload")
	return "auth-attachments-new", nil
}

func (p *Engine) createAttachment(c *gin.Context) (interface{}, error) {
	user := c.MustGet(CurrentUser).(*User)

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return nil, err
	}

	url, size, err := p.Uploader.Save(file, header)
	if err != nil {
		return nil, err
	}

	// http://golang.org/pkg/net/http/#DetectContentType
	buf := make([]byte, 512)
	file.Seek(0, 0)
	if _, err = file.Read(buf); err != nil {
		return nil, err
	}

	a := Attachment{
		Title:     header.Filename,
		URL:       url,
		UserID:    user.ID,
		MediaType: http.DetectContentType(buf),
		Length:    size / 1024,
	}
	if err := p.Db.Create(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

type fmAttachmentEdit struct {
	Title string `form:"title" binding:"required,max=255"`
}

func (p *Engine) updateAttachment(c *gin.Context, lang string, data gin.H) (string, error) {
	a := c.MustGet("attachment").(*Attachment)
	tpl := "auth-attachments-edit"
	data["title"] = p.I18n.T(lang, "buttons.edit")
	data["item"] = a

	if c.Request.Method == http.MethodPost {
		var fm fmAttachmentEdit
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		if err := p.Db.Model(a).Update("title", fm.Title).Error; err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/attachments")
		return "", nil
	}
	return tpl, nil
}

func (p *Engine) destroyAttachment(c *gin.Context) (interface{}, error) {
	a := c.MustGet("attachment").(*Attachment)
	err := p.Db.Delete(a).Error
	if err != nil {
		return nil, err
	}
	return a, p.Uploader.Remove(a.URL)
}
func (p *Engine) indexAttachments(c *gin.Context, lang string, data gin.H) (string, error) {
	user := c.MustGet(CurrentUser).(*User)
	isa := c.MustGet(IsAdmin).(bool)
	var items []Attachment
	qry := p.Db
	if !isa {
		qry = qry.Where("user_id = ?", user.ID)
	}
	err := qry.Order("updated_at DESC").Find(&items).Error
	data["attachments"] = items
	data["title"] = p.I18n.T(lang, "auth.attachments.index.title")
	return "auth-attachments-index", err
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
