package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	if S == "" {
		return 0
	}
	dic := make(map[string]int)
	for _, s := range S {
		dic[string(s)]++
	}

	res := 0
	for _, j := range J {
		if v, ok := dic[string(j)]; ok {
			res += v
		}
	}
	return res
}

func main() {
	//j=aA s=aAAbbb
	//输出3
	fmt.Println(numJewelsInStones("aAb", "aAAbb"))

}
