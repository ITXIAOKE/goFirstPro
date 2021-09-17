package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello 11111111111")
	r:=gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"pong",
		})
	})
	r.Run()
	fmt.Println("hello 2222222222")
}


