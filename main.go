package main

import (
	"fmt"
	"github.com/cns1rius/imgstore/admin"
	"github.com/cns1rius/imgstore/config"
	"github.com/cns1rius/imgstore/root"
	"github.com/cns1rius/imgstore/router"
	"github.com/cns1rius/imgstore/users"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	Config *viper.Viper
	db     *gorm.DB
)

func main() {
	Config = config.InitConfig()
	db = config.InitDB()

	// 加载多个APP的路由配置
	router.Include(root.Router, users.Router, admin.Router)
	// 初始化路由
	r := router.Init()
	if err := r.Run(":" + Config.GetString("server.port")); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	} else {
		fmt.Printf("service startup at 0.0.0.0:80")
	}
}
