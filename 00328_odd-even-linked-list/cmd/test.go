package main

type ListNode struct {
	Val int
	Next *ListNode
}


func oddEvenList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return  head
	}
	current := head
	evenList := &ListNode{}
	evenCurrnt := evenList
	for current.Next != nil {
		evenCurrnt.Next = current.Next
		evenCurrnt = evenCurrnt.Next
		current.Next = current.Next.Next
		if current.Next != nil {
			current = current.Next
		}
		//evenCurrnt.Next = nil
	}
	current.Next = evenList.Next
	return head
}


//func main(){
//	node7 := &ListNode{8,nil}
//	node6 := &ListNode{7,node7}
//	node5 := &ListNode{6,node6}
//	node4 := &ListNode{5,node5}
//	node3 := &ListNode{4,node4}
//	node2 := &ListNode{3,node3}
//	node1 := &ListNode{2,node2}
//	head := &ListNode{1,node1}
//
//	tmp := head
//
//	for tmp != nil {
//		fmt.Println(tmp.Val)
//		tmp = tmp.Next
//	}
//	fmt.Println("=============")
//	sorted := oddEvenList1( head )
//	tmp=sorted
//	for tmp != nil {
//		fmt.Println(tmp.Val)
//		tmp = tmp.Next
//
//	}
//}