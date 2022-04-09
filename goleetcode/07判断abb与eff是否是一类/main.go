package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return isIsomorphicHelper(s, t) && isIsomorphicHelper(t, s)
}

func isIsomorphicHelper(s string, t string) bool {
	dic := make(map[byte]byte)

	p := 0
	for p <= len(s)-1 {
		_, ok := dic[s[p]]
		if !ok {
			dic[s[p]] = t[p]
		} else {
			if dic[s[p]] != t[p] {
				return false
			}
		}
		p++
	}
	fmt.Println(dic)
	return true
}

func main() {
	fmt.Println(isIsomorphic("abb", "eff"))
}
