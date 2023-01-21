package Solution

func Solution(arr []int, queries [][]int) []int {
	l := len(arr)
	ans := make([]int, len(queries))
	xor := make([]int, len(arr))
	xor[0] = arr[0]
	for i := 1; i < l; i++ {
		xor[i] = xor[i-1] ^ arr[i]
	}
	for i, q := range queries {
		left, right := q[0], q[1]
		if left == right {
			ans[i] = arr[right]
			continue
		}
		if left == 0 {
			ans[i] = xor[right]
			continue
		}
		ans[i] = xor[right] ^ xor[left-1]
	}
	return ans
}
