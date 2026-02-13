# [3047.Find the Largest Area of Square Inside Two Rectangles][title]

## Description
There exist `n` rectangles in a 2D plane with edges parallel to the x and y axis. You are given two 2D integer arrays `bottomLeft` and `topRight` where `bottomLeft[i] = [a_i, b_i]` and `topRight[i] = [c_i, d_i]` represent the **bottom-left** and **top-right** coordinates of the `ith` rectangle, respectively.

You need to find the **maximum** area of a **square** that can fit inside the intersecting region of at least two rectangles. Return `0` if such a square does not exist.

**Example 1:**  

![1](./1.png)

```
Input: bottomLeft = [[1,1],[2,2],[3,1]], topRight = [[3,3],[4,4],[6,6]]

Output: 1

Explanation:

A square with side length 1 can fit inside either the intersecting region of rectangles 0 and 1 or the intersecting region of rectangles 1 and 2. Hence the maximum area is 1. It can be shown that a square with a greater side length can not fit inside any intersecting region of two rectangles.
```

**Example 2:**  

![2](./2.png)

```
Input: bottomLeft = [[1,1],[1,3],[1,5]], topRight = [[5,5],[5,7],[5,9]]

Output: 4

Explanation:

A square with side length 2 can fit inside either the intersecting region of rectangles 0 and 1 or the intersecting region of rectangles 1 and 2. Hence the maximum area is 2 * 2 = 4. It can be shown that a square with a greater side length can not fit inside any intersecting region of two rectangles.
```

**Example 3:**  

![3](./3.png)

```
Input: bottomLeft = [[1,1],[2,2],[1,2]], topRight = [[3,3],[4,4],[3,4]]

Output: 1

Explanation:

A square with side length 1 can fit inside the intersecting region of any two rectangles. Also, no larger square can, so the maximum area is 1. Note that the region can be formed by the intersection of more than 2 rectangles.
```

**Example 4:**  

![4](./4.png)

```
Input: bottomLeft = [[1,1],[3,3],[3,1]], topRight = [[2,2],[4,4],[4,2]]

Output: 0

Explanation:

No pair of rectangles intersect, hence, the answer is 0.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/find-the-largest-area-of-square-inside-two-rectangles/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
