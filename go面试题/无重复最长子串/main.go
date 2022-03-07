package main

import (
	"fmt"
)

func main() {

	fmt.Println(mySortLongRune("abcabcbb"))
	fmt.Println(mySortLongRune("bbbbbbb"))
	fmt.Println(mySortLongRune("a1234344"))
	fmt.Println(mySortLongRune("b"))
	fmt.Println(mySortLongRune(""))
	fmt.Println(mySortLongRune("我是小可"))
	fmt.Println(mySortLongRune("一二三二"))
	fmt.Println(mySortLongRune("黑话维护为化诶花费"))

}

//求最长不含有重复字符串的子串的长度
func mySortLongRune(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}
