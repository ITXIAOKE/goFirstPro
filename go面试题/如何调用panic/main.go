package main

import (
	"fmt"
	"time"
)

func proc() {
	panic("I am error!")
}

func main() {
	go func() {
		/**
		1,在这里需要你写算法
		2，要求每秒钟调用一次proc函数
		3，要求程序不能退出
		*/
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-t.C: //每秒钟输出
				go func() {
					defer func() { //defer捕获panic
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
			}
		}

	}()
	select {} //不能去掉，让主线程阻塞，等待goroutine执行完成
}
