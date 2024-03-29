package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Administrator struct {
	UserName string `form:"username"`
	Password string `form:"pwd"`
}

func Admin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.tmpl", gin.H{})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login.tmpl", gin.H{})
}

func Login(c *gin.Context) {
	var _admin Administrator
	_ = c.ShouldBind(&_admin)

}

func Manage(c *gin.Context) {

}
