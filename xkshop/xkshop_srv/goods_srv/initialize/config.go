package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"xkshop/v1/xkshop_srv/goods_srv/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	//从配置文件中读取对应的配置
	debug := GetEnvInfo("xkshop_debug") //在本机的环境变量中设置
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("xkshop_srv/goods_srv/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("xkshop_srv/goods_srv/%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName) //注意路径的问题

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&global.NacosConfig); err != nil {
		panic(err)
	}

	zap.S().Infof("配置信息：%v", global.NacosConfig)

	//viper的功能---动态监控变化
	//v.WatchConfig()
	//v.OnConfigChange(func(e fsnotify.Event) {
	//	zap.S().Infof("配置文件产生变化：%s", e.Name)
	//	_ = v.ReadInConfig()
	//	_ = v.Unmarshal(global.ServerConfig)
	//	zap.S().Infof("配置信息：%v", global.ServerConfig)
	//})

	//nacos的srv服务的配置
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         "a1f0e8ab-ffad-45f9-ac74-9c49bb17ab57", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
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
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group})

	if err != nil {
		panic(err)
	}
	//fmt.Println(content) //字符串 - yaml

	//想要将一个json字符串转换成struct，需要去设置这个struct的tag
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		zap.S().Fatalf("读取nacos配置失败：%s", err.Error())
	}
	fmt.Println(global.ServerConfig)

}
