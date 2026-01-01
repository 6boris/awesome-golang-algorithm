# [3668.Restore Finishing Order][title]

## Description
 are given an integer array `order` of length n and an integer array `friends`.

- `order` contains every integer from 1 to `n` **exactly once**, representing the IDs of the participants of a race in their **finishing** order.
- `friends` contains the IDs of your friends in the race **sorted** in strictly increasing order. Each ID in friends is guaranteed to appear in the `order` array.

Return an array containing your friends' IDs in their **finishing** order.

**Example 1:**

```


Input: order = [3,1,2,5,4], friends = [1,3,4]

Output: [3,1,4]

Explanation:

The finishing order is [3, 1, 2, 5, 4]. Therefore, the finishing order of your friends is [3, 1, 4].
```

**Example 2:**

```


Input: order = [1,4,5,3,2], friends = [2,5]

Output: [5,2]

Explanation:

The finishing order is [1, 4, 5, 3, 2]. Therefore, the finishing order of your friends is [5, 2].
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/restore-finishing-order/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
