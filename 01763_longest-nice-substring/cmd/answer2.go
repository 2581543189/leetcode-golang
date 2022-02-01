package main

import (
	"fmt"
	"unicode"
)

func main(){
	fmt.Println(longestNiceSubstring2("YazaAay"))
}

func longestNiceSubstring2(s string) string {
	ans:=""
	low,high:=0,0
	for i:=range s{
		if s[i] >= 'a'{
			low |= 1 << (s[i] - 'a')
		}else{
			high |= 1 << (s[i] - 'A')
		}
	}
	if low == high{
		return s
	}
	diff :=  low & high
	for i:=0;i < len(s);i++{
		start:=i
		for i < len(s) && (diff >> (unicode.ToLower(rune(s[i])) - 'a')) & 1 == 1{
			i++
		}
		if i > start {
			tmp:=longestNiceSubstring2(s[start:i])
			if len(tmp) > len(ans){
				ans = tmp
			}
		}
	}
	return ans
}

