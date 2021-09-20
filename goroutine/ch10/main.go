package main

import "fmt"

func main() {
	/**
	有缓冲通道
	*/
	ch := make(chan int, 1) //这个有缓冲通道只能放1个元素，要是多于一个元素就会发生deadlock
	ch <- 10
	//ch <- 11//deadlock
	fmt.Println("发送成功")
}
