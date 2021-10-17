package main

import "fmt"

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func main(){
	node7 := &ListNode{8,nil}
	node6 := &ListNode{7,node7}
	node5 := &ListNode{6,node6}
	node4 := &ListNode{5,node5}
	node3 := &ListNode{4,node4}
	node2 := &ListNode{3,node3}
	node1 := &ListNode{2,node2}
	head := &ListNode{1,node1}

	tmp := head

	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println("=============")
	sorted := oddEvenList( head )
	tmp=sorted
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next

	}
}