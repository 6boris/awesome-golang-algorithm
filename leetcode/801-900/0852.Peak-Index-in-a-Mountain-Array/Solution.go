package Solution

func Solution(arr []int) int {
	ll := len(arr)
	s := make([]int, ll+2)

	copy(s[1:], arr)
	s[0] = arr[0] - 1
	s[ll+1] = arr[ll-1] - 1

	l, r := 1, ll+2
	ans := 0
	for l < r {
		mid := r - (r-l)/2
		if s[mid] > s[mid+1] && s[mid] > s[mid-1] {
			ans = mid
			break
		}
		if s[mid] > s[mid+1] {
			r = mid - 1
			continue
		}
		l = mid
	}
	return ans - 1
}
