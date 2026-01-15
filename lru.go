package main

import "fmt"

type node struct {
	key, value int
	prev, next *node
}

type lrucache struct {
	cap        int
	cache      map[int]*node
	head, tail *node
}

func cons(cap int) *lrucache {
	l := &lrucache{
		cap:   cap,
		cache: make(map[int]*node),
		head:  &node{},
		tail:  &node{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l

}

func (this *lrucache) get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.move_to_head(node)
		return node.value
	}
	return -1
}

func (this *lrucache) put(key int, value int) {
	if node1, ok := this.cache[key]; ok {
		node1.value = value
		this.move_to_head(node1)
	} else {
		newnode := &node{
			key:   key,
			value: value,
		}
		this.cache[key] = newnode
		this.add_to_head(newnode)
		if len(this.cache) > this.cap {
			remove := this.remove_tail()
			delete(this.cache, remove.key)
		}
	}
}

func main() {
	fmt.Println("开始测试 LRU 缓存 (Go版):")

	// 初始化 capacity = 2
	obj := cons(2)
	fmt.Println("初始化 capacity = 2")

	// Put(1, 1)
	obj.put(1, 1)
	fmt.Println("Put(1, 1)")

	// Put(2, 2)
	obj.put(2, 2)
	fmt.Println("Put(2, 2)")

	// Get(1) -> 应该返回 1
	fmt.Printf("Get(1) -> %d\n", obj.get(1))

	// Put(3, 3) -> 导致 key 2 被逐出
	obj.put(3, 3)
	fmt.Println("Put(3, 3) -> 逐出 key 2")

	// Get(2) -> 应该返回 -1
	fmt.Printf("Get(2) -> %d\n", obj.get(2))

	// Put(4, 4) -> 导致 key 1 被逐出
	obj.put(4, 4)
	fmt.Println("Put(4, 4) -> 逐出 key 1")

	// Get(1) -> 应该返回 -1
	fmt.Printf("Get(1) -> %d\n", obj.get(1))

	// Get(3) -> 应该返回 3
	fmt.Printf("Get(3) -> %d\n", obj.get(3))

	// Get(4) -> 应该返回 4
	fmt.Printf("Get(4) -> %d\n", obj.get(4))

}

func (this *lrucache) move_to_head(node *node) {
	this.remove_node(node)
	this.add_to_head(node)
}

func (this *lrucache) remove_node(node *node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *lrucache) add_to_head(node *node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node

}

func (this *lrucache) remove_tail() *node {
	node := this.tail.prev
	this.remove_node(node)
	return node
}
