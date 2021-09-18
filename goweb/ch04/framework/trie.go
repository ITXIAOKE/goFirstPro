package framework

type Tree struct { //代表树结构
	root *node //根节点
}

//代表节点
type node struct {
	isLast  bool              //该节点是否能成为一个独立的uri，是否自身就是一个终极节点
	segment string            //uri中的字符串
	handler ControllerHandler //控制器
	childs  []*node           //子节点
}

func newNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		childs:  []*node{},
	}
}

func NewTree() *Tree {
	root:=newNode()
	return &Tree{root}
}
