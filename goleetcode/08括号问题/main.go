package main

import "fmt"

func generateParenthesis(n int) []string {
	var res []string
	if n == 0 {
		return res
	}
	backtrack(n, n, "", &res)
	return res
}

func backtrack(left int, right int, cur string, res *[]string) {
	if left < 0 || right < 0 {
		return
	}
	if right < left {
		return
	}
	if left == 0 && right == 0 {
		*res = append(*res, cur)
		return
	}
	//left
	cur = cur + "("
	backtrack(left-1, right, cur, res)
	cur = cur[:len(cur)-1]

	//right
	cur = cur + ")"
	backtrack(left, right-1, cur, res)
	cur = cur[:len(cur)-1]
}

func main() {
	fmt.Println(generateParenthesis(3))
}
