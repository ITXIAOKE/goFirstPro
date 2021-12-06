package main

import "fmt"

func main() {
	m := 10
	defer fmt.Println("first:", m)
	m = 100
	defer func() {
		fmt.Println("second:", m) //闭包，主函数执行完成后，最后的这个m值
	}()
	//defer fmt.Println("second:", m)//这个正常的显示值

	m *= 20
	defer fmt.Println("third:", m)

	funcVal := func1() //闭包
	funcVal()

}

func func1() func() {
	defer fmt.Println("before return ")
	return func() {
		defer fmt.Println("in the return")
	}
}

/**
结果：
before return
in the return
third: 2000
second: 2000
first: 10


*/
