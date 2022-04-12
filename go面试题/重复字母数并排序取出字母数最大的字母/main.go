package main

import (
	"fmt"
)

func main() {
	input := []string{"abcd", "abcbb", "abcccccc"}
	res := CountLetters(input)

	var myslice []int
	for _, value := range res { //取出所有的字母出现的次数value，存放到一个列表中
		myslice = append(myslice, value)
	}
	fmt.Println(myslice) //[3 5 8 1]


	for j := 0; j < len(myslice)-1; j++ {
		//多次循环遍历的时候i是越来越小，j是增大的 用len(arry)-i-j实现遍历
		for i := 0; i < len(myslice)-1-j; i++ {
			if myslice[i] > myslice[i+1] {
				myslice[i], myslice[i+1] = myslice[i+1], myslice[i]
			}
		}
	}
	fmt.Println(myslice) //冒泡排序后的结果，出现次数最多的是最后一个数[1 3 5 8]

	//拿到最多字母数对应的字母
	for key, value := range res {
		lena := myslice[len(myslice)-1]
		if value == lena {
			fmt.Println(string(byte(key))) //将98转化成b
		}
	}

}

type LetterFreq map[byte]int

func CountLetters(strs []string) LetterFreq { //遍历所有字符串，拿到每个字母对应的数字组成的map
	m := make(map[byte]int, 0)
	for _, str := range strs {
		for _, r := range str {
			m[byte(r)]++
		}
	}
	fmt.Println("map------", m)
	return m
}
