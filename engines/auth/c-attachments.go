package auth

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

// Uploader attachment uploader
type Uploader interface{}

// FileSystemUploader file-system storage
type FileSystemUploader struct {
}

func (p *Engine) getAttachment(c *gin.Context) {

}

func (p *Engine) showAttachment(c *gin.Context) (interface{}, error) {
	var a Attachment
	err := p.Db.Where("id = ?", c.Param("id")).First(&a).Error
	return a, err
}

func (p *Engine) createAttachment(c *gin.Context) (interface{}, error) {
	if err := c.Request.ParseMultipartForm(10 * 1024); err != nil {
		return nil, err
	}

	for _, fn := range c.Request.MultipartForm.File["files"] {
		fd, err := fn.Open()
		if err != nil {
			return nil, err
		}
		defer fd.Close()
		// out, err := os.Create("./tmp/" + filename + ".png")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer out.Close()
		// _, err = io.Copy(out, file)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}
	return gin.H{}, nil
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
