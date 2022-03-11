
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/palindrome-partitioning/

给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。
输入：
```
s = "aab"
```
输出：
```
wordDict = [["a","a","b"],["aa","b"]]
```

输入：
```
s = "a"
```
输出：
```
[["a"]]
```

提示：
    

* 1 <= s.length <= 16
* s 仅由小写英文字母组成

总结：

* 枚举的定义
```go
type State int
const (
	NIL State = iota
	TRUE
	FALSE
)
```
* slice 可以在递归中当成stack使用
```go
splits = append(splits, s[i:j+1])
dfs(j + 1)
splits = splits[:len(splits)-1]
```