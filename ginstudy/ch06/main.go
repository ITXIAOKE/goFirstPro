package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	//ID   string `uri:"id" binding:"required,uuid"`//这个是uuid的id
	ID   int    `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	//获取get和post的表单信息
	r := gin.Default()
	v1 := r.Group("/v1") //要符合restful设计风格
	{
		v1.GET("/welcome", welcome)
		v1.POST("/form_post", form_post)
	}
	r.Run(":8089")

}

func form_post(context *gin.Context) {
	message := context.PostForm("message")
	nick := context.DefaultPostForm("nick", "anannoymouse")
	context.JSON(http.StatusOK, gin.H{
		"message": message,
		"nick":    nick,
	})
}

func welcome(c *gin.Context) {
	firstName := c.DefaultQuery("firstname", "xiaoke")
	lastName := c.DefaultQuery("lastname", "imooc")
	c.JSON(http.StatusOK, gin.H{
		"first_name": firstName,
		"last_name":  lastName,
	})
}
