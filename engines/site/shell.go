package site

import (
	"crypto/x509/pkix"
	"encoding/xml"
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"github.com/google/uuid"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/kapmahc/sky"
	"github.com/spf13/viper"
	"github.com/steinbacher/goose"
	"github.com/unrolled/render"
	"github.com/urfave/cli"
	"golang.org/x/text/language"
	"golang.org/x/tools/blog/atom"
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
			Action:  sky.IocAction(p.runServer),
		},
		{
			Name:  "seo",
			Usage: "generate sitemap.xml.gz/rss.atom/robots.txt ...etc",
			Action: sky.IocAction(func(*cli.Context) error {
				root := "public"
				os.MkdirAll(root, 0755)
				if err := p.writeSitemap(root); err != nil {
					return err
				}
				for _, lang := range viper.GetStringSlice("languages") {
					if err := p.writeRssAtom(root, lang); err != nil {
						return err
					}
				}
				if err := p.writeRobotsTxt(root); err != nil {
					return err
				}
				if err := p.writeGoogleVerify(root); err != nil {
					return err
				}
				if err := p.writeBaiduVerify(root); err != nil {
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
			Action: sky.IocAction(p.runWorker),
		},
		{
			Name:    "redis",
			Aliases: []string{"re"},
			Usage:   "open redis connection",
			Action:  sky.CfgAction(p.connectRedis),
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
					Action:  sky.IocAction(p.listCacheItems),
				},
				{
					Name:    "clear",
					Usage:   "clear cache items",
					Aliases: []string{"c"},
					Action: sky.IocAction(func(*cli.Context) error {
						return p.CacheStore.Flush()
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
					Action:  sky.CfgAction(p.databaseExample),
				},
				{
					Name:    "migrate",
					Usage:   "migrate the DB to the most recent version available",
					Aliases: []string{"m"},
					Action:  sky.CfgAction(p.migrateDatabase),
				},
				{
					Name:    "rollback",
					Usage:   "roll back the version by 1",
					Aliases: []string{"r"},
					Action:  sky.CfgAction(p.rollbackDatabase),
				},
				{
					Name:    "version",
					Usage:   "dump the migration status for the current DB",
					Aliases: []string{"v"},
					Action:  sky.CfgAction(p.databaseVersion),
				},
				{
					Name:    "connect",
					Usage:   "connect database",
					Aliases: []string{"c"},
					Action:  sky.CfgAction(p.connectDatabase),
				},
				{
					Name:    "create",
					Usage:   "create database",
					Aliases: []string{"n"},
					Action:  sky.CfgAction(p.createDatabase),
				},
				{
					Name:    "drop",
					Usage:   "drop database",
					Aliases: []string{"d"},
					Action:  sky.CfgAction(p.dropDatabase),
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
					Action:  sky.CfgAction(p.generateNginxConf),
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
					Action: sky.CfgAction(p.generateMigration),
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
					Action: sky.CfgAction(p.generateLocale),
				},
			},
		},
		{
			Name:    "routes",
			Aliases: []string{"rt"},
			Usage:   "print out all defined routes",
			Action: func(*cli.Context) error {
				rt := sky.NewRouter(render.Options{})
				sky.Walk(func(en sky.Engine) error {
					en.Mount(rt)
					return nil
				})
				tpl := "%-24s %s\n"
				fmt.Printf(tpl, "NAME", "PATH")
				return rt.Walk(func(name, method, path string) error {
					fmt.Printf(tpl, name, path)
					// fmt.Println(runtime.FuncForPC(reflect.ValueOf(r.Handler).Pointer()).Name())
					return nil
				})

			},
		},
		{
			Name:  "i18n",
			Usage: "i18n operations",
			Subcommands: []cli.Command{
				{
					Name:    "sync",
					Aliases: []string{"s"},
					Usage:   "sync locales from files",
					Action: sky.IocAction(func(*cli.Context) error {
						return p.I18n.Load("locales")
					}),
				},
			},
		},
	}
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
	t, err := template.New("").Parse(nginxConf)
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
		Ssl     bool
	}{
		Name:    name,
		Port:    viper.GetInt("server.port"),
		Root:    path.Join(pwd, "public"),
		Version: "v1",
		Ssl:     viper.GetBool("server.ssl"),
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
		err = sky.Shell("psql",
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
		err = sky.Shell("psql",
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
		err = sky.Shell("psql",
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

func (p *Engine) listCacheItems(*cli.Context) error {
	keys, err := p.CacheStore.Keys()
	if err != nil {
		return err
	}
	for _, k := range keys {
		fmt.Println(k)
	}
	return nil
}

func (p *Engine) connectRedis(*cli.Context) error {
	return sky.Shell(
		"redis-cli",
		"-h", viper.GetString("redis.host"),
		"-p", viper.GetString("redis.port"),
		"-n", viper.GetString("redis.db"),
	)
}

func (p *Engine) runWorker(c *cli.Context) error {
	name := c.String("name")
	if name == "" {
		cli.ShowSubcommandHelp(c)
		return nil
	}
	sky.Walk(func(en sky.Engine) error {
		for k, v := range en.Workers() {
			p.Server.Register(k, v)
		}
		return nil
	})

	return p.Server.Do(name)
}

func (p *Engine) runServer(*cli.Context) error {
	port := viper.GetInt("server.port")
	log.Infof(
		"application starting in %s on http://localhost:%d",
		viper.GetString("env"),
		port,
	)

	// TODO

	rt := sky.NewRouter(render.Options{})

	rt.Use(
		p.Jwt.CurrentUserMiddleware,
	)
	sky.Walk(func(en sky.Engine) error {
		en.Mount(rt)
		return nil
	})

	// ---------------
	return rt.Start(port)
}

func (p *Engine) writeSitemap(root string) error {
	sm := stm.NewSitemap()
	sm.SetDefaultHost(viper.GetString("server.frontend"))
	sm.SetPublicPath(root)
	sm.SetCompress(true)
	sm.SetSitemapsPath("/")
	sm.Create()

	if err := sky.Walk(func(en sky.Engine) error {
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
	if sky.IsProduction() {
		sm.Finalize().PingSearchEngines()
	} else {
		sm.Finalize()
	}
	return nil
}

func (p *Engine) writeRssAtom(root string, lang string) error {

	feed := atom.Feed{
		// Title:   p.I18n.T(lang, "site.title"),
		ID:      uuid.New().String(),
		Updated: atom.Time(time.Now()),
		Author:  &atom.Person{
		// Name:  p.I18n.T(lang, "site.author.name"),
		// Email: p.I18n.T(lang, "site.author.email"),
		},
		Entry: make([]*atom.Entry, 0),
	}
	home := viper.GetString("server.frontend")
	if err := sky.Walk(func(en sky.Engine) error {
		items, err := en.Atom(lang)
		if err != nil {
			return err
		}
		for _, it := range items {
			for i := range it.Link {
				it.Link[i].Href = home + it.Link[i].Href
			}
			feed.Entry = append(feed.Entry, it)
		}
		return nil
	}); err != nil {
		return err
	}
	fn := path.Join(root, fmt.Sprintf("rss-%s.atom", lang))
	log.Infof("generate file %s", fn)
	fd, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()
	enc := xml.NewEncoder(fd)
	return enc.Encode(feed)

}

func (p *Engine) writeRobotsTxt(root string) error {
	t, err := template.New("").Parse(robotsTxt)
	if err != nil {
		return err
	}
	fn := path.Join(root, "robots.txt")
	log.Infof("generate file %s", fn)
	fd, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()
	return t.Execute(fd, struct {
		Home string
	}{Home: viper.GetString("server.frontend")})
}

func (p *Engine) writeGoogleVerify(root string) error {
	var code string
	if err := p.Settings.Get("site.google.verify.code", &code); err != nil {
		return err
	}
	fn := path.Join(root, fmt.Sprintf("google%s.html", code))
	log.Infof("generate file %s", fn)
	fd, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = fmt.Fprintf(fd, "google-site-verification: google%s.html", code)
	return err

}

func (p *Engine) writeBaiduVerify(root string) error {
	var code string
	if err := p.Settings.Get("site.baidu.verify.code", &code); err != nil {
		return err
	}
	fn := path.Join(root, fmt.Sprintf("baidu_verify_%s.html", code))
	log.Infof("generate file %s", fn)
	fd, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = fd.WriteString(code)
	return err
}
