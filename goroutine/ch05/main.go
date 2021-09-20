package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine:%d\n", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Printf("main goroutine%d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break //如果主协程退出了，其他任务还执行吗???答案是不执行
		}
	}

}
