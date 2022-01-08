package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("你好，世界"))
}

func main() {
	fmt.Println("第二种：路由和路由处理器")
	http.Handle( "/hello", &MyHandler{})
	addr := ":8080"
	http.ListenAndServe(addr, nil)

	//mux := http.NewServeMux()
	//mux.Handle()
	//mux.HandleFunc()
}
