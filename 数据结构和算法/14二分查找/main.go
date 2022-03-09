package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(binary_search(0, len(arr), arr, -1))

	arr2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(binary_search2(arr2, -1))
}

//递归方法

func binary_search(left int, right int, arr []int, data int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2

	if arr[mid] == data {
		return mid
	} else if arr[mid] > data {
		return binary_search(left, right-1, arr, data)
	} else {
		return binary_search(left+1, right, arr, data)
	}
}

func binary_search2(arr []int, data int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > data {
			right = mid - 1
		} else if arr[mid] < data {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
