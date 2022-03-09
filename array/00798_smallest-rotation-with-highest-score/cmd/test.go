package main

import "fmt"

func main(){
	fmt.Println(bestRotation1([]int{2,3,1,4,0}))
}

func bestRotation1(nums []int) int {
	n := len(nums)
	diffs := make([]int, n)
	for i, num := range nums {
		if i < num{
			// 首先走 i 到达 0 不记分
			// 然后从 i + 1 开始 可以回到最右侧，然后一直到回到num，全部都能 +1 分
			diffs[i + 1]++
			// 如果回到 num 刚好走了n步，差分数组就不用 -- 了
			if (i + 1  + (n - num)) < n {
				diffs[i + 1  + (n - num)]--
			}

		}else{
			// 从 0开始 走到 num 这个位置 全部都能 +1 分
			diffs[0]++
			// 如果从 0 走到 num 刚好走了n步，差分数组就不用 -- 了
			if i - num + 1 < n{
				diffs[i - num + 1]--
			}
			// 然后 一直走到 0 都不记分
			// 最后 从 i + 1 开始 可以回到最右侧，一直到回到原点，全部都能 +1 分
			if (i + 1) < n{
				diffs[i + 1]++
			}
			// 因为回到原点刚好走 n 步，所以差分数组不用 --
		}

		//为了方便理解记录日志：
		//[2,3,1,4,0]

		//2   [1 2 3]
		//3     [2 3]
		//1 [0 1   3 4]
		//4         [4]
		//0 [0 1 2 3 4]

		//0 1 1 1 0  0 1 0 0 -1
		//0 1 2 2 0  0 1 1 0 -2
		//1 2 2 3 1  1 1 0 1 -2
		//1 2 2 3 2  1 1 0 1 -1
		//2 3 3 4 3  2 1 0 1 -1

		// 下面的写法太过抽象，暂时先不这么写了

		// low := (i + 1) % n
		// high := (i - num + n + 1) % n
		// diffs[low]++
		// diffs[high]--
		// if low >= high {
		//     diffs[0]++
		// }
	}
	score, maxScore, idx := 0, 0, 0
	for i, diff := range diffs {
		score += diff
		if score > maxScore {
			maxScore, idx = score, i
		}
	}
	return idx
}