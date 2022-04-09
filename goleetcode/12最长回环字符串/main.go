package main

import "fmt"

//双指针解决回环问题
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isLetterOrNumber(s[left]) {
			left++
		}
		for left < right && !isLetterOrNumber(s[right]) {
			right--
		}
		if left <= len(s)-1 && right >= 0 {
			if toUpper(s[left]) != toUpper(s[right]) {
				return false
			}
		}
		left++
		right--
	}
	return true
}

func toUpper(char byte) byte {
	if isLower(char) {
		return byte(char - 32)
	}
	return char
}

func isLetterOrNumber(char byte) bool {
	//数字 48/0---57/9
	//大写 65A---90Z
	//小写 97a---122z
	return isUpper(char) || isLower(char) || isNumber(char)
}

func isNumber(char byte) bool {
	return 48 <= char && char <= 57
}

func isUpper(char byte) bool {
	return 65 <= char && char <= 90
}

func isLower(char byte) bool {
	return 97 <= char && char <= 122
}

func main() {
	fmt.Println(isPalindrome("abccbaa"))
	fmt.Println(isPalindrome("abccba"))
	fmt.Println(isPalindrome("a"))
	fmt.Println(isPalindrome("AA"))
	fmt.Println(isPalindrome("ABCCBA"))
}
