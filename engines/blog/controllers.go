package blog

import (
	"bufio"
	"os"
	"path/filepath"
	"sort"

	gin "gopkg.in/gin-gonic/gin.v1"

	"github.com/kapmahc/fly/web"
)

const (
	// MARKDOWN post ext
	MARKDOWN = ".md"
)

func (p *Engine) indexPosts(c *gin.Context) {
	data, err := p.getPosts()
	web.JSON(c, data, err)
}

func (p *Engine) showPost(c *gin.Context) {
	href := c.Param("href")[1:]
	posts, err := p.getPosts()
	if err == nil {
		for _, i := range posts {
			if i.Href == href {
				web.JSON(c, i, nil)
				return
			}
		}
	}
	web.JSON(c, Post{Href: href}, err)
}

// -------------

func (p *Engine) getPosts() ([]Post, error) {
	var items []Post
	root := p.root()
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
	return items, nil
}

func (p *Engine) root() string {
	return filepath.Join("tmp", "posts")
}
