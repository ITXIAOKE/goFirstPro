package main

import (
	"fmt"
	"time"
)

/*
使用channel完成消费者、生产者的例子，发现使用channel会非常的方便
*/

func producer(intChan chan int) {
	for i := 0; i < cap(intChan); i++ {
		fmt.Println("生产者：", i)
		intChan <- i
	}
	close(intChan)
}

func consumer(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if ok {
			fmt.Println("消费者：", v)
		} else { // 读完了
			break
		}
		time.Sleep(time.Second)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 10)  // “生产者”和“消费者”之间互相通信的桥梁，这里假设生产的元素就是int类型的数字
	exitChan := make(chan bool, 1) // 退出的channel，因为仅做为一个标志所以空间为一个元素就够了

	go producer(intChan)
	go consumer(intChan, exitChan)

	<-exitChan
	fmt.Println("主线程结束！")
}

/**
channel在没有被关闭的时候被遍历，此时会被当前线程阻塞，利用这个特性来实现同步的效果，更加的灵活和方便。
看到这里可能有的小伙伴会有疑问了，既然channel可以解决使用同步锁的阻塞问题，但是你使用了channel还是会阻塞啊，这不是很矛盾么？

说的没错，是的，使用channel可以方便的实现了同步锁的功能，但是我们的程序其实因为同步的关系目前仍然还是会产生阻塞，
不过既然go官方文档说了使用"锁"在go语言中是低级操作，那么官方肯定提供另外一种优雅的遍历的不阻塞的方法，
是的，就是这样，这里我们引入一个关键字select，这个关键字的存在就是为了解决我们的疑问的。
 */