package Solution

func Solution(nums []int) int {
	const MOD = 1000000007
	numCnt := make(map[int]int)
	numPartialCnt := make(map[int]int)

	for _, v := range nums {
		numCnt[v]++
	}

	ans := 0
	for _, v := range nums {
		target := v * 2
		lCnt := numPartialCnt[target]
		numPartialCnt[v]++
		rCnt := numCnt[target] - numPartialCnt[target]

		ans = (ans + lCnt*rCnt) % MOD
	}

	return ans
}
