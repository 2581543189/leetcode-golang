package main

import (
	"fmt"
	"sort"
)

func reversePairs1(nums []int) int {
	var mergeSort func(array []int,left,right int) int
	mergeSort = func(array []int,left,right int) int{
		if left >= right {
			return 0
		}
		mid := left + (right - left) >> 1
		ans:= mergeSort(array,left,mid) + mergeSort(array,mid + 1,right)

		i,j := left, mid + 1
		tmp:=make([]int,0,right - left + 1)
		for i <= mid && j <= right{
			if array[i] <= array[j]{
				tmp = append(tmp,array[i])
				i++
				ans += j - mid - 1
			}else{
				tmp = append(tmp,array[j])
				j++
			}
		}
		for i <= mid{
			tmp = append(tmp,array[i])
			ans += right - mid
			i++
		}
		for  j <= right{
			tmp = append(tmp,array[j])
			j++
		}
		for index:=left;index <= right;index++{
			nums[index] = tmp[index - left]
		}
		return ans
	}
	return mergeSort(nums,0,len(nums) - 1)
}

func reversePairs2(nums []int) int {
	n:= len(nums)
	bit := BIT{n,make([]int,n + 1)}
	tmp:= make([]int ,n)
	copy(tmp,nums)

	//离散化
	sort.Ints(tmp)
	for i:=0;i<n;i++{
		nums[i] = sort.SearchInts(tmp,nums[i]) + 1 //不能有0
	}
	ans:= 0
	for i:=n-1;i>=0;i--{
		bit.update(nums[i])
		ans += bit.query(nums[i] - 1)
	}

	return ans
}

type BIT struct {
	n int
	tree []int
}

func(bit *BIT) lowBit(x int)int {
	return x & -x
}


func(bit *BIT) update(x int){

	for x <= bit.n{
		bit.tree[x]++
		x += bit.lowBit(x)
	}
}

func(bit BIT) query(x int)int{
	ans:=0
	for x > 0{
		ans+=bit.tree[x]
		x -= bit.lowBit(x)

	}
	return ans
}


func main(){
	nums := []int{7,5,6,4}
	fmt.Println(reversePairs1(nums))
	nums = []int{7,5,6,4}
	fmt.Println(reversePairs2(nums))
}