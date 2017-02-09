package reading

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/epub"
	"github.com/kapmahc/fly/engines/base"
)

// ShowBookPage show book page
// @router /books/:id/* [get]
func (p *Controller) ShowBookPage() {
	id := p.Ctx.Input.Param(":id")
	name := p.Ctx.Input.Param(":splat")
	if err := p.readBookPage(id, name); err != nil {
		beego.Error(err)
		p.Abort(http.StatusInternalServerError)
	}

}

func (p *Controller) readBookPage(id string, name string) error {
	bk, err := p.readBook(id)
	if err != nil {
		return err
	}
	for _, fn := range bk.Files() {
		if strings.HasSuffix(fn, name) {
			for _, mf := range bk.Opf.Manifest {
				if mf.Href == name {
					rdr, err := bk.Open(name)
					if err != nil {
						return err
					}
					defer rdr.Close()
					body, err := ioutil.ReadAll(rdr)
					if err != nil {
						return err
					}
					p.Ctx.Output.ContentType(mf.MediaType)
					p.Ctx.Output.Body(body)
					return nil
				}
			}
		}
	}

	p.Abort(http.StatusNotFound)
	return nil
}

// ShowBook show books
// @router /books/:id [get]
func (p *Controller) ShowBook() {
	id := p.Ctx.Input.Param(":id")
	bk, err := p.readBook(id)
	if err != nil {
		beego.Error(err)
		p.Abort(http.StatusInternalServerError)
	}
	p.Data["book"] = bk
	href := p.URLFor("reading.Controller.ShowBook", ":id", id)
	p.Data["href"] = href

	var ncx bytes.Buffer
	p.writePoints(
		&ncx,
		href,
		bk.Ncx.Points,
	)
	p.Data["ncx"] = template.HTML(ncx.Bytes())
	p.HTML(strings.Join(bk.Opf.Metadata.Title, "|"), "reading/books/show.html")
}

func (p *Controller) writePoints(wrt io.Writer, href string, points []epub.NavPoint) {
	wrt.Write([]byte("<ol>"))
	for _, it := range points {
		wrt.Write([]byte("<li>"))
		fmt.Fprintf(
			wrt,
			`<a href="%s/%s" target="_blank">%s</a>`,
			href,
			it.Content.Src,
			it.Text,
		)
		p.writePoints(wrt, href, it.Points)
		wrt.Write([]byte("</li>"))
	}
	wrt.Write([]byte("</ol>"))
}

func (p *Controller) readBook(id string) (*epub.Book, error) {
	var book Book
	if err := orm.NewOrm().
		QueryTable(&book).
		Filter("id", id).One(&book); err != nil {
		return nil, err
	}
	return epub.Open(path.Join(booksRoot(), book.File))
}

// GetBooksScan scan books
// @router /books/scan [get]
func (p *Controller) GetBooksScan() {
	// TODO
	base.SendTask(scanBookTask)
	p.Data["json"] = map[string]interface{}{"ok": true}
	p.ServeJSON()
}

// IndexBooks books
// @router /books [get]
func (p *Controller) IndexBooks() {
	page, _ := p.GetInt64("page")
	size, _ := p.GetInt64("size")

	var err error
	p.Data["books"], err = p.getBooks(page, size)
	if err != nil {
		beego.Error(err)
		p.Abort(http.StatusInternalServerError)
	}
	p.HTML(p.T("reading.home.title"), "reading/books/index.html")
}

func (p *Controller) getBooks(page, size int64) (*base.Pagination, error) {

	o := orm.NewOrm()
	total, err := o.QueryTable(&Book{}).Count()
	if err != nil {
		return nil, err
	}
	pag := base.NewPagination(
		p.URLFor("reading.Controller.IndexBooks"),
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
