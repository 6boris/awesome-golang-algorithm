# [1475.Final Prices With a Special Discount in a Shop][title]

## Description
You are given an integer array `prices` where `prices[i]` is the price of the i<sup>th</sup> item in a shop.

There is a special discount for items in the shop. If you buy the i<sup>th</sup> item, then you will receive a discount equivalent to `prices[j]` where `j` is the minimum index such that `j > i` and `prices[j] <= prices[i]`. Otherwise, you will not receive any discount at all.

Return an integer array `answer` where answer[i] is the final price you will pay for the i<sup>th</sup> item of the shop, considering the special discount.

**Example 1:**

```
Input: prices = [8,4,6,2,3]
Output: [4,2,4,2,3]
Explanation: 
For item 0 with price[0]=8 you will receive a discount equivalent to prices[1]=4, therefore, the final price you will pay is 8 - 4 = 4.
For item 1 with price[1]=4 you will receive a discount equivalent to prices[3]=2, therefore, the final price you will pay is 4 - 2 = 2.
For item 2 with price[2]=6 you will receive a discount equivalent to prices[3]=2, therefore, the final price you will pay is 6 - 2 = 4.
For items 3 and 4 you will not receive any discount at all.
```

**Example 2:**

```
Input: prices = [1,2,3,4,5]
Output: [1,2,3,4,5]
Explanation: In this case, for all items, you will not receive any discount at all.
```

**Example 3:**

```
Input: prices = [10,1,1,6]
Output: [9,0,1,6]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop
[me]: https://github.com/kylesliu/awesome-golang-algorithm
