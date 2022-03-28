package main

import (
	"container/heap"
	"fmt"
)

func main(){
	sp:=Constructor2()
	sp.Update(1,10)
	sp.Update(2,5)
	fmt.Println(sp.Current())
	fmt.Println(sp.Maximum())
	sp.Update(1,3)
	fmt.Println(sp.Maximum())
	sp.Update(4,2)
	fmt.Println(sp.Minimum())
}

type StockPrice2 struct {
	maxPrice, minPrice hp
	timePriceMap       map[int]int
	maxTimestamp       int
}

func Constructor2() StockPrice2 {
	return StockPrice2{timePriceMap: map[int]int{}}
}

func (sp *StockPrice2) Update(timestamp, price int) {
	heap.Push(&sp.maxPrice, pair{-price, timestamp})
	heap.Push(&sp.minPrice, pair{price, timestamp})
	sp.timePriceMap[timestamp] = price
	if timestamp > sp.maxTimestamp {
		sp.maxTimestamp = timestamp
	}
}

func (sp *StockPrice2) Current() int {
	return sp.timePriceMap[sp.maxTimestamp]
}

func (sp *StockPrice2) Maximum() int {
	for {
		if p := sp.maxPrice[0]; -p.price == sp.timePriceMap[p.timestamp] {
			return -p.price
		}
		heap.Pop(&sp.maxPrice)
	}
}

func (sp *StockPrice2) Minimum() int {
	for {
		if p := sp.minPrice[0]; p.price == sp.timePriceMap[p.timestamp] {
			return p.price
		}
		heap.Pop(&sp.minPrice)
	}
}

type pair struct{ price, timestamp int }
type hp []pair
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
