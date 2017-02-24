package blog

import (
	"bufio"
	"net/http"
	"os"
	"path/filepath"
	"sort"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kapmahc/fly/web"
)

const (
	// MARKDOWN post ext
	MARKDOWN = ".md"
)

func (p *Engine) indexPosts(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value(web.DATA).(web.H)
	lang := r.Context().Value(web.LOCALE).(string)

	posts := p.getPosts()
	const size = 12

	if len(posts) > size {
		data["items"] = posts[:size]
	} else {
		data["items"] = posts[:]
	}
	data["posts"] = posts
	data["title"] = p.I18n.T(lang, "blog.list.title")
	p.Render.HTML(w, "blog/index", data)
}

func (p *Engine) showPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	data := r.Context().Value(web.DATA).(web.H)
	posts := p.getPosts()
	data["posts"] = posts
	for _, i := range posts {
		if i.Href == name {
			data["post"] = i
			data["title"] = i.Title
			p.Render.HTML(w, "blog/show", data)
			return
		}
	}

	p.Render.NotFound(w)

}

// -------------

func (p *Engine) getPosts() []Post {
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
		log.Error(err)
	}

	sort.Sort(ByPublished(items))
	return items
}

func (p *Engine) root() string {
	return filepath.Join("tmp", "posts")
}
