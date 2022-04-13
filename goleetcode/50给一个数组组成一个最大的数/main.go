package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	var res string

	if nums == nil || len(nums) == 0 {
		return ""
	}

	sort.Slice(nums, func(i, j int) bool {
		first, second := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		//"30"+"3"="303"
		//"3"+"30"="330"
		if first+second >= second+first {
			return true
		}
		return false
	})
	fmt.Println(nums)
	for _, num := range nums { //转换为字符串
		res += strconv.Itoa(num)
	}
	if res[0] == '0' {
		return res[:1]
	}

	return res
}

func main() {

	fmt.Println(largestNumber([]int{1, 2, 45, 9}))
	fmt.Println(largestNumber([]int{1, 9, 2, 45, 9}))
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{0, 0}))

}
