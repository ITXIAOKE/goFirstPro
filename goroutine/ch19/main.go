package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg     sync.WaitGroup
	rwLock sync.RWMutex
	lock   sync.Mutex
	x int64
)

func main() {
	/**
	读写锁
	读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁 会继续获得锁，
	如果是获取写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论 是获取读锁还是写锁都会等待。

	需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出 来。
	*/

	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()

	fmt.Println(end.Sub(start))
	fmt.Println(x)

}

func read() {
	rwLock.RLock()
	time.Sleep(time.Microsecond)
	rwLock.RUnlock()

	defer wg.Done()
}

func write() {
	rwLock.Lock()
	x=x+1
	time.Sleep(10*time.Millisecond)
	rwLock.Unlock()

	defer wg.Done()

}
