package main

import "fmt"

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	chars := []byte(s)
	first := -1 //代表是空格

	for i := len(chars) - 1; i >= 0; i-- {
		if first == -1 { //从后面往前找，首先发现是空格
			if chars[i] == ' ' { //继续找还是空格
				continue
			} else {
				first = i //找到了非空格的元素（单词的最后一位字母）对应的角标，赋值给当前的first
			}
		} else {
			if chars[i] == ' ' { //从非空格处往前找到空格的地方
				return first - i //找到单词后，最后一位的角标---再次发现空格的角标值
			}
		}
	}

	return first + 1 //没有单词，返回0

}

func main() {
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord("hello world"))
	fmt.Println(lengthOfLastWord("a"))
	fmt.Println(lengthOfLastWord(""))
	fmt.Println(lengthOfLastWord("  "))
	fmt.Println(lengthOfLastWord("aaa dfdfd h-g="))
	fmt.Println(lengthOfLastWord("aaa dfdfd h-g=    "))
}
