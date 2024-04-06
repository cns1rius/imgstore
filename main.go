package main

import (
	"embed"
	"github.com/cns1rius/imgstore/admin"
	"github.com/cns1rius/imgstore/config"
	"github.com/cns1rius/imgstore/root"
	"github.com/cns1rius/imgstore/users"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

//go:embed template/*
var tmpl embed.FS

//go:embed static/*
var Static embed.FS

//go:embed img/*
var Imgs embed.FS

func main() {
	config.InitConfig()
	config.InitDB()

	// 加载多个APP的路由配置
	config.Include(root.Router, users.Router, admin.Router)

	// 初始化gin
	r := gin.Default()
	// tmpl
	t, _ := template.ParseFS(tmpl, "template/**/*")
	r.SetHTMLTemplate(t)
	// static
	Static, _ := fs.Sub(Static, "static")
	r.StaticFS("/static", http.FS(Static))
	// img pictures
	Imgs, _ := fs.Sub(Imgs, "img")
	r.StaticFS("/img", http.FS(Imgs))
	// route init
	r = config.Init(r)

	// log
	log.Printf("service startup at http://localhost:%s\n", config.Conf.GetString("server.port"))
	log.Println(r.Run(":" + config.Conf.GetString("server.port")))
}
