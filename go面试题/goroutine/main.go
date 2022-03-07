package main

import (
	"fmt"
	"time"
)

//问题如下，会打印8  2  9 之类的多次，而不是123456789
//func main() {
//	data := make(map[int]int, 10)
//
//	for i := 1; i <= 10; i++ {
//		data[i] = i
//	}
//	fmt.Println(data)
//	fmt.Println("-------------")
//	for key, value := range data {
//		go func() {
//			fmt.Println(key, value)
//		}()
//	}
//	time.Sleep(time.Second * 1)
//}

//func main() {
//	data := make(map[int]int, 10)
//
//	for i := 1; i <= 10; i++ {
//		data[i] = i
//	}
//	fmt.Println(data)
//	fmt.Println("-------------")
//	for key, value := range data {
//		//第一种解决方法，显示使用内部变量接收key，value
//		key := key
//		value := value
//		go func() {
//			fmt.Println(key, value)
//		}()
//	}
//	time.Sleep(time.Second * 1)
//}

func main() {
	data := make(map[int]int, 10)

	for i := 1; i <= 10; i++ {
		data[i] = i
	}
	fmt.Println(data)
	fmt.Println("-------------")
	for key, value := range data {
		//第二种解决方法,给goroutine传参数，强制刷新goroutine栈中的值
		//fmt.Println(key, value)
		//开启10个goroutine并发执行
		go func(key, value int) {
			fmt.Println(key, value)
		}(key, value)
	}
	time.Sleep(time.Second * 1)
}

/**
goroutine栈最小空间1.4版本开始的时候为2KB
*/
