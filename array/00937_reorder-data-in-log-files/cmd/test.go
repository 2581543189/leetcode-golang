package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(reorderLogFiles1([]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}))
}
func reorderLogFiles1(logs []string) []string {
	letters := make([]string, 0)
	nums := make([]string, 0)
	for i := range logs {
		firstSpace := strings.Index(logs[i], " ")
		if isNum(logs[i][firstSpace+1:]) {
			nums = append(nums, logs[i])
		} else {
			letters = append(letters, logs[i])
		}
	}
	sort.Slice(letters, func(i, j int) bool {
		fsi := strings.Index(letters[i], " ")
		fsj := strings.Index(letters[j], " ")
		preI := letters[i][:fsi]
		suffixI := letters[i][fsi+1:]
		preJ := letters[j][:fsj]
		suffixJ := letters[j][fsj+1:]
		if suffixI == suffixJ {
			return preI < preJ
		}
		return suffixI < suffixJ
	})
	return append(letters, nums...)

}

func isNum(ch string) bool {
	return ch[0] >= '0' && ch[0] <= '9'
}
