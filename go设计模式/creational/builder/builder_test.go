package builder

import (
	"fmt"
	"testing"
)
//建造者模式
func TestConcreteBuilder_GetResult(t *testing.T) {
	builder := NewConcreteBuilder()
	director := NewDirector(&builder)
	director.Construct()
	product := builder.GetResult()
	fmt.Println(product.Built)
}
