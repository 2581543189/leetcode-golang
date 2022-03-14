package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ListNode struct {
	Val int
	Next *ListNode
}

type Solution struct {
	head *ListNode
	count int
}


func Constructor(head *ListNode) Solution {
	count:=0
	node:=head
	for node != nil{
		node = node.Next
		count++
	}
	return Solution{head,count}
}


func (this *Solution) GetRandom() int {
	n:= this.count
	node:=this.head
	ans:= node.Val
	node = node.Next
	rand.Seed(time.Now().UnixNano())
	for i:=1;i<n;i++{
		if rand.Intn(i + 1) == i {
			ans = node.Val
		}
		node = node.Next
	}
	return ans
}

func main(){
	third := &ListNode{3,nil}
	second := &ListNode{2,third}
	head:=&ListNode{1,second}

	s := Constructor(head)

	c1,c2,c3:=0,0,0
	for i:=0;i<1000;i++{
		r:=s.GetRandom()
		switch r {
			case 1:c1++
			case 2:c2++
			case 3:c3++
			default:
		}
	}

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)


}


