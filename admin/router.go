package admin

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	r.GET("/admin/*", func(c *gin.Context) {
		c.Request.URL.Path = "/library"
		r.HandleContext(c)
	})

	//r.GET("/admin", Admin)
	//r.POST("/admin", Manage)
	//
	//r.GET("/admin/login", LoginPage)
	//r.POST("/admin/login", Login)
	//
	//r.GET("/manage", func(c *gin.Context) {
	//	c.Request.URL.Path = "/admin"
	//	r.HandleContext(c)
	//})
}
