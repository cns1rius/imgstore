package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	Conf  *viper.Viper
	Types = []string{"人物", "风景", "动物", "植物", "动漫", "图像", "物品", "其他"}
)

func InitConfig() {
	// 设置配置文件路径
	config := "config.yaml"
	// 生产环境可以通过设置环境变量来改变配置文件路径
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	// 初始化 viper
	Conf = viper.New()
	Conf.SetConfigFile(config)
	Conf.SetConfigType("yaml")
	if err := Conf.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	// 初始化img文件夹
	_ = os.MkdirAll("./img/tmp/", os.ModePerm)
	for _, t := range Types {
		_ = os.MkdirAll("./img/"+t+"/", os.ModePerm)
	}
}
