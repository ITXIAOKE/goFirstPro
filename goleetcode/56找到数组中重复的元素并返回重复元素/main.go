package main

import (
	"fmt"
)

func removeElementAfter(nums []int) int {
	slow, fast := nums[0], nums[0]
	slow = nums[slow]
	fast = nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	pointer1 := nums[0]
	pointer2 := slow

	for pointer1 != pointer2 {
		pointer1 = nums[pointer1]
		pointer2 = nums[pointer2]
	}

	return pointer2
}

func main() {
	fmt.Println(removeElementAfter([]int{1, 3, 4, 2, 2}))
	fmt.Println(removeElementAfter([]int{3, 1, 3, 4, 2}))
}
