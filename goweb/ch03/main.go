package main

import (
	"goweb.com/ch03/framework"
	"net/http"
)

func main() {
	//context：请求控制器，让每个请求都在掌控之中
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    "localhost:8888",
	}
	server.ListenAndServe()
}
