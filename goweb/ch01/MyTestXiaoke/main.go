package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

/**
go doc net/http | grep "^func"
go doc context | grep "^func"

go doc net/http | grep "^type" | grep struct


*/

type Core struct {
	name int
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	fmt.Println("xiaoke----world")
}

func main() {
	fmt.Println("hello go web")
	// 创建一个Foo路由和处理函数
	http.Handle("/foo", NewCore())
	// 创建一个bar路由和处理函数
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	// 监听8080端口
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
