package shop

import (
	"time"

	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
)

// Catalog catalog
type Catalog struct {
	web.Model

	Name        string    `json:"name"`
	Description string    `json:"description"`
	ParentID    uint      `json:"parentId"`
	Products    []Product `json:"products" gorm:"many2many:shop_products_catalogs;"`
}

// TableName table name
func (Catalog) TableName() string {
	return "shop_catalogs"
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
	return "shop_vendors"
}

// Product product
type Product struct {
	web.Model

	Name        string    `json:"name"`
	Description string    `json:"description"`
	VendorID    uint      `json:"vendorId"`
	Vendor      Vendor    `json:"vendor"`
	Variants    []Variant `json:"variants"`
	Catalogs    []Catalog `json:"catalogs" gorm:"many2many:shop_products_catalogs;"`
}

// TableName table name
func (Product) TableName() string {
	return "shop_products"
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
	return "shop_variants"
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
	return "shop_properties"
}

// Address address
type Address struct {
	web.Model

	FirstName  string    `json:"firstName"`
	MiddleName string    `json:"middleName"`
	LastName   string    `json:"lastName"`
	Zip        string    `json:"zip"`
	Apt        string    `json:"apt"`
	Street     string    `json:"street"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	Country    string    `json:"country"`
	Phone      string    `json:"phone"`
	UserID     uint      `json:"userId"`
	User       auth.User `json:"user"`
}

// TableName table name
func (Address) TableName() string {
	return "shop_addresses"
}

// Order order
type Order struct {
	web.Model

	Number          string     `json:"number"`
	ItemTotal       float64    `json:"itemTotal"`
	AdjustmentTotal float64    `json:"adjustmentTotal"`
	PaymentTotal    float64    `json:"paymentTotal"`
	Total           float64    `json:"total"`
	State           string     `json:"state"`
	ShipmentState   string     `json:"shipmentState"`
	PaymentState    string     `json:"paymentState"`
	UserID          uint       `json:"userId"`
	User            auth.User  `json:"user"`
	LineItems       []LineItem `json:"lineItems"`
}

// TableName table name
func (Order) TableName() string {
	return "shop_orders"
}

// LineItem line-item
type LineItem struct {
	web.Model

	Price     float64 `json:"price"`
	Quantity  uint    `json:"quantity"`
	VariantID uint    `json:"variantId"`
	Variant   Variant `json:"variant"`
	OrderID   uint    `json:"orderId"`
	Order     Order   `json:"order"`
}

// TableName table name
func (LineItem) TableName() string {
	return "shop_line_items"
}

// PaymentMethod payment-method
type PaymentMethod struct {
	web.Model

	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

// TableName table name
func (PaymentMethod) TableName() string {
	return "shop_payment_methods"
}

// Payment payment
type Payment struct {
	web.Model

	Amount          float64       `json:"amount"`
	OrderID         uint          `json:"orderId"`
	Order           Order         `json:"order"`
	PaymentMethodID uint          `json:"paymentMethodId"`
	PaymentMethod   PaymentMethod `json:"paymentMethod"`
	State           string        `json:"state"`
	ResponseCode    string        `json:"responseCode"`
	AvsResponse     string        `json:"AvsResponse"`
}

// TableName table name
func (Payment) TableName() string {
	return "shop_payments"
}

// Zone zone
type Zone struct {
	web.Model
	Name   string  `json:"name"`
	States []State `json:"states"`
}

// TableName table name
func (Zone) TableName() string {
	return "shop_zones"
}

// Country country
type Country struct {
	web.Model
	Name   string  `json:"name"`
	States []State `json:"states"`
}

// TableName table name
func (Country) TableName() string {
	return "shop_countries"
}

// State state
type State struct {
	web.Model
	Name      string  `json:"name"`
	CountryID uint    `json:"countryId"`
	Country   Country `json:"country"`
	ZoneID    uint    `json:"zoneID"`
	Zone      Zone    `json:"zone"`
}

// TableName table name
func (State) TableName() string {
	return "shop_states"
}

// ShippingMethod shipping-method
type ShippingMethod struct {
	web.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	ZoneID      uint   `json:"zoneID"`
	Zone        Zone   `json:"zone"`
}

// TableName table name
func (ShippingMethod) TableName() string {
	return "shop_shipping_methods"
}

// Shipment shipment
type Shipment struct {
	web.Model
	Number           string         `json:"number"`
	Tracking         string         `json:"tracking"`
	Cost             float64        `json:"cost"`
	ShippedAt        *time.Time     `json:"shippedAt"`
	OrderID          uint           `json:"orderID"`
	Order            Order          `json:"order"`
	State            string         `json:"state"`
	ShippingMethodID uint           `json:"shippingMethoidID"`
	ShippingMethod   ShippingMethod `json:"shippingMethod"`
	AddressID        uint           `json:"addressID"`
	Address          Address        `json:"address"`
}

// TableName table name
func (Shipment) TableName() string {
	return "shop_shipments"
}

// ReturnAuthorization return-authorization
type ReturnAuthorization struct {
	web.Model
	Number    string     `json:"number"`
	State     string     `json:"state"`
	Amount    float64    `json:"amount"`
	OrderID   uint       `json:"orderID"`
	Order     Order      `json:"order"`
	Reason    string     `json:"reason"`
	EnterByID uint       `json:"enterById"`
	EnterBy   *auth.User `json:"enterBy"`
	EnterAt   time.Time  `json:"enterAt"`
}

// TableName table name
func (ReturnAuthorization) TableName() string {
	return "shop_return_authorizations"
}

// InventoryUnit inventory-unit
type InventoryUnit struct {
	web.Model
	State                 string              `json:"state"`
	Quantity              uint                `json:"quantity"`
	VariantID             uint                `json:"variantID"`
	Variant               Variant             `json:"variant"`
	OrderID               uint                `json:"orderID"`
	Order                 Order               `json:"order"`
	ShipmentID            uint                `json:"shipmentID"`
	Shipment              Shipment            `json:"shipment"`
	ReturnAuthorizationID uint                `json:"returnAuthorizationId"`
	ReturnAuthorization   ReturnAuthorization `json:"returnAuthorization"`
}

// TableName table name
func (InventoryUnit) TableName() string {
	return "shop_return_inventory_units"
}

// Chargeback chargeback
type Chargeback struct {
	web.Model
	State      string    `json:"state"`
	OrderID    uint      `json:"orderID"`
	Order      Order     `json:"order"`
	OperatorID uint      `json:"operatorID"`
	Operator   auth.User `json:"operator"`
	Amount     float64   `json:"amount"`
}

// TableName table name
func (Chargeback) TableName() string {
	return "shop_chargebacks"
}
