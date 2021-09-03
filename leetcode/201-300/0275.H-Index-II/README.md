# [275. H-Index II][title]

## Description

给定一个科研工作者的一系列论文的引用次数（引用次数是非负整数），已经从小到大排好序，请计算他的h因子。

h因子的定义：一个科学家如果有 hh 篇文章的引用次数至少是 hh，并且其他文章的引用次数不超过 hh，我们就说他的影响因子是 hh。

注意： 如果一个人有多个可能的 hh，返回最大的 hh 因子。
进一步：

这道题目是 LeetCode 274. H-Index的升级版，citations 保证是严格递增的；
你能否给出时间复杂度是 O(logn)O(logn) 级别的算法？

**Example 1:**

```
Input: citations = [0,1,3,5,6]
Output: 3 
Explanation: [0,1,3,5,6] means the researcher has 5 papers in total and each of them had 
             received 0, 1, 3, 5, 6 citations respectively. 
             Since the researcher has 3 papers with at least 3 citations each and the remaining 
             two with no more than 3 citations each, her h-index is 3.
```

**Tags:** Math, String

## 题意
> (二分) O(logn)O(logn)
  由于数组是从小到大排好序的，所以我们的任务是：
  在数组中找一个最大的 hh，使得后 hh 个数大于等于 hh。
  我们发现：如果 hh 满足，则小于 hh 的数都满足；如果 hh 不满足，则大于 hh 的数都不满足。所以具有二分性质。
  直接二分即可。
  时间复杂度分析：二分检索，只遍历 lognlogn 个元素，所以时间复杂度是 O(logn)O(logn)

## 题解

### 思路1
> 。。。。

```go
package main

func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}
	l, r := 0, len(citations)-1
	for l < r {
		mid := (l + r) / 2
		if len(citations)-mid <= citations[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if citations[l] <= len(citations) {
		return len(citations) - l
	}
	return 0
}
```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/h-index-ii/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
