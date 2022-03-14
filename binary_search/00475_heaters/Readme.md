
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/heaters

冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。

在加热器的加热半径范围内的每个房屋都可以获得供暖。

现在，给出位于一条水平线上的房屋 houses 和供暖器 heaters 的位置，请你找出并返回可以覆盖所有房屋的最小加热半径。

说明：所有供暖器都遵循你的半径标准，加热的半径也一样。



示例 1:
```

```
示例 2:
```

```

示例 3:
```

```

提示：

* 1 <= houses.length, heaters.length <= 3 * 104
* 1 <= houses[i], heaters[i] <= 109

总结：
* 排序数字，sort.Ints(heaters) 默认是升序
* 查找数字 j := sort.SearchInts(heaters, house) 返回被搜索的元素应该在目标slice中按升序排序该插入的位置