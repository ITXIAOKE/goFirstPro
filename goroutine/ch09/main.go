package main

import "fmt"

func main() {
	/**
	无缓冲通道：同步通道,使用无缓冲通道进行通信将导致发送和接收的goroutine同步化
	*/
	ch := make(chan int) //无缓冲通道，必须要有接收和发送，少一方，就会发生deadlock!!!
	go recoverMy(ch)
	ch <- 100
	fmt.Println("发送成功")
}

func recoverMy(ch chan int) {
	for {
		aa := <-ch
		fmt.Println("接收成功", aa)
	}
}
