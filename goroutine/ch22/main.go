package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//1，监控全局变量来完成  2，通过channel来完成,3,context

//注意：父的context取消，那么这个父context生成的子context也会取消
func cpuInfo(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cpu监控信息完成")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("获取cpu信息成功")
		}

	}

}

func main() {
	//context应用场景

	ctx, cancel := context.WithCancel(context.Background())
	//context.WithDeadline()
	//context.WithTimeout()
	//context.WithValue()
	wg.Add(1)

	go cpuInfo(ctx)
	time.Sleep(time.Second * 6)
	cancel()

	wg.Wait()

	fmt.Println("主--信息监控完结束")
}
