package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("init")
}
func main() {
	/**
	• 基于 Channel 编写一个简单的单线程生产者消费者模型
	• 队列：
	队列长度10，队列元素类型为 int
	• 生产者：
	每1秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
	• 消费者：
	每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
	*/
	//ch := make(chan string, 10)
	//go func() {
	//	fmt.Println("hello")
	//	ch <- "aa"
	//}()
	//i := <-ch
	//fmt.Println(i)
	//time.Sleep(time.Second)
	//fmt.Println("world")

	c := make(chan int, 10)
	go func() {
		for {
			time.Sleep(time.Second)
			c <- rand.Int()
		}
	}()
	//for v:=range c{
	//	time.Sleep(time.Second)
	//	println(v<-c)
	//}
	for{
		i := <-c
		fmt.Println(i)
	}


}
