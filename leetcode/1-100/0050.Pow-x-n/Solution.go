package Solution

/*
   X^27
   = X * X ^26
   = X * (X^2)^13
   = X * (X^2) * (X^2)^12
   = X * (X^2) * (X^4)^6
   = X * (X^2) * (X^8)^3
   = X * (X^2) * (X^8) * (X^8)^2
   = X * (X^2) * (X^8) * (X^16)
*/
func myPow(x float64, n int) float64 {
	//	判断递归结束
	if n == 0 {
		return 1
	}
	//	负数:直接求导
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	//	奇数
	if n%2 == 1 {
		return x * myPow(x, n-1)
	}
	//	偶数
	return myPow(x*x, n/2)
}

func myPow2(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	pow := 1.0
	for n != 0 {
		if n&1 != 0 {
			pow *= x
		}
		x *= x
		n >>= 1
	}
	return pow
}
