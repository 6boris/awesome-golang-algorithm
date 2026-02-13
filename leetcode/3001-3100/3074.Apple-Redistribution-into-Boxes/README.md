# [3074.Apple Redistribution into Boxes][title]

## Description
You are given an array `apple` of size `n` and an array `capacity` of size `m`.

There are `n` packs where the `ith` pack contains `apple[i]` apples. There are m boxes as well, and the `ith` box has a capacity of `capacity[i]` apples.

Return the **minimum** number of boxes you need to select to redistribute these `n` packs of apples into boxes.

**Note** that, apples from the same pack can be distributed into different boxes.

**Example 1:**

```
Input: apple = [1,3,2], capacity = [4,3,1,5,2]
Output: 2
Explanation: We will use boxes with capacities 4 and 5.
It is possible to distribute the apples as the total capacity is greater than or equal to the total number of apples.
```

**Example 2:**

```
Input: apple = [5,5,5], capacity = [2,4,2,7]
Output: 4
Explanation: We will need to use all the boxes.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/apple-redistribution-into-boxes/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
