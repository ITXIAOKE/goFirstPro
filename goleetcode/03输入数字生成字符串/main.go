package main

import "fmt"

func generateTheString(n int) string {
	if n == 0 {
		return ""
	}
	if n%2 == 0 {
		return buildString(n-1) + "b"
	} else {
		return buildString(n)
	}
}

func buildString(n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += "a"
	}
	return res
}

func main() {
	fmt.Println(generateTheString(6))
}
