package site

import (
	"fmt"
	"runtime"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) _osStatus() web.H {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return web.H{
		"go version": runtime.Version(),
		"go root":    runtime.GOROOT(),
		"go runtime": runtime.NumGoroutine(),
		"go last gc": time.Unix(int64(mem.LastGC), 0),
		"cpu":        runtime.NumCPU(),
		"memory":     fmt.Sprintf("%dM/%dM", mem.Alloc/1024/1024, mem.Sys/1024/1024),
		"now":        time.Now(),
		"version":    fmt.Sprintf("%s(%s)", runtime.GOOS, runtime.GOARCH),
	}
}
func (p *Engine) _cacheStatus() string {
	c := p.Redis.Get()
	defer c.Close()
	s, err := redis.String(c.Do("INFO"))
	if err != nil {
		log.Error(err)
	}
	return s
}

func (p *Engine) _dbStatus() web.H {

	val := web.H{}
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
			log.Error(err)
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
	return val

}

func (p *Engine) _jobsStatus() web.H {
	return web.H{
		"tasks": p.Server.GetRegisteredTaskNames(),
	}
}
func (p *Engine) getAdminSiteStatus(c *gin.Context) {

	web.JSON(c, gin.H{
		"os":       p._osStatus(),
		"cache":    p._cacheStatus(),
		"database": p._dbStatus(),
		"jobs":     p._jobsStatus(),
	}, nil)
}

type fmSiteInfo struct {
	Title       string `form:"title"`
	SubTitle    string `form:"subTitle"`
	Keywords    string `form:"keywords"`
	Description string `form:"description"`
	Copyright   string `form:"copyright"`
}

func (p *Engine) postAdminSiteInfo(c *gin.Context) {
	lang := c.MustGet(web.LOCALE).(string)

	var fm fmSiteInfo
	err := c.Bind(&fm)
	if err == nil {
		for k, v := range map[string]string{
			"title":       fm.Title,
			"sub-title":   fm.SubTitle,
			"keywords":    fm.Keywords,
			"description": fm.Description,
			"copyright":   fm.Copyright,
		} {
			if err = p.I18n.Set(lang, "site."+k, v); err != nil {
				break
			}
		}
	}

	web.JSON(c, nil, err)
}

type fmSiteAuthor struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

func (p *Engine) postAdminSiteAuthor(c *gin.Context) {

	lang := c.MustGet(web.LOCALE).(string)

	var fm fmSiteAuthor
	err := c.Bind(&fm)
	if err == nil {
		for k, v := range map[string]string{
			"name":  fm.Name,
			"email": fm.Email,
		} {
			if err = p.I18n.Set(lang, "site.author."+k, v); err != nil {
				break
			}
		}
	}

	web.JSON(c, nil, err)
}

type fmSiteSeo struct {
	GoogleVerifyCode string `form:"googleVerifyCode"`
	BaiduVerifyCode  string `form:"baiduVerifyCode"`
}

func (p *Engine) getAdminSiteSeo(c *gin.Context) {
	var gc string
	var bc string
	p.Settings.Get("site.google.verify.code", &gc)
	p.Settings.Get("site.baidu.verify.code", &bc)
	web.JSON(c, gin.H{
		"googleVerifyCode": &gc,
		"baiduVerifyCode":  &bc,
	}, nil)
}

func (p *Engine) postAdminSiteSeo(c *gin.Context) {
	var fm fmSiteSeo
	err := c.Bind(&fm)
	if err == nil {
		for k, v := range map[string]string{
			"google.verify.code": fm.GoogleVerifyCode,
			"baidu.verify.code":  fm.BaiduVerifyCode,
		} {
			if err = p.Settings.Set("site."+k, v, true); err != nil {
				break
			}
		}
	}
	web.JSON(c, nil, err)
}

type fmSiteSMTP struct {
	Host                 string `form:"host"`
	Port                 int    `form:"port"`
	Ssl                  bool   `form:"ssl"`
	Username             string `form:"username"`
	Password             string `form:"password"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

func (p *Engine) getAdminSiteSMTP(c *gin.Context) {
	var smtp map[string]interface{}
	p.Settings.Get("site.smtp", &smtp)
	web.JSON(c, smtp, nil)
}

func (p *Engine) postAdminSiteSMTP(c *gin.Context) {
	var fm fmSiteSMTP
	err := c.Bind(&fm)
	if err == nil {

		val := map[string]interface{}{
			"host":     fm.Host,
			"port":     fm.Port,
			"username": fm.Username,
			"password": fm.Password,
			"ssl":      fm.Ssl,
		}
		err = p.Settings.Set("site.smtp", val, true)
	}

	web.JSON(c, nil, err)
}

func (p *Engine) getAdminLocales(c *gin.Context) {

	lang := c.MustGet(web.LOCALE).(string)

	var items []web.Locale
	err := p.Db.Select([]string{"code", "message"}).
		Where("locale = ?", lang).
		Order("code ASC").Find(&items).Error

	web.JSON(c, items, err)
}

type fmLocale struct {
	Code    string `form:"code" validate:"required,max=255"`
	Message string `form:"message" validate:"required"`
}

func (p *Engine) postAdminLocales(c *gin.Context) {

	lang := c.MustGet(web.LOCALE).(string)

	var fm fmLocale
	err := c.Bind(&fm)
	if err == nil {
		err = p.I18n.Set(lang, fm.Code, fm.Message)
	}

	web.JSON(c, nil, err)
}

func (p *Engine) getAdminUsers(c *gin.Context) {
	var items []auth.User
	err := p.Db.
		Order("last_sign_in_at DESC").Find(&items).Error
	web.JSON(c, items, err)
}
