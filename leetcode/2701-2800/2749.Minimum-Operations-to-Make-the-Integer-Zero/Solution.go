package Solution

import (
	"math/bits"
)

func canBeSumOfPowersOfTwo(n int64, k int64) bool {
	// 1. 如果 n < k，则无法组成。因为每个 2^i 至少是 1 (2^0)。
	if n < k {
		return false
	}

	// 2. 计算 n 的二进制表示中 1 的个数，即所需的最小 2^i 数量。
	minK := int64(bits.OnesCount64(uint64(n)))

	// 3. 检查 k 是否在 [minK, n] 范围内。
	// 任何在 minK 和 n 之间的 k 值都是可达的，
	// 因为每次拆分一个 2^i 都会让总数增加 1，直到达到 n。
	return k >= minK && k <= n
}

func Solution(num1 int, num2 int) int {
	// 这个可定是无法完成的
	if num2 >= num1 {
		return -1
	}
	an, bn := int64(num1), int64(num2)
	i := int64(1)
	for ; i < 61; i++ {
		an -= bn
		if canBeSumOfPowersOfTwo(an, i) {
			return int(i)
		}
	}
	return -1
}
