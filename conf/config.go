package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	config *viper.Viper
}

func InitConfigure() *Config {
	v := viper.New()
	mode := gin.Mode()
	// 设置文件名称（无后缀）
	v.SetConfigName("application-" + mode)
	// 设置后缀名
	v.SetConfigType("yaml")
	// 设置文件所在路径
	v.AddConfigPath("./")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(" Config file not found; ignore error if desired")
		} else {
			panic("Config file was found but another error was produced")
		}
	}
	v.WatchConfig() // 监控配置和重新获取配置

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	return &Config{
		config: v,
	}
}

// 获取服务端口号，例如“:8080”
func (c *Config) GetServePort() string {
	return ":" + c.config.GetString("serve.port")
}
