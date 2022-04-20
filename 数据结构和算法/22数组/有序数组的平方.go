package _2数组

/**
双指针
创建一个和原来数组一样大小的数组res
比较头指针和尾指针指向数的平方和，将大的数存在res[k--]中

给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1：
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]


示例 2：
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]


*/
func sortedSquares(nums []int) []int {
	l := len(nums)
	res := make([]int, l)

	i, j := 0, l-1
	k := j
	for i <= j {
		if nums[i]*nums[i] >= nums[j]*nums[j] {
			res[k] = nums[i] * nums[i]
			i++
		} else {
			res[k] = nums[j] * nums[j]
			j--
		}
		k--
	}
	return res
}

