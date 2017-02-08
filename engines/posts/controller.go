package posts

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"

	"github.com/astaxie/beego"
	"github.com/kapmahc/fly/engines/auth"
)

const (
	// MARKDOWN post ext
	MARKDOWN = ".md"
)

// Controller controller
type Controller struct {
	auth.UserController
}

// Prepare prepare
func (p *Controller) Prepare() {
	p.UserController.Prepare()
	p.Data["posts"] = p.getPosts()
}

// GetHome home
// @router / [get]
func (p *Controller) GetHome() {
	posts := p.Data["posts"].([]Post)
	const size = 12
	if len(posts) > 12 {
		p.Data["items"] = posts[:12]
	} else {
		p.Data["items"] = posts[:]
	}
	p.HTML(p.T("posts.home.title"), "posts/home.html")
}

// GetShow show
// @router /* [get]
func (p *Controller) GetShow() {
	name := p.Ctx.Input.Param(":splat")

	switch filepath.Ext(name) {
	case MARKDOWN:
		for _, i := range p.Data["posts"].([]Post) {
			if i.Href == name {
				p.Data["post"] = i
			}
		}
		if p.Data["post"] == nil {
			p.Abort(http.StatusNotFound)
		}

		p.HTML(name, "posts/show.html")
	case ".png":
		body, err := ioutil.ReadFile(filepath.Join(p.postsRoot(), name))
		if err != nil {
			beego.Error(err)
			p.Abort(http.StatusInternalServerError)
		}
		p.Ctx.Output.ContentType("image/png")
		p.Ctx.Output.Body(body)
	default:
		p.Abort(http.StatusNotFound)
	}
}

func (p *Controller) getPosts() []Post {
	var items []Post
	root := p.postsRoot()
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
		beego.Error(err)
	}

	sort.Sort(ByPublished(items))
	return items
}

func (p *Controller) postsRoot() string {
	return filepath.Join("tmp", "posts")
}
