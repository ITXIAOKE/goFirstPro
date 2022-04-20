package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	res := strs[0] //默认设置字符串列表的第0个字符串
	for i := 1; i < len(strs); i++ {
		resLength := len(res)
		tempLength := len(strs[i])
		length := min(resLength, tempLength) //遍历所有字符串列表，拿到长度最小的字符串

		j := 0
		for j < length { //根据最小的字符串长度确定最小遍历的次数
			if res[j] != strs[i][j] { //注意：这种是相同的字符都在字符串的前面，那问题在字符串后面呢？？？
				break
			}
			j++
		}
		res = res[:j] //20行判断了，最后一个值是否相等
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"abllllllll", "abcdff", "abfef"}))
	fmt.Println(longestCommonPrefix([]string{"abllllllllab", "abcdff", "abfef"}))

	fmt.Println(longestCommonPrefix2([]string{"llllllllab", "abcdff", "abfef", "abgggg"}))
	fmt.Println(longestCommonPrefix2([]string{"llllllllab", "abcdff", "abfef", "abgggg"}))
	fmt.Println(longestCommonPrefix2([]string{"abllllllll", "abcdff", "abfef", "ggabgg"}))

}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	res := strs[0] //默认设置字符串列表的第0个字符串

	m := make(map[string]int, 0)
	for i := 1; i < len(strs); i++ {
		resLength := len(res)
		tempLength := len(strs[i])
		length := min(resLength, tempLength) //遍历所有字符串列表，拿到长度最小的字符串,比如拿到第三个

		j := 0
		for j < length { //根据最小的字符串长度确定最小遍历的次数
			for _, v := range res {
				if string(v) == string(strs[i][j]) {
					m[string(strs[i][j])]++ //去重效果
				}
			}
			j++
		}

	}

	//去重,链接字符串
	var str string
	//m := make(map[string]int, 0)
	//for _, b := range result {
	//	m[b]++
	//}
	//for v, _ := range m {
	//	str+=v
	//}

	//str = strings.Join(result, "") //拼接字符串

	fmt.Println(m)
	for v, _ := range m {
		str += v
	}

	return str
}

func in(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

/**
有一个排序模块sort，它里面有一个sort.Strings()函数，可以对字符串数组进行排序。
同时，还有一个sort.SearchStrings()[1]函数，会用二分法在一个有序字符串数组中寻找特定字符串的索引。

结合两个函数，我们可以实现一个更高效的算法
*/
func in2(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}
