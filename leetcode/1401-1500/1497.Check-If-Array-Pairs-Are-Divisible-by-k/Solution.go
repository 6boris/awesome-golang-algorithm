package Solution

func Solution(arr []int, k int) bool {
	count := make([]int, k)
	for _, n := range arr {
		t := n % k
		if n < 0 {
			t = k - (-n % k)
			if t == k {
				t = 0
			}
		}
		count[t]++
	}
	for i := 1; i < k; i++ {
		if count[i] != count[k-i] {
			return false
		}
	}
	return count[0]%2 == 0
}
