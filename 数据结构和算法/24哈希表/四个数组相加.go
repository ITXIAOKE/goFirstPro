package _4哈希表
/**
给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。
首先将四个数组分割为两两数组，前两个数组值相加，后两个数组相加，如果前两个数组相加和 与 后两个数组相加和 正好为相反数，四个元素之和为0.
 */
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	h := make(map[int]int)
	var t int

	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			h[v1+v2]++
		}
	}

	//再次遍历另两个数组，将两个数组的元素进行相加，取和的相反数，通过使用相反数在map中查找，
	//如果没出现，所指向的数是0，如果出现过这个数的相反数，则所指向的数大于一。
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			t += h[0-v3-v4]
		}
	}

	return t
}
