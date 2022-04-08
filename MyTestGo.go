package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-1, -3, 7, 5, 11, 15}
	arrLength := len(arr) //计算数组长度

	myMap := make(map[int]int)

	for i := 0; i <= arrLength-1; i++ {
		for j := i; j < arrLength-i-1; j++ {
			if math.Abs(float64(arr[i]+arr[j])) < math.Abs(float64(arr[i+1]+arr[j+1])) {
				myMap[arr[i]] = arr[j]
			}
		}
	}
	fmt.Println(myMap)

	myarr := make([]int, 0)
	for k, v := range myMap {
		myarr = append(myarr, k+v)
	}

	fmt.Println(myarr)
}
