package main

import "fmt"

func twoSum(nums []int, target int) []int {

	h := make(map[int]int, len(nums))

	for k, v := range nums {
		if p, ok := h[target-v]; ok {
			return []int{p, k} //走到4的时候返现7-4=3在map中，那就直接返回，程序结束
		}
		h[v] = k
	}
	return []int{}
}
func main() {
	//fmt.Println(twoSum([]int{1, 3, 2, 4, 9, 16}, 7))

	//fmt.Println(fourSumCount([]int{1,3}, []int{4,2}, []int{-3,1}, []int{-2,4}))


	fmt.Println(fourSum([]int{1}, []int{2}))

}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	des := map[int]int{}
	ans:=0
	for _,v :=range A{//遍历两个数组，将两个数组的和作为一个索引，进行+1操作
		for _,w:=range B{
			des[v+w]++
		}
	}
	for _,v :=range C{//遍历另两个数组，如果这两个数组进行相加的和的相反数在map中不为1,则证明出现过
		for _,w:=range D{
			ans +=des[-v-w]
		}
	}
	return ans//返回总数
}



func fourSum(A []int, B []int) int {
	des := make([]int,0)

	for _,v :=range A{//遍历两个数组，将两个数组的和作为一个索引，进行+1操作
		des=append(des,v)
	}

	for _,w:=range B{
		des=append(des,w)
	}

	var res int
	for _, v := range des {
		res+=v
	}


	return res
}
