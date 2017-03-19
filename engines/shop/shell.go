package shop

import (
	log "github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	"github.com/kapmahc/fly/engines/auth"
	"github.com/urfave/cli"
)

// Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:  "shop",
			Usage: "shop operations",
			Subcommands: []cli.Command{
				{
					Name:    "seed",
					Usage:   "loads the seed data",
					Aliases: []string{"s"},
					Action:  auth.Action(p.loadSeed),
				},
			},
		},
	}
}
func (p *Engine) isTablesEmpty(args ...interface{}) (bool, error) {
	for _, arg := range args {
		var count int
		if err := p.Db.Model(arg).Count(&count).Error; err != nil {
			return false, err
		}
		if count > 0 {
			log.Warningf("ingnore %T", arg)
			return false, nil
		}
	}
	return true, nil
}
func (p *Engine) initStates() error {
	if ok, err := p.isTablesEmpty(&Zone{}, &Country{}, &State{}); err != nil || !ok {
		return err
	}
	na := Zone{
		Name: "North America",
	}
	if err := p.Db.Create(&na).Error; err != nil {
		return err
	}
	usa := Country{
		Name: "United States of America",
	}
	if err := p.Db.Create(&usa).Error; err != nil {
		return err
	}
	for _, s := range []string{
		"Alabama", "Alaska", "Arizona", "Arkansas",
		"California", "Colorado", "Connecticut",
		"Delaware", "Florida", "Georgia", "Hawaii",
		"Idaho", "Illinois", "Indiana", "Iowa",
		"Kansas", "Kentucky", "Louisiana",
		"Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana",
		"Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", " North Dakota",
		"Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island",
		"South Carolina", "South Dakota",
		"Tennessee", "Texas", "Utah",
		"Vermont", "Virginia",
		"Washington", "West Virginia", "Wisconsin", "Wyoming",
	} {
		if err := p.Db.Create(&State{
			Name:      s,
			CountryID: usa.ID,
			ZoneID:    na.ID,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (p *Engine) initPaymentMethods() error {
	if ok, err := p.isTablesEmpty(&PaymentMethod{}); err != nil || !ok {
		return err
	}
	for _, pm := range []PaymentMethod{
		{
			Type:        "paypal",
			Name:        "Paypal",
			Description: "https://www.paypal.com/us/webapps/mpp/about",
			Active:      true,
		},
		{
			Type:        "alipay",
			Name:        "支付宝",
			Description: "https://www.alipay.com/",
			Active:      true,
		},
		{
			Type:        "weixin",
			Name:        "微信支付",
			Description: "https://pay.weixin.qq.com/wxzf_guide/index.shtml",
			Active:      true,
		},
	} {
		if err := p.Db.Create(&pm).Error; err != nil {
			return err
		}
	}
	return nil
}

func (p *Engine) initShippingMethods() error {
	if ok, err := p.isTablesEmpty(&ShippingMethod{}); err != nil || !ok {
		return err
	}
	var zones []Zone
	if err := p.Db.Find(&zones).Error; err != nil {
		return err
	}
	var items []interface{}
	for _, z := range zones {
		items = append(items, z)
	}
	for _, sm := range []ShippingMethod{
		{
			Name:        "USPS",
			Description: "https://www.usps.com/",
			Logo:        "https://www.usps.com/global-elements/header/images/utility-header/logo-sb.svg",
			Tracking:    "https://tools.usps.com/go/TrackConfirmAction_input",
			Active:      true,
		},
		{
			Name:        "UPS",
			Description: "https://www.ups.com/",
			Logo:        "https://www.ups.com/img/glo_ups_brandmark.gif",
			Tracking:    "https://www.ups.com/WebTracking/track",
			Active:      true,
		},
		{
			Name:        "FedEx",
			Description: "https://www.fedex.com/",
			Logo:        "https://images.fedex.com/images/c/t1/gh/logo-header-fedex.png",
			Tracking:    "https://www.fedex.com/apps/fedextrack",
			Active:      true,
		},
	} {
		if err := p.Db.Create(&sm).Error; err != nil {
			return err
		}
		if err := p.Db.Model(&sm).Association("Zones").Append(items...).Error; err != nil {
			return err
		}
	}
	return nil
}

func (p *Engine) loadSeed(*cli.Context, *inject.Graph) error {
	if err := p.initStates(); err != nil {
		return err
	}
	if err := p.initPaymentMethods(); err != nil {
		return err
	}
	if err := p.initShippingMethods(); err != nil {
		return err
	}
	return nil
}
