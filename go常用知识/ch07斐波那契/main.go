package main

import "fmt"

func main() {
	f := fanacc()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func fanacc() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
