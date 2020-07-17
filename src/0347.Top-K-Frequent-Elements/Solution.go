package Solution

import "sort"

func topKFrequent(nums []int, k int) []int {
	tmp := make(map[int]int)
	res := []int{}
	s := [][]int{}
	//	统计每个字符出现次数
	for _, v := range nums {
		tmp[v]++
	}
	//	保存 数字/次数 映射  i,v 数字，次数
	for i, v := range tmp {
		s = append(s, []int{i, v})
	}
	//	按照次数排序
	sort.Slice(s, func(a, b int) bool {
		return s[a][1] > s[b][1]
	})
	//	截取前k个
	for i := 0; i < k; i++ {
		res = append(res, s[i][0])
	}
	return res
}
