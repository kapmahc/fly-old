package erp

import (
	"github.com/facebookgo/inject"
	"github.com/kapmahc/fly/engines/mall"
	"github.com/kapmahc/fly/web"
	"github.com/urfave/cli"
)

func (p *Engine) isTablesEmpty(args ...interface{}) (bool, error) {
	for _, arg := range args {
		var count int
		if err := p.Db.Model(arg).Count(&count).Error; err != nil {
			return false, err
		}
		if count > 0 {
			return false, nil
		}
	}
	return true, nil
}

func (p *Engine) initTags() error {
	if ok, err := p.isTablesEmpty(
		&mall.Tag{},
	); err != nil || !ok {
		return err
	}

	for _, n := range []string{
		"Books",
		"Musics",
		"Videos",
		"Tools",
	} {
		if err := p.Db.Create(&mall.Tag{
			Model: mall.Model{
				Name: n,
				Type: web.TypeMARKDOWN,
			},
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (p *Engine) loadSeed(*cli.Context, *inject.Graph) error {
	if err := p.initTags(); err != nil {
		return err
	}
	return nil
}
