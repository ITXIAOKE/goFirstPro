package main

import "fmt"

func isSubString(s string, t string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) > len(t) {
		return false
	}
	first := 0
	second := 0
	for first < len(s) && second < len(t) {
		for second < len(t) && s[first] != t[second] {
			second++
		}
		if first < len(s) && second < len(t) && s[first] == t[second] {
			first++
			second++
		}
	}
	if first < len(s) && second == len(t) { //s没走完，t已经走完，也就是s中的一个字母没有在t中
		return false
	}
	return first == len(s) //最后判断这个子字符串是否走完，走完代表true，没走完是false
}

func main() {
	bol := isSubString("abc", "xiaoabcke")
	fmt.Println(bol)
}
