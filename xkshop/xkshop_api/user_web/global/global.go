package global

import (
	ut "github.com/go-playground/universal-translator"
	"xkshop/v1/xkshop_api/user_web/config"
	"xkshop/v1/xkshop_api/user_web/proto"
)

var (
	ServerConfig  = &config.ServerConfig{}
	Trans         ut.Translator //常量首字母切记要大写
	UserSrvClient proto.UserClient
)
