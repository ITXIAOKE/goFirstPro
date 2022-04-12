package main

import "fmt"

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1
	res[2] = 1
	for i := 3; i <= n; i++ {
		res[i] = res[i-1] + res[i-2] + res[i-3]
	}
	return res[n]
}

func main() {
	fmt.Println(tribonacci(4))
}
