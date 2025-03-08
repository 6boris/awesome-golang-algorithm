package Solution

func Solution(blocks string, k int) int {
	w, b := 0, 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			w++
			continue
		}
		b++
	}
	ans := w
	start, end := 0, k
	for ; end < len(blocks); end++ {
		if blocks[end] == 'W' {
			w++
		} else {
			b++
		}
		if blocks[start] == 'W' {
			w--
		} else {
			b--
		}
		ans = min(ans, w)
		start++
	}
	return ans

}
