package Solution

const MOD = 1e9 + 7
const L = 26

func Solution(s string, t int, nums []int) int {
	T := NewMat()
	for i := 0; i < L; i++ {
		for j := 1; j <= nums[i]; j++ {
			T.a[(i+j)%L][i] = 1
		}
	}

	res := quickMul(T, t)
	f := make([]int, L)
	for _, ch := range s {
		f[ch-'a']++
	}

	ans := 0
	for i := 0; i < L; i++ {
		for j := 0; j < L; j++ {
			ans = (ans + res.a[i][j]*f[j]) % MOD
		}
	}
	return ans
}

type Mat struct {
	a [L][L]int
}

func NewMat() *Mat {
	return &Mat{}
}

func NewMatCopy(from *Mat) *Mat {
	m := &Mat{}
	for i := 0; i < L; i++ {
		for j := 0; j < L; j++ {
			m.a[i][j] = from.a[i][j]
		}
	}
	return m
}

func (m *Mat) Mul(other *Mat) *Mat {
	result := NewMat()
	for i := 0; i < L; i++ {
		for j := 0; j < L; j++ {
			for k := 0; k < L; k++ {
				result.a[i][j] = (result.a[i][j] + m.a[i][k]*other.a[k][j]) % MOD
			}
		}
	}
	return result
}

/* identity matrix */
func I() *Mat {
	m := NewMat()
	for i := 0; i < L; i++ {
		m.a[i][i] = 1
	}
	return m
}

/* matrix exponentiation by squaring */
func quickMul(x *Mat, y int) *Mat {
	ans := I()
	cur := NewMatCopy(x)
	for y > 0 {
		if y&1 == 1 {
			ans = ans.Mul(cur)
		}
		cur = cur.Mul(cur)
		y >>= 1
	}
	return ans
}
