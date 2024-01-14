package Solution

func Solution(s string, k int) int {
	if k > len(s) {
		return 0
	}

	length := len(s)
	cache := make([][26]int, length)
	indies := [26][]int{}
	cache[0][s[0]-'a'] = 1
	indies[s[0]-'a'] = append(indies[s[0]-'a'], 0)

	for i := 1; i < length; i++ {
		indies[s[i]-'a'] = append(indies[s[i]-'a'], i)
		for j := 0; j < 26; j++ {
			cache[i][j] = cache[i-1][j]
		}
		cache[i][s[i]-'a']++
	}
	ans := 0
	for i := 0; i < 26; i++ {
		l := len(indies[i])
		if l < k {
			continue
		}
		for j := 0; j <= l-k; j++ {
			start := indies[i][j]
			end := length - 1
			for end >= indies[i][j+k-1] {
				cur := -1
				for kk := 0; kk < 26; kk++ {
					cut := 0
					if start > 0 {
						cut = cache[start-1][kk]
					}
					diff := cache[end][kk] - cut
					if diff == 0 || diff >= k {
						continue
					}
					at := indies[kk][cut]
					if cur == -1 || at < cut {
						cur = at
						if cur < indies[i][j+k-1] {
							cur = indies[i][j+k-1]
						}
					}
				}
				if cur == -1 {
					if ans == 0 || end-start+1 > ans {
						ans = end - start + 1
					}
					break
				}
				end = cur - 1
			}
		}
	}

	return ans
}
