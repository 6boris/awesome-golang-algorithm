package Solution

func Solution(nums []int, k int) int {
	count := make(map[int][]int)
	for i, n := range nums {
		count[n] = append(count[n], i)
	}
	ans := 0
	for _, list := range count {
		l := len(list)
		for i := 0; i < l-1; i++ {
			for j := i + 1; j < l; j++ {
				if (list[i]*list[j])%k == 0 {
					ans++
				}
			}
		}
	}
	return ans
}
