package main

import "fmt"

func main(){
	var s string = "ac"
	var t string = "bb"
	fmt.Println(isAnagram(s,t))
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