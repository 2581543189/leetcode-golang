package main

import (
	"container/heap"
	"fmt"
	"sort"
)


type hp struct {sort.IntSlice}
func (this *hp) Push(v interface{}){
	this.IntSlice = append(this.IntSlice,v.(int))
}

func (this *hp) Pop()interface{}{
	v := this.IntSlice[len(this.IntSlice) - 1]
	this.IntSlice = this.IntSlice[:len(this.IntSlice) - 1]
	return v
}
func main()  {
	//创建一个heap
	h := &hp{}
	heap.Init(h)

	//向heap中插入元素
	h.Push(5)
	h.Push(2)
	h.Push(1)
	h.Push(8)
	h.Push(4)
	h.Push(6)
	h.Push(2)

	//输出heap中的元素，相当于一个数组，原始数组
	fmt.Println(h)

	//这里必须要reheapify，建立好堆了
	heap.Init(h)

	//小顶堆对应的元素在数组中的位置
	fmt.Println(h)

	//移除下标为5的元素，下标从0开始
	//h.Remove(0)

	//按照堆的形式输出
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()
}
