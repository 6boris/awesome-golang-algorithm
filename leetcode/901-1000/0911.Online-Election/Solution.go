package Solution

import (
	"sort"
)

type TopVotedCandidate struct {
	matrix [][]int
	mm     []int
	times  []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	l := len(times)

	matrix := make([][]int, l)
	for i := 0; i < l; i++ {
		matrix[i] = make([]int, l)
	}
	mm := make([]int, l)
	voteTime := make([]int, l)
	for i := range l {
		voteTime[i] = -1
	}
	matrix[0][persons[0]]++
	mm[0] = persons[0]
	voteTime[persons[0]] = times[0]

	for i := 1; i < l; i++ {
		matrix[i][persons[i]]++
		voteTime[persons[i]] = times[i]

		selectuser := 0
		maxvalue := 0
		for c := range matrix[i] {
			matrix[i][c] += matrix[i-1][c]
			if matrix[i][c] == maxvalue && maxvalue != 0 {
				if voteTime[c] > voteTime[selectuser] {
					selectuser = c
				}
			}

			if matrix[i][c] > maxvalue {
				selectuser = c
				maxvalue = matrix[i][c]
			}
		}
		mm[i] = selectuser
	}
	return TopVotedCandidate{matrix: matrix, times: times, mm: mm}
}

func (this *TopVotedCandidate) Q(t int) int {
	l := len(this.times)
	index := sort.Search(l, func(i int) bool {
		return this.times[i] >= t
	})
	if index == l || this.times[index] != t {
		// 没有
		index--
	}
	return this.mm[index]
}

func Solution(persons, times []int, ops []int) []int {
	ans := make([]int, len(ops))
	c := Constructor(persons, times)
	for i := range ops {
		ans[i] = c.Q(ops[i])
	}
	return ans
}
