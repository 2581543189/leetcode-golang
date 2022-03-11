package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func wordBreak1(s string, wordDict []string) []string {
	//返回值
	ret := []string{}

	n := len(s)

	//使用动态规划计算计算结果集
	wordDictMap := map[string]string{}
	for _, str := range wordDict {
		wordDictMap[str] = str
	}

	cache := make([][][]string, n)


	//定义dfs
	var dfs func( x int)[][]string
	dfs = func(x int) [][]string{

		words := [][]string {}

		if cache[x] != nil {
			return cache[x]
		}

		for i := x + 1; i < n; i++ {
			sub := s[x : i]
			if _,has := wordDictMap[sub]; has {
				for _,value := range dfs(i){
					words = append(words,append([]string{sub},value...))
				}
			}
		}

		sub := s[x:]
		if _,has := wordDictMap[sub]; has {
			words = append(words,[]string{sub})
		}

		cache[x] = words
		return words
	}
	dfs(0)

	if len(cache[0]) > 0{
		for _,value := range cache[0]{
			ret = append(ret,strings.Join(value, " "))
		}
	}

	return ret
}

func main() {
	//s := "pineapplepenapple"
	//wordDict := []string{"apple", "pen", "applepen", "pine", "pineapple"}
	s := "abcd"
	wordDict := []string{"a", "abc", "b", "cd"}
	breaked := wordBreak1(s, wordDict)
	byteArray, err := json.Marshal(breaked)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteArray))
}
