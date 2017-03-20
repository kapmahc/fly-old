package erp

import (
	"github.com/kapmahc/fly/engines/auth"
	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	ag := rt.Group("/erp", p.Jwt.MustAdminMiddleware)

	ag.GET("/payment-methods", auth.HTML(p.indexPaymentMethods))
	ag.GET("/payment-methods/new", auth.HTML(p.createPaymentMethod))
	ag.POST("/payment-methods/new", auth.HTML(p.createPaymentMethod))
	ag.GET("/payment-methods/edit/:id", auth.HTML(p.updatePaymentMethod))
	ag.POST("/payment-methods/edit/:id", auth.HTML(p.updatePaymentMethod))
	ag.DELETE("/payment-methods/:id", web.JSON(p.destroyPaymentMethod))

	ag.GET("/shipping-methods", auth.HTML(p.indexShippingMethods))
	ag.GET("/shipping-methods/new", auth.HTML(p.createShippingMethod))
	ag.POST("/shipping-methods/new", auth.HTML(p.createShippingMethod))
	ag.GET("/shipping-methods/edit/:id", auth.HTML(p.updateShippingMethod))
	ag.POST("/shipping-methods/edit/:id", auth.HTML(p.updateShippingMethod))
	ag.DELETE("/shipping-methods/:id", web.JSON(p.destroyShippingMethod))
}
