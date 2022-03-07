package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//获取get和post的表单信息
	r := gin.Default()
	v1 := r.Group("/v1") //要符合restful设计风格
	{
		v1.GET("/welcome", welcome)
		v1.POST("/form_post", form_post) //数据放body中，最终数据在body上
		v1.POST("/post", getPost)        //混合，数据放param中，最终数据在url上
	}
	r.Run(":8089")

}

func getPost(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.DefaultPostForm("message", "信息")
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})
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
