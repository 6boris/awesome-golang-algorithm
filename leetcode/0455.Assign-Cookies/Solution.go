package Solution

import "sort"

//	先排序，然后直接贪心每一步即可
//	需求 g[i] 满足s[j]
func findContentChildren(g []int, s []int) int {
	//	排序准备使用贪心
	sort.Ints(g)
	sort.Ints(s)
	i := 0
	for j := 0; i < len(g) && j < len(s); j++ {
		//	第j个糖果可以满足第i个孩子
		if g[i] <= s[j] {
			i++
		}
		//	糖果已经检查完毕直接结束
		if j == len(s) {
			break
		}

	}
	return i
}
