package _2数组

func search(nums []int, target int) int {

	r, l, mid := 0, len(nums)-1, (len(nums)-1)/2
	for l >= r {
		mid = (r + l) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			l = l - 1
		} else {
			r = r + 1
		}
	}

	return -1
}
