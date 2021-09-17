package main

import (
	"fmt"
	"strconv"
)

func main() {
	aa:=[]string{"xiaoke","ww","23"}
	//fmt.Println(aa)
	//fmt.Println(len(aa))
	//fmt.Println(aa[len(aa)-1])
	intval, err := strconv.Atoi(aa[len(aa)-1])
	fmt.Println(intval)
	fmt.Println(err)
	fmt.Println(strconv.Atoi("11"))

	//fmt.Println(aa[0])
	//fmt.Println(aa[1])
	//fmt.Println(aa[2])
	//fmt.Println(aa[3])
	//fmt.Println(aa[4])
	//fmt.Println(aa[5])


	//i, err := strconv.Atoi("12345")
	//if err != nil {
	//	panic(err)
	//}
	//i += 3
	//println(i)
	//
	//s := strconv.Itoa(12345)
	//s += "3"
	//println(s)

}
