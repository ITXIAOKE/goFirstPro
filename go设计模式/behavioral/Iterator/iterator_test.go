package Iterator

import (
	"fmt"
	"testing"
)

//迭代器设计模式（类比for循环）
func TestIterator(t *testing.T) {
	//只能是这样的[]interface{}{1,2,3,4}，不能是[]interface{1,2,3,4}
	array := []interface{}{1, 3, 9, 2, 8,7777,"abc","sdfasfasdf"}
	a := 0
	iterator := ArrayIterator{array: array, index: &a}
	for it := iterator; iterator.HasNext(); iterator.Next() {
		index, value := it.Index(), it.Value()
		if value != array[*index] {
			fmt.Println("error...")
		}
		fmt.Println(*index, value)
	}
}
