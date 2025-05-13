package Solution

const mod = 1000000007

func Solution(s string, t int) int {
	count := [26]int{}
	for _, b := range s {
		count[b-'a']++
	}
	for ; t > 0; t-- {
		tmp := [26]int{}
		// a, b, c, d .... z
		// count是前一轮的变化
		z := count[25]
		for i := 25; i > 0; i-- {
			tmp[i] = count[i-1]
		}
		tmp[0] = z
		tmp[1] = (tmp[1] + z) % mod
		for i := range 26 {
			count[i] = tmp[i]
		}
	}
	ans := 0
	for i := range 26 {
		ans = (ans + count[i]) % mod
	}
	return ans
}
