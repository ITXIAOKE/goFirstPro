package main

import "fmt"

//分析冒泡排序

func BubbleSort(arr []int) {
	fmt.Println("排序前", arr)

	for j := 0; j < len(arr)-1; j++ {
		for i := 0; i < len(arr)-1-j; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}

	fmt.Println("排序后", arr)
}

var arr = []int{1, -2, 241, 12, 69, 80, 57, -13}

func main() {
	BubbleSort(arr) //传入数组的地址

}
