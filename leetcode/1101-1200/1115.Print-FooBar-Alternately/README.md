# [1115. Print FooBar Alternately][title]

## Description

Suppose you are given the following code:

```cpp
class FooBar {
  public void foo() {
    for (int i = 0; i < n; i++) {
      print("foo");
    }
  }

  public void bar() {
    for (int i = 0; i < n; i++) {
      print("bar");
    }
  }
}
```

The same instance of `FooBar` will be passed to two different threads:

- thread `A` will call `foo()`, while
- thread `B` will call `bar()`.

Modify the given program to output "`foobar`" `n` times.


**Example 1:**

```
Input: n = 1
Output: "foobar"
Explanation: There are two threads being fired asynchronously. One of them calls foo(), while the other calls bar().
"foobar" is being output 1 time.
```

**Example 2:**

```
Input: n = 2
Output: "foobarfoobar"
Explanation: "foobar" is being output 2 times.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/print-foobar-alternately
[me]: https://github.com/kylesliu/awesome-golang-algorithm
