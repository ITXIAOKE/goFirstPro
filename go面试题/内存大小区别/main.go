package main

import (
	"fmt"
	"unsafe"
)

type DemoA struct {
	A int8  //1字节，填充7字节
	B int64 //8字节
	C int16 //2字节，填充6字节
}

type DemoB struct {
	A int8 //第一个元素int8与第二个元素加起来是3个字节，填充5字节
	C int16
	B int64 //8字节
}

func main() {
	fmt.Println(unsafe.Sizeof(DemoA{})) //24
	fmt.Println(unsafe.Sizeof(DemoB{})) //16
}
