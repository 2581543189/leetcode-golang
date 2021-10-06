
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/word-break-ii/

给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。

输入：
```
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
```
输出：
```
[
  "cats and dog",
  "cat sand dog"
]
```

输入：
```
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
```
输出：
```
[
  "pine apple pen apple",
  "pineapple pen apple",
  "pine applepen apple"
]
```

输入：
```
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
```
输出：
```
[]
```

提示：
    

* 分隔时可以重复使用字典中的单词。  
* 你可以假设字典中没有重复的单词。

总结：

* 初始化动态大小的数组
```go
	cache := make([][]string,n,n)
```

* 状态转移方程错误，实际上一维数组即可，没有这么条件分支
* 判断集合中是否存在元素的写法
```go
if _, has := wordSet[word]; has {
	
}
```
* slice 第0项插入的写法
```go
append([]string{word}, nextWords...)
```
* 字符串拼接不需要fmt，strings.Join(value, " ") 即可