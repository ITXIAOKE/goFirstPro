package main

import (
	"fmt"
	"log"
)

func main() {
	input := []string{"abcd", "abc", "ab"}
	res := CountLetters(input)
	if res['a'] != 4 {
		log.Fatalf("get %d\n", res['a'])
	} else {
		fmt.Println("success")
	}
}

type LetterFreq map[rune]int

func CountLetters(strs []string) LetterFreq {
	m := make(map[rune]int, 0)
	for _, str := range strs {
		for _, r := range str {
			m[r]++
		}
	}
	fmt.Println(m)
	return m
}
