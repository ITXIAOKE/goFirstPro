package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
问题：
1，使用三个协程，每秒钟打印cat、dog、fish
2，顺序不能变化
3，无限循环即可
*/

func main() {
	var wg sync.WaitGroup

	var catCounter uint64
	var dogCounter uint64
	var finishCounter uint64

	catChan := make(chan struct{}, 1)
	dogChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)

	wg.Add(3)

	go cat(catChan, dogChan, &wg, catCounter)
	go dog(dogChan, fishChan, &wg, dogCounter)
	go fish(fishChan, catChan, &wg, finishCounter)

	fishChan <- struct{}{} //这是触发条件,触发有谁开头，否则会触发deadlock

	wg.Wait()
}

func fish(fishChan chan struct{}, catChan chan struct{}, w *sync.WaitGroup, counter uint64) {
	for {
		if counter >= uint64(3) {
			w.Done()
			return
		}
		<-fishChan
		fmt.Println("finish")
		atomic.AddUint64(&counter, 1)
		catChan <- struct{}{}
	}
}

func dog(dogChan chan struct{}, fishChan chan struct{}, w *sync.WaitGroup, counter uint64) {
	for {
		if counter >= uint64(3) {
			w.Done()
			return
		}
		<-dogChan
		fmt.Println("dog")
		atomic.AddUint64(&counter, 1)
		fishChan <- struct{}{}
	}
}

func cat(catChan chan struct{}, dogChan chan struct{}, w *sync.WaitGroup, counter uint64) {
	for {
		if counter >= uint64(3) {
			w.Done()
			return
		}
		<-catChan
		fmt.Println("cat")
		atomic.AddUint64(&counter, 1)
		dogChan <- struct{}{}
		//<-catChan
	}
}
