package users

import (
	"fmt"
	"github.com/cns1rius/imgstore/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	UserName string `form:"username"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func Login(c *gin.Context) {
	var _user User
	_ = c.ShouldBind(&_user)

	if _user.UserName == "" || _user.Password == "" {
		c.HTML(http.StatusOK, "login.html", gin.H{"error": "用户名或密码不能为空"})
		//return
	}

	// 查询user表
	DB := db
	// 分配cookie
	//err := config.SetCookie(c, _user.UserName)
	//
	//// 成功则跳转
	//if err == 0 {
	//	c.Redirect(http.StatusMovedPermanently, "/")
	//}
}

func Register(c *gin.Context) {
	var _user User
	_ = c.ShouldBind(&_user)

	if _user.UserName == "" || _user.Password == "" || _user.Phone == "" {
		c.HTML(http.StatusOK, "login.html", gin.H{"error": "输入不能为空"})
		return
	}

	// 添加到user表 分配私有图片表

	c.Redirect(http.StatusMovedPermanently, "/login")
}

func Upload(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)

		// 上传文件至指定目录
		// c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
