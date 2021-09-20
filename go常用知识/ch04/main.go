package main

import (
	"fmt"
	"os"
)

func main() {
	//第一个版本
	//var s, temp string
	//for i := 1; i < len(os.Args); i++ {
	//	s += temp + os.Args[i]
	//	temp = " "
	//}
	//fmt.Println(s)

	//第二个版本
	//s, temp := "", ""
	//for _, arg := range os.Args[1:] {
	//	s += temp + arg
	//	temp = " "
	//}
	//fmt.Println(s)
	//
	//ff1 := [...]string{"aa", "bb"} //数组
	//ff := ff1[:]                   //切片
	//fmt.Println(strings.Join(ff, "+"))
	//
	//g := make([]string, 0) //这里的长度和容量都是0，如果设置为非0，则表示有几个空字符串
	//g = append(g, "mm")
	//g = append(g, "nn")
	//fmt.Println(g)
	//fmt.Println(strings.Join(g, "-"))


	fmt.Println(os.Args[0:])
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[:])
}
