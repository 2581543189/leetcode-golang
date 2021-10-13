
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/top-k-frequent-elements/

给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。


输入：
```
matrix = nums = [1,1,1,2,2,3], k = 2
```
输出：
```
[1,2]
```

输入：
```
nums = [1], k = 1
```
输出：
```
[1]
```

提示：
    

* 1 <= nums.length <= 105
* k 的取值范围是 [1, 数组中不相同的元素的个数]
* 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的

总结：
* go 的heap 需要实现多个方法，其中改变堆内容的 比如 push pop 是指针方法，不改变的比如len less 是非指针方法
* heap.Push 或者pop 因为要改变堆的内容，所以入参是指针
* less 方法 的i 和 j 是 heap 数组的下标。如果堆本身存的就是下标，不要搞混了
* index = 0  是根，默认是小根堆
* heap.pop 方法，会把根交换到 len - 1,所以pop 需要把 len - 1 返回