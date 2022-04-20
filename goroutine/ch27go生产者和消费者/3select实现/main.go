package main

import (
	"fmt"
	"time"
)

/*
使用channel完成消费者、生产者的例子，发现使用channel会非常的方便
*/

func producer(intChan chan int) {
	for i := 0; i < cap(intChan); i++ {
		fmt.Println("生产者：", i)
		intChan <- i
	}
	// 写完后关闭掉
	close(intChan)
}

func consumer(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if ok {
			fmt.Println("消费者：", v)
		} else { // 读完了
			break
		}
		time.Sleep(time.Second)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	go producer(intChan)
	go consumer(intChan, exitChan)

	for {
		select {
		case _, ok := <-exitChan:
			if ok {
				fmt.Println("执行完毕！")
				return
			}
		default:
			fmt.Println("读不到，执行其他的！")
			time.Sleep(time.Second) // 此处添加Sleep才会看到效果，否则打印太多了找不到输出
			// break // break只是跳出select循环，可配合lable跳出
			// return
		}
	}

}
