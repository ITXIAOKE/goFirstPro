package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var stop bool //1，监控全局变量来完成  2，通过channel来完成 ，3，context

func cpuInfo() {
	defer wg.Done()
	for {
		if stop {
			break
		}
		time.Sleep(time.Second * 2)
		fmt.Println("CPU信息读取完成")
	}

}

func main() {
	//context应用场景
	wg.Add(1)
	go cpuInfo()

	time.Sleep(time.Second * 6)
	stop = true

	wg.Wait()

	fmt.Println("信息监控完成")
}
