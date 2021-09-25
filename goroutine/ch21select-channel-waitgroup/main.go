package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//1，监控全局变量来完成  2，通过channel来完成,3,context
var stop = make(chan bool)

func cpuInfo() {
	defer wg.Done()
	for {
		select {
		case <-stop: //这个默认会阻塞，等待发送进来的消息，只要有信息发送到这个channel中，就开始执行了
			fmt.Println("退出cpu监控")
			return
		default: //上面的case阻塞后，会执行这个default逻辑
			time.Sleep(time.Second * 2)
			fmt.Println("CPU信息读取完成")
		}

	}

}

func main() {
	//context应用场景
	wg.Add(1)
	go cpuInfo()

	time.Sleep(time.Second * 6)
	stop <- true

	wg.Wait()

	fmt.Println("信息监控完成")
}
