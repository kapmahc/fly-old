package erp

import (
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	ag := rt.Group("/erp", p.Jwt.MustAdminMiddleware)
	// ---------
	// ---------
	ag.GET("/countries", auth.HTML(p.indexCountries))
	ag.GET("/countries/new", auth.HTML(p.createCountry))
	ag.POST("/countries/new", auth.HTML(p.createCountry))
	ag.GET("/countries/edit/:id", auth.HTML(p.updateCountry))
	ag.POST("/countries/edit/:id", auth.HTML(p.updateCountry))
	ag.DELETE("/countries/:id", web.JSON(p.destroyCountry))
	// ---------
	ag.GET("/states", auth.HTML(p.indexStates))
	ag.GET("/states/new", auth.HTML(p.createState))
	ag.POST("/states/new", auth.HTML(p.createState))
	ag.GET("/states/edit/:id", auth.HTML(p.updateState))
	ag.POST("/states/edit/:id", auth.HTML(p.updateState))
	ag.DELETE("/states/:id", web.JSON(p.destroyState))
	// ---------
	ag.GET("/zones", auth.HTML(p.indexZones))
	ag.GET("/zones/new", auth.HTML(p.createZone))
	ag.POST("/zones/new", auth.HTML(p.createZone))
	ag.GET("/zones/edit/:id", auth.HTML(p.updateZone))
	ag.POST("/zones/edit/:id", auth.HTML(p.updateZone))
	ag.DELETE("/zones/:id", web.JSON(p.destroyZone))
	// ---------
	ag.GET("/payment-methods", auth.HTML(p.indexPaymentMethods))
	ag.GET("/payment-methods/new", auth.HTML(p.createPaymentMethod))
	ag.POST("/payment-methods/new", auth.HTML(p.createPaymentMethod))
	ag.GET("/payment-methods/edit/:id", auth.HTML(p.updatePaymentMethod))
	ag.POST("/payment-methods/edit/:id", auth.HTML(p.updatePaymentMethod))
	ag.DELETE("/payment-methods/:id", web.JSON(p.destroyPaymentMethod))
	// ---------
	ag.GET("/shipping-methods", auth.HTML(p.indexShippingMethods))
	ag.GET("/shipping-methods/new", auth.HTML(p.createShippingMethod))
	ag.POST("/shipping-methods/new", auth.HTML(p.createShippingMethod))
	ag.GET("/shipping-methods/edit/:id", auth.HTML(p.updateShippingMethod))
	ag.POST("/shipping-methods/edit/:id", auth.HTML(p.updateShippingMethod))
	ag.DELETE("/shipping-methods/:id", web.JSON(p.destroyShippingMethod))
}
