package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := []byte{} //存放左边括号
	for i := 0; i < len(s); i++ {
		if isLeft(s[i]) {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == findLeft(s[i]) {
				stack = stack[:len(stack)-1] //依次销掉存放左边括号切片的最后一个
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func isLeft(c byte) bool {
	if c == '(' || c == '{' || c == '[' {
		return true
	}
	return false
}

func findLeft(c byte) byte {
	switch c {
	case ')':
		return '('
	case '}':
		return '{'
	case ']':
		return '['
	default:
		return ' '
	}
}
func main() {
	fmt.Println(isValid("({[]})"))
	fmt.Println(isValid("(){}[]"))
	fmt.Println(isValid("()"))
}
