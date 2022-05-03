package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(reorderLogFiles([]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}))
}

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		s, t := logs[i], logs[j]
		s1 := strings.SplitN(s, " ", 2)[1]
		s2 := strings.SplitN(t, " ", 2)[1]
		isDigit1 := unicode.IsDigit(rune(s1[0]))
		isDigit2 := unicode.IsDigit(rune(s2[0]))
		if isDigit1 && isDigit2 {
			return false
		}
		if !isDigit1 && !isDigit2 {
			return s1 < s2 || s1 == s2 && s < t
		}
		return !isDigit1
	})
	return logs
}
