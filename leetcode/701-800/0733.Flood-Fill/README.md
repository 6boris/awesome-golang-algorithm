# [733.Flood Fill][title]

## Description
An image is represented by an `m x n` integer grid `image` where `image[i][j]` represents the pixel value of the image.

You are also given three integers `sr`, `sc`, and `newColor`. You should perform a __flood fill__ on the image starting from the pixel `image[sr][sc]`.

To perform a __flood fill__, consider the starting pixel, plus any pixels connected __4-directionally__ to the starting pixel of the same color as the starting pixel, plus any pixels connected __4-directionally__ to those pixels (also with the same color), and so on. Replace the color of all of the aforementioned pixels with `newColor`.

Return the modified image after performing the flood fill.



**Example 1:**

![flood1-grid](flood1-grid.jpg)

```
Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, newColor = 2
Output: [[2,2,2],[2,2,0],[2,0,1]]
Explanation: From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel (i.e., the blue pixels) are colored with the new color.
Note the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.
```

__Example 2:__
```
Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, newColor = 2
Output: [[2,2,2],[2,2,2]]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/flood-fill/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
