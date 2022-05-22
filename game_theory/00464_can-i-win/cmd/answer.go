package main

import "fmt"

func main() {
	fmt.Println(canIWin(10, 11))
}

func canIWin(maxChoosableInteger, desiredTotal int) bool {
	// 所有数字加起来都不够，双输
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	// 构造缓存
	dp := make([]int8, 1<<maxChoosableInteger)
	for i := range dp {
		dp[i] = -1
	}
	// 该方法返回当前玩家的结果，返回 1 表示必胜, 0 表示非必胜
	// 当前玩家不一定是整个游戏的先手，但一定是本轮游戏的先手
	var dfs func(int, int) int8
	dfs = func(usedNum, curTot int) (res int8) {
		// 有缓存用缓存
		dv := &dp[usedNum]
		if *dv != -1 {
			return *dv
		}
		// 使用 defer 优化代码
		defer func() { *dv = res }()
		// 随机选择一个数字
		for i := 0; i < maxChoosableInteger; i++ {
			// 已经选过的不再重复选择
			if choosen(usedNum, i) {
				continue
			}
			chooseNum := i + 1
			newTotal := curTot + chooseNum
			// 如果本轮能赢，当前玩家肯定会选择让自己赢的数字
			if newTotal >= desiredTotal {
				return 1
			}
			newUsedNum := usedNum | (1 << i)
			dfsResult := dfs(newUsedNum, newTotal)
			// 如果本轮玩家无法确定是否能赢，就要保证，下一轮对方无论怎么选都是输
			if dfsResult == 0 {
				return 1
			}
		}
		// 当权玩家无论选择什么，都无法赢，那么当前玩家在此场景下必输
		return 0
	}

	return dfs(0, 0) == 1
}

// 判断当前第index 个数字数否被选中
func choosen(usedNum, index int) bool {
	return (usedNum>>index)&1 == 1
}
