package main

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
