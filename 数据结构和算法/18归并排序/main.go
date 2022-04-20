package main

/**
归并排序思想
归并排序（MERGE-SORT）是利用归并的思想实现的排序方法，该算法采用经典的分治策略
（分治法将问题分成一些小的问题然后递归求解，而治的阶段则将分的阶段得到的各答案"修补"在一起，即分而治之)。

归并排序是用分治思想，分治模式在每一层递归上有三个步骤：

分解（Divide）：将n个元素分成个含n/2个元素的子序列。
解决（Conquer）：用合并排序法对两个子序列递归的排序。
合并（Combine）：合并两个已排序的子序列已得到排序结果。

复杂度
最理想情况：O(nlogn)
最坏情况：O(nlogn)

*/

import (
	"fmt"
)

func mergeSort(nums []int) []int {
	if len(nums) < 2 { // 分治，两两拆分，一直拆到基础元素才向上递归。
		return nums
	}
	i := len(nums) / 2
	left := mergeSort(nums[0:i]) // 左侧数据递归拆分
	right := mergeSort(nums[i:]) // 右侧数据递归拆分
	result := merge(left, right) // 排序 & 合并
	return result
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	l, r := len(left), len(right)
	for i < l && j < r {
		if left[i] > right[j] { //从小到小大排序
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}
	result = append(result, right[j:]...)
	result = append(result, left[i:]...)
	return result
}

func main() {
	arr := []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	result := mergeSort(arr)
	fmt.Println(result)
}
