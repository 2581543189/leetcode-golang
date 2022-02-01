package main

import "fmt"

func main(){
	fmt.Println(longestNiceSubstring1("YazaAay"))
}

func longestNiceSubstring1(s string) string {
	ans:= ""
	for l:=range s{
		maxR:= len(s) - 1
		if maxR - l <= len(ans) - 1{
			continue
		}
		for r:=maxR;r >= l;r--{
			if r - l <= len(ans) - 1{
				break
			}
			if isNice(s[l:r+1]){
				ans = s[l:r+1]
			}
		}
	}
	return ans

}

func isNice(s string)bool{
	low,high:=0,0
	for i:=range s{
		if s[i] >= 'a'{
			low |= 1 << (s[i] - 'a')
		}else{
			high |= 1 << (s[i] - 'A')
		}
	}
	return low == high
}

