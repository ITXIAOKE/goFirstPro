package main

import (
	"fmt"
	"sync"
)

type SliceNum []int

func NewSlice() SliceNum {
	return make(SliceNum, 0)

}

func (s *SliceNum) Add(elem int) *SliceNum {
	*s = append(*s, elem)
	fmt.Println("add", elem)
	fmt.Println("add SliceNum end", s)
	fmt.Println("add SliceNum end", *s) //取值
	fmt.Println("add SliceNum end", &s) //取地址
	return s
}

func main() {
	var once sync.Once
	s := NewSlice()
	// 看源代码理解once的行为，once只执行一次
	once.Do(func() {
		s.Add(16)
	})
	once.Do(func() {
		s.Add(17)
	})
	once.Do(func() {
		s.Add(18)
	})
}
