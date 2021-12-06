//package main
//
//import "fmt"
//
////闭包测试
//func func1() func(int) int {
//	sum := 0
//	return func(val int) int {
//		sum += val
//		return sum
//	}
//}
//
//func printFunc() {
//	mysum := func1()
//	fmt.Println(mysum(1))
//	fmt.Println(mysum(1))
//	fmt.Println(mysum(1))
//}
//
//func main() {
//	printFunc()
//
//}

//结果：1   2   3
//-------------------------------------------------------

//package main
//
//import "fmt"
//
////闭包测试，有变量
//func mysum() (val int) {
//	val = 10
//	defer func() {
//		val += 1
//	}()
//	//return  //结果11
//	return 100 //结果101
//}
//
//func printFunc() {
//	fmt.Println(mysum())
//}
//
//func main() {
//	printFunc()
//
//}

package main

import "fmt"

//闭包测试,没有变量
func mysum() int {
	val := 10
	defer func() {
		val += 1
	}()
	return val //结果10
	//return 100 //结果100
}

func printFunc() {
	fmt.Println(mysum())
}

func main() {
	printFunc()

}
