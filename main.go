package main

import (
	"github.com/cns1rius/imgstore/admin"
	"github.com/cns1rius/imgstore/config"
	"github.com/cns1rius/imgstore/root"
	"github.com/cns1rius/imgstore/users"
	"log"
)

func main() {
	config.InitConfig()
	config.InitDB()

	// 加载多个APP的路由配置
	config.Include(root.Router, users.Router, admin.Router)
	// 初始化路由
	r := config.Init()
	log.Printf("service startup at http://localhost:%s\n", config.Conf.GetString("server.port"))
	log.Println(r.Run(":" + config.Conf.GetString("server.port")))
	//if err := r.Run(":" + config.Conf.GetString("server.port")); err != nil {
	//	fmt.Printf("startup service failed, err:%v\n\n", err)
	//} else {
	//	fmt.Printf("service startup at 0.0.0.0:80")
	//}
}
