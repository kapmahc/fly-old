package shop

import "github.com/kapmahc/fly/engines/base"

// https://baya.github.io/2015/09/17/%E7%94%B5%E5%AD%90%E5%95%86%E5%8A%A1%E7%B3%BB%E7%BB%9F%E5%9F%BA%E7%A1%80%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E5%92%8C%E6%B5%81%E7%A8%8B.html

// Catalog catalog
type Catalog struct {
	base.Model
	Name   string
	Parent uint
}

// TableName table name
func (Catalog) TableName() string {
	return "shop_catalogs"
}

// Country country
type Country struct {
	base.Model
	Name string
}

// TableName table name
func (Country) TableName() string {
	return "shop_countries"
}

// State state
type State struct {
	base.Model
	Country *Country
	Name    string
}
