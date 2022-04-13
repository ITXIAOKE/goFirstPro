package main

import "fmt"

func MySqrt(x int) int {
	if x <= 1 {
		return x
	}
	left, right := 1, x
	for left < right {
		mid := left + (right-left)/2
		if mid == x/mid {
			return mid
		} else if mid > x/mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left - 1
}

func main() {
	fmt.Println(MySqrt(16))
	fmt.Println(MySqrt(9))
	fmt.Println(MySqrt(8))
	fmt.Println(MySqrt(4))
	fmt.Println(MySqrt(3))
	fmt.Println(MySqrt(1))

}
