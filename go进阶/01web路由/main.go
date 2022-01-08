package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("第一种：路由和路由处理器")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Host)
		fmt.Println(request.Header)
		fmt.Println(request.Method)
		fmt.Println(request.URL.Path)
		fmt.Println(request.URL.RawQuery)

	})

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("hello world!!"))
	})

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm() //将字符串数据转换为map类型

		if request.Method == http.MethodGet {
			fmt.Println(request.Form.Get("name"))
			fmt.Println(request.Form.Get("password"))
		}

		if request.Method == http.MethodPost {
			fmt.Println(request.FormValue("name"))
			fmt.Println(request.FormValue("password"))
		}

	})

	addr := ":8080"
	http.ListenAndServe(addr, nil)

}
