package main

import "fmt"

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct {
	list [] *NestedInteger
	num int
	isInteger bool
}

func(n NestedInteger) IsInteger()bool{
	return n.isInteger
}

func(n NestedInteger) GetList()[] *NestedInteger{
	return n.list
}

func(n NestedInteger) GetInteger()int{
	return n.num
}


type NestedIterator struct {
	// 将列表视作一个队列，栈中直接存储该队列
	stack []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	//反转
	if len(nestedList) > 0{
		copy:= nestedList
		left:=0
		right:=len(copy) - 1
		for left < right {
			copy[left],copy[right] = copy[right],copy[left]
			left++
			right--
		}
	}
	return &NestedIterator{nestedList}
}

func (it *NestedIterator) Next() int {
	v:=it.stack[len(it.stack) - 1].GetInteger()
	it.stack = it.stack[:len(it.stack) - 1]
	return v
}


func (it *NestedIterator) HasNext() bool {
	if len(it.stack) == 0{
		return false
	}
	for len(it.stack) > 0 && !it.stack[len(it.stack) - 1].IsInteger() {
		tmpList:= it.stack[len(it.stack) - 1].GetList()
		it.stack = it.stack[:len(it.stack) - 1]
		if len(tmpList) > 0{
			for i:=len(tmpList) - 1;i >=0;i-- {
				it.stack = append(it.stack,tmpList[i])
			}
		}
	}
	return len(it.stack) > 0
}

func main(){
	//nestedList = [[1,1],2,[1,1]]
	//ni:= NestedInteger {[] *NestedInteger{
	//	&NestedInteger{nil,1,true},
	//	&NestedInteger{[] *NestedInteger{
	//		&NestedInteger{nil,4,true},
	//		&NestedInteger{[] *NestedInteger{
	//			&NestedInteger{nil,6,true},
	//		},0,false},
	//	},0,false},
	//},0,false}
	ni:= NestedInteger {[] *NestedInteger{
		&NestedInteger{[] *NestedInteger{},0,false},
	},0,false}


	ni2 :=Constructor(ni.list)
	for ni2.HasNext(){
		fmt.Println(ni2.Next())
	}

}