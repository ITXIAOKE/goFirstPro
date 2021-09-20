package main

import (
	"fmt"
	"time"
)

func myPrint() {
	i := 1
	for {
		i++
		fmt.Println(i)
	}

}

func main() { //这个是主线程，或者叫主协程
	/**
	协程：用户态的线程
	遵从主死从随
	*/
	//go myPrint() //开启协程
	go func() {
		i := 1
		for {
			i++
			fmt.Println(i)
		}
	}()
	time.Sleep(time.Second)

}
