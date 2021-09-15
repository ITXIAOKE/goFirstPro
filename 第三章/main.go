package main

import "fmt"

func appendSlice(s []int) {
	s = append(s, 5)
	for i := range s {
		s[i]++
	}
	fmt.Println(s)
}

func main() {
	//s1 := []int{1, 2}
	//s2 := s1
	//fmt.Println(s2)
	//s2 = append(s2, 3)
	//
	//fmt.Println( s2)
	//appendSlice(s1)
	//appendSlice(s2)
	//fmt.Println(s1, s2)

	//a:=make([]int,2)
	//b :=[]int{1,2,3}
	//fmt.Printf("%T",a)
	//fmt.Println("-----")
	//fmt.Println(a,b)

	a:=10
	b:=20
	a,b=b,a
	fmt.Println(a,b)
}
