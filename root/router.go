package root

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/library"
		r.HandleContext(c)
	})
	r.GET("/library/*path", IndexPage)
	r.GET("/delete", DeleteImg)
}
