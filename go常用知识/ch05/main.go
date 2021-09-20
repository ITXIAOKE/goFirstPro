package main

import "fmt"

func main() {
	/**
	ch这个channel的buffer大小是1，所以会交替的为空或为满，所以只有一个 case可以进行下去，无论i是奇数或者偶数，它都会打印0 2 4 6 8。
	 */
	ch := make(chan int,1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
			fmt.Println("done")
		}
	}
}
