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
	fmt.Println(res)
}
