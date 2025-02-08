package Solution

import "sort"

type NumberContainers struct {
	indies map[int]int
	list   map[int][]int
}

func Constructor() NumberContainers {
	return NumberContainers{
		indies: map[int]int{},
		list:   map[int][]int{},
	}
}

func (this *NumberContainers) Change(index int, number int) {
	v, ok := this.indies[index]
	this.indies[index] = number
	if !ok {
		list := this.list[number]
		idx := sort.Search(len(list), func(i int) bool {
			return list[i] > index
		})
		if idx == len(list) {
			this.list[number] = append(this.list[number], index)
		} else {
			b := append([]int{index}, list[idx:]...)
			a := append(list[:idx], b...)
			this.list[number] = a
		}
		return
	}
	if v != number {
		list := this.list[v]
		idx := sort.Search(len(list), func(i int) bool {
			return list[i] >= index
		})
		this.list[v] = append(this.list[v][:idx], this.list[v][idx+1:]...)

		list = this.list[number]
		idx = sort.Search(len(list), func(i int) bool {
			return list[i] > index
		})
		if idx == len(list) {
			this.list[number] = append(this.list[number], index)
		} else {
			b := append([]int{index}, list[idx:]...)
			a := append(list[:idx], b...)
			this.list[number] = a
		}
	}

}

func (this *NumberContainers) Find(number int) int {
	v := this.list[number]
	if len(v) == 0 {
		return -1
	}
	return v[0]
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */

type input struct {
	name          string
	index, number int
}

func Solution(inputs []input) []int {
	c := Constructor()
	ans := make([]int, 0)
	for _, in := range inputs {
		if in.name == "change" {
			c.Change(in.index, in.number)
			continue
		}
		ans = append(ans, c.Find(in.number))
	}
	return ans
}
