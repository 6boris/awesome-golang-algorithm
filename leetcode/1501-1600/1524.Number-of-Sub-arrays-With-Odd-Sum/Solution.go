package Solution

const mod1524 = 1000000007

func Solution(arr []int) int {
	ans := 0
	e, o := 1, 0
	sum := 0
	for _, n := range arr {
		sum += n
		if sum&1 == 0 {
			ans = (ans + o) % mod1524
			e++
		} else {
			ans = (ans + e) % mod1524
			o++
		}
	}
	return ans
}
