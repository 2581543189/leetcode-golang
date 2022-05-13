package main

import "fmt"

func main() {
	fmt.Println(oneEditAway("ab", "bc"))
}

func oneEditAway(first, second string) bool {
	m, n := len(first), len(second)
	if m < n {
		return oneEditAway(second, first)
	}
	if m-n > 1 {
		return false
	}
	for i, ch := range second {
		if first[i] != byte(ch) {
			if m == n {
				return first[i+1:] == second[i+1:]
			}
			return first[i+1:] == second[i:]
		}
	}
	return true
}
