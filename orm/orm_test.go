package orm_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/migrate"
	"github.com/kapmahc/fly/orm"
	_ "github.com/lib/pq"
)

func TestOrm(t *testing.T) {
	db, err := gorm.Open("postgres", "user=postgres dbname=fly_test sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	db.LogMode(true)
	var dia migrate.Dialect
	dia = &orm.MigrateDialect{Db: db}
	dia.Check()
}
