package vpn

import (
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexLogs(c *gin.Context) (interface{}, error) {
	data["title"] = p.I18n.T(lang, "ops.vpn.logs.index.title")
	tpl := "ops-vpn-logs-index"
	var total int64
	if err := p.Db.Model(&Log{}).Count(&total).Error; err != nil {
		return nil, err
	}
	pag := web.NewPagination(c.Request, total)

	var items []Log
	if err := p.Db.
		Limit(pag.Limit()).Offset(pag.Offset()).
		Find(&items).Error; err != nil {
		return nil, err
	}
	for _, b := range items {
		pag.Items = append(pag.Items, b)
	}
	data["pager"] = pag
	return tpl, nil
}
