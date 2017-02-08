package reading

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/epub"
	"github.com/kapmahc/fly/engines/base"
)

const (
	// EPUB epub book ext
	EPUB = ".epub"

	scanBookTask = "reading.books.scan"

	// SEP sep
	SEP = ";"
)

func booksRoot() string {
	return filepath.Join("tmp", "books")
}

func scanBook() (int, error) {
	o := orm.NewOrm()
	root := booksRoot()
	count := 0
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(info.Name()) != EPUB {
			return nil
		}
		beego.Info("find book", path)

		bk, err := epub.Open(path)
		if err != nil {
			return err
		}
		mt := bk.Opf.Metadata
		var authors []string
		for _, a := range mt.Creator {
			authors = append(authors, a.Data)
		}

		book := Book{
			Title:       strings.Join(mt.Title, SEP),
			Lang:        strings.Join(mt.Language, SEP),
			Author:      strings.Join(authors, SEP),
			Publisher:   strings.Join(mt.Publisher, SEP),
			Subject:     strings.Join(mt.Subject, SEP),
			Description: strings.Join(mt.Description, SEP),
			Type:        bk.Mimetype,
			File:        path[len(root)+1:],
		}
		ct, err := o.QueryTable(&Book{}).Filter("file", book.File).Count()
		if err != nil {
			return err
		}
		if ct > 0 {
			return nil
		}
		if len(mt.Date) > 0 {
			date := mt.Date[0].Data
			layout := "2006-01-02"
			if len(date) > len(layout) {
				date = date[:len(layout)]
			}
			book.PublishedAt, err = time.Parse(layout, date)
			if err != nil {
				return err
			}
		}
		if len(mt.Coverage) > 0 {
			book.Cover = mt.Coverage[0]
		}
		if _, err = o.Insert(&book); err != nil {
			return err
		}
		count++
		return nil
	})
	return count, err
}

func init() {
	base.RegisterTask(scanBookTask, scanBook)
}
