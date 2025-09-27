package Solution

import (
	"container/heap"
)

type task3408 struct {
	UserID, ID, Priority int
	index                int
}

type TaskList struct {
	items []*task3408
}

func (t *TaskList) Len() int {
	return len(t.items)
}
func (t *TaskList) Swap(i, j int) {
	t.items[i], t.items[j] = t.items[j], t.items[i]
	t.items[i].index = i
	t.items[j].index = j
}

func (t *TaskList) Less(i, j int) bool {
	a, b := t.items[i], t.items[j]
	if a.Priority == b.Priority {
		return a.ID > b.ID
	}
	return a.Priority > b.Priority
}

func (t *TaskList) Push(x any) {
	item := x.(*task3408)
	item.index = len(t.items)
	t.items = append(t.items, item)
}

func (t *TaskList) Pop() any {
	old := t.items
	l := len(old)
	x := old[l-1]
	t.items = old[:l-1]
	return x
}

type TaskManager struct {
	taskHeap  *TaskList
	taskIndex map[int]*task3408
}

func Constructor(tasks [][]int) TaskManager {
	th := &TaskList{items: make([]*task3408, 0)}
	tm := TaskManager{taskHeap: th, taskIndex: make(map[int]*task3408)}
	// userId, taskId, priority
	for _, task := range tasks {
		tm.Add(task[0], task[1], task[2])
	}
	return tm
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	item := &task3408{
		UserID:   userId,
		ID:       taskId,
		Priority: priority,
	}
	this.taskIndex[taskId] = item
	heap.Push(this.taskHeap, item)
	//this.taskHeap.Push(item)
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	item := this.taskIndex[taskId]
	item.Priority = newPriority
	heap.Fix(this.taskHeap, item.index)
}

func (this *TaskManager) Rmv(taskId int) {
	item := this.taskIndex[taskId]
	delete(this.taskIndex, taskId)
	heap.Remove(this.taskHeap, item.index)
}

func (this *TaskManager) ExecTop() int {
	if len(this.taskHeap.items) == 0 {
		return -1
	}
	top := heap.Pop(this.taskHeap).(*task3408)
	return top.UserID
}

type op struct {
	name                     string
	userId, taskId, priority int
}

func Solution(input [][]int, ops []op) []int {
	c := Constructor(input)
	var ret []int
	for _, o := range ops {
		if o.name == "add" {
			c.Add(o.userId, o.taskId, o.priority)
			continue
		}
		if o.name == "edit" {
			c.Edit(o.taskId, o.priority)
			continue
		}
		if o.name == "exec" {
			ret = append(ret, c.ExecTop())
			continue
		}
		if o.name == "rmv" {
			c.Rmv(o.taskId)
		}
	}
	return ret
}
