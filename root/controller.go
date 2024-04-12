package root

import (
	"github.com/cns1rius/imgstore/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexPage(c *gin.Context) {
	if err := config.Verify(c); err != "" { // 应该为cookie鉴权
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})

	// todo 根据cookie显示所能访问的图片及信息
	// todo 显示typeTags“文件夹” 点击整框url get 跳转
	// todo 公有文件夹进去再有分类
}
