package main

import (
	"fmt"
)

func selectSort(arr *[9]int) {
	fmt.Println("排序前", *arr)

	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] { //从大到小排序
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}

	fmt.Println("排序后", *arr)
}

func main() {
	var arr = [9]int{1, 23, 123, 123, 321, 22, 33, 12, 31}
	selectSort(&arr)
}
