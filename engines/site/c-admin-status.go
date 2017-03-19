package site

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"os/user"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) _osStatus() (gin.H, error) {
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
	return gin.H{
		"app version":          fmt.Sprintf("%s(%s) - %s", web.Version, web.BuildTime, viper.GetString("env")),
		"app root":             pwd,
		"who-am-i":             fmt.Sprintf("%s@%s", hu.Username, hn),
		"go version":           runtime.Version(),
		"go root":              runtime.GOROOT(),
		"go runtime":           runtime.NumGoroutine(),
		"go last gc":           time.Unix(0, int64(mem.LastGC)).String(),
		"os cpu":               runtime.NumCPU(),
		"os ram(free/total)":   fmt.Sprintf("%dM/%dM", ifo.Freeram/1024/1024, ifo.Totalram/1024/1024),
		"os swap(free/total)":  fmt.Sprintf("%dM/%dM", ifo.Freeswap/1024/1024, ifo.Totalswap/1024/1024),
		"go memory(alloc/sys)": fmt.Sprintf("%dM/%dM", mem.Alloc/1024/1024, mem.Sys/1024/1024),
		"os time":              time.Now(),
		"os arch":              fmt.Sprintf("%s(%s)", runtime.GOOS, runtime.GOARCH),
		"os uptime":            (time.Duration(ifo.Uptime) * time.Second).String(),
		"os loads":             ifo.Loads,
		"os procs":             ifo.Procs,
	}, nil
}
func (p *Engine) _networkStatus() (gin.H, error) {
	sts := gin.H{}
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
func (p *Engine) _cacheStatus() ([]string, error) {
	c := p.Redis.Get()
	defer c.Close()
	sts, err := redis.String(c.Do("INFO"))
	if err != nil {
		return nil, err
	}
	return strings.Split(sts, "\n"), nil
}

func (p *Engine) _dbStatus() (gin.H, error) {
	val := gin.H{
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

func (p *Engine) _jobsStatus() (gin.H, error) {
	cfg := p.Server.GetConfig()
	return gin.H{
		"tasks":             p.Server.GetRegisteredTaskNames(),
		"broker":            cfg.Broker,
		"result broker":     cfg.ResultBackend,
		"results expire in": cfg.ResultsExpireIn,
		"default queue":     cfg.DefaultQueue,
	}, nil
}
func (p *Engine) getAdminSiteStatus(c *gin.Context, lang string, data gin.H) (string, error) {
	tpl := "site-admin-status"
	var err error
	data["title"] = p.I18n.T(lang, "site.admin.status.title")
	data["os"], err = p._osStatus()
	if err != nil {
		return tpl, err
	}
	data["network"], err = p._networkStatus()
	if err != nil {
		return tpl, err
	}
	data["jobs"], err = p._jobsStatus()
	if err != nil {
		return tpl, err
	}
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
