package users

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/login", LoginPage)
	r.POST("/login", Login)

	r.GET("/register", RegisterPage)
	r.POST("/register", Register)

	//r.GET("/invoke", test)

}
