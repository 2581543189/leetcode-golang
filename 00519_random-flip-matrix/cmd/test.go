package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution1 struct {
	m int
	n int
	total int
	cache map[int]int
}


func Constructor1(m int, n int) Solution1 {
	rand.Seed(time.Now().UnixNano())
	return Solution1{
		m,
		n,
		m * n,
		make(map[int]int),
	}
}


func (this *Solution1) Flip() []int {
	target:=0
	if this.total != 0{
		target= rand.Intn(this.total)
	}

	this.total--
	if target <= this.total{
		// 计算realTotal
		realTotal :=  this.total
		if val,ok:=this.cache[this.total];ok{
			realTotal =val
		}

		if val,ok:=this.cache[target];ok{
			this.cache[target] = realTotal
			return []int{val / this.n,val % this.n}
		}
		this.cache[target] = realTotal
		return []int{target / this.n,target % this.n}
	}
	return []int{target / this.n,target % this.n}

}


func (this *Solution1) Reset()  {
	this.cache = make(map[int]int)
	this.total = this.m * this.n
}


func main(){
	s:=Constructor1(2,2)
	fmt.Println(s.Flip())
	fmt.Println(s.Flip())
	s.Reset()
	fmt.Println(s.Flip())
}