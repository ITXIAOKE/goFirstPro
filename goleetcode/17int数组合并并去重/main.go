package main

import "fmt"

func main() {
	v1 := []int{1, 2, 3, 44, 22, 33}
	v2 := []int{11, 2, 33, 44, 220, 33}

	//合并
	res := make([]int, 0)
	res = append(res, v1...)
	res = append(res, v2...)
	fmt.Println(res)

	//排序
	sortslice := quickSort(res)
	fmt.Println(sortslice)

	//去重
	myslice:=make([]int,0)
	mymap:=make(map[int]interface{})
	for _,val:=range sortslice{
		if _,ok:=mymap[val];!ok{
			myslice=append(myslice,val)
			mymap[val]=nil
		}
	}
	fmt.Println(myslice)

}

func quickSort(res []int) []int {
	if len(res) <= 1 {
		return res
	}
	low := make([]int, 0)
	mid := make([]int, 0)
	hight := make([]int, 0)
	flag := res[0]
	mid = append(mid, flag)

	for i := 1; i < len(res); i++ {//这里一定要从第一个开始，否则就会重复的打印数字
		if res[i] < flag {
			low = append(low, res[i])
		} else if res[i] > flag {
			hight = append(hight, res[i])
		} else {
			mid = append(mid, res[i])
		}
	}
	low, hight = quickSort(low), quickSort(hight)
	return append(append(low, mid...), hight...)

}
