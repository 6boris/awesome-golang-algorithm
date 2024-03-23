package Solution

import "sort"

type unionFinc2092 struct {
	father []int
}

func (u *unionFinc2092) find(x int) int {
	if u.father[x] != x {
		u.father[x] = u.find(u.father[x])
	}
	return u.father[x]
}

func (u *unionFinc2092) union(x, y int) {
	fx := u.find(x)
	fy := u.find(y)
	if fx < fy {
		u.father[fy] = fx
		return
	}
	u.father[fx] = fy
}

func Solution(n int, meetings [][]int, firstPerson int) []int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	u := unionFinc2092{father: make([]int, n)}
	for i := 0; i < n; i++ {
		u.father[i] = i
	}
	u.father[firstPerson] = 0
	startTime := -1
	for cur, meet := range meetings {
		if meet[2] == startTime {
			u.union(meet[0], meet[1])
			continue
		}
		for i := cur - 1; i >= 0; i-- {
			if meetings[i][2] != startTime {
				break
			}
			fx := u.find(meetings[i][0])
			fy := u.find(meetings[i][1])
			if fx == 0 || fy == 0 {
				u.union(meetings[i][0], meetings[i][1])
			} else {
				u.father[meetings[i][0]] = meetings[i][0]
				u.father[meetings[i][1]] = meetings[i][1]
			}
		}

		startTime = meet[2]
		u.union(meet[0], meet[1])
	}
	for i := len(meetings) - 1; i >= 0; i-- {
		if meetings[i][2] != startTime {
			break
		}
		fx := u.find(meetings[i][0])
		fy := u.find(meetings[i][1])
		if fx == 0 || fy == 0 {
			u.union(meetings[i][0], meetings[i][1])
		} else {
			u.father[meetings[i][0]] = meetings[i][0]
			u.father[meetings[i][1]] = meetings[i][1]
		}
	}
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		if u.find(i) == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}
