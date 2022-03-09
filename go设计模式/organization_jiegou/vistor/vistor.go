package vistor

import "fmt"

type IVisitor interface {
	Visit()
}

type WeiBoVisitor struct {
}

func (w *WeiBoVisitor) Visit() {
	fmt.Println("WeiBoVisitor")
}

type IQIYIVisitor struct {
}

func (i *IQIYIVisitor) Visit() {
	fmt.Println("IQIYIVisitor")
}

type IElement interface {
	Accept(visitor IVisitor)
}

type Element struct {
}

func (e *Element) Accept(v IVisitor) {
	v.Visit()
}
