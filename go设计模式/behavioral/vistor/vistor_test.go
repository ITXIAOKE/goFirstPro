package vistor

import "testing"

//访问者模式
func TestElement_Accept(t *testing.T) {
	e := new(Element)
	e.Accept(new(WeiBoVisitor))
	e.Accept(new(IQIYIVisitor))
}
