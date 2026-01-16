package main

import (
	"fmt"
)

type listNode struct {
	value int
	next  *listNode
}

func mergetwolist(l1 *listNode, l2 *listNode) *listNode {
	dummy := &listNode{value: -1}
	cur := dummy
	p1, p2 := l1, l2

	for p1 != nil && p2 != nil {
		if p1.value < p2.value {
			cur.next = p1
			p1 = p1.next
		} else {
			cur.next = p2
			p2 = p2.next
		}
		cur = cur.next
	}

	if p1 != nil {
		cur.next = p1
	} else {
		cur.next = p2
	}

	return dummy.next

}

func createlist(nums []int) *listNode {
	dummy := &listNode{}
	cur := dummy
	for _, num := range nums {
		cur.next = &listNode{value: num}
		cur = cur.next
	}

	return dummy.next
}

func output(head *listNode) {
	for head != nil {
		fmt.Println("%d->", head.value)
		head = head.next
	}

}

func main() {
	l1 := createlist([]int{1, 2, 4})
	l2 := createlist([]int{1, 3, 4})
	result := mergetwolist(l1, l2)

	output(result)

}
