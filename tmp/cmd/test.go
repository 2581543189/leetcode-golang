package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch:= make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(){
		defer wg.Done()
		for i:=0;i<5;i++{
			ch <- rand.Intn(100)
		}
		close(ch)
	}()
	go func(){
		defer wg.Done()
		for val:=range ch{
			fmt.Println(val)
		}
	}()
	wg.Wait()
}



func rotate(nums []int, k int)  {
	n:=len(nums)

	first := nums[0]
	second := nums[0]
	for i := 0; i < n - 1; i++{
		second = nums[(i + 1)  * k % n]
		nums[(i + 1) * k % n] = first
		first = second
	}
	nums[0]=first
}

func isAnagram(s string, t string) bool {
	s1:= []byte(s)
	t1:= []byte(t)
	cache:= make(map[byte] int64)

	for _,ch := range s1{
		cache[ch]++
	}
	for _,ch := range t1{
		cache[ch]--
	}

	for key,_ := range cache{
		if cache[key] != 0{
			return false
		}
	}
	return true
}