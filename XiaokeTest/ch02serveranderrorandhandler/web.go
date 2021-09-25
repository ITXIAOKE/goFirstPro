package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"XiaokeTest.com/ch02serveranderrorandhandler/filelisting"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		// panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request) //做统一的业务逻辑处理

		if err != nil {
			log.Printf("Error occurred "+"handling request: %s", err.Error())

			// user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest) //给用户的message
				return
			}

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error            //给系统看的
	Message() string //给用户看的
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

