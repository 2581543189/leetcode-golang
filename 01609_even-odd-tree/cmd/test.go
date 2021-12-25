package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Wrapper struct{
	node *TreeNode
	level int
}

func main(){
	root := TreeNode{1,&TreeNode{10,nil,nil},&TreeNode{4,nil,nil}}
	root.Left.Left = &TreeNode{3,nil,nil}
	root.Left.Left.Left = &TreeNode{12,nil,nil}
	root.Left.Left.Right = &TreeNode{8,nil,nil}

	root.Right.Left = &TreeNode{7,nil,nil}
	root.Right.Right = &TreeNode{9,nil,nil}
	root.Right.Left.Left = &TreeNode{6,nil,nil}
	root.Right.Right.Right = &TreeNode{2,nil,nil}
	fmt.Println(isEvenOddTree(&root))
}

func isEvenOddTree(root *TreeNode) bool {
	if root == nil || isEven(root.Val){
		return false
	}

	queue:=make([]Wrapper,0)
	queue = append(queue,Wrapper{root,0})

	lastLevel := -1;
	lastValue := 0;

	for len(queue) > 0{
		curNode := queue[0].node
		level := queue[0].level
		queue = queue[1:]


		if isEven(level){
			// 偶数层
			if isEven(curNode.Val){
				return false
			}
			if level == lastLevel && curNode.Val <= lastValue{
				return false
			}
			if curNode.Left != nil{
				queue = append(queue,Wrapper{curNode.Left,level + 1})
			}
			if curNode.Right != nil{
				queue = append(queue,Wrapper{curNode.Right,level + 1})
			}
		}else{
			// 奇数层
			if !isEven(curNode.Val){
				return false
			}
			if level == lastLevel && curNode.Val >= lastValue{
				return false
			}
			if curNode.Left != nil{
				queue = append(queue,Wrapper{curNode.Left,level + 1})
			}
			if curNode.Right != nil{
				queue = append(queue,Wrapper{curNode.Right,level + 1})
			}

		}
		lastLevel = level
		lastValue = curNode.Val
	}
	return true
}

func isEven(val int)bool {
	return val % 2 == 0
}