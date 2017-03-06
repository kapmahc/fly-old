package site

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
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
func (p *Engine) adminSiteStatus(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)

	data["title"] = p.I18n.T(lang, "site.admin.status.title")
	data["os"] = p._osStatus()
	data["cache"] = p._cacheStatus()
	data["database"] = p._dbStatus()
	data["jobs"] = p._jobsStatus()
	p.Ctx.HTML(w, "site/admin/status", data)
}

func (p *Engine) adminUsers(w http.ResponseWriter, r *http.Request) {
	lang := r.Context().Value(web.LOCALE).(string)
	data := r.Context().Value(web.DATA).(web.H)
	var users []auth.User
	if err := p.Db.
		Select([]string{"name", "logo", "home"}).
		Order("last_sign_in_at DESC").
		Find(&users).Error; err != nil {
		log.Error(err)
	}
	data["users"] = users
	data["title"] = p.I18n.T(lang, "auth.users.index.title")
	p.Ctx.HTML(w, "site/admin/users", data)
}
