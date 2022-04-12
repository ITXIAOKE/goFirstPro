package main

import "fmt"

/**
LRU是一个老生常谈的问题，即最近最少使用，LRU是Least Recently Used的缩写，是一种操作系统中常用的页面置换算法，选择最近最久未使用的页面予以淘汰。
该算法赋予每个页面一个访问字段，用来记录一个页面自上次被访问以来所经历的时间 t，
当须淘汰一个页面时，选择现有页面中其 t 值最大的，即最近最少使用的页面予以淘汰

实现LRU基本的数据结构：Map+LinkedList

一般规则：
添加数据时，将新增数据节点放在头指针，尾结点部分大于最大长度时删除。
删除数据时，先按照Map的规则进行查找，再根据链表规则进行删除。
查找数据时，按照Map进行查找，没有则返回空，有则返回该数据的值并移动到头节点。
*/
//import "fmt"
//
//var head *Node
//var end *Node
//
//type Node struct {
//	Key   string
//	Value string
//	pre   *Node
//	next  *Node
//}
//
//func (n *Node) Init(key string, value string) {
//	n.Key = key
//	n.Value = value
//}
//
//type LRUCache struct {
//	Capacity int              //页面初始化大小
//	Size     int              //页面实际大小
//	Map      map[string]*Node //具体的cache
//}
//
//func GetLRUCache(capacity int) *LRUCache {
//	lruCache := LRUCache{Capacity: capacity}
//	lruCache.Map = make(map[string]*Node, capacity)
//	return &lruCache
//}
//
//func (l *LRUCache) get(key string) string {
//	if v, ok := l.Map[key]; ok {
//		l.refreshNode(v)
//		return v.Value
//	} else {
//		return "null"
//	}
//}
//
//func (l *LRUCache) put(key, value string) {
//	if v, ok := l.Map[key]; !ok {
//		if len(l.Map) >= l.Capacity {
//			oldKey := l.removeNode(head)
//			delete(l.Map, oldKey)
//		}
//		node := Node{Key: key, Value: value}
//		l.addNode(&node)
//		l.Map[key] = &node
//	} else {
//		v.Value = value
//		l.refreshNode(v)
//	}
//}
//
//func (l *LRUCache) refreshNode(node *Node) {
//	if node == end {
//		return
//	}
//	l.removeNode(node)
//	l.addNode(node)
//}
//
//func (l *LRUCache) removeNode(node *Node) string {
//	if node == end {
//		end = end.pre
//	} else if node == head {
//		head = head.next
//	} else {
//		node.pre.next = node.next
//		node.next.pre = node.pre
//	}
//	return node.Key
//}
//
//func (l *LRUCache) addNode(node *Node) {
//	if end != nil {
//		end.next = node
//		node.pre = end
//		node.next = nil
//	}
//	end = node
//	if head == nil {
//		head = node
//	}
//}
//
//func main() {
//	lruCache := GetLRUCache(3)
//	lruCache.put("001", "1")
//	lruCache.put("002", "2")
//	lruCache.put("003", "3")
//	lruCache.put("004", "4")
//	lruCache.put("005", "5")
//	lruCache.put("006", "6")
//
//	lruCache.get("005")
//
//	fmt.Println(lruCache.get("003"))
//	fmt.Println(lruCache.get("006"))
//
//	fmt.Println(lruCache.Map)
//}



type DLinkNode struct {
	key, val  int
	pre, next *DLinkNode
}

type LRUCache struct {
	size, capacity int
	cache          map[int]*DLinkNode
	head, tail     *DLinkNode
}

func NewDLinkNode(key, val int) *DLinkNode {
	return &DLinkNode{key: key, val: val}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		size:     0,
		cache:    map[int]*DLinkNode{},
		head:     NewDLinkNode(0, 0),
		tail:     NewDLinkNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func removeNode(node *DLinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) removeTail() *DLinkNode {
	node := this.tail.pre
	removeNode(node)
	return node
}

func (this *LRUCache) addHead(node *DLinkNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *DLinkNode) {
	removeNode(node)
	this.addHead(node)
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.moveToHead(node)
	} else {
		newNode := NewDLinkNode(key, value)
		this.addHead(newNode)
		this.cache[key] = newNode
		this.size++
		if this.size > this.capacity {
			tail := this.removeTail()
			this.size--
			delete(this.cache, tail.key)
		}
	}
}

func main() {
	lruCache := Constructor(3)
	lruCache.Put(1,11)
	lruCache.Put(2,22)
	lruCache.Put(3,33)
	lruCache.Put(4,44)
	lruCache.Put(5,55)


	lruCache.Get(3)

	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(4))

	fmt.Println(lruCache.cache)
}