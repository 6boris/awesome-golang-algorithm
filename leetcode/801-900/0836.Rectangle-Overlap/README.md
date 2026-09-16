# [836.Rectangle Overlap][title]

## Description
An axis-aligned rectangle is represented as a list `[x1, y1, x2, y2]`, where (`x1, y1`) is the coordinate of its bottom-left corner, and (`x2, y2`) is the coordinate of its top-right corner. Its top and bottom edges are parallel to the X-axis, and its left and right edges are parallel to the Y-axis.

Two rectangles overlap if the area of their intersection is **positive**. To be clear, two rectangles that only touch at the corner or edges do not overlap.

Given two axis-aligned rectangles `rec1` and `rec2`, return `true` if they overlap, otherwise return `false`.

**Example 1:**

```
Input: rec1 = [0,0,2,2], rec2 = [1,1,3,3]
Output: true
```

**Example 2:**

```
Input: rec1 = [0,0,1,1], rec2 = [1,0,2,1]
Output: false
```

**Example 3:**

```
Input: rec1 = [0,0,1,1], rec2 = [2,2,3,3]
Output: false
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/rectangle-overlap/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
