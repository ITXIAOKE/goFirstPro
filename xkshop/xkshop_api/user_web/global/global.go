package global

import (
	ut "github.com/go-playground/universal-translator"
	"xkshop/v1/xkshop_api/user_web/config"
)

var (
	ServerConfig = &config.ServerConfig{}
	Trans        ut.Translator //常量受字母切记要大写
)
