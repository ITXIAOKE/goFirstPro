package _4哈希表

func twoSum(nums []int, target int) []int {

	h := make(map[int]int, len(nums))

	for k, v := range nums {
		if p, ok := h[target-v]; ok {
			return []int{k, p}
		}
		h[v] = k
	}
	return []int{}
}


