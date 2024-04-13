package users

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/cns1rius/imgstore/config"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

type UserForm struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "user/login.tmpl", gin.H{})
}

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "user/register.tmpl", gin.H{})
}

func Login(c *gin.Context) {
	var UserTable config.User
	_userForm := UserForm{
		UserName: c.PostForm("username"),
		Password: c.PostForm("password"),
	}

	if _userForm.UserName == "" || _userForm.Password == "" {
		c.HTML(http.StatusOK, "user/login.tmpl", gin.H{"error": "用户名或密码不能为空"})
		return
	}

	_userForm.Password = hex.EncodeToString(md5.New().Sum([]byte(_userForm.Password)))

	// 查询user表
	DB, cookie := config.DB, ""
	//ctx := fmt.Sprintf("user_name= %v and password= %v", _userForm.UserName, _userForm.Password)
	err := DB.Where(fmt.Sprintf("user_name = '%v' AND password = '%v'", _userForm.UserName, _userForm.Password)).First(&UserTable).Error

	// 分配cookie
	if err == nil {
		cookie, err = config.GenJWT(UserTable.ID, false)
	} else {
		c.HTML(http.StatusOK, "user/login.tmpl", gin.H{"error": err})
		return
	}
	//// 成功则跳转
	if err == nil {
		c.SetCookie("gin_cookie", cookie, 3600, "/", config.Conf.GetString("server.domain"), true, true)
		c.Redirect(http.StatusFound, "/")
		//c.HTML(http.StatusOK, "index.tmpl", gin.H{"error": cookie})
	} else {
		c.HTML(http.StatusOK, "user/login.tmpl", gin.H{"error": err})
	}
}

func Register(c *gin.Context) {
	_userForm := UserForm{
		UserName: c.PostForm("username"),
		Password: c.PostForm("password"),
	}

	if _userForm.UserName == "" || _userForm.Password == "" {
		c.HTML(http.StatusOK, "user/register.tmpl", gin.H{"error": "输入不能为空"})
		return
	}

	_userForm.Password = hex.EncodeToString(md5.New().Sum([]byte(_userForm.Password)))

	UserTable := config.User{
		UserName: _userForm.UserName,
		Password: _userForm.Password,
	}

	// 添加到user表 分配私有图片表
	DB := config.DB
	err := DB.Create(&UserTable).Error

	if err == nil {
		c.Redirect(http.StatusFound, "/login")
	} else {
		switch err.(*mysql.MySQLError).Number {
		case 1062:
			c.HTML(http.StatusOK, "user/register.tmpl", gin.H{"error": "该用户名已存在"})
		default:
			c.HTML(http.StatusOK, "user/register.tmpl", gin.H{"error": "注册失败"})
		}
	}
}
