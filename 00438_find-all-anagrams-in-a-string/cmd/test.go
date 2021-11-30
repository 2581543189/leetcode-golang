package main

import "fmt"

func findAnagrams1(s string, p string) []int {
	ls:=len(s)
	lp:=len(p)
	ans := []int{}
	if lp > ls {
		return ans
	}
	pMap:=make(map[byte]int)
	for _,cha:= range p{
		pMap[byte(cha) - 'a'] ++
	}


	sWindow:=make(map[byte]int)

	for i:=0; i<lp - 1; i++{
		sWindow[s[i] - 'a'] ++
	}

	var equal func ()bool
	equal = func ()bool{
		for k,v:=range pMap{
			if v1,ok:=sWindow[k];ok {
				if v1 != v{
					return false
				}
			}else{
				return false
			}
		}
		return true
	}

	for i:=lp - 1;i< ls;i++{
		if i - lp >= 0{
			sWindow[s[i - lp] - 'a'] --
		}
		sWindow[s[i] - 'a'] ++
		if equal(){
			ans= append(ans,i - lp + 1)
		}
	}
	return ans
}

func main(){
	fmt.Println(findAnagrams1("cbaebabacd","abc"))
}
