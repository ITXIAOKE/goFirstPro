package global

import (
	"gorm.io/gorm"
	"xkshop/v1/xkshop_srv/user_srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)
