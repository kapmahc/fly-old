package auth

import (
	"crypto/x509/pkix"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	"github.com/fvbock/endless"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/kapmahc/fly/web"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/spf13/viper"
	"github.com/steinbacher/goose"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
	"golang.org/x/text/language"
)

const (
	postgresqlDriver = "postgres"
)

// Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start the app server",
			Action:  Action(p.runServer),
		},
		{
			Name:  "seo",
			Usage: "generate sitemap.xml.gz/rss.atom/robots.txt ...etc",
			Action: Action(func(*cli.Context, *inject.Graph) error {
				if err := p.writeSitemap(); err != nil {
					return err
				}
				return nil
			}),
		},
		{
			Name:    "worker",
			Aliases: []string{"w"},
			Usage:   "start the worker progress",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Value: viper.GetString("app.name"),
					Usage: "worker's name",
				},
			},
			Action: Action(p.runWorker),
		},
		{
			Name:    "redis",
			Aliases: []string{"re"},
			Usage:   "open redis connection",
			Action:  web.Action(p.connectRedis),
		},
		{
			Name:    "cache",
			Aliases: []string{"c"},
			Usage:   "cache operations",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Usage:   "list all cache keys",
					Aliases: []string{"l"},
					Action:  Action(p.listCacheItems),
				},
				{
					Name:    "clear",
					Usage:   "clear cache items",
					Aliases: []string{"c"},
					Action: Action(func(*cli.Context, *inject.Graph) error {
						return p.Cache.Flush()
					}),
				},
			},
		},
		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "database operations",
			Subcommands: []cli.Command{
				{
					Name:    "example",
					Usage:   "scripts example for create database and user",
					Aliases: []string{"e"},
					Action:  web.Action(p.databaseExample),
				},
				{
					Name:    "migrate",
					Usage:   "migrate the DB to the most recent version available",
					Aliases: []string{"m"},
					Action:  web.Action(p.migrateDatabase),
				},
				{
					Name:    "rollback",
					Usage:   "roll back the version by 1",
					Aliases: []string{"r"},
					Action:  web.Action(p.rollbackDatabase),
				},
				{
					Name:    "version",
					Usage:   "dump the migration status for the current DB",
					Aliases: []string{"v"},
					Action:  web.Action(p.databaseVersion),
				},
				{
					Name:    "connect",
					Usage:   "connect database",
					Aliases: []string{"c"},
					Action:  web.Action(p.connectDatabase),
				},
				{
					Name:    "create",
					Usage:   "create database",
					Aliases: []string{"n"},
					Action:  web.Action(p.createDatabase),
				},
				{
					Name:    "drop",
					Usage:   "drop database",
					Aliases: []string{"d"},
					Action:  web.Action(p.dropDatabase),
				},
			},
		},
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate file template",
			Subcommands: []cli.Command{
				{
					Name:    "config",
					Aliases: []string{"c"},
					Usage:   "generate config file",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "environment, e",
							Value: "development",
							Usage: "environment, like: development, production, stage, test...",
						},
					},
					Action: p.generateConfig,
				},
				{
					Name:    "nginx",
					Aliases: []string{"ng"},
					Usage:   "generate nginx.conf",
					Action:  web.Action(p.generateNginxConf),
				},
				{
					Name:    "openssl",
					Aliases: []string{"ssl"},
					Usage:   "generate ssl certificates",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name, n",
							Usage: "name",
						},
						cli.StringFlag{
							Name:  "country, c",
							Value: "Earth",
							Usage: "country",
						},
						cli.StringFlag{
							Name:  "organization, o",
							Value: "Mother Nature",
							Usage: "organization",
						},
						cli.IntFlag{
							Name:  "years, y",
							Value: 1,
							Usage: "years",
						},
					},
					Action: p.generateSsl,
				},
				{
					Name:    "migration",
					Usage:   "generate migration file",
					Aliases: []string{"m"},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name, n",
							Usage: "name",
						},
					},
					Action: web.Action(p.generateMigration),
				},
				{
					Name:    "locale",
					Usage:   "generate locale file",
					Aliases: []string{"l"},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name, n",
							Usage: "locale name",
						},
					},
					Action: web.Action(p.generateLocale),
				},
			},
		},
		{
			Name:    "routes",
			Aliases: []string{"rt"},
			Usage:   "print out all defined routes",
			Action:  Action(p.printRoutes),
		},
		{
			Name:  "i18n",
			Usage: "i18n operations",
			Subcommands: []cli.Command{
				{
					Name:    "sync",
					Aliases: []string{"s"},
					Usage:   "sync locales from files",
					Action: Action(func(*cli.Context, *inject.Graph) error {
						return p.I18n.Sync("locales")
					}),
				},
			},
		},
	}
}

func (p *Engine) printRoutes(*cli.Context, *inject.Graph) error {

	tpl := "%-24s %s\n"
	fmt.Printf(tpl, "NAME", "PATH")
	p.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pat, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		hnd := route.GetHandler()
		if hnd == nil {
			return nil
		}
		fmt.Printf(
			tpl,
			route.GetName(),
			pat,
			// runtime.FuncForPC(reflect.ValueOf(route.GetHandler()).Pointer()).Name(),
		)
		return nil
	})
	return nil
}

func (p *Engine) generateConfig(c *cli.Context) error {
	const fn = "config.toml"
	if _, err := os.Stat(fn); err == nil {
		return fmt.Errorf("file %s already exists", fn)
	}
	fmt.Printf("generate file %s\n", fn)

	viper.Set("env", c.String("environment"))
	args := viper.AllSettings()
	fd, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer fd.Close()
	end := toml.NewEncoder(fd)
	err = end.Encode(args)

	return err

}

func (p *Engine) generateNginxConf(*cli.Context) error {
	const tpl = `
		server {
		  listen 80;
		  server_name {{.Name}};
		  rewrite ^(.*) https://$host$1 permanent;
		}
		upstream {{.Name}}_prod {
		  server localhost:{{.Port}} fail_timeout=0;
		}
		server {
		  listen 443;
		  ssl  on;
		  ssl_certificate  /etc/ssl/certs/{{.Name}}.crt;
		  ssl_certificate_key  /etc/ssl/private/{{.Name}}.key;
		  ssl_session_timeout  5m;
		  ssl_protocols  SSLv2 SSLv3 TLSv1;
		  ssl_ciphers  RC4:HIGH:!aNULL:!MD5;
		  ssl_prefer_server_ciphers  on;
		  client_max_body_size 4G;
		  keepalive_timeout 10;
		  proxy_buffers 16 64k;
		  proxy_buffer_size 128k;
		  server_name {{.Name}};
		  root {{.Root}}/public;
		  index index.html;
		  access_log /var/log/nginx/{{.Name}}.access.log;
		  error_log /var/log/nginx/{{.Name}}.error.log;
		  location ~* \.(?:css|js)$ {
		    gzip_static on;
		    expires max;
		    access_log off;
		    add_header Cache-Control "public";
		  }
		  location ~* \.(?:jpg|jpeg|gif|png|ico|cur|gz|svg|svgz|mp4|ogg|ogv|webm|htc)$ {
		    expires 1M;
		    expires max;
		    access_log off;
		    add_header Cache-Control "public";
		  }
		  location ~* \.(?:rss|atom)$ {
		    expires 12h;
		    access_log off;
		    add_header Cache-Control "public";
		  }
		  location / {
		    proxy_set_header X-Forwarded-Proto https;
		    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		    proxy_set_header Host $http_host;
		    proxy_set_header X-Real-IP $remote_addr;
		    proxy_redirect off;
		    proxy_pass http://{{.Name}}_prod;
		    # limit_req zone=one;
		  }
		}
		`
	t, err := template.New("").Parse(tpl)
	if err != nil {
		return err
	}
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	name := viper.GetString("server.name")
	fn := path.Join("etc", "nginx", "sites-enabled", name+".conf")
	if err = os.MkdirAll(path.Dir(fn), 0700); err != nil {
		return err
	}
	fmt.Printf("generate file %s\n", fn)
	fd, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()

	return t.Execute(fd, struct {
		Name    string
		Port    int
		Root    string
		Version string
	}{
		Name:    name,
		Port:    viper.GetInt("http.port"),
		Root:    pwd,
		Version: "v1",
	})
}

func (p *Engine) generateSsl(c *cli.Context) error {
	name := c.String("name")
	if len(name) == 0 {
		cli.ShowCommandHelp(c, "openssl")
		return nil
	}
	root := path.Join("etc", "ssl", name)

	key, crt, err := CreateCertificate(
		true,
		pkix.Name{
			Country:      []string{c.String("country")},
			Organization: []string{c.String("organization")},
		},
		c.Int("years"),
	)
	if err != nil {
		return err
	}

	fnk := path.Join(root, "key.pem")
	fnc := path.Join(root, "crt.pem")

	fmt.Printf("generate pem file %s\n", fnk)
	err = WritePemFile(fnk, "RSA PRIVATE KEY", key, 0600)
	fmt.Printf("test: openssl rsa -noout -text -in %s\n", fnk)

	if err == nil {
		fmt.Printf("generate pem file %s\n", fnc)
		err = WritePemFile(fnc, "CERTIFICATE", crt, 0444)
		fmt.Printf("test: openssl x509 -noout -text -in %s\n", fnc)
	}
	if err == nil {
		fmt.Printf(
			"verify: diff <(openssl rsa -noout -modulus -in %s) <(openssl x509 -noout -modulus -in %s)",
			fnk,
			fnc,
		)
	}
	fmt.Println()
	return err
}

func (p *Engine) generateMigration(c *cli.Context) error {
	name := c.String("name")
	if len(name) == 0 {
		cli.ShowCommandHelp(c, "migration")
		return nil
	}
	cfg, err := dbConf()
	if err != nil {
		return err
	}
	if err = os.MkdirAll(cfg.MigrationsDir, 0700); err != nil {
		return err
	}
	file, err := goose.CreateMigration(name, "sql", cfg.MigrationsDir, time.Now())
	if err != nil {
		return err
	}

	fmt.Printf("generate file %s\n", file)
	return nil
}

func (p *Engine) generateLocale(c *cli.Context) error {
	name := c.String("name")
	if len(name) == 0 {
		cli.ShowCommandHelp(c, "locale")
		return nil
	}
	lng, err := language.Parse(name)
	if err != nil {
		return err
	}
	const root = "locales"
	if err = os.MkdirAll(root, 0700); err != nil {
		return err
	}
	file := path.Join(root, fmt.Sprintf("%s.ini", lng.String()))
	fmt.Printf("generate file %s\n", file)
	fd, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()
	return err
}

func (p *Engine) databaseExample(*cli.Context) error {
	drv := viper.GetString("database.driver")
	args := viper.GetStringMapString("database.args")
	var err error
	switch drv {
	case postgresqlDriver:
		fmt.Printf("CREATE USER %s WITH PASSWORD '%s';\n", args["user"], args["password"])
		fmt.Printf("CREATE DATABASE %s WITH ENCODING='UTF8';\n", args["dbname"])
		fmt.Printf("GRANT ALL PRIVILEGES ON DATABASE %s TO %s;\n", args["dbname"], args["user"])
	default:
		err = fmt.Errorf("unknown driver %s", drv)
	}
	return err
}
func (p *Engine) migrateDatabase(*cli.Context) error {
	conf, err := dbConf()
	if err != nil {
		return err
	}

	target, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
	if err != nil {
		return err
	}

	return goose.RunMigrations(conf, conf.MigrationsDir, target)
}

func (p *Engine) rollbackDatabase(*cli.Context) error {
	conf, err := dbConf()
	if err != nil {
		return err
	}

	current, err := goose.GetDBVersion(conf)
	if err != nil {
		return err
	}

	previous, err := goose.GetPreviousDBVersion(conf.MigrationsDir, current)
	if err != nil {
		return err
	}

	return goose.RunMigrations(conf, conf.MigrationsDir, previous)
}

func (p *Engine) databaseVersion(*cli.Context) error {
	conf, err := dbConf()
	if err != nil {
		return err
	}

	// collect all migrations
	migrations, err := goose.CollectMigrations(conf.MigrationsDir)
	if err != nil {
		return err
	}

	db, err := goose.OpenDBFromDBConf(conf)
	if err != nil {
		return err
	}
	defer db.Close()

	// must ensure that the version table exists if we're running on a pristine DB
	if _, err = goose.EnsureDBVersion(conf, db); err != nil {
		return err
	}

	fmt.Println("    Applied At                  Migration")
	fmt.Println("    =======================================")
	for _, m := range migrations {
		if err = printMigrationStatus(db, m.Version, filepath.Base(m.Source)); err != nil {
			return err
		}
	}
	return nil
}

func (p *Engine) connectDatabase(*cli.Context) error {
	drv := viper.GetString("database.driver")
	args := viper.GetStringMapString("database.args")
	var err error
	switch drv {
	case postgresqlDriver:
		err = web.Shell("psql",
			"-h", args["host"],
			"-p", args["port"],
			"-U", args["user"],
			args["dbname"],
		)
	default:
		err = fmt.Errorf("unknown driver %s", drv)
	}
	return err
}

func (p *Engine) createDatabase(*cli.Context) error {
	drv := viper.GetString("database.driver")
	args := viper.GetStringMapString("database.args")
	var err error
	switch drv {
	case postgresqlDriver:
		err = web.Shell("psql",
			"-h", args["host"],
			"-p", args["port"],
			"-U", "postgres",
			"-c", fmt.Sprintf(
				"CREATE DATABASE %s WITH ENCODING='UTF8'",
				args["dbname"],
			),
		)
	default:
		err = fmt.Errorf("unknown driver %s", drv)
	}
	return err
}

func (p *Engine) dropDatabase(*cli.Context) error {
	drv := viper.GetString("database.driver")
	args := viper.GetStringMapString("database.args")
	var err error
	switch drv {
	case postgresqlDriver:
		err = web.Shell("psql",
			"-h", args["host"],
			"-p", args["port"],
			"-U", "postgres",
			"-c", fmt.Sprintf("DROP DATABASE %s", args["dbname"]),
		)
	default:
		err = fmt.Errorf("unknown driver %s", drv)
	}
	return err
}

func (p *Engine) listCacheItems(*cli.Context, *inject.Graph) error {
	keys, err := p.Cache.Keys()
	if err != nil {
		return err
	}
	for _, k := range keys {
		fmt.Println(k)
	}
	return nil
}

func (p *Engine) connectRedis(*cli.Context) error {
	return web.Shell(
		"redis-cli",
		"-h", viper.GetString("redis.host"),
		"-p", viper.GetString("redis.port"),
		"-n", viper.GetString("redis.db"),
	)
}

func (p *Engine) runWorker(c *cli.Context, _ *inject.Graph) error {
	name := c.String("name")
	if name == "" {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	web.Walk(func(en web.Engine) error {
		en.Do()
		return nil
	})

	return p.Server.NewWorker(name).Launch()
}

func (p *Engine) runServer(*cli.Context, *inject.Graph) error {
	port := viper.GetInt("server.port")
	log.Infof(
		"application starting in %s on http://localhost:%d",
		viper.GetString("env"),
		port,
	)

	ng := negroni.New()
	ng.Use(negroni.NewRecovery())
	ng.Use(negronilogrus.NewMiddleware())
	ng.Use(negroni.NewStatic(http.Dir(path.Join("themes", viper.GetString("server.theme"), "assets"))))
	ng.Use(negroni.HandlerFunc(p.I18n.Middleware))
	ng.Use(negroni.HandlerFunc(p.layoutMiddleware))
	ng.UseHandler(csrf.Protect(
		[]byte(viper.GetString("secrets.csrf")),
		csrf.Secure(web.IsProduction()),
		csrf.CookieName("_csrf_token_"),
		csrf.FieldName("authenticity_token"),
		csrf.Path("/"),
	)(p.Router))

	addr := fmt.Sprintf(":%d", port)
	if web.IsProduction() {
		srv := endless.NewServer(addr, ng)
		srv.BeforeBegin = func(add string) {
			fd, err := os.OpenFile(path.Join("tmp", "pid"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
			if err != nil {
				log.Error(err)
				return
			}
			defer fd.Close()
			pid := syscall.Getpid()
			log.Printf("pid is %d", pid)
			fmt.Fprint(fd, pid)
		}
		return srv.ListenAndServe()
	}

	return http.ListenAndServe(addr, ng)
}

func (p *Engine) writeSitemap() error {
	sm := stm.NewSitemap()
	sm.SetDefaultHost(web.Home())
	sm.SetPublicPath(path.Join("themes", viper.GetString("server.theme"), "assets"))
	sm.SetCompress(true)
	sm.SetSitemapsPath("/")
	sm.Create()

	if err := web.Walk(func(en web.Engine) error {
		urls, err := en.Sitemap()
		if err != nil {
			return err
		}
		for _, u := range urls {
			sm.Add(u)
		}
		return nil
	}); err != nil {
		return err
	}
	if web.IsProduction() {
		sm.Finalize().PingSearchEngines()
	} else {
		sm.Finalize()
	}
	return nil
}
