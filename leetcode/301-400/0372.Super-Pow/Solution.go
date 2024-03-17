package Solution

// 因为a可以到2^31 所以直接计算数据会越界
// (a^b)%mod -> (a * a^b-1) % mod -> (a%mod)^b)%mod
func pow372(a, b, mod int) int {
	if b == 0 {
		return 1
	}
	//(a^b)%mod -> ((a%mod)^b)%mod
	return ((a % mod) * (pow372(a, b-1, mod)) % mod) % mod
}

// (a^b) % m的结果等于(((a % m)^b) % m)
// (a * b) % m的结果等于((a % m) * (b % m)) % m
// b的长度可以到2000，数字是无法表示的
func Solution(a int, b []int) int {
	l := len(b)
	if l == 0 {
		return 1
	}
	/*
			3, [1, 2, 3]
		   3^123 % 1337
			(3^120 * 3^3) % 1337 -> ((3^120%1337) * (3^3%1337))%1337

			然后计算(3^120)%1337 -> ((3^12)^10)%1337
			-> (( (3^12)^10 )%1337) * (3^3)%1337)%1337
			所以可以看到第一部分可以递归计算
	*/
	cp := b[:l-1]
	return (pow372(Solution(a, cp), 10, 1337) * pow372(a, b[l-1], 1337)) % 1337
}
