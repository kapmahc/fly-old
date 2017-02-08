package reading

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/engines/base"
)

// Controller controller
type Controller struct {
	auth.UserController
}

// GetHome get home
// @router / [get]
func (p *Controller) GetHome() {
	p.GetBooks()
}

// GetBooks get books
// @router /books [get]
func (p *Controller) GetBooks() {
	page, _ := p.GetInt64("page")
	size, _ := p.GetInt64("size")

	var err error
	p.Data["books"], err = p.getBooks(page, size)
	if err != nil {
		beego.Error(err)
		p.Abort(http.StatusInternalServerError)
	}
	p.HTML(p.T("reading.home.title"), "reading/home.html")
}

func (p *Controller) getBooks(page, size int64) (*base.Pagination, error) {

	o := orm.NewOrm()
	total, err := o.QueryTable(&Book{}).Count()
	if err != nil {
		return nil, err
	}
	pag := base.NewPagination(
		p.URLFor("reading.Controller.GetBooks"),
		page, size, total,
	)

	var books []Book
	_, err = o.QueryTable(&Book{}).
		Limit(pag.Limit()).Offset(pag.Offset()).
		All(&books, "title", "id", "author", "subject", "description", "published_at")
	if err != nil {
		return nil, err
	}
	for _, b := range books {
		pag.Items = append(pag.Items, b)
	}
	return pag, nil
}
