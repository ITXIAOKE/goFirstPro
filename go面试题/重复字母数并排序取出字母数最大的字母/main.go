package main

import (
	"fmt"
)

func main() {
	input := []string{"abcd", "abcbb", "ab"}
	res := CountLetters(input)

	var myslice []int
	for _, value := range res {
		myslice = append(myslice, value)
	}

	fmt.Println(myslice)
	tmp := 0 //定义临时变量
	for j := 0; j < len(myslice)-1; j++ {
		//多次循环遍历的时候i是越来越小，j是增大的 用len(arry)-i-j实现遍历
		for i := 0; i < len(myslice)-1-j; i++ {
			if myslice[i] > myslice[i+1] {
				tmp = myslice[i]
				myslice[i] = myslice[i+1]
				myslice[i+1] = tmp
			}
		}
	}
	fmt.Println(myslice)

	//拿到最多字母数对应的字母
	for key, value := range res {
		lena := myslice[len(myslice)-1]
		if value == lena {
			fmt.Println(string(byte(key))) //将98转化成b
		}
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
