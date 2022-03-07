package main

import (
	"fmt"
	"sync"
	"time"
)

/**
问题：
1，使用三个协程，每秒钟打印cat、dog、fish
2，顺序不能变化
3，无限循环即可
*/

func main() {
	wg := sync.WaitGroup{}

	catChan := make(chan bool) //这里放缓冲一样
	dogChan := make(chan bool)
	fishChan := make(chan bool)

	cat(fishChan, catChan, &wg)
	dog(catChan, dogChan, &wg)
	fish(dogChan, fishChan, &wg)

	wg.Wait()
}

func fish(dogChan chan bool, fishChan chan bool, w *sync.WaitGroup) {
	w.Add(1)
	go func() {
		for {
			<-dogChan
			fmt.Println("fish")
			time.Sleep(time.Second)
			fmt.Println(time.Now())
			fishChan <- true

		}
		w.Done() //必须放在协程里面
	}()
	//w.Done()//不能放这里，否则程序啥都不输出
}

func dog(catChan chan bool, dogChan chan bool, w *sync.WaitGroup) {
	w.Add(1)
	go func() {
		for {
			<-catChan
			fmt.Println("dog")
			dogChan <- true
		}
		w.Done()
	}()
}

func cat(fishChan chan bool, catChan chan bool, w *sync.WaitGroup) {
	w.Add(1)
	go func() {
		for {
			//<-fishChan//注意这里不能先这样，按照程序运行的先后顺序，否则deadlock
			fmt.Println("cat")
			catChan <- true
			<-fishChan
		}
		w.Done()
	}()
}
