package main

import (
	"fmt"

	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type LoginForm struct {
	User     string `json:"user"  binding:"required,min=3,max=10"`
	Password string `json:"password" binding:"required"`
}

type SignUpForm struct {
	Age        uint8  `json:"age" binding:"required,gte=1,lte=130"`
	Name       string `json:"name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` //跨字段
}

var trans ut.Translator

func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		uni := ut.New(enT, zhT, enT)
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			en_translations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

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

func main() {
	////中间件功能相关

	r := gin.Default()

	r.Use(MyLogger())

	//r.Static()//加载静态文件

	r.GET("/ping", ping)
	r.Run(":8082")

	//r := gin.New()
	//r.Use(gin.Logger(), gin.Recovery())
	//
	//authrized := r.Group("/goods")
	//authrized.Use(AuthRequired)

}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

}

func AuthRequired(c *gin.Context) {

}
