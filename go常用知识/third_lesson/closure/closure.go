package main

import (
	"fmt"
	"time"
)

func main() {

	i := 13
	a := func() {
		fmt.Printf("i is %d \n", i)
	}
	a()
	fmt.Println(ReturnClosure("Tom")())

	Delay()
	Delay2()
	time.Sleep(time.Second)
}

func ReturnClosure(name string) func() string {
	return func() string {
		return "Hello, " + name
	}
}

//闭包延时绑定的问题

func Delay() {
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		fns = append(fns, func() {
			fmt.Printf("hello, this is : %d \n", i)
		})
	}

	for _, fn := range fns {
		fn()
	}
}

func Delay2() {
	fns := make([]func(), 0, 10) //初始化一个10个空间的函数数组

	for i := 0; i < 10; i++ {
		a := i
		fns = append(fns, func() {
			fmt.Printf("hello, this is : %d \n", a)
		})
	}

	for _, fn := range fns {
		fn()
	}
}
