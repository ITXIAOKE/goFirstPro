package main

import "fmt"

func main() {
	/**
	goroutine池
	本质上是生产者消费者模型
	可以有效控制goroutine数量，防止暴涨
	需求：
	计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6 随机生成数字进行计算
	*/
	fmt.Println(123456 % 10)
	fmt.Println(123456 / 10)

	fmt.Println(12345 % 10)
	fmt.Println(12345 / 10)

	fmt.Println(1234 % 10)
	fmt.Println(1234 / 10)

	fmt.Println(123 % 10)
	fmt.Println(123 / 10)
	fmt.Println("--------------------")
	var sum int
	var res int = 123
	for res != 0 {
		tmp := res % 10
		sum += tmp
		res /= 10
	}
	fmt.Println(sum)
}
