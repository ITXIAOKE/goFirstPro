package main

import (
	"goweb.com/ch03/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    "localhost:8888",
	}
	server.ListenAndServe()
}
