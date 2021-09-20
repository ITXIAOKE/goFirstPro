package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //控制协程退出，实现goroutine的同步机制

/**
add
done
wait

add的数量和done的数量一致
*/

func main() {
	//wg.Add(5) //知道有5个goroutine，所以可以直接填写5
	for i := 0; i < 10; i++ {
		wg.Add(1) //每次循环都添加一个
		go func(i int) {
			defer wg.Done() //这个结束忘记写的话，会发生死锁   fatal error: all goroutines are asleep - deadlock!
			fmt.Println(i)
		}(i)
	}
	wg.Wait()

	//多次执行上面的代码，会发现每次打印的数字的顺序都不一致。这是因为10个goroutine是并发执行 的，而goroutine的调度是随机的。
}
