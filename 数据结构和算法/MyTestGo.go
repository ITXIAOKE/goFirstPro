package main

import "fmt"

func main(){
	x:=10
	defer func (x int){
		fmt.Println(x)
	}(x)
	x++




}
