package main

import "fmt"

func main() {
	var a uint = 1
	var b uint = 2
	fmt.Println(a - b)
}

//1-2=-1可以转换为0-1，也就是0+（-1），-1在计算机中也就是所有位数都是1，用一个无符号数区识别，则就是该操作系统的最大值
//结果：32位操作系统2的32次方-1，64位操作系统2的64次方-1

//rune   int32
//rune更多用来做一个字符长度计算（计算中文）
//int32是字节长度（计算中文）
