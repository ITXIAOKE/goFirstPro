package main

import (
	"fmt"
)

func main() {
	a := "123abcd爱"
	b := "abxiaoke我可爱"

	//testStr := make([]string, 0)
	//testStr = append(testStr, "haha", "hehe", "hoho", "hehe")
	//testStr = append(testStr, a)
	//testStr = append(testStr, b)
	//fmt.Println(strings.Join([]string{},"/"))

	resStr := []rune(a + b) //这一步将字符串变成rune切片，是为了不乱码进行后面的处理
	fmt.Println(resStr)

	//遍历字节切片，变成字符串切片
	var res []string
	for i := 0; i < len(resStr); i++ {
		res = append(res, string(resStr[i]))
	}
	fmt.Println(res)

	//字符串切片排序
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res)-i-1; j++ {
			if res[j] > res[j+1] {
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
	}
	fmt.Println(res)

	//字符串切片去重
	//var rStr []string
	//for i := 0; i < len(res); i++ {
	//	if i < len(res)-1 && res[i] == res[i+1] {
	//		continue
	//	}
	//	rStr = append(rStr, res[i])
	//}
	//fmt.Println(rStr)

	//var rStr []string
	rStr := make([]string, 0)
	mymap := make(map[string]interface{})
	for _, val := range res {
		if _, ok := mymap[val]; !ok {
			rStr = append(rStr, val)
			mymap[val] = nil
		}
	}
	fmt.Println(rStr)

	//将字符串切片变成一个字符串
	var code string
	for i := 0; i < len(rStr); i++ {
		code = code + rStr[i]
	}
	fmt.Println(code)

}
