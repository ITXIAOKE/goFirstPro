package main

import "fmt"

/**
iface 和 eface 都是 Go 中描述接口的底层结构体，区别在于
iface 描述的接口包含方法，
而 eface 则是不包含任何方法的空接口：interface{}。


type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type itab struct {
	inter  *interfacetype
	_type  *_type
	link   *itab
	hash   uint32 // copy of _type.hash. Used for type switches.
	bad    bool   // type does not implement interface
	inhash bool   // has this itab been added to hash?
	unused [2]byte
	fun    [1]uintptr // variable sized
}

type interfacetype struct {
	typ     _type
	pkgpath name
	mhdr    []imethod
}

iface 内部维护两个指针，
tab 指向一个 itab 实体， 它表示接口的类型以及赋给这个接口的实体类型。
data 则指向接口具体的值，一般而言是一个指向堆内存的指针。

再来仔细看一下 itab 结构体：
_type 字段描述了实体的类型，包括内存对齐方式，大小等；
inter 字段则描述了接口的类型。
fun 字段放置和接口方法对应的具体数据类型的方法地址，实现接口调用方法的动态分派，一般在每次给接口赋值发生转换时会更新此表，或者直接拿缓存的 itab。



type eface struct {
    _type *_type
    data  unsafe.Pointer
}

相比 iface，eface 就比较简单了。
只维护了一个 _type 字段，表示空接口所承载的具体的实体类型。
data 描述了具体的值。


*/

type INodeImpl interface {
	SetValue(value int)
	Traverse() //遍历方法
}

type TreeNode struct {
	Value     int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		Value: value,
	}
}

func (tree *TreeNode) SetValue(value int) {
	tree.Value = value
}

func (tree *TreeNode) Traverse() {
	if tree == nil {
		return
	}

	fmt.Println(tree.Value) //前序遍历  根左右
	tree.LeftNode.Traverse()
	//fmt.Println(tree.Value) //中  左根右
	tree.RightNode.Traverse()
	//fmt.Println(tree.Value) //后  左右根

}

func main() {

	root := new(TreeNode)
	root.SetValue(3)

	root.LeftNode = NewTreeNode(2)
	root.RightNode = NewTreeNode(4)

	root.Traverse()

}
