package Solution

func Solution(arr []int, target int) int {
	pos := map[int]int{0: -1}
	n := len(arr)
	s, ans, minL := 0, n+1, n
	for i, x := range arr {
		s += x
		if j, ok := pos[s-target]; ok {
			length := i - j
			prev := n
			if j != -1 {
				prev = arr[j]
			}
			if length+prev < ans {
				ans = length + prev
			}
			if length < minL {
				minL = length
			}
		}
		arr[i] = minL
		pos[s] = i
	}
	if ans == n+1 {
		return -1
	}
	return ans
}
