package main

import (
	"fmt"
)

func intersect(num1 []int, num2 []int) []int {
	var res []int
	if num1 == nil || len(num1) == 0 || num2 == nil || len(num2) == 0 {
		return res
	}
	m := make(map[int]int)
	for i := 0; i < len(num1); i++ {
		m[num1[i]]++
	}

	for i := 0; i < len(num2); i++ {
		if val, ok := m[num2[i]]; ok {
			if val > 0 {
				res = append(res, num2[i])
			}
			m[num2[i]]--
		}
	}
	return res
}

func main() {

	fmt.Println(intersect([]int{1, 2, 45, 9}, []int{2, 3, 5, 9}))

}
