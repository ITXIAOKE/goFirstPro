package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

type ServerConfig struct {
	ServiceName string      `mapstructure:"name"`  //将yaml文件name的值映射到servicename上
	MysqlInfo   MysqlConfig `mapstructure:"mysql"` //嵌套配置
}

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

//如何将线上和线下的配置文件隔离？？？？？
//解决办法就是本地环境中设置一个参数env

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func GetEnvAll(env string) interface{} {
	viper.AutomaticEnv()
	return viper.Get(string(env))
}

func main() {
	fmt.Println(GetEnvAll("GOROOT"))
	fmt.Println(GetEnvInfo("GOROOT"))
	v := viper.New()
	v.SetConfigFile("xkshop_api/user_web/viper_test/ch02/config.yaml") //注意路径的问题
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)

	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("mysql"))
	fmt.Println("============================")

	//viper的功能---动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file channed", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&serverConfig)
		fmt.Println(serverConfig)
	})

	time.Sleep(time.Second * 30)

}
