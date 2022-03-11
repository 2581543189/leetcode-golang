package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next, slow = nil, slow.Next
	return merge(sortList(head), sortList(slow))
}

func merge(n1, n2 *ListNode) *ListNode {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	p1 := n1
	p2 := n2
	queue := &ListNode{}
	curr := queue
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			curr.Next = p1
			p1 = p1.Next
		} else {
			curr.Next = p2
			p2 = p2.Next
		}
		curr = curr.Next
	}
	if p1 != nil {
		curr.Next = p1
	}
	if p2 != nil {
		curr.Next = p2
	}
	return queue.Next
}

func main(){
	node3 := &ListNode{3,nil}
	node1 := &ListNode{1,node3}
	node2 := &ListNode{2,node1}
	head := &ListNode{4,node2}
	tmp:=head
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println("=============")
	sorted := sortList(head)
	tmp=sorted
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next

	}
}