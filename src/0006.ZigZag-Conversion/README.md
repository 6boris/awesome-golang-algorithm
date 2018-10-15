# [6. ZigZag Conversion][title]

## Description

The string `"PAYPALISHIRING"` is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```
P   A   H   N
A P L S I I G
Y   I   R
```

And then read line by line: `"PAHNAPLSIIGYIR"`

Write the code that will take a string and make this conversion given a number of rows:

```
string convert(string s, int numRows);
```

**Example 1:**

```
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
```

**Example 2:**

```
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
```

**Tags:** String

## 题意
>题意是让你把字符串按波浪形排好，然后返回横向读取的字符串。听不懂的话，看下面的表示应该就明白了：
```go
0                           2n-2                        4n-4
1                    2n-3   2n-1                 4n-5   4n-5
2              2n-4         2n               4n-6       .
.           .               .             .             .
.       n+1                 .          3n-1             .
n-2   n                     3n-4   3n-2                 5n-6
n-1                         3n-3                        5n-5
```

## 题解

### 思路1
> 那么我们可以根据上面找规律，可以看到波峰和波谷是单顶点的，它们周期是 2 * (n - 1)，单独处理即可；中间的部分每个周期会出现两次，规律很好找，留给读者自己想象，不懂的可以结合以下代码。

```go

```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/zigzag-conversion/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
