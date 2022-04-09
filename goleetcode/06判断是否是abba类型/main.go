package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	patterns := strings.Split(pattern, "") //a b b a  一个中间没有空格的字符串分割成切片的时候，分割符是""
	fmt.Println(strs)
	fmt.Println(patterns)

	if len(pattern) != len(strs) { //按分割符分成的切片长度
		return false
	}
	return checkPattern(patterns, strs) && checkPattern(strs, patterns)
}

func checkPattern(a []string, b []string) bool {
	dic := make(map[string]string)
	p := 0
	for p < len(a) {
		val, ok := dic[a[p]]
		if !ok {
			dic[a[p]] = b[p]
		} else {
			if val != b[p] {
				return false
			}
		}
		p++
	}
	fmt.Println(dic)
	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}
