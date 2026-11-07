package Solution

import "container/heap"

type free1882 struct {
	weight, sourceIndex int
}
type used1882 struct {
	sourceIndex, end int
}

type usedList []*used1882
type freeList []*free1882

func (f *freeList) Len() int {
	return len(*f)
}

func (f *freeList) Less(i, j int) bool {
	a, b := (*f)[i], (*f)[j]
	if a.weight == b.weight {
		return a.sourceIndex < b.sourceIndex
	}
	return a.weight < b.weight
}

func (f *freeList) Swap(i, j int) {
	(*f)[i], (*f)[j] = (*f)[j], (*f)[i]
}

func (f *freeList) Push(x any) {
	(*f) = append((*f), x.(*free1882))
}

func (f *freeList) Pop() any {
	n := len(*f)
	x := (*f)[n-1]
	(*f) = (*f)[:n-1]
	return x
}

func (u *usedList) Len() int {
	return len(*u)
}

func (u *usedList) Less(i, j int) bool {
	a, b := (*u)[i], (*u)[j]
	return a.end < b.end
}

func (u *usedList) Swap(i, j int) {
	(*u)[i], (*u)[j] = (*u)[j], (*u)[i]
}

func (u *usedList) Push(x any) {
	(*u) = append((*u), x.(*used1882))
}

func (u *usedList) Pop() any {
	n := len(*u)
	x := (*u)[n-1]
	(*u) = (*u)[:n-1]
	return x
}

func Solution(servers []int, tasks []int) []int {
	m := len(tasks)
	ret := make([]int, m)
	free := &freeList{}
	used := &usedList{}

	for i := range servers {
		heap.Push(free, &free1882{
			weight:      servers[i],
			sourceIndex: i,
		})
	}

	cur := 0
	for i := 0; i < m; i++ {
		cur = max(cur, i)
		for used.Len() > 0 {
			top := (*used)[0]
			if top.end > cur {
				break
			}
			heap.Pop(used)
			heap.Push(free, &free1882{
				sourceIndex: top.sourceIndex,
				weight:      servers[top.sourceIndex],
			})
		}

		// 2. 如果此时没有空闲服务器，时间需要快进到最早有服务器释放的时刻
		if free.Len() == 0 {
			cur = (*used)[0].end
			for used.Len() > 0 {
				top := (*used)[0]
				if top.end > cur {
					break
				}
				heap.Pop(used)
				heap.Push(free, &free1882{
					sourceIndex: top.sourceIndex,
					weight:      servers[top.sourceIndex],
				})
			}
		}

		pop := heap.Pop(free).(*free1882)
		ret[i] = pop.sourceIndex
		heap.Push(used, &used1882{
			sourceIndex: pop.sourceIndex,
			end:         cur + tasks[i],
		})
	}
	return ret
}
