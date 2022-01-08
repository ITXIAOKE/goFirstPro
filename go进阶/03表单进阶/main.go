package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type User struct {
	Id       int    `json:"id" validate:"gt=0"` //validate,表单验证
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func main() {
	fmt.Println("表单进阶")

	mux := http.NewServeMux()
	//json请求处理
	handler := func(writer http.ResponseWriter, request *http.Request) {
		//request.ParseForm()
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}
		var user User
		err = json.Unmarshal(data, &user) //反序列化
		if err != nil {
			panic(err)
		}
		log.Printf("User = %#v", user)
		//验证
		validate := validator.New()
		if err = validate.Struct(user); err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
		} else {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte("成功了"))
		}

	}
	mux.HandleFunc("/json", handler)

	//文件处理
	mux.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
		const maxFileSize = 10 << 20 //10M文件
		err := request.ParseMultipartForm(maxFileSize)
		if err != nil {
			panic(err)
		}
		for _, headers := range request.MultipartForm.File {
			for _, file := range headers {
				saveFilePath := filepath.Join("./", file.Filename)

				//原始上传的文件
				src, err := file.Open()
				if err != nil {
					panic(err)
				}
				defer src.Close()

				//创建一个空文件
				dst, err := os.Create(saveFilePath)
				if err != nil {
					panic(err)
				}
				defer dst.Close()

				if _, err = io.Copy(dst, src); err != nil {
					panic(err)
				}

			}
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("上传完成"))
	})

	_ = http.ListenAndServe(":8080", mux)

}
