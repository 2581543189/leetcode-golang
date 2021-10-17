package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	var doReverse func(node *ListNode) (*ListNode, *ListNode)
	doReverse =  func(node *ListNode) (*ListNode, *ListNode){
		if node == nil || node.Next == nil {
			return node,node
		}
		h,t := doReverse(node.Next)
		node.Next = nil
		t.Next = node
		return h,node
	}
	ans,_ := doReverse(head)
	return ans
}

func main(){
	node3 := &ListNode{4,nil}
	node1 := &ListNode{3,node3}
	node2 := &ListNode{2,node1}
	head := &ListNode{1,node2}
	tmp:=head
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println("=============")
	sorted := reverseList(head)
	tmp=sorted
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next

	}
}