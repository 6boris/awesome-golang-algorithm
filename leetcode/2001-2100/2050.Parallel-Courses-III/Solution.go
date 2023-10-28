package Solution

import "container/heap"

type courses [][2]int

func (c *courses) Len() int {
	return len(*c)
}

func (c *courses) Less(i, j int) bool {
	return (*c)[i][1] < (*c)[j][1]
}

func (c *courses) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func (c *courses) Push(x interface{}) {
	*c = append(*c, x.([2]int))
}

func (c *courses) Pop() interface{} {
	old := *c
	l := len(old)
	x := old[l-1]
	*c = old[:l-1]
	return x
}

func Solution(n int, relations [][]int, time []int) int {
	cost := 0
	in := make(map[int]int)
	unlockCourses := make(map[int][]int)
	for _, rel := range relations {
		pre, next := rel[0], rel[1]
		in[next]++
		if _, ok := unlockCourses[pre]; !ok {
			unlockCourses[pre] = make([]int, 0)
		}
		unlockCourses[pre] = append(unlockCourses[pre], next)
	}
	queue := courses{}
	for i := 1; i <= n; i++ {
		if in[i] == 0 {
			heap.Push(&queue, [2]int{i, time[i-1]})
		}
	}
	for len(queue) > 0 {
		top := heap.Pop(&queue).([2]int)
		endCourse := make(map[int]struct{})
		endCourse[top[0]] = struct{}{}
		for len(queue) > 0 {
			x := heap.Pop(&queue).([2]int)
			if x[1] != top[1] {
				heap.Push(&queue, x)
				break
			}
			endCourse[x[0]] = struct{}{}
		}
		for i := 0; i < len(queue); i++ {
			queue[i][1] -= top[1]
		}

		cost += top[1]
		for ec := range endCourse {
			for _, next := range unlockCourses[ec] {
				in[next]--
				if in[next] == 0 {
					heap.Push(&queue, [2]int{next, time[next-1]})
				}
			}
		}
	}
	return cost
}
