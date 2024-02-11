package Solution

func Solution(n int) int {
	if n <= 3 {
		return 1
	}
	if n == 4 {
		return 2
	}
	if n <= 6 {
		return 3
	}
	ans := 3
	cur := 2
	index := 4
	buf := []byte("122112")

	//1 22 11 2 1 22 1 22 |11 2 11 22
	//1 2  2  1 1 2  1 2   2  1 2  2
	// 122112   1 22 1 22 11
	for len(buf) < n {
		cnt := buf[index] - '0'
		index++
		if cur == 2 {
			for ; cnt > 0 && len(buf) < n; cnt-- {
				buf = append(buf, '1')
				ans++
			}
			cur = 1
			continue
		}
		if cur == 1 {
			for ; cnt > 0 && len(buf) < n; cnt-- {
				buf = append(buf, '2')
			}
			cur = 2
			continue
		}
	}
	return ans
}
