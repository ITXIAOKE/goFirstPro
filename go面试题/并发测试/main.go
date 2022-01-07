package main

import (
	"fmt"
	"sync"
	"time"
)

//互斥锁、不可重入锁，在一个goroutine下，重复上锁的话，就会发生panic
var mx sync.Mutex

func main() {
	mx.Lock()
	go func() {
		mx.Lock()
		fmt.Println("T am oK!")
		defer mx.Unlock()
	}()
	//defer mx.Unlock()//不能这样写，因为这样会一直处于上锁的状态
	mx.Unlock()
	time.Sleep(time.Second)
	return
}
