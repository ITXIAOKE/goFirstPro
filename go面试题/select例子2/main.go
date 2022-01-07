package main

import (
	"log"
	"sync"
	"time"
)

var (
	wg       sync.WaitGroup
	execTime time.Duration = time.Second
)

func main() {
	timeout := 50 * time.Millisecond
	log.Printf("result:%d", finishReq(timeout))

	wg.Wait()
}

func finishReq(timeout time.Duration) int {
	//ch := make(chan int)//这里没有加buffer缓冲，结果就会造成死锁
	ch := make(chan int, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(execTime)
		ch <- 100
	}()

	select {
	case result := <-ch:
		return result
	case <-time.After(timeout):
		log.Print("Timeout")
		return -1
	}
}
