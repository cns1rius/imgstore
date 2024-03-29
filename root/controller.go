package root

import (
	"github.com/cns1rius/imgstore/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexPage(c *gin.Context) {
	if err := config.Verify(c); err != "" { // 应该为cookie鉴权
		c.Redirect(http.StatusFound, "/login")
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})

	// 根据cookie显示所能访问的图片及信息
}
