package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"xkshop/v1/xkshop_srv/user_srv/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	//从配置文件中读取对应的配置
	debug := GetEnvInfo("xkshop_debug") //在本机的环境变量中设置
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("xkshop_srv/user_srv/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("xkshop_srv/user_srv/%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName) //注意路径的问题

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}

	zap.S().Infof("配置信息：%v", global.ServerConfig)

	//viper的功能---动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件产生变化：%s", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
		zap.S().Infof("配置信息：%v", global.ServerConfig)
	})
}
