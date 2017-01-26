package migrate_test

import (
	"os"
	"testing"

	"github.com/kapmahc/fly"
	"github.com/kapmahc/fly/migrate"
)

func testReadSql(file string, t *testing.T) {
	m := migrate.Model{File: file}
	if up, err := m.Up(); err == nil {
		t.Logf(up)
	} else {
		t.Fatal(err)
	}
	if down, err := m.Down(); err == nil {
		t.Logf(down)
	} else {
		t.Fatal(err)
	}
}

func TestMigrate(t *testing.T) {
	const root = "db"
	const driver = "postgres"
	os.RemoveAll(root)
	st := migrate.Store{Logger: &fly.ConsoleLogger{}}
	if err := st.Generate(root, driver, "test"); err != nil {
		t.Fatal(err)
	}
	items, err := st.Load(root, driver)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", items)

	for _, f := range items {
		testReadSql(f, t)
	}
}
