package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	/**
	goexit():退出当前协程
	*/
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			//结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	i := 0
	for {
		fmt.Println(i)
		time.Sleep(time.Second)
		i++
		if i == 10 {
			break
		}
	}
}
