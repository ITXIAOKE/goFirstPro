package main

import "fmt"

//分析冒泡排序

func BubbleSort(arr *[7]int) {
	fmt.Println("排序前", arr)

	tmp := 0 //定义临时变量
	for j := 0; j < len(arr)-1; j++ {
		//多次循环遍历的时候i是越来越小，j是增大的 用len(arry)-i-j实现遍历
		for i := 0; i < len(arr)-1-j; i++ {
			if arr[i] > arr[i+1] {
				tmp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
			}
		}
	}

	fmt.Println("排序后", arr)
}

var arr = [7]int{1, -2, 241, 69, 80, 57, -13}

func main() {
	BubbleSort(&arr) //传入数组的地址

}
