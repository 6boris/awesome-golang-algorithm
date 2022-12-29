package Solution

import (
	"container/heap"
	"sort"
)

/*
这道题开始是这样定义的结构体, 想用Pop后，更新costTime，然后在Fix，后面发现Fix的不太对

	type task1 struct {
		costTime
		tasks []task
	}
*/
type task struct {
	index int
	cost  []int
}
type taskHeap []task

func (t *taskHeap) Len() int {
	return len(*t)
}

func (t *taskHeap) Less(i, j int) bool {
	a := (*t)[i]
	b := (*t)[j]
	if a.cost[1] == b.cost[1] {
		return a.index < b.index
	}
	return a.cost[1] < b.cost[1]
}

func (t *taskHeap) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t *taskHeap) Push(x interface{}) {
	*t = append(*t, x.(task))
}

func (t *taskHeap) Pop() interface{} {
	old := *t
	l := len(old)
	x := old[l-1]
	*t = old[:l-1]
	return x
}

func Solution(tasks [][]int) []int {
	tl := len(tasks)
	ans := make([]int, tl)
	taskList := make([]task, tl)
	for i := 0; i < len(tasks); i++ {
		taskList[i] = task{index: i, cost: tasks[i]}
	}
	h := taskHeap{}
	sort.Slice(taskList, func(i, j int) bool {
		return taskList[i].cost[0] < taskList[j].cost[0]
	})

	// wait first task
	costTime, done, i := 0, 0, 0
	for done < tl {
		if taskList[done].cost[0] > costTime {
			costTime = taskList[done].cost[0]
		}
		for ; done < tl && taskList[done].cost[0] <= costTime; done++ {
			heap.Push(&h, taskList[done])
		}

		for h.Len() > 0 && done < tl && costTime < taskList[done].cost[0] {
			x := heap.Pop(&h).(task)
			ans[i] = x.index
			i++
			costTime += x.cost[1]
		}
	}
	for h.Len() > 0 {
		x := heap.Pop(&h).(task)
		ans[i] = x.index
		i++
	}
	return ans
}
