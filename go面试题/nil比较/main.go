package main

import "fmt"

/**
runtime2源码接口如下：

type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type eface struct {
	_type *_type
	data  unsafe.Pointer
}
*/
func main() {
	var x *int = nil
	var y interface{} = x //空接口    eface源码发现   type不同

	fmt.Println(x == y)
	fmt.Println(x == nil)
	fmt.Println(y == nil) //既比较类型，又比较值

	var z interface{} = nil //空接口    eface源码发现   type不同
	var m interface{} = nil //空接口    eface源码发现   type不同
	fmt.Println(z == nil)
	fmt.Println(z == y) //既比较类型，又比较值
	fmt.Println(m == nil)
	fmt.Println(m == z) //既比较类型，又比较值
}
