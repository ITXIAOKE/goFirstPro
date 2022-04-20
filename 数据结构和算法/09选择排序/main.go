package main

import (
	"fmt"
)

/**
思想：

选择排序是不稳定的排序算法，不稳定发生在最小元素与A[i]交换的时刻

最优时间复杂度：O(n^2)
最坏时间复杂度：O(n^2)
稳定性：不稳定（考虑升序每次选择最大的情况）

选择排序（Selection sort）是一种简单直观的排序算法。它的工作原理如下：

首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，
然后，再从剩余未排序元素中继续寻找最小（大）元素，
然后放到已排序序列的末尾。
以此类推，直到所有元素均排序完毕。

*/

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
		//交换。相邻的两个元素再次进行交互
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
