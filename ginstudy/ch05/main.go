package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	//ID   string `uri:"id" binding:"required,uuid"`//这个是uuid的id
	ID   int    `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	//获取url中的变量
	r := gin.Default()
	v1 := r.Group("/v1") //要符合restful设计风格
	{
		v1.GET("", goodslist)
		v1.GET("/:name/:id", goodsdetail)
		v1.POST("", creategoods)
	}
	r.Run(":8087")

}

func creategoods(context *gin.Context) {

}

func goodslist(c *gin.Context) {

}

func goodsdetail(c *gin.Context) {
	var person Person
	if err := c.ShouldBindUri(&person); err != nil {
		c.Status(404)
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"name": person.Name,
		"id":   person.ID,
	})

}
