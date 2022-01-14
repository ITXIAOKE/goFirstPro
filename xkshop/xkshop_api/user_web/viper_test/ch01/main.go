package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	ServiceName string `mapstructure:"name"` //将yaml文件name的值映射到servicename上
	Port        int    `mapstructure:"port"`
}

func main() {
	v := viper.New()
	v.SetConfigFile("xkshop_api/user_web/viper_test/ch01/config.yaml") //注意路径的问题
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)

	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("port"))

}
