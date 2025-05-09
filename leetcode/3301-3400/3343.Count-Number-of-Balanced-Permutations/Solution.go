package Solution

const MOD = 1_000_000_007

func Solution(num string) int {
	tot, n := 0, len(num)
	cnt := make([]int, 10)
	for _, ch := range num {
		d := int(ch - '0')
		cnt[d]++
		tot += d
	}
	if tot%2 != 0 {
		return 0
	}

	target := tot / 2
	maxOdd := (n + 1) / 2
	comb := make([][]int, maxOdd+1)
	for i := range comb {
		comb[i] = make([]int, maxOdd+1)
		comb[i][i], comb[i][0] = 1, 1
		for j := 1; j < i; j++ {
			comb[i][j] = (comb[i-1][j] + comb[i-1][j-1]) % MOD
		}
	}

	f := make([][]int, target+1)
	for i := range f {
		f[i] = make([]int, maxOdd+1)
	}
	f[0][0] = 1

	psum, totSum := 0, 0
	for i := 0; i <= 9; i++ {
		/* Sum of the number of the first i digits */
		psum += cnt[i]
		/* Sum of the first i numbers */
		totSum += i * cnt[i]
		for oddCnt := min(psum, maxOdd); oddCnt >= max(0, psum-(n-maxOdd)); oddCnt-- {
			/* The number of bits that need to be filled in even numbered positions */
			evenCnt := psum - oddCnt
			for curr := min(totSum, target); curr >= max(0, totSum-target); curr-- {
				res := 0
				for j := max(0, cnt[i]-evenCnt); j <= min(cnt[i], oddCnt) && i*j <= curr; j++ {
					/* The current digit is filled with j positions at odd positions, and cnt[i] - j positions at even positions */
					ways := comb[oddCnt][j] * comb[evenCnt][cnt[i]-j] % MOD
					res = (res + ways*f[curr-i*j][oddCnt-j]%MOD) % MOD
				}
				f[curr][oddCnt] = res % MOD
			}
		}
	}

	return f[target][maxOdd]
}
