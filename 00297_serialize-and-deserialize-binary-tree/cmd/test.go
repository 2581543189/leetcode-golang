package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil{
		return ""
	}
	ret := [] string{}
	var dfs func (node *TreeNode)
	dfs = func ( node *TreeNode) {
		if node == nil {
			ret = append(ret,"null")
			return
		}else{
			ret = append(ret,strconv.Itoa(node.Val))
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(ret,",")

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0{
		return nil
	}
	array := strings.Split(data,",")
	root := TreeNode{}
	var dfs func ( node *TreeNode)*TreeNode
	dfs= func(node *TreeNode)*TreeNode{
		if len(array) == 0 {
			return nil
		}
		str := array[0]
		array = array[1:]
		if str == "null" {
			return nil
		}
		node.Val, _ = strconv.Atoi(str)
		node.Left = dfs(&TreeNode{})
		node.Right = dfs(&TreeNode{})
		return node
	}
	return dfs(&root)
}

func main()  {
	codec:=Constructor()
	root := TreeNode{0,&TreeNode{1,nil,nil},&TreeNode{2,nil,nil}}
	root.Left.Left = &TreeNode{3,nil,nil}
	root.Left.Right = &TreeNode{4,nil,nil}
	root.Right.Left = &TreeNode{5,nil,nil}
	root.Right.Right = &TreeNode{6,nil,nil}
	s:= codec.serialize(&root)
	fmt.Println(s)
	o:= codec.deserialize(s)
	var dfs func (node *TreeNode)
	dfs = func ( node *TreeNode) {
		if node == nil {
			fmt.Print("null")
			fmt.Print(",")
			return
		}else{
			fmt.Print(node.Val)
			fmt.Print(",")
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(o)

}
