package initialize

import (
	"github.com/gin-gonic/gin"
	"xkshop/v1/xkshop_api/user_web/middlewares"
	newRouter "xkshop/v1/xkshop_api/user_web/router"
)

func Routers() *gin.Engine {
	gin.SetMode(gin.DebugMode) //这里要设置以下gin的模式，否则有告警
	engine := gin.Default()
	//配置跨域
	engine.Use(middlewares.Cors())

	ApiGroup := engine.Group("/v1") //全局的routerGroup
	newRouter.InitUserRouter(ApiGroup)
	newRouter.InitBaseRouter(ApiGroup)

	return engine
}
