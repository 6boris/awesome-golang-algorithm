package Solution

func Solution(banned []int, n int, maxSum int) int {
	sum := 0
	bm := make(map[int]struct{})
	for _, x := range banned {
		if x >= 1 && x <= n {
			bm[x] = struct{}{}
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if _, ok := bm[i]; ok {
			continue
		}
		sum += i
		if sum > maxSum {
			break
		}
		ans++
	}
	return ans
}
