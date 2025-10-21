package Solution

func Solution(nums []int, k int, numOperations int) int {
	const N = 100000 + 2

	freq := make([]int, N)
	sweep := make([]int, N)

	mm := N
	MM := 0

	for _, x := range nums {
		if x >= N {
			panic("Value in nums exceeds the array limit N")
		}
		freq[x]++

		x0 := max(1, x-k)

		xN := min(x+k+1, N-1)

		sweep[x0]++
		sweep[xN]--

		mm = min(mm, x0)
		MM = max(MM, xN)
	}

	ans := 0
	cnt := 0

	for x := mm; x <= MM; x++ {
		cnt += sweep[x]

		max_addable := min(cnt-freq[x], numOperations)
		ans = max(ans, freq[x]+max_addable)
	}

	return ans
}
