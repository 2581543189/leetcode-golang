
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/maximum-product-subarray/

给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

输入：
```
[2,3,-2,4]
```
输出：
```
6
```

输入：
```
[-2,0,-1]
```
输出：
```
0
```
提示：
    

* 

总结：

* 已经四道动态规划的题目了，状态转移方程首先要考虑一个参数行不行，比如，如果是数组截取的话，考虑以n为start 的数组作为状态n
* math.MinInt32
* const INT_MAX = int(^uint(0) >> 1)
* const INT_MIN = ^INT_MAX
* Math.max(a,b)