package global

import (
	"gorm.io/gorm"
	"xkshop/v1/xkshop_srv/goods_srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)
