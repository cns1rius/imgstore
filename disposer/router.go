package disposer

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	//r.GET("/upload", UploadPage)
	r.POST("/upload", Upload)

	//r.GET("/spider", SpiderPage)
	r.POST("/spider", Spider)
}
