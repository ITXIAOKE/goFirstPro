package main

import "fmt"

func isPerfectSquare(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 {
		return true
	}
	left, right := 0, num
	for left <= right {
		mid := left + (right-left)/2

		if mid == num/mid && num%mid == 0 {
			return true
		} else if mid == num/mid && num%mid > 0 {
			left = mid + 1
		} else if mid < num/mid {
			left = mid + 1
		} else if mid > num/mid {
			right = mid - 1
		}
	}

	return false
}

func main() {
	fmt.Println(isPerfectSquare(14))
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(7))
	fmt.Println(isPerfectSquare(5))
	fmt.Println(isPerfectSquare(8))
	fmt.Println(isPerfectSquare(9))
	fmt.Println(isPerfectSquare(4))

}
