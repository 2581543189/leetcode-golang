package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode2 struct {
	Val int
	Next *ListNode2
}


func sortList1(head *ListNode2) *ListNode2 {
	if(head==nil || head.Next == nil){
		return head
	}
	var merge func(left,right *ListNode2)*ListNode2
	merge = func(left,right *ListNode2)*ListNode2{
		if left == nil {
			return right
		}
		if right == nil {
			return left
		}
		var ret *ListNode2 = nil
		var retTail *ListNode2 = nil
		for left!= nil && right != nil {
			var target *ListNode2 = nil
			if left.Val <= right.Val {
				target = left
				left = left.Next
			}else{
				target = right
				right = right.Next
			}
			if ret == nil{
				ret = target
				retTail = target
			}else{
				retTail.Next = target
				retTail = target
			}
		}
		if left != nil{
			retTail.Next = left

		}else if right != nil {
			retTail.Next = right
		}
		return ret
	}

	fast:= head.Next
	slow := head
	for fast !=nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}

	list1:= sortList1(slow.Next)
	slow.Next = nil
	list2:= sortList1(head)
	return merge(list1,list2)
}

func main(){
	node3 := &ListNode2{3,nil}
	node1 := &ListNode2{1,node3}
	node2 := &ListNode2{2,node1}
	head := &ListNode2{4,node2}
	tmp:=head
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println("=============")
	sorted := sortList1(head)
	tmp=sorted
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next

	}
}