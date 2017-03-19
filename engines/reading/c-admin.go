package reading

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexAdminBooks(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "reading.admin.books.index.title")
	tpl := "reading-admin-books-index"
	var total int64
	if err := p.Db.Model(&Book{}).Count(&total).Error; err != nil {
		return tpl, err
	}
	pag := web.NewPagination(c.Request, total)

	var books []Book
	if err := p.Db.
		Select([]string{"id", "title", "author"}).
		Limit(pag.Limit()).Offset(pag.Offset()).
		Find(&books).Error; err != nil {
		return tpl, err
	}

	for _, b := range books {
		pag.Items = append(pag.Items, b)
	}
	data["pager"] = pag
	return tpl, nil
}

func (p *Engine) destroyAdminBook(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(Book{}).Error
	return gin.H{}, err
}
