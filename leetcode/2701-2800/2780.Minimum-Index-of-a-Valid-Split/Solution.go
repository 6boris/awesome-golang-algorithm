package Solution

func Solution(nums []int) int {
	l := len(nums)
	leftCnt := make(map[int]int)
	left := make([]int, l)
	m := -1
	for i, n := range nums {
		left[i] = -1
		leftCnt[n]++
		if m == -1 {
			m = n
			left[i] = m
			continue
		}
		selected := m
		if leftCnt[n] > leftCnt[m] {
			selected = n
		}

		if leftCnt[selected] > (i+1)/2 {
			left[i] = selected
		}
		m = selected
	}

	rightCnt := make(map[int]int)
	right := make([]int, l)
	m = -1
	for i := l - 1; i >= 0; i-- {
		n := nums[i]
		right[i] = -1
		rightCnt[n]++
		if m == -1 {
			m = n
			right[i] = m
			continue
		}
		selected := m
		if rightCnt[n] > rightCnt[m] {
			selected = n
		}
		if rightCnt[selected] > (l-i)/2 {
			right[i] = selected
		}
		m = selected
	}
	for i := 0; i < l-1; i++ {
		if left[i] == -1 || right[i] == -1 {
			continue
		}
		if left[i] == right[i+1] {
			return i
		}
	}
	return -1
}
