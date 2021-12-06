package main

import "fmt"

func main() {
	//defer fmt.Println("main:", func2())
	defer fmt.Println("------------")
	defer fmt.Println("main:", func3())
}

func func2() (sum int) {
	sumA := 100
	sumB := 200
	sum = sumA + sumB
	defer func() {
		fmt.Println("first:", sum)//闭包，这个sum值是函数执行完后，有这个返回值sum,最后的值
	}()
	defer fmt.Println("second:", sum)//函数里面执行完后，显示的值
	return sum + 10
}

func func3() int{
	sumA := 100
	sumB := 200
	sum := sumA + sumB
	defer func() {
		fmt.Println("first:", sum)//闭包，这个sum值是函数执行完后，没有这个返回值sum，
	}()
	defer fmt.Println("second:", sum)//函数里面执行完后，显示的值
	return sum + 10
}
