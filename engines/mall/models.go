package mall

import (
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
)

// Address address
type Address struct {
	web.Model

	FirstName  string
	MiddleName string
	LastName   string
	Zip        string
	Street     string
	City       string
	State      string
	Country    string
	Phone      string
	Email      string

	UserID uint
	User   auth.User
}

// TableName table name
func (Address) TableName() string {
	return "mall_addresses"
}

// Model base model
type Model struct {
	web.Model

	Name        string
	Type        string
	Description string
}

// Store store
type Store struct {
	Model

	Currency string

	OwnerID   uint
	Owner     auth.User
	AddressID uint
	Address   Address
}

// TableName table name
func (Store) TableName() string {
	return "mall_stores"
}

// Product product
type Product struct {
	Model

	VendorID uint
	Vendor   Vendor
	Tags     []Tag `gorm:"many2many:shop_products_tags;"`
	Variants []Variant
}

// TableName table name
func (Product) TableName() string {
	return "mall_products"
}

// Tag tag
type Tag struct {
	Model

	Products []Product `gorm:"many2many:shop_products_tags;"`
}

// TableName table name
func (Tag) TableName() string {
	return "mall_tags"
}

// Vendor vendor
type Vendor struct {
	Model

	Products []Product
}

// TableName table name
func (Vendor) TableName() string {
	return "mall_vendors"
}

// Variant variant
type Variant struct {
	web.Model

	Cost  float64
	Price float64

	ProductID  uint
	Product    Product
	Properties []Property
}

// TableName table name
func (Variant) TableName() string {
	return "mall_variants"
}

// Property property
type Property struct {
	web.Model

	Name  string
	Value string

	VariantID uint
	Variant   Variant
}

// TableName table name
func (Property) TableName() string {
	return "mall_properties"
}
