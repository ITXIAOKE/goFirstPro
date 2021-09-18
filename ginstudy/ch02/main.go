package main

import "github.com/gin-gonic/gin"

func main() {
	//7中restful风格方式
	//r:=gin.New()

	//default方式添加了日志log和故障恢复recover两个功能，方便调试和运行
	router := gin.Default()

	router.GET("/someGet", gettting)
	//router.POST("/somePost", posting)
	//router.PUT("/somePut", puttting)
	//router.DELETE("/someDelete", deletting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	router.Run()

}

func gettting(c *gin.Context) {
	//c.JSON(200, gin.H{"name": "xiaoke"})
	c.JSON(200, map[string]string{"name": "xiaoke"})
}
