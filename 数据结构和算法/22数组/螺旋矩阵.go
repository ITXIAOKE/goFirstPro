package _2数组

/**
给你一个正整数n ，生成一个包含 1 到n2所有元素，且元素按顺时针顺序螺旋排列的n x n 正方形矩阵 matrix 。

示例 1：
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]

示例 2：
输入：n = 1
输出：[[1]]

提示：
1 <= n <= 20
*/

func generateMatrix(n int) [][]int {

	arr := make([][]int, n)

	for k, _ := range arr {
		arr[k] = make([]int, n)
	}

	top, bottom := 0, n-1
	left, right := 0, n-1
	m := 1

	for m <= n*n {
		for i := left; i <= right; i++ {
			arr[top][i] = m
			m++
		}
		top++

		for i := top; i <= bottom; i++ {
			arr[i][right] = m
			m++
		}
		right--

		for i := right; i >= left; i-- {
			arr[bottom][i] = m
			m++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			arr[i][left] = m
			m++
		}
		left++
	}
	return arr

}
