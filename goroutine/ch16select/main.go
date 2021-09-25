package main

import (
	"fmt"
)

func main() {
	//go语言提供了一个select的功能，作用与channel之上的，多路复用
	//select会随机公平的选择一个case语句执行
	//select的应用场景：1，timeout超时机制  2,判断某个channel是否阻塞
	output1 := make(chan string, 1)
	output2 := make(chan string, 1)

	//以下这种方式不行
	go test1(output1)
	go test2(output2)

	//output1 <- "test1"
	//output2 <- "test2"

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	default:
		fmt.Println("default")
	}
	fmt.Println("main结束")
}

func test1(output1 chan string) {
	output1 <- "test1"
	//time.Sleep(time.Second)
}

func test2(output2 chan string) {
	output2 <- "test2"
	//time.Sleep(time.Second)
}
