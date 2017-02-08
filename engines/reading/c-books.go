package reading

import "github.com/kapmahc/fly/engines/base"

// GetBooksScan scan books
// @router /books/scan [get]
func (p *Controller) GetBooksScan() {
	base.SendTask(scanBookTask)
	p.Data["json"] = map[string]interface{}{"ok": true}
	p.ServeJSON()
}
