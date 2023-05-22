package Solution

import (
	"container/heap"
	"sort"
)

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

type fq struct {
	v, c int
}

type fqs []fq

func (f *fqs) Len() int {
	return len(*f)
}

func (f *fqs) Less(i, j int) bool {
	return (*f)[i].c > (*f)[j].c
}

func (f *fqs) Swap(i, j int) {
	(*f)[i], (*f)[j] = (*f)[j], (*f)[i]
}

func (f *fqs) Push(x interface{}) {
	*f = append(*f, x.(fq))
}

func (f *fqs) Pop() interface{} {
	old := *f
	l := len(old)
	x := old[l-1]
	*f = old[:l-1]
	return x
}

func topKFrequent1(nums []int, k int) []int {
	keys := make(map[int]*fq)
	for _, n := range nums {
		if _, ok := keys[n]; !ok {
			keys[n] = &fq{v: n}
		}
		keys[n].c++
	}
	h := fqs([]fq{})
	for _, f := range keys {
		heap.Push(&h, *f)
	}
	ans := make([]int, k)
	index := 0
	for k > 0 && h.Len() > 0 {
		x := heap.Pop(&h).(fq)
		ans[index] = x.v
		index++
		k--
	}
	return ans
}
