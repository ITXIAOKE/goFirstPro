package _4哈希表

func intersection(nums1 []int, nums2 []int) []int {
	var arr []int
	h := make(map[int]int, 1)
	for _, k := range nums1 {
		h[k]++
	}

	for _, v := range nums2 {
		if h[v] > 0 {
			arr = append(arr, v)
			h[v] = 0
		}
	}

	return arr
}
