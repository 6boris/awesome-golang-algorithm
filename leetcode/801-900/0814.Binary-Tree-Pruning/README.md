# [814.Binary Tree Pruning][title]


## Description

Given the `root` of a binary tree, return the same tree where every subtree (of the given tree) not containing a `1` has been removed.

A subtree of a node `node` is `node` plus every node that is a descendant of `node`.

**Example 1:**

```
Input: root = [1,null,0,0,1]
Output: [1,null,0,null,1]
Explanation:
Only the red nodes satisfy the property "every subtree not containing a 1".
The diagram on the right represents the answer.
```

__Example 2:__

```
Input: root = [1,0,1,0,0,0,1]
Output: [1,null,1,null,1]
```

__Example 3:__

```
Input: root = [1,1,0,1,1,0,1,0]
Output: [1,1,0,1,1,null,1]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/binary-tree-pruning/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
