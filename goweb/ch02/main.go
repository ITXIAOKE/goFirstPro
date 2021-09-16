package main

import "net/http"

// 框架核心结构

type Core struct {
}

// 初始化框架核心结构

func NewCore() *Core {
	return &Core{}
}

// 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO
}

func main() {
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: NewCore(),
		// 请求监听地址
		Addr: ":8080",
	}
	server.ListenAndServe()
}
