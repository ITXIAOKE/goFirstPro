package main

import (
	"fmt"
	"time"
)

func main() {
	//timer1 := time.NewTimer(2 * time.Second)
	//t1 := time.Now()
	//fmt.Println(t1)
	//
	//t2 := <-timer1.C
	//fmt.Println(t2)

	//timer1 := time.NewTimer(time.Second)
	//for {
	//	<-timer1.C
	//	fmt.Println("时间到")
	//}
	//
	//<-time.After(2 * time.Second)
	//fmt.Println("2秒到。。。")

	timer4 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer4.C
		fmt.Println("定时器执行")
	}()
	b := timer4.Stop()
	if b {
		fmt.Println("timer4已关闭")
	}
	time.Sleep(time.Second*3)
	for{}
}
