# [56. Merge Intervals][title]

## Description

Given a collection of intervals, merge all overlapping intervals.

**Example 1:**

```
Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
```

**Example 2:**

```
Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considerred overlapping.
```

**Tags:** Array, Sort

## 题意
>题意是给你一组区间，让你把区间合并成没有交集的一组区间。

## 题解

### 思路1
> 我们可以把区间按 `start` 进行排序，然后遍历排序后的区间，如果当前的 `start` 小于前者的 `end`，那么说明这两个存在交集，我们取两者中较大的 `end` 即可；否则的话直接插入到结果序列中即可。

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
