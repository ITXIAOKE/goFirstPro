package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
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
	http.HandleFunc("/", errWrapper(HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

const prefix = "/list/"

type userErrorMy string

func (e userErrorMy) Error() string {
	return e.Message()
}

func (e userErrorMy) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userErrorMy(fmt.Sprintf("path %s must start "+"with %s", request.URL.Path, prefix))
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
