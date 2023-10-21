package Solution

func Solution(arr []int, k int, threshold int) int {
	ans := 0
	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	if sum/k >= threshold {
		ans++
	}
	start := 0
	// 1, 2, 3
	for i := k; i < len(arr); i++ {
		sum -= arr[start]
		sum += arr[i]
		if sum/k >= threshold {
			ans++
		}
		start++
	}
	return ans
}
