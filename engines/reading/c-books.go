package reading

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/kapmahc/epub"
	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexBooks(c *gin.Context) {
	data, err := p.getBooks(c.Request)
	web.JSON(c, data, err)
}

func (p *Engine) showBook(c *gin.Context) {
	id := c.Param("id")
	bk, err := p.readBook(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	p.writePoints(
		c.Writer,
		fmt.Sprintf("%s/reading/pages/%s", viper.GetString("server.backend"), id),
		bk.Ncx.Points,
	)
}

func (p *Engine) showPage(c *gin.Context) {
	log.Printf("%+v\n", c.Params)
	err := p.readBookPage(c.Writer, c.Param("id"), c.Param("href")[1:])
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

// -----------------------

func (p *Engine) readBookPage(w http.ResponseWriter, id string, name string) error {
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
					w.Header().Set("Content-Type", mf.MediaType)
					w.Write(body)
					return nil
				}
			}
		}
	}
	return errors.New("not found")
}

func (p *Engine) getBooks(r *http.Request) (*web.Pagination, error) {
	var total int64
	if err := p.Db.Model(&Book{}).Count(&total).Error; err != nil {
		return nil, err
	}
	pag := web.NewPagination(r, total)

	var books []Book
	if err := p.Db.
		Limit(pag.Limit()).Offset(pag.Offset()).
		Find(&books).Error; err != nil {
		return nil, err
	}
	for _, b := range books {
		pag.Items = append(pag.Items, b)
	}
	return pag, nil
}

func (p *Engine) writePoints(wrt io.Writer, href string, points []epub.NavPoint) {
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

func (p *Engine) readBook(id string) (*epub.Book, error) {
	var book Book
	if err := p.Db.
		Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}
	return epub.Open(path.Join(p.root(), book.File))
}
