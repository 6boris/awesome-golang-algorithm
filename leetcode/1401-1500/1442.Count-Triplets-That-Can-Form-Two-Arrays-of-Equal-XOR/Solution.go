package Solution

// max(n)=300, So O(n^3) is ok.
func Solution(arr []int) int {
	ans, xor := 0, arr[0]
	l := len(arr)
	cache := make([]int, l)
	cache[0] = xor
	for k := 1; k < l; k++ {
		xor ^= arr[k]
		now := 0
		for j := k; j > 0; j-- {
			now ^= arr[k]
			left := xor ^ now
			if left == now {
				ans++
			}
			for i := j - 1; i > 0; i-- {
				if left^cache[i-1] == now {
					ans++
				}
			}
		}
		cache[k] = xor
	}
	return ans
}
