package blog

import (
	"bufio"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

const (
	// MARKDOWN post ext
	MARKDOWN = ".md"
)

func (p *Engine) indexPosts(c *gin.Context) (interface{}, error) {
	items, err := p.getPosts(c.MustGet(web.LOCALE).(string))
	return items, err
}

func (p *Engine) showPost(c *gin.Context) (interface{}, error) {
	lang := c.MustGet(web.LOCALE).(string)
	href := c.Param("href")[1:]

	posts, err := p.getPosts(lang)
	if err != nil {
		return nil, err
	}

	for _, i := range posts {
		if i.Href == href {
			return i, nil
		}
	}

	return nil, p.I18n.E(lang, "errors.no-found")
}

// -------------

func (p *Engine) getPosts(lang string) ([]Post, error) {
	key := "blogs://" + lang
	var items []Post
	if err := p.Cache.Get(key, &items); err == nil {
		return items, nil
	}
	root := p.root(lang)
	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		name := info.Name()
		if info.IsDir() || filepath.Ext(name) != MARKDOWN {
			return nil
		}

		fd, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fd.Close()
		san := bufio.NewScanner(fd)
		var title, body string
		for san.Scan() {
			line := san.Text()
			if title == "" && line != "" {
				title = line
				continue
			}
			body += line + "\n"
		}

		items = append(items, Post{
			Href:      path[len(root)+1:],
			Title:     title,
			Body:      body,
			Published: info.ModTime(),
		})
		return nil
	}); err != nil {
		return nil, err
	}

	sort.Sort(ByPublished(items))
	p.Cache.Set(key, items, time.Hour*24)
	return items, nil
}

func (p *Engine) root(lang string) string {
	return filepath.Join("tmp", "posts", lang)
}
