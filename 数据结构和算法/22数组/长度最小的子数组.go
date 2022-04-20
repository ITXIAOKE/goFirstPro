package _2数组

import "math"
/**
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，
并返回其长度。如果不存在符合条件的子数组，返回 0。

1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 105
 */

func minSubArrayLen(s int, nums []int) int {
	sum := 0
	start := 0
	end := 0
	a := math.MaxInt32

	for end < len(nums) {
		sum = sum + nums[end]
		for sum >= s {
			sum = sum - nums[start]
			a = min(a, end-start+1)
			start++
		}
		end++
	}
	if a == math.MaxInt32 {
		return 0
	} else {
		return a
	}
}
func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
