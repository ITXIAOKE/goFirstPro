package global

import (
	ut "github.com/go-playground/universal-translator"
	"xkshop/v1/xkshop_api/goods-web/config"
	"xkshop/v1/xkshop_api/goods-web/proto"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	GoodsSrvClient proto.GoodsClient
)
