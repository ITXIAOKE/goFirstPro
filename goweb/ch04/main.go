package main

import (
	"goweb.com/ch04/framework"
	"net/http"
)

func main() {
	core:=framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8877",
	}
	server.ListenAndServe()
}
