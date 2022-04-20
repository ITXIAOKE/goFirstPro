package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	product = 0
	lock    sync.Mutex
	cond    = sync.NewCond(&lock)
)

func producer() {
	for {
		cond.L.Lock() // 先加锁

		for product > 10 {
			fmt.Println("生产完了！")
			cond.Wait()
		}
		fmt.Println("生产中...", product)
		product += 1

		cond.L.Unlock()
		cond.Broadcast()
	}
}

func consumer() {
	for {
		cond.L.Lock()

		for product <= 0 {
			fmt.Println("消费完了！")
			cond.Wait()
		}
		fmt.Println("消费中...", product)
		product -= 1

		cond.L.Unlock()
		cond.Broadcast()
	}
}

func main() {
	go producer()
	go consumer()

	time.Sleep(time.Second * 2)
	fmt.Println("主线程结束！")
}
