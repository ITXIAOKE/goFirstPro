package main

import (
	"fmt"
	"sync"
)

func main() {
	PrintOnce()
	PrintOnce()
	PrintOnce()
}

var once sync.Once //这个sync.Once只能在全局中初始化才能生效，在局部函数中初始化不能生效

// 这个方法，不管调用几次，只会输出一次

func PrintOnce() {
	once.Do(func() {
		fmt.Println("只输出一次")
	})
}
