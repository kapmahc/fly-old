package site

import (
	"text/template"

	"github.com/astaxie/beego"
)

// GetNginxConf nginx.conf
// @router /nginx.conf [get]
func (p *Controller) GetNginxConf() {
	tpl := `
server {
  listen       443;
  server_name  {{.Domain}};
  ssl on;
  ssl_certificate /etc/ssl/certs/{{.Domain}}.crt;
  ssl_certificate_key /etc/ssl/private/{{.Domain}}.key;
  charset utf-8;
  access_log  /var/www/{{.Domain}}/logs/access.log;
  error_log  /var/www/{{.Domain}}/logs/error.log;
  location /(css|js|fonts|img)/ {
    access_log off;
    expires 1d;
    root "/var/www/{{.Domain}}/public";
    try_files $uri @backend;
  }
  location / {
    try_files /_not_exists_ @backend;
  }
  location @backend {
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header Host $http_host;
    proxy_redirect off;
    proxy_set_header X-Forwarded-Proto https;
    proxy_pass http://127.0.0.1:{{.Port}};
  }
}
`
	t := template.Must(template.New("").Parse(tpl))
	t.Execute(p.Ctx.ResponseWriter, struct {
		Port   int
		Static string
		Domain string
	}{
		Port:   beego.AppConfig.DefaultInt("httpport", 8080),
		Static: beego.AppConfig.String("staticdir"),
		Domain: p.GetString("domain", "www.change-me.com"),
	})
}
