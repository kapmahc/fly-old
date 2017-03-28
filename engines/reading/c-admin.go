package reading

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getAdminStatus(c *gin.Context) (interface{}, error) {
	data["title"] = p.I18n.T(lang, "reading.admin.status.title")
	tpl := "reading-admin-status"
	var bc int
	if err := p.Db.Model(&Book{}).Count(&bc).Error; err != nil {
		return nil, err
	}
	data["book"] = gin.H{
		p.I18n.T(lang, "reading.admin.status.book-count"): bc,
	}

	dict := gin.H{}
	for _, dic := range dictionaries {
		dict[dic.GetBookName()] = dic.GetWordCount()
	}
	data["dict"] = dict
	return tpl, nil
}

func (p *Engine) indexAdminBooks(c *gin.Context) (interface{}, error) {
	data["title"] = p.I18n.T(lang, "reading.admin.books.index.title")
	tpl := "reading-admin-books-index"
	var total int64
	if err := p.Db.Model(&Book{}).Count(&total).Error; err != nil {
		return nil, err
	}
	pag := web.NewPagination(c.Request, total)

	var books []Book
	if err := p.Db.
		Select([]string{"id", "title", "author"}).
		Limit(pag.Limit()).Offset(pag.Offset()).
		Find(&books).Error; err != nil {
		return nil, err
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
