package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ss, tt := []byte(s), []byte(t)
	m := make(map[byte]int)
	for i := 0; i < len(ss); i++ {
		m[ss[i]]++
	}

	for i := 0; i < len(tt); i++ {
		m[tt[i]]--
		if m[tt[i]] < 0 {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println(isAnagram("abc", "cab"))
	fmt.Println(isAnagram("aba", "caab"))
	fmt.Println(isAnagram("ab1", "1ab"))

}
