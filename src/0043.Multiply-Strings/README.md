# [1. Add Sum][title]

## Description

1. Given two non-negative integers `num1` and `num2` represented as strings, return the product of `num1` and `num2`, also represented as a string.

   **Example 1:**

   ```
   Input: num1 = "2", num2 = "3"
   Output: "6"
   ```

   **Example 2:**

   ```
   Input: num1 = "123", num2 = "456"
   Output: "56088"
   ```

   **Note:**

   1. The length of both `num1` and `num2` is < 110.
   2. Both `num1` and `num2` contain only digits `0-9`.
   3. Both `num1` and `num2` do not contain any leading zero, except the number 0 itself.
   4. You **must not use any built-in BigInteger library** or **convert the inputs to integer** directly.

**Tags:** Math, String

## 题意
>题意是让你计算两个非负字符串的乘积

## 题解

### 思路1
> 我们模拟小学数学的方式来做，一位一位模拟计算，再各位累加。下面是一位大佬给的图

<div align=center>
<img src="https://github.com/kylesliu/awesome-golang-leetcode/blob/master/assets/images/0043-Solution.jpg"></img>
</div>

```go
func multiply(num1 string, num2 string) string {
	pos := make([]int64, len(num1)+len(num2))
	ans := ""

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			mul := (num1[i] - '0') * (num2[j] - '0')
			p1, p2 := i+j, i+j+1
			sum := int64(mul) + pos[p2]

			pos[p1] += sum / 10
			pos[p2] = sum % 10

		}
	}


	for _, v := range pos {
		ans +=  strconv.Itoa(int(v))
	}
	ans = strings.TrimLeft(ans,"0")
	if len(ans) == 0 {
		return "0"
	}

	return ans
}
```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/two-sum/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
