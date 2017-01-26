package migrate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kapmahc/fly"
)

// Store store
type Store struct {
	Logger  fly.Logger `inject:""`
	Dialect Dialect    `inject:""`
}

// Migrate migrate database to leatest version
func (p *Store) Migrate() error {
	items, err := p.Dialect.All()
	if err != nil {
		return err
	}
	for _, m := range items {
		if m.Applied {
			p.Logger.Info(m.File, "was applied")
			continue
		}
		p.Logger.Info("migrate", m.File)
		if err := p.Dialect.Up(m.File); err != nil {
			return err
		}
	}
	return nil
}

// Rollback rollback to previous version
func (p *Store) Rollback() error {
	cur, err := p.Dialect.Version()
	if err != nil {
		return err
	}
	p.Logger.Info("rollback", cur.File)

	return p.Dialect.Down(cur.File)
}

// Clear clear database
func (p *Store) Clear() error {
	items, err := p.Dialect.All()
	if err != nil {
		return err
	}
	for _, m := range items {
		if m.Applied {
			p.Logger.Info("rollback", m.File)
			up, err := m.Up()
			if err != nil {
				return err
			}
			if err := p.Dialect.Exec(up); err != nil {
				return err
			}
		}
	}

	return p.Dialect.Exec(fmt.Sprintf("drop table %s", Model{}.TableName()))
}

// Load load migrations
func (p *Store) Load(root, driver string) ([]string, error) {
	var items []string
	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		name := info.Name()
		if filepath.Ext(name) != EXT && len(name) <= len(FORMAT)+len(EXT)+1 {
			p.Logger.Warn("ingnore file", path)
			return nil
		}
		p.Logger.Info("find file", path)
		items = append(items, path)
		return nil
	}); err != nil {
		return nil, err
	}
	return items, nil
}

// Generate generate migration file
func (p *Store) Generate(root, driver, name string) error {
	root = filepath.Join(root, driver)
	os.MkdirAll(root, 0700)
	fn := filepath.Join(root, time.Now().Format(FORMAT)+"_"+name+EXT)
	p.Logger.Info("generate file", fn)
	fd, err := os.OpenFile(
		fn,
		os.O_WRONLY|os.O_EXCL|os.O_CREATE,
		0600,
	)
	if err != nil {
		return err
	}
	defer fd.Close()

	fd.WriteString(
		strings.Join([]string{
			UP,
			fmt.Sprintf("create table %s(id int primary key);", name),
			END,
			DOWN,
			fmt.Sprintf("drop table %s;", name),
			END,
		},
			"\n",
		),
	)
	return nil
}
