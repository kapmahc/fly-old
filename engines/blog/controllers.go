package blog

import (
	"bufio"
	"net/http"
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
	data["items"] = items
	data["title"] = p.I18n.T(lang, "blog.index.title")
	return "blog-index", err
}

func (p *Engine) showPost(c *gin.Context) (interface{}, error) {
	href := c.Param("href")[1:]
	tpl := "blog-show"
	posts, err := p.getPosts(c.MustGet(web.LOCALE).(string))
	if err != nil {
		return nil, err
	}
	data["items"] = posts
	for _, i := range posts {
		if i.Href == href {
			data["title"] = i.Title
			data["item"] = i
			return tpl, nil
		}
	}

	c.AbortWithStatus(http.StatusNotFound)
	return "", nil
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
