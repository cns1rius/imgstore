/*
 * @Author: s1rius
 * @Date: 2024-03-28 21:32:39
 * @LastEditTime: 2024-04-06 14:58:11
 * @Description: https://s1rius.space/
 */

package main

import (
	"log"

	"github.com/cns1rius/imgstore/admin"
	"github.com/cns1rius/imgstore/config"
	"github.com/cns1rius/imgstore/root"
	"github.com/cns1rius/imgstore/router"
	"github.com/cns1rius/imgstore/users"
)

func main() {
	config.InitConfig()
	config.InitDB()

	// 加载多个APP的路由配置
	router.Include(root.Router, users.Router, admin.Router)
	// 初始化路由
	r := router.Init()
	log.Printf("service startup at http://localhost:%s\n", config.Conf.GetString("server.port"))
	log.Println(r.Run(":" + config.Conf.GetString("server.port")))
	//if err := r.Run(":" + config.Conf.GetString("server.port")); err != nil {
	//	fmt.Printf("startup service failed, err:%v\n\n", err)
	//} else {
	//	fmt.Printf("service startup at 0.0.0.0:80")
	//}
}
