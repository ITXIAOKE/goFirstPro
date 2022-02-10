package main

import "fmt"

func main() {
	fmt.Println(fibonacci(6))
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}


//func main() {
//	f := fanacc()
//	fmt.Println(f())
//	fmt.Println(f())
//	fmt.Println(f())
//	fmt.Println(f())
//	fmt.Println(f())
//	fmt.Println(f())
//	fmt.Println(f())
//}
//
//func fanacc() func() int {
//	a, b := 0, 1
//	return func() int {
//		a, b = b, a+b
//		return a
//	}
//}
