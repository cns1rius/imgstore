package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	Conf *viper.Viper
)

func InitConfig() *viper.Viper {
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
	return Conf
}
