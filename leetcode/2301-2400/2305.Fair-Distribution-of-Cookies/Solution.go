package Solution

// 这道题看数据量很小，尝试了一下暴力算法，还行，过了
func Solution(cookies []int, k int) int {
	ans := -1
	var lookup func(int, []int)
	lookup = func(index int, child []int) {
		if index == len(cookies) {
			m := -1
			for i := 0; i < k; i++ {
				if child[i] > m {
					m = child[i]
				}
			}
			if ans == -1 || ans > m {
				ans = m
			}
			return
		}
		for i := 0; i < k; i++ {
			// 第i个孩子得到这个饼干
			child[i] += cookies[index]
			lookup(index+1, child)
			child[i] -= cookies[index]
		}
	}
	lookup(0, make([]int, k))
	return ans
}
