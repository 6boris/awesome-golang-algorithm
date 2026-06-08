package Solution

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func Solution(employees []*Employee, id int) int {
	id2employe := make(map[int]*Employee)
	for idx := range employees {
		id2employe[employees[idx].Id] = employees[idx]
	}
	importance := 0
	queue := []int{id}
	for len(queue) > 0 {
		nq := make([]int, 0)
		for _, item := range queue {
			importance += id2employe[item].Importance
			nq = append(nq, id2employe[item].Subordinates...)
		}
		queue = nq
	}
	return importance
}
