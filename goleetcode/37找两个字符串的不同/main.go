package main

import "fmt"

func findTheDifference(s string, t string) byte {
	var res byte
	if len(s) > len(t) {
		return res
	}
	T := []byte(t)
	S := []byte(s)

	dic := make(map[byte]int)
	for i := 0; i < len(S); i++ {
		dic[S[i]]++
	}

	for i := 0; i < len(T); i++ {
		val, ok := dic[T[i]]
		if !ok {
			res = T[i]
			break
		} else {
			if val == 0 {//这个是为了都是一样字母的，比如aaa和aaaa
				res = T[i]
				break
			}
			dic[T[i]]--
		}
	}
	return res
}

func main() {
	fmt.Println(string(findTheDifference("aaa","aaab")))
}
