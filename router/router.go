package router

import (
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options []Option

// Include 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// Init 初始化
func Init() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("template/**/*")
	r.Static("/static", "./static")
	r.Static("/img", "./pictures")

	for _, opt := range options {
		opt(r)
	}
	return r
}
