# [1944.Number of Visible People in a Queue][title]

## Description
There are `n` people standing in a queue, and they numbered from `0` to `n - 1` in **left to right** order. You are given an array `heights` of **distinct** integers where `heights[i]` represents the height of the i<sup>th</sup> person.

A person can **see** another person to their right in the queue if everybody in between is **shorter** than both of them. More formally, the i<sup>th</sup> person can see the j<sup>th</sup> person if `i < j` and `min(heights[i], heights[j]) > max(heights[i+1], heights[i+2], ..., heights[j-1])`.

Return an array `answer` of length n where `answer[i]` is the **number of people** the i<sup>th</sup> person can **see** to their right in the queue.

**Example 1:**  

![example](./queue-plane.jpg)

```
Input: heights = [10,6,8,5,11,9]
Output: [3,1,2,1,1,0]
Explanation:
Person 0 can see person 1, 2, and 4.
Person 1 can see person 2.
Person 2 can see person 3 and 4.
Person 3 can see person 4.
Person 4 can see person 5.
Person 5 can see no one since nobody is to the right of them.
```

**Example 2:**

```
Input: heights = [5,1,2,3,10]
Output: [4,1,1,1,0]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/number-of-visible-people-in-a-queue/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
