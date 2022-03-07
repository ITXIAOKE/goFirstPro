package main

import (
	"fmt"
	"reflect"
)

type UserInfo struct {
	Name     string `bilibili:"bilibili_name" abc:"name"`
	PublicWX string `bilibili:"bilibili_publicwx" abc:"publicwx"`
}

func PrintTag(myptr interface{}) {
	reType := reflect.TypeOf(myptr)
	if reType.Kind() != reflect.Ptr || reType.Elem().Kind() != reflect.Struct {
		panic("传入的参数不是结构体/指针")
		return
	}

	v := reflect.ValueOf(myptr).Elem()
	fmt.Println(v)
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag
		labeTag := tag.Get("abc")
		fmt.Println(labeTag)
	}

}

func main() {
	userInfo := &UserInfo{
		Name:     "小可",
		PublicWX: "xiaoke",
	}
	PrintTag(userInfo)
}
