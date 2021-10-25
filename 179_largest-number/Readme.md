
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/largest-number/

给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

输入：
```
nums = [10,2]
```
输出：
```
"210"
```

输入：
```
nums = [3,30,34,5,9]
```
输出：
```
"9534330"
```

输入：
```
nums = [1]
```
输出：
```
"1"
```

输入：
```
nums = [10]
```
输出：
```
"10"
```


提示：
    

* 1 <= nums.length <= 100
  0 <= nums[i] <= 10^9


总结：
* golang 排序 sort.Slice(nums, func(i, j int) bool {
* 从数组顺序上说，j在i 的前面 因为 data.Less(j, j-1)
* 从方便理解上看，i 在 j 的前面 比如，如果希望大的放在前面，就返回 i > j
* ans = append(ans, strconv.Itoa(val)...) 字符串作为 append 的参数，加上省略号，会被自动拆分