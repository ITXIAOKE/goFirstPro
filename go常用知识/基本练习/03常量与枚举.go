package main

import (
	"fmt"
	"math"
)

const filename11 string = "abx.txt"

func consts() {
	const filename string = "ajb.txt"

	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func consts2() {
	const (
		filename string = "ajb2.txt"
		a, b            = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		java
		python
	)
	fmt.Println(cpp, java, python)

	const(
		b=1<<(10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

func main() {
	consts()
	consts2()
	fmt.Println(filename11)
	enums()

}
