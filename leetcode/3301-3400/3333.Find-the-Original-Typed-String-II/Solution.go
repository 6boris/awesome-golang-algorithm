package Solution

const mod = 1000000007

func Solution(word string, k int) int {
	n := len(word)
	cnt := 1
	var freq []int

	for i := 1; i < n; i++ {
		if word[i] == word[i-1] {
			cnt++
		} else {
			freq = append(freq, cnt)
			cnt = 1
		}
	}
	freq = append(freq, cnt)

	ans := 1
	for _, o := range freq {
		ans = ans * o % mod
	}

	if len(freq) >= k {
		return ans
	}

	f := make([]int, k)
	g := make([]int, k)
	f[0] = 1
	for i := range g {
		g[i] = 1
	}

	for i := 0; i < len(freq); i++ {
		f_new := make([]int, k)
		for j := 1; j < k; j++ {
			f_new[j] = g[j-1]
			if j-freq[i]-1 >= 0 {
				f_new[j] = (f_new[j] - g[j-freq[i]-1] + mod) % mod
			}
		}
		g_new := make([]int, k)
		g_new[0] = f_new[0]
		for j := 1; j < k; j++ {
			g_new[j] = (g_new[j-1] + f_new[j]) % mod
		}
		f, g = f_new, g_new
	}

	return (ans - g[k-1] + mod) % mod
}
