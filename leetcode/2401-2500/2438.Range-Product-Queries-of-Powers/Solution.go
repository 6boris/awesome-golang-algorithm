package Solution

import "sort"

func construct(n int, powers [33]int) []int {
	if n&(n-1) == 0 {
		return []int{n}
	}
	index := sort.Search(33, func(i int) bool {
		return powers[i] > n
	})
	set := []int{powers[index-1]}
	left := construct(n-powers[index-1], powers)
	set = append(set, left...)
	return set
}

func modPow(x int64, n int64, m int64) int64 {
	result := int64(1)
	x %= m
	for n > 0 {
		if n&1 == 1 {
			result = (result * x) % m
		}
		x = (x * x) % m
		n >>= 1
	}
	return result
}

func modInverse(x int64, m int64) int64 {
	return modPow(x, m-2, m)
}

const mod = 1000000007

func Solution(n int, queries [][]int) []int {

	powers := [33]int{}
	powers[0] = 1
	for i := 1; i < 33; i++ {
		powers[i] = powers[i-1] * 2
	}
	arr := construct(n, powers)
	for s, e := 0, len(arr)-1; s < e; s, e = s+1, e-1 {
		arr[s], arr[e] = arr[e], arr[s]
	}
	ans := make([]int, len(queries))
	prefix := make([]int64, len(arr)+1)
	index := 0
	prefix[index] = 1
	for i := 0; i < len(arr); i++ {
		prefix[i+1] = (prefix[index] * int64(arr[i])) % mod
		index++
	}
	for i, q := range queries {
		l, r := q[0], q[1]
		res := (prefix[r+1] * modInverse(prefix[l], mod)) % mod
		ans[i] = int(res)
	}
	return ans
}
