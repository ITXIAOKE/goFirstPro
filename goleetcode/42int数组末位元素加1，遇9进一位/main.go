package main

import "fmt"

func plusElementOne(nums []int) []int {
	res := make([]int, 0)
	if nums == nil || len(nums) == 0 {
		return res
	}

	carry := 1 //题意加1
	for i := len(nums) - 1; i >= 0; i-- {
		sum := nums[i] + carry //取出最后一位加1
		carry = sum / 10       //取出加1后的进位值
		nums[i] = sum % 10     //更新最后一位的值
	}

	if carry == 1 { //遇到999999这类的，进一位
		res = []int{1}
		nums = append(res, nums...)
	}

	return nums

}

func main() {
	fmt.Println(plusElementOne([]int{9, 9, 9}))
	fmt.Println(plusElementOne([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(plusElementOne([]int{2, 2, 3, 3, 4, 5, 6, 2, 2, 2}))
	fmt.Println(plusElementOne([]int{4, 9}))
	fmt.Println(plusElementOne([]int{5}))
	fmt.Println(plusElementOne([]int{5, 5, 5, 5, 5, 1, 2, 3, 4, 6}))
}
