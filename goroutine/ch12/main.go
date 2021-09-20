package main

import (
	"fmt"
)

func main() {
	/**
	从通道循环取值
	判断通道是否关闭

	两种方式(data, ok := <-ch1和for range )在接收值的时候判断通道是否被关闭，我们通常使用的是for range的方式。

	1.chan<- int是一个只能发送的通道，可以发送但是不能接收；
	2.<-chan int是一个只能接收的通道，可以接收但是不能发送。

	注意：对数据来说，可以是发送或者接收通道
	在函数传参及任何赋值操作中将双向通道转换为单向通道是可以的，但反过来是不可以的

	关闭已经关闭的channel也会引发panic

	*/
	ch1 := make(chan int)
	ch2 := make(chan int)
	go mypro1(ch1)
	go mypro2(ch1, ch2)

	for i := range ch2 {
		fmt.Println(i)
	}
}

//限制单向通道
func mypro2(ch1 <-chan int, ch2 chan<- int) { //ch1这个时候只是接收channel
	for x := range ch1 {
		ch2 <- x * x
	}
	close(ch2)
}

func mypro1(ch1 chan<- int) { //ch1这个时候只是发送channel
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
}

//最原始的通道，没限制
//func mypro2(ch1 chan int, ch2 chan int) {
//	for {
//		data, ok := <-ch1
//		if !ok {
//			break
//		}
//		ch2 <- data * data
//	}
//	close(ch2)
//}
//
//func mypro1(ch1 chan int) {
//	for i := 0; i < 10; i++ {
//		ch1 <- i
//	}
//	close(ch1)
//}
