package main

import "fmt"

func main() {
	fmt.Println(maxConsecutiveAnswers("TTFTTFTT", 2))
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsecutiveChar(answerKey, k, 'T'), maxConsecutiveChar(answerKey, k, 'F'))
}

func maxConsecutiveChar(answerKey string, k int, ch byte) (ans int) {
	left, lsum, rsum := 0, 0, 0
	for i := range answerKey {
		if answerKey[i] != ch {
			rsum++
		}
		for lsum < rsum-k {
			if answerKey[left] != ch {
				lsum++
			}
			left++
		}
		ans = max(ans, i-left+1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
