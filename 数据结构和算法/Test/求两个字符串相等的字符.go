package main

import "fmt"

func main() {
	res := make([]string, 0)
	str1 := "abcdf"
	str2 := "egab"

	for _, v1 := range str1 {
		for _, v2 := range str2 {
			if v1 == v2 {
				res = append(res, string(v1))
			}
		}
	}
	//fmt.Println(res)

	fmt.Println(isAnagram("abbc", "bca"))
}

func isAnagram(s string, t string) bool {
	var m [26]int
	for _, v := range s {
		m[v-'a']++
	}
	fmt.Println(m)

	for _, k := range t {
		m[k-'a']--
	}
	fmt.Println(m)

	for _, w := range m {
		if w != 0 {
			return false
		}
	}
	fmt.Println(m)

	return true
}
