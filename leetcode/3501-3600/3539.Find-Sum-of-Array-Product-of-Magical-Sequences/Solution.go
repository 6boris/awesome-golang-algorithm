package Solution

import "math/bits"

func quickmul(x, y, mod int64) int64 {
	res, cur := int64(1), x%mod
	for y > 0 {
		if y&1 == 1 {
			res = res * cur % mod
		}
		y >>= 1
		cur = cur * cur % mod
	}
	return res
}

func Solution(m int, k int, nums []int) int {
	mod := int64(1000000007)
	fac := make([]int64, m+1)
	fac[0] = 1
	for i := 1; i <= m; i++ {
		fac[i] = fac[i-1] * int64(i) % mod
	}
	ifac := make([]int64, m+1)
	ifac[0] = 1
	ifac[1] = 1
	for i := 2; i <= m; i++ {
		ifac[i] = quickmul(int64(i), mod-2, mod)
	}
	for i := 2; i <= m; i++ {
		ifac[i] = ifac[i-1] * ifac[i] % mod
	}

	numsPower := make([][]int64, len(nums))
	for i := range nums {
		numsPower[i] = make([]int64, m+1)
		numsPower[i][0] = 1
		for j := 1; j <= m; j++ {
			numsPower[i][j] = numsPower[i][j-1] * int64(nums[i]) % mod
		}
	}

	f := make([][][][]int64, len(nums))
	for i := range nums {
		f[i] = make([][][]int64, m+1)
		for j := 0; j <= m; j++ {
			f[i][j] = make([][]int64, m*2+1)
			for p := 0; p <= m*2; p++ {
				f[i][j][p] = make([]int64, k+1)
			}
		}
	}

	for j := 0; j <= m; j++ {
		f[0][j][j][0] = numsPower[0][j] * ifac[j] % mod
	}
	for i := 0; i+1 < len(nums); i++ {
		for j := 0; j <= m; j++ {
			for p := 0; p <= m*2; p++ {
				for q := 0; q <= k; q++ {
					q2 := p%2 + q
					if q2 > k {
						break
					}
					for r := 0; r+j <= m; r++ {
						p2 := p/2 + r
						f[i+1][j+r][p2][q2] += f[i][j][p][q] * numsPower[i+1][r] % mod * ifac[r] % mod
						f[i+1][j+r][p2][q2] %= mod
					}
				}
			}
		}
	}

	res := int64(0)
	for p := 0; p <= m*2; p++ {
		for q := 0; q <= k; q++ {
			if bits.OnesCount(uint(p))+q == k {
				res = (res + f[len(nums)-1][m][p][q]*fac[m]%mod) % mod
			}
		}
	}
	return int(res)
}
