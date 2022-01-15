package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"xkshop/v1/xkshop_api/user_web/global"
	"xkshop/v1/xkshop_api/user_web/initialize"
	myValidator "xkshop/v1/xkshop_api/user_web/validator"
)

func main() {
	//port := 8021
	//zap.L() //全局的logger
	//zap.S() //全局的global
	//zap.ReplaceGlobals(logger) //让日志打印输出

	//1,初始化log
	initialize.InitLogger()

	//2，初始化配置文件
	initialize.InitConfig()

	//3,初始化routers
	engine := initialize.Routers()

	//4,初始化翻译器
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	//5，注册自定义的手机验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myValidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码！", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	//日志相关debug  info warn error fetal
	zap.S().Debugf("启动服务器，端口：%d", global.ServerConfig.Port)

	if err := engine.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败：", err.Error())
	}

}
