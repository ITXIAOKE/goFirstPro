package main

import (
	"fmt"
	"net"
	"runtime"
	"runtime/pprof"
	"sync"
)

var threadProfile = pprof.Lookup("threadcreate")

func main() {
	fmt.Printf("协程执行之前的线程数：%d\n", threadProfile.Count())
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			runtime.LockOSThread() //强制回收线程
			for j := 0; j < 10; j++ {
				_, err := net.LookupHost("www.baidu.com")
				if err != nil {
					panic(err)
				}
			}
		}()
	}
	wg.Wait()
	fmt.Printf("协程执行之后的线程数：%d\n", threadProfile.Count())

}
