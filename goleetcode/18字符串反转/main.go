package main

import "fmt"

func main() {
	str := "abcdef" //纯字母字符串，没有中文
	var code string
	for i := len(str); i > 0; i-- {
		code = code + str[i-1:i]
	}
	fmt.Printf(code)
	fmt.Println()

	//1. 利用双指针，一个指针从第一个元素开始，另一个指针从最后一个元素开始，相互交换.
	//2. 循环移动这两个指针,到他们两个重合为止.
	fmt.Printf(string(reverseString([]byte(str))))
	fmt.Println()

	str2 := "xiao小可爱ke"
	fmt.Printf(reverseString2([]rune(str2))) //有中文的字符串, 因为byte是 uint8 遇到中文会有截断。 换成 rune 完美解决

	fmt.Println()
	fmt.Printf(reverseString3([]rune(str2))) //有中文的字符串

}

func reverseString(s []byte) []byte {
	var i, j = 0, len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}

func reverseString2(myrune []rune) string {
	var temp []rune
	for i, _ := range myrune {
		temp = append(temp, myrune[len(myrune)-i-1])
	}
	return string(temp)
}

func reverseString3(s []rune) string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
