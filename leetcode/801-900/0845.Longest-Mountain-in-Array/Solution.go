package Solution

func Solution(arr []int) int {
	left := make([]int, len(arr))
	left[0] = 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	ans := 0
	cnt := 0
	for i := len(arr) - 2; i > 0; i-- {
		if arr[i] > arr[i+1] {
			cnt++
		} else {
			cnt = 0
		}
		if left[i] != 0 && cnt != 0 {
			ans = max(ans, left[i]+cnt)
		}
	}
	if ans > 0 {
		ans++
	}
	return ans
}
