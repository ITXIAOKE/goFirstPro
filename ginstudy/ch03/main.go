package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//路由分组
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/list", list)
		v1.POST("/add", add)
	}

	r.Run(":8085")
}

func add(context *gin.Context) {
	fmt.Println("add")
	context.JSON(http.StatusOK, gin.H{"name": "add"})
}

func list(context *gin.Context) {
	fmt.Println("list")
	context.JSON(http.StatusOK, gin.H{"name": "list"})
}
