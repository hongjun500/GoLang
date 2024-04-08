// @author hongjun500
// @date 2024-04-08-21:04
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: config_properties.go

package conf

var (
	GlobalApplicationConfigProperties = &ApplicationConfigProperties{}
	GlobalGinConfigProperties         = &GinConfigProperties{}
)

const (
	DefaultEnabled         = true
	DefaultApplicationName = "common-web"
	DefaultProfile         = "dev"
	DefaultPort            = "8080"
	DefaultReadTimeout     = 60

	DefaultGinMode = "debug"
)

// ApplicationConfigProperties 应用配置属性
type ApplicationConfigProperties struct {
	Enabled         bool   // 是否启用
	ApplicationName string // 应用名称
	Profile         string // 环境
	Port            string // 端口
	ReadTimeout     int    // 读取超时时间
}

type GinConfigProperties struct {
	GinMode string // gin 运行模式
}
