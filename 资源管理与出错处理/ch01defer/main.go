package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//defer 确保调用 在函数结束时 发生
	//defer 不管函数中有return还是panic，defer都会执行
	//defer 列表为后进先出
	//tryDefer()
	//writeFile("fib.txt")
	tryDeferJisuan() //参数在defer语句时计算
}

func tryDeferJisuan() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i) //这个i参数就是在语句中计算的，先进后出，数据倒着输出，结果不会是都打印30行30，而是打印30行从0到30的数据
		if i == 30 {
			panic("print too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
		//return
	}
	defer file.Close() //打开文件最后一定要记得close文件

	writer := bufio.NewWriter(file)
	defer writer.Flush() //写文件最后一定要flush，否则数据不会写到文件中去，刷新之后数据才会写到文件中

	f := Fibonacci()
	for i := 0; i < 20; i++ { //打印前20个数列
		fmt.Fprintln(writer, f())
	}
}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println(4)
	//return
	//panic("error xiaoke")
	fmt.Println(5)
}
