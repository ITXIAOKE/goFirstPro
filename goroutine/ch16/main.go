package main

import (
	"fmt"
)

func main() {
	output1 := make(chan string,1)
	output2 := make(chan string,1)
	go test1(output1)
	go test2(output2)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
	fmt.Println("main结束")
}

func test1(output1 chan string) {
	output1 <- "test1"
}

func test2(output2 chan string) {
	output2 <- "test2"
}
