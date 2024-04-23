package root

import (
	"fmt"
	"github.com/cns1rius/imgstore/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func IndexPage(c *gin.Context) {
	var (
		ImgTable []config.Img
		li       string
	)
	reqPath := c.Param("path")
	if err := config.Verify(c); err != "" { // 应该为cookie鉴权
		c.Redirect(http.StatusFound, "/login")
		return
	}
	ImgTable = config.ImgPermissionImg(strconv.Itoa(config.GetCookieId(c)))

	if reqPath == "" {
		ImgTable = append(ImgTable, config.ImgPermissionImg("0")...)
		li = Sprintf(ImgTable, "")
	} else if reqPath == "私有" {
		li = Sprintf(ImgTable, "")
	} else {
		li = Sprintf(ImgTable, reqPath)
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{"li": li})

	// todo 根据cookie显示所能访问的图片及信息
	// todo 显示typeTags“文件夹” 点击整框url get 跳转
	// todo 公有文件夹进去再有分类
}

func Sprintf(ImgTable []config.Img, reqPath string) string {
	var li string
	for _, Img := range ImgTable {
		if strings.Contains(Img.Path, reqPath) {
			li += fmt.Sprintf("%s", Img.Path)
		}
	}
	return li
}

func DeleteImg(c *gin.Context) {
	imgIDs := strings.Split(c.Query("id"), ",")
	imgIDList := make([]int, len(imgIDs))
	for i, imgID := range imgIDs {
		imgIDList[i], _ = strconv.Atoi(imgID)
	}
	userID := config.GetCookieId(c)
	config.ImgDelete(imgIDList, strconv.Itoa(userID))
	c.Redirect(http.StatusFound, "/")
}
