package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkshop/v1/xkshop_api/user_web/middlewares"
	newRouter "xkshop/v1/xkshop_api/user_web/router"
)

func Routers() *gin.Engine {
	//gin.SetMode(gin.DebugMode) //这里要设置以下gin的模式，否则有告警
	engine := gin.Default()
	engine.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"code":http.StatusOK,
			"success":true,
		})
	})
	//配置跨域
	engine.Use(middlewares.Cors())

	ApiGroup := engine.Group("/v1") //全局的routerGroup
	newRouter.InitUserRouter(ApiGroup)
	newRouter.InitBaseRouter(ApiGroup)

	return engine
}
