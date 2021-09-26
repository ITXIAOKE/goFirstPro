package main

import (
	"goweb.com/ch05/framework"
	"goweb.com/ch05/framework/middleware"
	"net/http"
)

func main() {
	//增加中间件：如何提高框架的可拓展性
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
