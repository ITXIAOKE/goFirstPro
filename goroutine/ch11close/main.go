package main

import "fmt"

func main() {
	/**
	可以通过内置的close()函数关闭channel，（如果你的管道不往里存值或者取值的时候一定哟记得关闭管道）
	*/
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) //没有关闭channel，发生deadlock
	}()

	for {
		if data, ok := <-ch; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main 结束")
}
