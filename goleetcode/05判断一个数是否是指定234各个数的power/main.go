package main

import "fmt"

func isPowerOfTwo(num int) bool {
	return isPowerOfN(num, 2)
	//return isPowerOfN(num, 3)
	//return isPowerOfN(num, 4)
}

func isPowerOfN(num int, n int) bool {
	if num < 1 {
		return false
	}
	for num%n == 0 {
		num = num / n
	}
	return num == 1
}

func main() {
	fmt.Println(isPowerOfTwo(8))
}
