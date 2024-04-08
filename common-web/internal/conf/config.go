// @author hongjun500
// @date 2024-04-08-21:03
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: config.go

package conf

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	// 向下查找配置文件
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./config.docker")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("load config file error: %s\n", err)
		// use default config
		initDefaultConfigProperties()
		initDefaultGinConfigProperties()
	} else {
		log.Println("load config file success")
		viper.WatchConfig()
		_ = viper.UnmarshalKey("server", &GlobalApplicationConfigProperties)
		_ = viper.UnmarshalKey("gin", &GlobalGinConfigProperties)
		log.Printf("global config: %+v\n", GlobalApplicationConfigProperties)
	}
}

func initDefaultConfigProperties() {
	log.Println("init default config properties")
	GlobalApplicationConfigProperties.Enabled = DefaultEnabled
	GlobalApplicationConfigProperties.ApplicationName = DefaultApplicationName
	GlobalApplicationConfigProperties.Profile = DefaultProfile
	GlobalApplicationConfigProperties.Port = DefaultPort
	GlobalApplicationConfigProperties.ReadTimeout = DefaultReadTimeout
}

func initDefaultGinConfigProperties() {
	GlobalGinConfigProperties.GinMode = DefaultGinMode
}

// todo other config properties
