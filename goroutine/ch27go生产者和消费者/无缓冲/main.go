package main

import (
	"fmt"
	"time"
)

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Send:", i)
	}
}

func consumer(ch <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-ch
		fmt.Println("Receive:", v)
	}
}

// 因为channel没有缓冲区，所以当生产者给channel赋值后，
// 生产者线程会阻塞，直到消费者线程将数据从channel中取出
// 消费者第一次将数据取出后，进行下一次循环时，消费者的线程
// 也会阻塞，因为生产者还没有将数据存入，这时程序会去执行
// 生产者的线程。程序就这样在消费者和生产者两个线程间不断切换，直到循环结束。
func main() {
	ch := make(chan int)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}
