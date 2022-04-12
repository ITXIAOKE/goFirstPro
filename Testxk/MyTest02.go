package main

import "fmt"

func main() {

	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(operMySile3(nums))
}

func operMySile3(nums []int) [][]int {
	var sli [][]int
	var sli1 []int
	for i := 0; i < len(nums)-1; i++ {
		var sli2 []int
		if nums[i+1]-nums[i] == 1 {
			sli1 = append(sli1, nums[i])
			sli = append(sli, sli1)
		} else {
			sli2 = append(sli2, nums[i+1])
			sli = append(sli, sli2)
		}
	}
	//
	return sli
}

//func operMySile2(nums []int) [][]int {
//	var sli [][]int
//	var sli1 []int
//
//	for i := 0; i < len(nums)-1; i++ {
//
//		if nums[i+1]-nums[i] == 1 {
//			sli1 = append(sli1, nums[i])
//		} else {
//			sli1 = append(sli1, nums[i])
//
//			var sli2 []int
//			sli2 = append(sli2, nums[i+1])
//			sli = append(sli, sli2)
//		}
//	}
//
//	sli = append(sli, sli1)
//	return sli
//}

func operMySile(nums []int) []int {
	//sli:=[][]int{}
	//var sli [][]int
	var sli1 []int
	mymap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mymap[nums[i]] = i
	}
	fmt.Println(mymap)
	return sli1
}

//func main() {
//	//定义一个二维数组
//	var arr = [2][3]int{{1, 4, 3},{7, 5, 6}}
//
//	//方式1. 用for循环来遍历
//	for i := 0; i < len(arr); i++ {
//		for j := 0; j < len(arr[i]); j++ {
//			fmt.Printf("%v ",arr[i][j])
//		}
//		fmt.Println()
//	}
//
//	//方式2. for-range 遍历
//	for i, v := range arr {
//		for j, v2 := range v {
//			fmt.Printf("arr[%v][%v]=%v \t \n", i, j, v2)
//		}
//		fmt.Println()
//	}
//}
