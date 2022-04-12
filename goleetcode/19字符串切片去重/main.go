package main

import (
	"fmt"
	"sort"
)

//去除重复字符串和空格

func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return ret
}

func main() {
	//当数据量小和数据量大时分别考虑用双重for循环方法和map主键唯一方法
	a := []string{"hello", "", "world", "yes", "hello", "nihao", "hello", "yes", "nihao", "good", "可以", "可以"}
	sort.Strings(a) //排序
	fmt.Println(a)

	fmt.Println(RemoveDuplicatesAndEmpty(a))

	fmt.Println(removeDuplicate(a))
}

// 通过map主键唯一的特性过滤重复元素
func removeDuplicate(arr []string) []string {
	resArr := make([]string, 0)
	tmpMap := make(map[string]interface{})
	for _, val := range arr {
		//判断主键为val的map是否存在
		if _, ok := tmpMap[val]; !ok {
			resArr = append(resArr, val)
			tmpMap[val] = nil //将val这个值放入新的切片中后，同时把这个val对应的value值设置为nil，这样下次遍历的时候，就不会再次放入这个值了
		}
	}
	fmt.Println(tmpMap)
	return resArr
}
