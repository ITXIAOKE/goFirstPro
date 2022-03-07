package main

import "fmt"

func main() {
	fmt.Println(reverse("可小崔-=CBA@$321"))
}

func reverse(str string) string {
	s := []rune(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)

}
