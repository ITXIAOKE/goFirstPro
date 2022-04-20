package _4哈希表

import "sort"

/**
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

示例 2：
输入：nums = []
输出：[]

示例 3：
输入：nums = [0]
输出：[]

两层循环+双指针
*/

func threeSum(nums []int) [][]int {
	var res [][]int

	sort.Ints(nums)//排序

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for right > left {
			n2 := nums[left]
			n3 := nums[right]

			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for right > left && nums[left] == n2 {
					left++
				}
				for right > left && nums[right] == n3 {
					right--
				}
			} else if n1+n2+n3 > 0 {
				right--
			} else if n1+n2+n3 < 0 {
				left++
			}
		}
	}
	return res
}
