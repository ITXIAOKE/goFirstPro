package framework

import (
	"errors"
	"strings"
)

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
	root := newNode()
	return &Tree{root}
}

//判断一个segment是否是通用segment，即以：开头
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

//过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	if isWildSegment(segment) { //如果segment是通配符，则所有下一层子节点都满足需求
		return n.childs
	}

	nodes := make([]*node, 0, len(n.childs))
	//过滤所有的下一层子节点
	for _, cnode := range n.childs {
		if isWildSegment(cnode.segment) {
			//如果下一层子节点有通配符，则满足需求
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment {
			//如果下一层子节点没有通配符，但是文本完全匹配，则满足需求
			nodes = append(nodes, cnode)
		}
	}
	return nodes
}

//判断路由是否已经在节点的所有子节点树中存在
func (n *node) matchNode(uri string) *node {
	//使用分隔符将uri切割为两个部分
	segments := strings.SplitN(uri, "/", 2)
	//第一个部分用于匹配下一层子节点
	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	//匹配符合的下一层子节点
	cnodes := n.filterChildNodes(segment)
	//如果当前子节点没有一个符合，那么说明这个uri一定是之前存在，直接返回nil
	if cnodes == nil || len(cnodes) == 0 {
		return nil
	}
	if len(segments) == 1 {
		for _, tn := range cnodes {
			if tn.isLast {
				return tn
			}
		}
		return nil
	}
	for _, tn := range cnodes {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil
}

// 增加路由节点, 路由节点有先后顺序
/*
/book/list
/book/:id (冲突)
/book/:id/name
/book/:student/age
/:user/name(冲突)
/:user/name/:age
*/

func (tree *Tree) AddRouter(uri string, handler ControllerHandler) error {
	n:=tree.root
	if n.matchNode(uri)!=nil{
		return errors.New("route exist:"+uri)
	}
	segments:=strings.Split(uri,"/")
	for index,segment :=range segments{
		if !isWildSegment(segment){
			segment=strings.ToUpper(segment)
		}
		isLast:=index==len(segments)-1

		var objNode *node
		childNodes:=n.filterChildNodes(segment)
		if len(childNodes)>0{
			for _,cnode:=range childNodes{
				if cnode.segment==segment{
					objNode=cnode
					break
				}
			}
		}
		if objNode==nil{
			cnode:=newNode()
			cnode.segment=segment
			if isLast{
				cnode.isLast=true
				cnode.handler=handler
			}
			n.childs=append(n.childs,cnode)
			objNode=cnode
		}
		n=objNode
	}
	return nil
}

//匹配uri

func (tree *Tree) FindHandler(uri string) ControllerHandler {
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handler
}
