package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type MedianFinder struct {
	minHeap,maxHeap hp
}


func Constructor() MedianFinder {
	return MedianFinder{}
}


func (this *MedianFinder) AddNum(num int)  {
	min,max := &this.minHeap, &this.maxHeap
	if max.Len() == 0 || num > -max.IntSlice[0] {
		heap.Push(min,num)
		if min.Len() - max.Len() > 1{
			heap.Push(max,-heap.Pop(min).(int))
		}
	}else{
		heap.Push(max,-num)
		heap.Push(min,-heap.Pop(max).(int))
		if min.Len() - max.Len() > 1{
			heap.Push(max,-heap.Pop(min).(int))
		}
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() == 0{
		return 0
	}
	if this.maxHeap.Len() == 0 && this.minHeap.Len()==1{
		return float64(this.minHeap.IntSlice[0])
	}
	if this.maxHeap.Len() == 0 && this.minHeap.Len()==2{
		return float64((this.minHeap.IntSlice[0] - this.minHeap.IntSlice[1]) / 2)
	}
	if this.minHeap.Len() > this.maxHeap.Len() {
		return float64(this.minHeap.IntSlice[0])
	}else{
		return float64((this.minHeap.IntSlice[0] - this.maxHeap.IntSlice[0])) / 2
	}
}

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
	mf := Constructor()
	array :=[] int{6,10,2,6,5,0,6,3,1,0,0}
	for _,value := range array{
		mf.AddNum(value)
		fmt.Println(mf.FindMedian())
	}
	fmt.Println(mf.FindMedian())

}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */