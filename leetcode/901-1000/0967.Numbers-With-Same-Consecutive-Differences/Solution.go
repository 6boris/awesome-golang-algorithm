package Solution

import "fmt"

func Solution(n int, k int) []int {

	ans := make([]int, 0)
	bs := make([]byte, n)
	for idx := 1; idx <= 9; idx++ {
		ans = append(ans, helper967(0, idx, k, n, bs)...)
	}
	return ans
}

func helper967(deep, now, k, n int, bs []byte) []int {

	ans := make([]int, 0)
	bs[deep] = uint8(now) + '0'
	if deep == n-1 {
		var num int
		fmt.Sscanf(string(bs), "%d", &num)
		return []int{num}
	}
	if r := now - k; r >= 0 && r <= 9 {
		ans = append(ans, helper967(deep+1, r, k, n, bs)...)
	}
	if r := now + k; r >= 0 && r <= 9 && k != 0 {
		ans = append(ans, helper967(deep+1, r, k, n, bs)...)
	}

	return ans
}
