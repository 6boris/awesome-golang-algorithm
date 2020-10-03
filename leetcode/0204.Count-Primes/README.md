# [204. Count Primes][title]

## Description

Count the number of prime numbers less than a non-negative number, n.

**Example 1:**

```
Input: 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
```

**Tags:** Math, String

## 题意
> 求一个数以内的质数的个数

## 题解

### 思路1
> 求质数可以用 `Sieve of Eratosthenes` 算法将时间复杂度降到

```json
（1）先把1删除（现今数学界1既不是质数也不是合数）
（2）读取队列中当前最小的数2，然后把2的倍数删去
（3）读取队列中当前最小的数3，然后把3的倍数删去
（4）读取队列中当前最小的数5，然后把5的倍数删去
（5）读取队列中当前最小的数7，然后把7的倍数删去
（6）如上所述直到需求的范围内所有的数均删除或读取
```

```go
func countPrimes(num int) int {
	isPrime := make([]bool, num)
	//	初始化数组 true表示都是质数
	for i := 0; i < num; i++ {
		isPrime[i] = true
	}
	//	循环检查质数 i<sqrt(n)可以转换为i*i < n
	//	因为提劲的进行标记所以时间复杂可以到O(log N)
	for i := 2; i*i < num; i++ {
		if isPrime[i] == false {
			continue
		}
		for j := i * i; j < num; j = j + i {
			isPrime[j] = false
		}
	}
	ans := 0
	for i := 2; i < num; i++ {
		if isPrime[i] {
			ans++
		}
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

[title]: https://leetcode.com/problems/count-primes/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
