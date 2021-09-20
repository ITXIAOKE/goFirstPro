package main

import (
	"fmt"
	"time"
)

func main() {
	//for i := 0; i < 100; i++ {
	//	//注意闭包的问题
	//	go func() {
	//		for {
	//			fmt.Println(i)
	//			time.Sleep(time.Second)
	//		}
	//	}()
	//}
	//time.Sleep(time.Second)

	for i := 0; i < 100; i++ {
		//注意闭包的问题,传参数解决
		go func(i int) {
			for {
				fmt.Println(i)
				time.Sleep(time.Second)
			}
		}(i)
	}
	time.Sleep(time.Second)
}
