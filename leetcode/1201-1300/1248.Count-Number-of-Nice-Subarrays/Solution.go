package Solution

type tmp1248 struct {
	left, right int
}

func Solution(nums []int, k int) int {
	ans := 0
	l := len(nums)
	indies := make([]tmp1248, 0)
	even := 0

	for i := 0; i < l; i++ {
		if nums[i]&1 == 0 {
			even++
			continue
		}
		// left can start with itself so, left = even+1
		indies = append(indies, tmp1248{left: even + 1})
		li := len(indies)
		if li > 1 {
			indies[li-2].right = even
		}
		even = 0
	}
	if len(indies) == 0 || k > len(indies) {
		return 0
	}
	li := len(indies)
	indies[li-1].right = even

	start := 0
	for ; start <= len(indies)-k; start++ {
		ans += indies[start].left * (indies[start+k-1].right + 1)
	}
	return ans
}
