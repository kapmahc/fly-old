package erp

import "github.com/kapmahc/fly/web"

// Catalog catalog
type Catalog struct {
	web.Model

	Name        string    `json:"name"`
	Description string    `json:"description"`
	ParentID    uint      `json:"parentId"`
	Products    []Product `json:"products" gorm:"many2many:erp_products_catalogs;"`
}

// TableName table name
func (Catalog) TableName() string {
	return "erp_catalogs"
}

// Vendor vendor
type Vendor struct {
	web.Model

	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"products"`
}

// TableName table name
func (Vendor) TableName() string {
	return "erp_vendors"
}

// Product product
type Product struct {
	web.Model

	Name        string    `json:"name"`
	Description string    `json:"description"`
	VendorID    uint      `json:"vendorId"`
	Vendor      Vendor    `json:"vendor"`
	Variants    []Variant `json:"variants"`
	Catalogs    []Catalog `json:"catalogs" gorm:"many2many:erp_products_catalogs;"`
}

// TableName table name
func (Product) TableName() string {
	return "erp_products"
}

// Variant variant
type Variant struct {
	web.Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Cost        float64    `json:"cost"`
	Sku         string     `json:"sku"`
	Weight      float64    `json:"weight"`
	Height      float64    `json:"height"`
	Length      float64    `json:"length"`
	Width       float64    `json:"width"`
	ProductID   uint       `json:"productId"`
	Product     Product    `json:"product"`
	Properties  []Property `json:"properties"`
}

// TableName table name
func (Variant) TableName() string {
	return "erp_variants"
}

// Property property
type Property struct {
	web.Model

	Key       string  `json:"key"`
	Val       string  `json:"val"`
	VariantID uint    `json:"variantId"`
	Variant   Variant `json:"variant"`
}

// TableName table name
func (Property) TableName() string {
	return "erp_properties"
}
