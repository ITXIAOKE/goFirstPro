package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	const filename string = "/Users/xiaoke/Desktop/GoProject/goFirstPro/基本练习/abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func readFile2() {

	const filename22 = "/Users/xiaoke/Desktop/GoProject/goFirstPro/基本练习/xiaoke.txt"
	if contents,err:=ioutil.ReadFile(filename22);err!=nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n", contents)
	}


}

func main() {
	fmt.Println("---")
	readFile()
	fmt.Println("---")
	readFile2()
	fmt.Println("---")
	fmt.Println("---")
	fmt.Println("---")

}
