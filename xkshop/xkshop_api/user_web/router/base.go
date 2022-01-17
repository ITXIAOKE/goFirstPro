package router

import (
	"github.com/gin-gonic/gin"
	"xkshop/v1/xkshop_api/user_web/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha) //验证码
		BaseRouter.POST("send_sms", api.SendSms)  //手机发送短信
	}
}
