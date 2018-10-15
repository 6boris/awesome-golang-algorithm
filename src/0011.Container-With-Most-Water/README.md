# [1. Add Sum][title]

## Description

Given *n* non-negative integers *a1*, *a2*, ..., *an*, where each represents a point at coordinate (*i*, *ai*). *n* vertical lines are drawn such that the two endpoints of line *i* is at (*i*, *ai*) and (*i*, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and *n* is at least 2.

**Tags:** Array, Two Pointers




## 思路

题意是给你 *a1*, *a2*, ..., *an* 这 *n* 个数，代表 (*i*, *ai*) 坐标，让你从中找两个点与 x 轴围成的容器可以容纳最多的水。

## 题意
>题意是给你 *a1*, *a2*, ..., *an* 这 *n* 个数，代表 (*i*, *ai*) 坐标，让你从中找两个点与 x 轴围成的容器可以容纳最多的水。

## 题解

### 思路1
> 如果用暴力法求每种情况的结果，其时间复杂度为 O(n^2)，相信肯定会超时，
我们可以探索下是否有更巧妙的办法呢，题目的标签有双指针，是否就可以想到首尾各放一指针，
然后根据条件来收缩。首先计算一次首尾构成的最大面积，然后分析下该移动哪个指针，
如果移动大的那个指针的话，那样只会减小面积，所以我们要移动小的那个指针，
小的那个指针移动到哪呢？当然是移动到大于之前的值的地方，否则面积不都比之前小么，
然后继续更新最大值即可，借助如上分析写出如下代码应该不是什么难事了吧。
<div align=center>
<img src="https://github.com/kylesliu/awesome-golang-leetcode/blob/master/assets/images/0003-BenchMark.png"></img>
</div>

[官方演示][soltion]

```go

```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/two-sum/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
[soltion]: https://leetcode.com/problems/container-with-most-water/solution/
