package main

/**
将用链表实现栈，并使用 interface{} 来代替 int实现多种类型的兼容

由于任何的变量都实现了空接口，所以我们可以通过传递空接口来实现在栈中压入不同元素的目的
*/
type node struct {
	Val  interface{}
	Prev *node
}

type stackLink struct {
	Top    *node
	Length int
}

func (sl *stackLink) push(value interface{}) {
	newNode := &node{
		Val:  value,
		Prev: sl.Top}
	sl.Top = newNode
	sl.Length++
}

func (sl *stackLink) pop() interface{} {
	topNodeVal := sl.Top.Val
	sl.Top = sl.Top.Prev
	sl.Length--
	return topNodeVal
}

func (sl stackLink) length() int {
	return sl.Length
}

func (sl stackLink) isEmpty() bool {
	return sl.Length == 0
}

func (sl stackLink) peer() interface{} {
	return sl.Top.Val
}
