package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := [...]student{
		{name: "a", age: 10},
		{name: "b", age: 20},
		{name: "c", age: 30},
	}

	for k, stu := range stus {
		/**
		b &{c 30}
		c &{c 30}
		a &{c 30}
		*/
		//m[stu.name] =&stu//注意，这个结果最后三个值都是C：30


		m[stu.name] = &stus[k] //这个才是对的
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
