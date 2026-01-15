package main

import "fmt"

type listNode struct {
	value int
	next  *listNode
}

func reversek(head *listNode, k int) *listNode {
	dummy := &listNode{0, head}
	pre := dummy

	for {
		end := pre
		for i := 0; i < k; i++ {
			end = end.next
			if end == nil {
				return dummy.next
			}
		}

		start := pre.next
		nextnode := end.next

		end.next = nil
		pre.next = reverse(start)
		start.next = nextnode
		pre = start
	}
}

func reverse(head *listNode) *listNode {
	var pre *listNode
	cur := head
	for cur != nil {
		next := cur.next
		cur.next = pre
		pre = cur
		cur = next
	}

	return pre

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

	head := createlist([]int{1, 2, 3, 4, 5})
	k := 2

	head = reversek(head, k)
	output(head)

}
