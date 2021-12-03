package abstract_factory

import "fmt"

type Lunch interface {
	Cook()
}

type rise struct {
}

func (c *rise) Cook() {
	fmt.Println("it is a rise.")
}

type Tomato struct {
}

func (c *Tomato) Cook() {
	fmt.Println("it is Tomato.")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type simpleLunchFactory struct {
}

//当一个结构体具备接口的所有的方法的时候，它就实现了这个接口

func NewSimpleShapeFactory() LunchFactory {
	return &simpleLunchFactory{}
}

func (s *simpleLunchFactory) CreateFood() Lunch {
	return &rise{}
}

func (s *simpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}
