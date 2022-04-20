package main

import "errors"

type node struct {
	Val  interface{}
	Next *node
	Pre  *node
}

type queue struct {
	First *node
	Last  *node
	Len   int
}

func (qu *queue) enqueue(data interface{}) {
	nNode := &node{
		Val:  data,
		Pre:  qu.First,
		Next: nil,
	}
	if qu.First == nil {
		qu.First = nNode
	} else {
		qu.First.Next = nNode
		qu.First = nNode
	}
	if qu.Last == nil {
		qu.Last = nNode
	}
	qu.Len++
}

func (qu *queue) dequeue() interface{} {
	if qu.Len > 0 {
		nNode := qu.Last.Val
		if qu.Last.Next != nil {
			qu.Last.Next.Pre = nil
		}
		qu.Last = qu.Last.Next
		qu.Len--
		return nNode
	}
	return errors.New("error")
}

func (qu queue) isEmpty() bool {
	return qu.Len <= 0
}

func (qu queue) getLength() int {
	return qu.Len
}
