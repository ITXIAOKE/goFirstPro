package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"xkshop/v1/xkshop_api/user_web/config"
)

func main() {

	sc := []constant.ServerConfig{
		{
			IpAddr: "192.168.1.101",
			Port:   8848,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         "80317df8-5a77-4615-83b6-555905b2078c", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",   //日志地址
		CacheDir:            "tmp/nacos/cache", //缓存地址
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-web-dev.json", //手动把yaml文件通过网络工具转换为json文件，然后采用json格式放入nacos中
		Group:  "dev"})

	if err != nil {
		panic(err)
	}
	//fmt.Println(content) //字符串 - yaml

	serverConfig := config.ServerConfig{}
	//想要将一个json字符串转换成struct，需要去设置这个struct的tag
	json.Unmarshal([]byte(content), &serverConfig)
	fmt.Println(serverConfig)

	//动态监听配置文件的变动
	//err = configClient.ListenConfig(vo.ConfigParam{
	//	DataId: "user-web-dev.yaml",
	//	Group:  "dev",
	//	OnChange: func(namespace, group, dataId, data string) {
	//		fmt.Println("配置文件变化")
	//		fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
	//	},
	//})
	//time.Sleep(3000 * time.Second)
}
