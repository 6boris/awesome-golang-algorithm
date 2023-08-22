# [1237.Find Positive Integer Solution for a Given Equation][title]

## Description
Given a callable function `f(x, y)` **with a hidden formula** and a value `z`, reverse engineer the formula and return all positive integer pairs `x` and y where `f(x,y) == z`. You may return the pairs in any order.

While the exact formula is hidden, the function is monotonically increasing, i.e.:

- `f(x, y) < f(x + 1, y)`
- `f(x, y) < f(x, y + 1)`

The function interface is defined like this:

```
interface CustomFunction {
public:
  // Returns some positive integer f(x, y) for two positive integers x and y based on a formula.
  int f(int x, int y);
};
```

We will judge your solution as follows:

- The judge has a list of `9` hidden implementations of `CustomFunction`, along with a way to generate an **answer key** of all valid pairs for a specific `z`.
- The judge will receive two inputs: a `function_id` (to determine which implementation to test your code with), and the target `z`.
- The judge will call your `findSolution` and compare your results with the **answer key**.
- If your results match the **answer key**, your solution will be `Accepted`.

**Example 1:**

```
Input: function_id = 1, z = 5
Output: [[1,4],[2,3],[3,2],[4,1]]
Explanation: The hidden formula for function_id = 1 is f(x, y) = x + y.
The following positive integer values of x and y make f(x, y) equal to 5:
x=1, y=4 -> f(1, 4) = 1 + 4 = 5.
x=2, y=3 -> f(2, 3) = 2 + 3 = 5.
x=3, y=2 -> f(3, 2) = 3 + 2 = 5.
x=4, y=1 -> f(4, 1) = 4 + 1 = 5.
```

**Example 2:**

```
Input: function_id = 2, z = 5
Output: [[1,5],[5,1]]
Explanation: The hidden formula for function_id = 2 is f(x, y) = x * y.
The following positive integer values of x and y make f(x, y) equal to 5:
x=1, y=5 -> f(1, 5) = 1 * 5 = 5.
x=5, y=1 -> f(5, 1) = 5 * 1 = 5.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/find-positive-integer-solution-for-a-given-equation/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
