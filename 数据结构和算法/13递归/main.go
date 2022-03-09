package main

import "fmt"

func main() {

	//for x := range Fib(30) {
	//	fmt.Println(x)
	//}

	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}

//斐波那契数列go实现

func Fib(n int) chan int { //返回n以内的数列
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()
	return out
}

func Fibonacci() func() int { //返回一个函数，然后循环几次就得到几个数列
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
