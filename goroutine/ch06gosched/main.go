package main

import (
	"fmt"
	"runtime"
)

func main() {
	/**
	gosched:让出CPU时间片，重新等待安排任务
	*/
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	for i := 0; i < 2; i++ {
		runtime.Gosched() /*切一下，再次分配任务*/
		fmt.Println("hello")
	}
}
