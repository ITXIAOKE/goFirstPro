package main

import (
	"goweb.com/ch04/framework"
	"net/http"
)

func main() {
	//增加路由：如何让请求更快寻找到目标函数
	core:=framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8877",
	}
	server.ListenAndServe()
}
