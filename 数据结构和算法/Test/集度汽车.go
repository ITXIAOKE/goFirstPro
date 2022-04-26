package main

import (
	"fmt"
)

func FirstNotRepeatingChar(str string) int {
	if len(str) < 0 || len(str) > 10000 {
		return -1
	}

	myMap := make(map[string]int)

	for i := 0; i < len(str); i++ {
		if _, ok := myMap[string(str[i])]; !ok {
			myMap[string(str[i])] = 1
		} else {
			myMap[string(str[i])]++
		}
	}

	fmt.Println(myMap)

	for i := 0; i < len(str); i++ {
		for k, v := range myMap {
			if v == 1 && string(str[i]) == k {
				return i
			}
		}
	}

	return -1
}

func main() {
	var str = "google"
	fmt.Println(FirstNotRepeatingChar(str))
}
