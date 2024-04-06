package config

import (
	"embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Option func(*gin.Engine)

//go:embed static/*
var static embed.FS

//go:embed img/*
var imgs embed.FS

var options []Option

// Include 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// Init 初始化
func Init() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("template/**/*")
	r.StaticFS("/static", http.FS(static))
	r.StaticFS("/img", http.FS(imgs))

	for _, opt := range options {
		opt(r)
	}
	return r
}
