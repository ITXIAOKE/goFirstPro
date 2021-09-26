package main

import (
	"goweb.com/ch06/framework"
	"goweb.com/ch06/framework/middleware"
	"net/http"
)

func main() {
	//封装：如何让你的框架更好用？？
	core := framework.NewCore()

	// core.Use(
	// 	middleware.Test1(),
	// 	middleware.Test2())
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())
	// core.Use(middleware.Timeout(1 * time.Second))

	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
