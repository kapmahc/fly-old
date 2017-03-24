package erp

import gin "gopkg.in/gin-gonic/gin.v1"

// Mount web mount-points
func (p *Engine) Mount(rt *gin.Engine) {
	// rg := rt.Group("/erp", p.Jwt.MustSignInMiddleware, p.mustSellerMiddleware)
}
