package main

import "fmt"

func combine(n int, k int) [][]int {
	res := [][]int{}
	if k == 0 || n == 0 {
		return res
	}
	cur := []int{}
	backtrack(n, k, 1, cur, &res)
	return res
}

func backtrack(n int, k int, start int, cur []int, res *[][]int) { //start代表n以下的数字从1开始
	if len(cur) == k {
		temp := make([]int, k)
		copy(temp, cur)
		*res = append(*res, temp)
		return
	}
	for i := start; i <= n; i++ {
		cur = append(cur, i)
		backtrack(n, k, i+1, cur, res)
		cur = cur[:len(cur)-1]
	}
}

func main() {
	//n=4   k=2
	//1234
	//输出 【【34】【24】【23】【12】【13】【14】】
	fmt.Println(combine(4, 2))
}
