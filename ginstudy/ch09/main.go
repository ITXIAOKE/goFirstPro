package main

import (
	"fmt"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "123456")
		//c.Abort()//终止之后，后面的代码不会再运行
		c.Next()

		end := time.Since(t)
		fmt.Println("耗时：%V\n", end)

		status := c.Writer.Status()
		fmt.Println("状态", status)

	}
}

func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for k, v := range c.Request.Header {
			if k == "X-Token" {
				token = v[0]
			} else {
				fmt.Println(k, v)
			}
		}

		if token != "bobby" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录",
			})
			//挑战 为什么连return都阻止不了后续逻辑的执行,只能是c.abort（）才能执行阻止的功能
			c.Abort()
		}
		c.Next()
	}
}

func main() {
	////中间件功能相关

	r := gin.Default()

	r.Use(MyLogger())
	r.Use(TokenRequired())

	//r.Static()//加载静态文件

	r.GET("/ping", ping)
	r.Run(":8082")

	//r := gin.New()
	//r.Use(gin.Logger(), gin.Recovery())//全局所有
	//
	//authrized := r.Group("/goods")
	//authrized.Use(AuthRequired)//针对/goods这个url生效

}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

}

func AuthRequired(c *gin.Context) {

}
