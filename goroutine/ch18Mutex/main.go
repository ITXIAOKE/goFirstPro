package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var x int64
var lock sync.Mutex

func main() {
	/**
	互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资 源。Go语言中使用sync包的Mutex类型来实现互斥锁。

	使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；
	当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时， 唤醒的策略是随机的。

	下面代码结果为200000
	*/

	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

func add() {
	for i := 0; i < 10000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
