package main

import "fmt"

func main() {
	fmt.Println(findKthNumber(9, 4))
	fmt.Println(findKthNumber(122, 16))
}

// 从 1 ~ end 以 pre 开头的数字有多少个（含 pre）
func count(pre, end int) int {
	ans := 0
	s, e := pre, pre
	for s <= end {
		ans += min(e, end) - s + 1
		s *= 10
		e = e*10 + 9
	}
	return ans
}

func findKthNumber(n, k int) int {
	// need 表示还差几个数
	need := k
	cur := 1
	for need > 1 { // 还差1个的时候，就找到了要找的数字
		// 当前前缀有几个数字
		c := count(cur, n)

		if c >= need {
			// 当前前缀数字足够的话，就继续找
			cur *= 10
			need--
		} else {
			// 当前前缀不够，跳过所有数字，去找下一个前缀
			cur++
			need -= c
		}
	}
	return cur
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
