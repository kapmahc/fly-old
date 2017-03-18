package site

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) _osStatus() gin.H {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return gin.H{
		"go version": runtime.Version(),
		"go root":    runtime.GOROOT(),
		"go runtime": runtime.NumGoroutine(),
		"go last gc": time.Unix(int64(mem.LastGC), 0).String(),
		"cpu":        runtime.NumCPU(),
		"memory":     fmt.Sprintf("%dM/%dM", mem.Alloc/1024/1024, mem.Sys/1024/1024),
		"now":        time.Now(),
		"version":    fmt.Sprintf("%s(%s)", runtime.GOOS, runtime.GOARCH),
	}
}
func (p *Engine) _cacheStatus() (string, error) {
	c := p.Redis.Get()
	defer c.Close()
	return redis.String(c.Do("INFO"))
}

func (p *Engine) _dbStatus() (gin.H, error) {
	val := gin.H{}
	switch viper.GetString("database.driver") {
	case postgresqlDriver:
		// http://blog.javachen.com/2014/04/07/some-metrics-in-postgresql.html
		row := p.Db.Raw("select pg_size_pretty(pg_database_size('postgres'))").Row()
		var size string
		row.Scan(&size)
		val["size"] = size
		if rows, err := p.Db.
			Raw("select pid,current_timestamp - least(query_start,xact_start) AS runtime,substr(query,1,25) AS current_query from pg_stat_activity where not pid=pg_backend_pid()").
			Rows(); err == nil {
			defer rows.Close()
			for rows.Next() {
				var pid int
				var ts time.Time
				var qry string
				row.Scan(&pid, &ts, &qry)
				val[fmt.Sprintf("pid-%d", pid)] = fmt.Sprintf("%s (%v)", ts.Format("15:04:05.999999"), qry)
			}
		} else {
			return nil, err
		}
		val["url"] = fmt.Sprintf(
			"%s://%s@%s:%d/%s",
			viper.GetString("database.driver"),
			viper.GetString("database.args.user"),
			viper.GetString("database.args.host"),
			viper.GetInt("database.args.port"),
			viper.GetString("database.args.dbname"),
		)

	}
	return val, nil
}

func (p *Engine) _jobsStatus() gin.H {
	return gin.H{
		"tasks": p.Server.GetRegisteredTaskNames(),
	}
}
func (p *Engine) getAdminSiteStatus(c *gin.Context, lang string, data gin.H) (string, error) {
	tpl := "site-admin-status"
	data["title"] = p.I18n.T(lang, "site.admin.status.title")
	data["os"] = p._osStatus()
	data["jobs"] = p._jobsStatus()
	var err error
	data["cache"], err = p._cacheStatus()
	if err != nil {
		return tpl, err
	}
	data["database"], err = p._dbStatus()
	if err != nil {
		return tpl, err
	}
	return tpl, nil
}

type fmSiteInfo struct {
	Title       string `form:"title"`
	SubTitle    string `form:"subTitle"`
	Keywords    string `form:"keywords"`
	Description string `form:"description"`
	Copyright   string `form:"copyright"`
}

func (p *Engine) formAdminSiteInfo(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "site.admin.info.title")
	tpl := "site-admin-info"
	if c.Request.Method == http.MethodPost {
		var fm fmSiteInfo
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		for k, v := range map[string]string{
			"title":       fm.Title,
			"subTitle":    fm.SubTitle,
			"keywords":    fm.Keywords,
			"description": fm.Description,
			"copyright":   fm.Copyright,
		} {
			if err := p.I18n.Set(lang, "site."+k, v); err != nil {
				return tpl, err
			}
		}
	}

	return tpl, nil
}

type fmSiteAuthor struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

func (p *Engine) formAdminSiteAuthor(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "site.admin.author.title")
	tpl := "site-admin-author"
	if c.Request.Method == http.MethodPost {

		var fm fmSiteAuthor
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		for k, v := range map[string]string{
			"name":  fm.Name,
			"email": fm.Email,
		} {
			if err := p.I18n.Set(lang, "site.author."+k, v); err != nil {
				return tpl, err
			}
		}
	}
	return tpl, nil
}

type fmSiteSeo struct {
	GoogleVerifyCode string `form:"googleVerifyCode"`
	BaiduVerifyCode  string `form:"baiduVerifyCode"`
}

func (p *Engine) formAdminSiteSeo(c *gin.Context, lang string, data gin.H) (string, error) {

	data["title"] = p.I18n.T(lang, "site.admin.seo.title")
	tpl := "site-admin-seo"
	if c.Request.Method == http.MethodPost {
		var fm fmSiteSeo
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		for k, v := range map[string]string{
			"google.verify.code": fm.GoogleVerifyCode,
			"baidu.verify.code":  fm.BaiduVerifyCode,
		} {
			if err := p.Settings.Set("site."+k, v, true); err != nil {
				return tpl, err
			}
		}
	}

	var gc string
	var bc string
	p.Settings.Get("site.google.verify.code", &gc)
	p.Settings.Get("site.baidu.verify.code", &bc)
	data["googleVerifyCode"] = gc
	data["baiduVerifyCode"] = bc
	return tpl, nil
}

type fmSiteSMTP struct {
	Host                 string `form:"host"`
	Port                 int    `form:"port"`
	Ssl                  bool   `form:"ssl"`
	Username             string `form:"username"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) formAdminSiteSMTP(c *gin.Context, lang string, data gin.H) (string, error) {

	data["title"] = p.I18n.T(lang, "site.admin.smtp.title")
	tpl := "site-admin-smtp"
	if c.Request.Method == http.MethodPost {
		var fm fmSiteSMTP
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}

		val := map[string]interface{}{
			"host":     fm.Host,
			"port":     fm.Port,
			"username": fm.Username,
			"password": fm.Password,
			"ssl":      fm.Ssl,
		}
		if err := p.Settings.Set("site.smtp", val, true); err != nil {
			return tpl, err
		}
	}

	smtp := make(map[string]interface{})
	if err := p.Settings.Get("site.smtp", &smtp); err == nil {
		smtp["password"] = ""
	} else {
		smtp["host"] = "localhost"
		smtp["port"] = 25
		smtp["ssl"] = false
		smtp["username"] = "no-reply@change-me.com"
		smtp["password"] = ""
	}
	data["smtp"] = smtp
	data["ports"] = []int{25, 465, 587, 2525, 2526}
	return tpl, nil
}

func (p *Engine) getAdminLocales(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "site.admin.locales.index.title")
	var items []web.Locale
	err := p.Db.Select([]string{"id", "code", "message"}).
		Where("lang = ?", lang).
		Order("code ASC").Find(&items).Error
	data["items"] = items
	return "site-admin-locales-index", err
}

func (p *Engine) deleteAdminLocales(c *gin.Context) (interface{}, error) {
	err := p.Db.
		Where("id = ?", c.Param("id")).
		Delete(&web.Locale{}).Error

	return gin.H{}, err
}

type fmLocale struct {
	Code    string `form:"code" binding:"required,max=255"`
	Message string `form:"message" binding:"required"`
}

func (p *Engine) formAdminLocales(c *gin.Context, lang string, data gin.H) (string, error) {
	data["title"] = p.I18n.T(lang, "buttons.edit")
	tpl := "site-admin-locales-edit"
	if c.Request.Method == http.MethodPost {
		var fm fmLocale
		if err := c.Bind(&fm); err != nil {
			return tpl, err
		}
		data["code"] = fm.Code
		if err := p.I18n.Set(lang, fm.Code, fm.Message); err != nil {
			return tpl, err
		}
		c.Redirect(http.StatusFound, "/admin/locales")
		return "", nil
	}

	data["code"] = c.Request.URL.Query().Get("code")
	return tpl, nil
}

func (p *Engine) getAdminUsers(c *gin.Context, lang string, data gin.H) (string, error) {
	var items []auth.User
	err := p.Db.
		Order("last_sign_in_at DESC").Find(&items).Error
	data["users"] = items
	data["title"] = p.I18n.T(lang, "site.admin.users.index.title")
	return "site-admin-users-index", err
}
