package main

import (
	"encoding/json"
	"fmt"
)

type State int
const (
	NIL State = iota
	TRUE
	FALSE
)

func partition1(s string) [][]string {

	n := len(s)

	palindromeCache := make([][]State,n)
	for i := 0; i < n; i++ {
		palindromeCache[i] = make([]State, n)
	}

	// 判断是否是回文的方法
	var isPalindrome func(i,j int)bool
	isPalindrome = func(i,j int)bool{

		if palindromeCache[i][j] != NIL {
			return  palindromeCache[i][j] == TRUE
		}
		if i == j || j < i{
			palindromeCache[i][j] = TRUE
			return true
		}
		if s[i:i+1] == s[j:j+1]  && isPalindrome(i+1,j-1){
			palindromeCache[i][j] = TRUE
			return true
		}
		palindromeCache[i][j] = FALSE
		return false
	}

	// 回溯所有可能的情况
	partitionCache := make([][][]string,n)
	var doPattitions func(i int)[][]string
	doPattitions = func(i int) [][]string {
		if partitionCache[i] != nil {
			return partitionCache[i]
		}

		partitions := [][]string{}

		for x:= i+1; x<n; x++{
			if isPalindrome(i,x-1){
				sub := s[i:x]
				for _,pattition := range doPattitions(x){
					partitions = append(partitions,append([]string{sub},pattition...))
				}
			}
		}
		if isPalindrome(i,n-1) {
			sub := s[i:n]
			partitions = append(partitions,append([]string{sub}))
		}

		partitionCache[i] = partitions
		return partitions
	}

	doPattitions(0)

	return partitionCache[0]
}


func main() {

	s := "aab"
	partitioned := partition1(s)
	byteArray, err := json.Marshal(partitioned)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteArray))
}