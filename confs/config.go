package confs

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	config *viper.Viper
}

var c = new(Config)

func init() {
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
	// 监控配置和重新获取配置
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	c.config = v
}

// GetServePort 获取服务端口号，例如 “:8080”
func GetServePort() string {
	return ":" + c.config.GetString("serve.port")
}

func GetLogFilePath() string {
	return c.config.GetString("log.path")
}

type Mysql struct {
	username string
	password string
	host     string //数据库地址，可以是Ip或者域名
	port     string //数据库端口
	dbname   string //数据库名
	timeout  string //连接超时
}

func GetMysqlConf() *Mysql {
	return &Mysql{
		c.config.GetString("mysql.username"),
		c.config.GetString("mysql.password"),
		c.config.GetString("mysql.host"),
		c.config.GetString("mysql.port"),
		c.config.GetString("mysql.dbname"),
		c.config.GetString("mysql.timeout"),
	}
}

type Redis struct {
	addr     string
	password string
	db       int
}

func GetRedisConf() *Redis {
	return &Redis{
		c.config.GetString("redis.addr"),
		c.config.GetString("redis.password"),
		c.config.GetInt("redis.db"),
	}
}
