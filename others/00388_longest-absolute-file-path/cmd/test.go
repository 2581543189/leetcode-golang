package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthLongestPath1("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"))
}

func lengthLongestPath1(input string) (ans int) {
	list := strings.Split(input, "\n")
	stack := make([]string, 0)
	lastCount := 0
	ansStr := ""
	for i := range list {
		count, s := analysis(list[i])
		if find := strings.Contains(s, "."); find {
			for len(stack) > 0 && len(stack) > count {
				stack = stack[:len(stack)-1]
			}
			tmp := s
			if len(stack) > 0 {
				tmp = strings.Join(stack, "/") + "/" + s
			}

			if len(tmp) > len(ansStr) {
				ansStr = tmp
			}

		} else {
			if count > lastCount {
				stack = append(stack, s)
			} else {
				for len(stack) > 0 && len(stack) > count {
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, s)
			}
		}
		lastCount = count
	}
	return len(ansStr)

}
func analysis(s string) (int, string) {
	//s = strings.Replace(s," ","",-1)
	count := 0
	for len(s) > 1 {
		if s[:1] == "\t" {
			count++
			s = s[1:]
		} else {
			break
		}
	}
	return count, s

}
