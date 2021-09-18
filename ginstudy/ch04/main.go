package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//获取url中的变量
	r := gin.Default()
	v1 := r.Group("/v1") //要符合restful设计风格
	{
		v1.GET("", goodslist)
		v1.GET(":id", goodsdetail)
		v1.POST("", creategoods)
	}
	r.Run(":8086")

}

func creategoods(context *gin.Context) {

}

func goodslist(c *gin.Context) {

}

func goodsdetail(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}
