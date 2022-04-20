package Solution

/*

1: 10
        末尾是0
2: 10 + 1*9 + 8*9 = 91
        末尾是0   末尾非0，
3: 91 + 1*9*8 +   9*8*8 = 739

4: 739 + 1*9*8*7 +  9*8*8*7 = 5275
*/
func Solution(n int) int {
	ans := make([]int, 9)
	ans[0], ans[1] = 1, 10
	base1, base2 := 1, 8
	for idx := 2; idx <= 8; idx++ {
		diff := 11 - idx
		base1 *= diff
		base2 *= diff
		ans[idx] = ans[idx-1] + base1 + base2
	}
	return ans[n]
}
