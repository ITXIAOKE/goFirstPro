package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  @params ch 双向通道，接收int型；
  @params wg 等待协程完成，注意这里传的是指针值（还是值传递）
*/
func producer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		i := 1
		for {
			ch <- i
			fmt.Println("p_in: ", i)
			// 等待10秒钟，如果想看明显的读阻塞，这里的沉睡时间就小很多倍，反之下面消费者里面就长很多倍；
			//time.Sleep(time.Second)
			time.Sleep(time.Microsecond * 10)
			// 这里是控制到20次循环时就关闭通道
			if i == 20 {
				close(ch)
				fmt.Println("close ch")
				break
			}
			i++
		}
		fmt.Println("producer done")
		// 这句代表当前这次整体操作ok，不用再等待他了。
		wg.Done()
	}()
}

/**
@params: ch 双向通道 同上；
   @params: wg 同上
*/
func consumer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			// 如果通道没有被关闭，正常接收值，通过ok来进行判断
			if value, ok := <-ch; ok {
				fmt.Println(value, ok)
				// 这里没有满20次时，每次沉睡1秒，这样可以很明显看到写操作的阻塞情况
				if value != 20 {
					time.Sleep(time.Second * 1)
				}
			} else {
				// 通道关闭，ok 为false，就走到这里了
				break
			}
		}
		fmt.Println("consumer done")
		wg.Done()
	}()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	// 这里说的是我要开始执行一个新的大任务了；
	wg.Add(1)
	// 指针值传递，是为了在函数里面可以修改对应的内部值，通过wg.Done()里面自动实现减1
	producer(ch, &wg)
	wg.Add(1)
	consumer(ch, &wg)
	wg.Wait()

	//time.Sleep(time.Second * 10) 使用 wg 就不需要执行sleep了
}

/*
 可能碰到的问题：

 // 在一段时间内，生产者可以不断往 channel 写入数据，消费者进行处理，一段时间后 channel 关闭了，
 // 这个时候如果还有数据往 channel 发送，程序就会报错。：panic: send on closed channel
 //close(ch)

 //fatal error: all goroutines are asleep - deadlock! 等待其他线程放入数据
*/