# [1669.Merge In Between Linked Lists][title]

## Description
You are given two linked lists: `list1` and `list2` of sizes `n` and `m` respectively.

Remove `list1`'s nodes from the a<sup>th</sup> node to the b<sup>th</sup> node, and put `list2` in their place.

The blue edges and nodes in the following figure indicate the result:

![1](./fig1.png)

**Example 1:**  

![2](./ll.png)

```
Input: list1 = [10,1,13,6,9,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
Output: [10,1,13,1000000,1000001,1000002,5]
Explanation: We remove the nodes 3 and 4 and put the entire list2 in their place. The blue edges and nodes in the above figure indicate the result.
```

**Example 2:**  

![3](./merge_linked_list_ex2.png)

```
Input: list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1000003,1000004]
Output: [0,1,1000000,1000001,1000002,1000003,1000004,6]
Explanation: The blue edges and nodes in the above figure indicate the result.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/merge-in-between-linked-lists/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
