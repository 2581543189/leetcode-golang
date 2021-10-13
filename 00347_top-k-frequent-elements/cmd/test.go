package main

import (
	"container/heap"
	"fmt"
)

func main()  {
	nums :=[]int{-1,-1}
	fmt.Println(topKFrequent1(nums,1))
}

var countMap = make(map[int]int)

func topKFrequent1(nums []int, k int) []int {
	countMap = make(map[int]int)
	for _,value := range nums {
		countMap[value]++
	}
	//使用堆
	h := &Heap{}
	for key,_ := range countMap {
		heap.Push(h,key)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ret:=make([]int,k)

	for i :=k-1;i>=0;i--{
		ret[i] = heap.Pop(h).(int)
	}

	return ret

}

type Heap struct {

array [] int
}

func (h *Heap) Push(x interface{}){
	h.array = append(h.array,x.(int))
}

func (h *Heap) Pop() interface{} {
	val := h.array[len(h.array) - 1]
	h.array = h.array[:len(h.array) - 1]
	return val
}

func (h *Heap) Swap(i,j int){
	h.array[i],h.array[j] = h.array[j],h.array[i]
}

func (h Heap) Less(i,j int)bool {
	return countMap[h.array[i]] < countMap[h.array[j]]
}

func (h Heap) Len() int {
	return len(h.array)
}