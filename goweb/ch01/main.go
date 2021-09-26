package main

import (
	"goweb.com/ch01/framework"
	"net/http"
)

func main() {
	//net/http：使用标准库搭建server
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr:    "localhost:8080",
	}
	server.ListenAndServe()
}
