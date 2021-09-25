package main

import (
	"fmt"
)

func main() {
	//panic :1，停止当前函数执行   2，一直向上返回，执行每一层的defer   3，如果没有遇见recover，程序退出，遇见的话，错误异常信息就被recover给收集到了
	//recover:1，仅在defer调用中使用  2，获取panic的值    3，如果无法处理，可重新panic

	tryRecover()
}

func tryRecover() {
	defer func() {
		r := recover() //接收程序的错误
		if err, ok := r.(error); ok {
			fmt.Println("抓到异常：", err)
		} else {
			panic(r)
		}
	}()
	//panic(errors.New("this is an error"))

	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	panic(123) //这种的最后还是panic出来，因为这个123不是直接的错误
}
