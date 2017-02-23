package reading

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kapmahc/epub"
	"github.com/kapmahc/fly/web"
)

func (p *Engine) indexBooks(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value(web.DATA).(web.H)
	lang := r.Context().Value(web.LOCALE).(string)
	books, err := p.getBooks(r)
	if err != nil {
		log.Error(err)
		p.Render.Text(w, http.StatusInternalServerError, err.Error())
		return
	}
	data["books"] = books
	data["title"] = p.I18n.T(lang, "reading.books.index.title")
	p.Render.HTML(w, http.StatusOK, "reading/books/index", data)
}

func (p *Engine) showBook(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value(web.DATA).(web.H)

	vars := mux.Vars(r)
	id := vars["id"]

	bk, err := p.readBook(id)
	if err != nil {
		log.Error(err)
		p.Render.Text(w, http.StatusInternalServerError, err.Error())
		return
	}
	data["book"] = bk
	href := fmt.Sprintf("/reading/books/%s", id)
	data["href"] = href

	var ncx bytes.Buffer
	p.writePoints(
		&ncx,
		href,
		bk.Ncx.Points,
	)
	data["ncx"] = template.HTML(ncx.Bytes())
	data["title"] = strings.Join(bk.Opf.Metadata.Title, "|")
	p.Render.HTML(w, http.StatusOK, "reading/books/show", data)
}

func (p *Engine) showBookPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Debugf("VARS %+v", vars)
	if err := p.readBookPage(w, vars["id"], vars["name"]); err != nil {
		p.Render.Text(w, http.StatusInternalServerError, err.Error())
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
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))

	var total int64
	if err := p.Db.Model(&Book{}).Count(&total).Error; err != nil {
		return nil, err
	}
	pag := web.NewPagination(
		"/reading/books",
		int64(page), int64(size), total,
	)

	var books []Book
	if err := p.Db.Select([]string{"title", "id", "author", "subject", "description", "published_at"}).
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
