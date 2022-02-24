package memorandum

import (
	"fmt"
	"testing"
)
//备忘录模式
func TestNumber_ReinstateMemoran(t *testing.T) {
	n := NewNumber(7)
	fmt.Println(n.value)
	n.Double()
	fmt.Println(n.value)
	n.Double()
	fmt.Println(n.value)
	mem := n.CreateMemoran()
	n.Half()
	fmt.Println(n.value)
	n.ReinstateMemoran(mem)
	fmt.Println(n.value)
}
