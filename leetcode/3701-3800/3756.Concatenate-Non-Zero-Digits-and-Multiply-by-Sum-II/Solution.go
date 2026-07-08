package Solution

const mod3756 = 1_000_000_007

func Solution(s string, queries [][]int) []int {
	n := len(s)
	pow10 := make([]int64, n+1)
	pow10[0] = 1
	for i := 1; i <= n; i++ {
		pow10[i] = (pow10[i-1] * 10) % mod3756
	}

	pSum := make([]int64, n+1)
	pCnt := make([]int64, n+1)
	pVal := make([]int64, n+1)
	for i := range s {
		digit := int64(s[i] - '0')
		if digit != 0 {
			pSum[i+1] = pSum[i] + digit
			pCnt[i+1] = pCnt[i] + 1
			pVal[i+1] = (pVal[i]*10 + digit) % mod3756
			continue
		}
		pSum[i+1] = pSum[i]
		pCnt[i+1] = pCnt[i]
		pVal[i+1] = pVal[i]
	}

	ret := make([]int, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		sum := pSum[r+1] - pSum[l]
		nzCount := pCnt[r+1] - pCnt[l]
		if nzCount == 0 {
			ret[i] = 0
			continue
		}

		x := (pVal[r+1] - pVal[l]*pow10[nzCount]) % mod3756
		if x < 0 {
			x = (x + mod3756) % mod3756
		}
		ret[i] = int((x * (sum % mod3756)) % mod3756)
	}

	return ret
}
