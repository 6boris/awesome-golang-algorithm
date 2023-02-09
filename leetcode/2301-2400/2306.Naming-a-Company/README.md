# [2306.Naming a Company][title]

## Description
You are given an array of strings `ideas` that represents a list of names to be used in the process of naming a company. The process of naming a company is as follows:

Choose 2 **distinct** names from `ideas`, call them idea<sub>A</sub> and idea<sub>B</sub>.

Swap the first letters of idea<sub>A</sub> and idea<sub>B</sub> with each other.

If **both** of the new names are not found in the original `ideas`, then the name idea<sub>A</sub> idea<sub>B</sub> (the **concatenation** of idea<sub>A</sub> and idea<sub>B</sub>, separated by a space) is a valid company name.
Otherwise, it is not a valid name.
Return the number of **distinct** valid names for the company.

**Example 1:**

```
Input: ideas = ["coffee","donuts","time","toffee"]
Output: 6
Explanation: The following selections are valid:
- ("coffee", "donuts"): The company name created is "doffee conuts".
- ("donuts", "coffee"): The company name created is "conuts doffee".
- ("donuts", "time"): The company name created is "tonuts dime".
- ("donuts", "toffee"): The company name created is "tonuts doffee".
- ("time", "donuts"): The company name created is "dime tonuts".
- ("toffee", "donuts"): The company name created is "doffee tonuts".
Therefore, there are a total of 6 distinct company names.

The following are some examples of invalid selections:
- ("coffee", "time"): The name "toffee" formed after swapping already exists in the original array.
- ("time", "toffee"): Both names are still the same after swapping and exist in the original array.
- ("coffee", "toffee"): Both names formed after swapping already exist in the original array.

```

**Example 2**

```
Input: ideas = ["lack","back"]
Output: 0
Explanation: There are no valid selections. Therefore, 0 is returned.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/naming-a-company/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
