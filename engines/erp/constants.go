package erp

const (
	// RoleSeller seller role name
	RoleSeller = "seller"
)

var (
	// iso4217
	// https://en.wikipedia.org/wiki/ISO_4217
	currencyCodes = map[string]string{
		"usd": "$",
		"eur": "€",
		"cny": "￥",
		"gbp": "￡",
	}
)
