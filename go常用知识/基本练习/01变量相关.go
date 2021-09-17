package main

import (
	"fmt"
)

func main2() {
	str := "xiaoke"
	fmt.Println(str)
}

func abs() {

}

func variable() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

const (
	a = iota
	b
	c
)

func variableInitialValue() {
	var a int = 3
	var s string = "abc"
	fmt.Println(a, s)
}
func variableTypeDeduction() {
	var a, b, c, d = 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	//:=这种声明变量的方式只能在函数内使用，不能在函数外使用
	a, b, c, d := 3, 4, false, "func"
	b = 6777
	fmt.Println(a, b, c, d)
}

func main() {
	//variableInitialValue()
	//main2()
	//variableTypeDeduction()
	//fmt.Println("hello go!! ")
	variableShorter()
	//var color11 = []color2.Color{color2.White, color2.Black}
	//
	//for i, v := range color11 {
	//	fmt.Println(i, v)
	//}

	//var gender bool = true
	//m := 5
	//n := strconv.Itoa(m)
	//fmt.Printf("%T,%s",n,n)
	//variable()

}
