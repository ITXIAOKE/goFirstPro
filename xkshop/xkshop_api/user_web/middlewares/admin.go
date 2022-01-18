package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkshop/v1/xkshop_api/user_web/models"
)

//判断是否是管理员

func IsAdminAuth() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*models.CustomClaims)

		if currentUser.AuthorityId == 2 { //根据role的角色，1表示普通用户，2表示管理员
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
