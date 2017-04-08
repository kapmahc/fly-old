package site

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/user"
	"reflect"
	"runtime"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/kapmahc/fly-bak/web"
	"github.com/kapmahc/sky"
	"github.com/spf13/viper"
)

func (p *Engine) _osStatus() (sky.H, error) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	hn, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	hu, err := user.Current()
	if err != nil {
		return nil, err
	}
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	var ifo syscall.Sysinfo_t
	if err := syscall.Sysinfo(&ifo); err != nil {
		return nil, err
	}
	return sky.H{
		"app version":          fmt.Sprintf("%s(%s) - %s", web.Version, web.BuildTime, viper.GetString("env")),
		"app root":             pwd,
		"who-am-i":             fmt.Sprintf("%s@%s", hu.Username, hn),
		"go version":           runtime.Version(),
		"go root":              runtime.GOROOT(),
		"go runtime":           runtime.NumGoroutine(),
		"go last gc":           time.Unix(0, int64(mem.LastGC)).Format(time.ANSIC),
		"os cpu":               runtime.NumCPU(),
		"os ram(free/total)":   fmt.Sprintf("%dM/%dM", ifo.Freeram/1024/1024, ifo.Totalram/1024/1024),
		"os swap(free/total)":  fmt.Sprintf("%dM/%dM", ifo.Freeswap/1024/1024, ifo.Totalswap/1024/1024),
		"go memory(alloc/sys)": fmt.Sprintf("%dM/%dM", mem.Alloc/1024/1024, mem.Sys/1024/1024),
		"os time":              time.Now().Format(time.ANSIC),
		"os arch":              fmt.Sprintf("%s(%s)", runtime.GOOS, runtime.GOARCH),
		"os uptime":            (time.Duration(ifo.Uptime) * time.Second).String(),
		"os loads":             ifo.Loads,
		"os procs":             ifo.Procs,
	}, nil
}
func (p *Engine) _networkStatus() (sky.H, error) {
	sts := sky.H{}
	ifs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, v := range ifs {
		ips := []string{v.HardwareAddr.String()}
		adrs, err := v.Addrs()
		if err != nil {
			return nil, err
		}
		for _, adr := range adrs {
			ips = append(ips, adr.String())
		}
		sts[v.Name] = ips
	}
	return sts, nil
}

func (p *Engine) _dbStatus() (sky.H, error) {
	val := sky.H{
		"drivers": sql.Drivers(),
	}
	switch viper.GetString("database.driver") {
	case postgresqlDriver:
		row := p.Db.Raw("select version()").Row()
		var version string
		row.Scan(&version)
		val["version"] = version
		// http://blog.javachen.com/2014/04/07/some-metrics-in-postgresql.html
		row = p.Db.Raw("select pg_size_pretty(pg_database_size('postgres'))").Row()
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

func (p *Engine) _routes() map[string]string {
	val := make(map[string]string)
	p.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		name := route.GetName()
		if name == "" {
			if hnd := route.GetHandler(); hnd != nil {
				val[tpl] = runtime.FuncForPC(reflect.ValueOf(hnd).Pointer()).Name()
			}
		} else {
			val[tpl] = name
		}
		return nil
	})
	return val
}

func (p *Engine) getAdminSiteStatus(c *sky.Context) error {
	lang := c.Get(sky.LOCALE).(string)
	data := c.Get(sky.DATA).(sky.H)

	var err error
	data["title"] = p.I18n.T(lang, "site.admin.status.title")
	data["os"], err = p._osStatus()
	if err != nil {
		return err
	}
	data["network"], err = p._networkStatus()
	if err != nil {
		return err
	}
	data["jobs"] = p.Server.Status()
	data["routes"] = p._routes()

	data["cache"], err = p.Cache.Store.Status()
	if err != nil {
		return err
	}
	data["database"], err = p._dbStatus()
	if err != nil {
		return err
	}
	c.HTML(http.StatusOK, "site/admin/status", data)
	return nil
}
