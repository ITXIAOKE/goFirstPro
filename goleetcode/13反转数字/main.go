package main

import "fmt"

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	isPos := true
	if x < 0 {
		isPos = false
		x = 0 - x
	}
	var num int
	for x > 0 {
		temp := x % 10      //123->3  \\12->2  \\1->1
		x = x / 10          //123->12  \\1\\0
		num = num*10 + temp //3||32||321

		if num < -2147483648 || num > 2147483647 { //2的32次方
			return 0
		}
	}
	if !isPos {
		num = 0 - num
	}
	return num
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}
