package main

import (
	"fmt"
)

func main() {
	//	给定一个字符串数组
	//["I","am","stupid","and","weak"]
	//	用 for 循环遍历该数组并修改为
	//	["I","am","smart","and","strong"]
	myArray:=[5]string{"I","am","stupid","and","weak"}
	fmt.Println(myArray)
	for k,v :=range myArray{

		switch v {
		case "stupid":
			myArray[k]="smart"
		case "weak":
			myArray[k]="strong"
		}

	}
	fmt.Println(myArray)




}
