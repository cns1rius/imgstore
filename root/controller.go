package root

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexPage(c *gin.Context) {
	cookie, _ := c.Cookie("gin_cookie")
	if cookie == "" { // 应该为cookie鉴权
		c.Redirect(http.StatusMovedPermanently, "/login")
	}

	// 根据cookie显示所能访问的图片及信息
}
