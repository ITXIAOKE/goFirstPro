package main

import (
	"ginstudy.com/ch07/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	//ID   string `uri:"id" binding:"required,uuid"`//这个是uuid的id
	ID   int    `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	//返回json文件和protobuf文件，供客户端调用，做数据交互
	r := gin.Default()
	v1 := r.Group("/v1") //要符合restful设计风格
	{
		v1.GET("/morejson", morejson)
		v1.GET("/someProtoBuf", returnProto) //返回protobuf文件
		//v1.POST("/form_post", form_post)
	}
	r.Run(":8089")

}

func returnProto(c *gin.Context) {
	course := []string{"python", "go", "微服务"}
	user := &proto.Teacher{
		Name:   "xiaoke",
		Course: course,
	}
	//c.JSON(http.StatusOK, user)
	c.ProtoBuf(http.StatusOK, user)
}

func morejson(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"` //加上tag后，这个字段最后以tag里面的user展示信息
		Message string
		Number  int
	}
	msg.Name = "xiaoke"
	msg.Message = "我是测试json文件"
	msg.Number = 100

	c.JSON(http.StatusOK, msg)
}
