package Solution

import "math/bits"

func Solution(k int64, operations []int) byte {
	ans := 0
	for k != 1 {
		t := bits.Len64(uint64(k)) - 1
		if (1 << t) == k {
			t--
		}
		k -= (1 << t)
		if operations[t] != 0 {
			ans++
		}
	}
	return byte('a' + (ans % 26))
}
