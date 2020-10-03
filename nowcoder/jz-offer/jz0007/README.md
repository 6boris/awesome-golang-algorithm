# [JZ7.斐波那契数列][title]


## 题目描述
大家都知道斐波那契数列，现在要求输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0，第1项是1）。
n<=39

**范例 :**

```
Input: n=1
Output: 1

Input: n = 6
Output: 8

```

## 题意
> 计算斐波那契数列

## 题解

### 思路1
> 普通递归版求法，这种方法通常和汉诺塔一起被放在课本的递归教学部分，应该是面试官不希望看到的算法。

```go
package main
/**
 * 
 * @param n int整型 
 * @return int整型
*/
func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
```


### 思路2
> 我们可以将递推式的求解从自顶向下改为自底向上（循环实现）。简而言之，我们已知前两项的值，然后我们就可以用前两项的值求出第3项的值，接着求第4、第5、……，直到求出第n项的值。（废话）

```go
package main
/**
 * 
 * @param n int整型 
 * @return int整型
*/
func Fibonacci2(n int) int {
	a, b := 0, 1

	for i := 1; i <= n; i++ {
		a = a + b
		b = a - b
	}
	return a
}
```



## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://www.nowcoder.com/practice/c6c7742f5ba7442aada113136ddea0c3/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
