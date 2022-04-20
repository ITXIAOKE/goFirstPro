package _4哈希表

import "sort"

/**
给你一个由 n 个整数组成的数组nums ，和一个目标值 target 。请你找出并返回满足下述
全部条件且不重复的四元组[nums[a], nums[b], nums[c], nums[d]]（若两个四元组元素一一对应，则认为两个四元组重复）：

0 <= a, b, c, d< n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。

示例 1：
输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

示例 2：
输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]

*/

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] {
				continue
			}
			l := j + 1
			r := len(nums) - 1
			for l < r {
				n3 := nums[l]
				n4 := nums[r]

				if n1+n2+n3+n4 == target {
					res = append(res, []int{n1, n2, n3, n4})
					for l < r && nums[l+1] == nums[l] {
						l++
					}
					for l < r && nums[r-1] == nums[r] {
						r--
					}
					l++
					r--
				} else if n1+n2+n3+n4 > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return res
}
