package main

import (
	"fmt"
)

func main() {
	/**
	//变量
	var (
		foo    string
		foo2  string
		_foo1 string
		//break string
	)
	//常量
	const aa=11
	const (
		b=4
		c=0.4
	)
	*/
	var foo string
	var foo2 string="xiaoke2"
	var (
		foo3,foo4 string
	)
	foo5:="xiaoke5"
	fmt.Println(foo)
	fmt.Println(&foo)
	fmt.Println((foo))
	fmt.Println(&foo)
	fmt.Println(len(foo))

	fmt.Println(foo,foo2,foo3,foo4,foo5)
	fmt.Println("=======================")

	const xiaoke1 = 11
	const xiaoke2 int =33
	const (
		a=2
		b=8
	)


}
