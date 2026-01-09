# [1116.Print Zero Even Odd][title]

## Description
You have a function `printNumber` that can be called with an integer parameter and prints it to the console.

- For example, calling `printNumber(7)` prints `7` to the console.

You are given an instance of the class `ZeroEvenOdd` that has three functions: `zero`, `even`, and `odd`. The same instance of `ZeroEvenOdd` will be passed to three different threads:

- **Thread A**: calls `zero()` that should only output `0`'s.
- **Thread B**: calls `even()` that should only output even numbers.
- **Thread C**: calls `odd()` that should only output odd numbers.

Modify the given class to output the series `"010203040506..."` where the length of the series must be `2n`.

Implement the `ZeroEvenOdd` class:

- `ZeroEvenOdd(int n)` Initializes the object with the number `n` that represents the numbers that should be printed.
- `void zero(printNumber)` Calls `printNumber` to output one zero.
- `void even(printNumber)` Calls `printNumber` to output one even number.
- `void odd(printNumber)` Calls `printNumber` to output one odd number.

**Example 1:**

```
Input: n = 2
Output: "0102"
Explanation: There are three threads being fired asynchronously.
One of them calls zero(), the other calls even(), and the last one calls odd().
"0102" is the correct output.
```

**Example 2:**

```
Input: n = 5
Output: "0102030405"
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/print-zero-even-odd/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
