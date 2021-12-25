package main

import (
	"fmt"
	"math"
)

type TreeNode2 struct {
	Val int
	Left *TreeNode2
	Right *TreeNode2
}


func main(){
	root := TreeNode2{1,&TreeNode2{10,nil,nil},&TreeNode2{4,nil,nil}}
	root.Left.Left = &TreeNode2{3,nil,nil}
	root.Left.Left.Left = &TreeNode2{12,nil,nil}
	root.Left.Left.Right = &TreeNode2{8,nil,nil}

	root.Right.Left = &TreeNode2{7,nil,nil}
	root.Right.Right = &TreeNode2{9,nil,nil}
	root.Right.Left.Left = &TreeNode2{6,nil,nil}
	root.Right.Right.Right = &TreeNode2{2,nil,nil}
	fmt.Println(isEvenOddTree2(&root))
}

func isEvenOddTree2(root *TreeNode2) bool {
	q := []*TreeNode2{root}
	for level := 0; len(q) > 0; level++ {
		prev := 0
		if level%2 == 1 {
			prev = math.MaxInt32
		}
		size := len(q)
		for _, node := range q {
			val := node.Val
			if val%2 == level%2 || level%2 == 0 && val <= prev || level%2 == 1 && val >= prev {
				return false
			}
			prev = val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[size:]
	}
	return true
}


