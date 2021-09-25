package main

import (
	"fmt"
)

func main() {
	//select  判断管道是否存满
	ch := make(chan string, 10)
	go writer(ch)
	for x := range ch {
		fmt.Println(x)
		if x == "channel full" {
			break
		}
	}
}

func writer(ch chan string) {
	for {
		select {
		case ch <- "hello":
		default:
			ch <- "channel full"
			fmt.Println("channel full")
		}
	}

}
