package Solution

import "sort"

type unionFind721 struct {
	father []int
}

func (u *unionFind721) find(x int) int {
	if u.father[x] != x {
		u.father[x] = u.find(u.father[x])
	}
	return u.father[x]
}
func (u *unionFind721) union(x, y int) {
	fx := u.find(x)
	fy := u.find(y)
	if fx < fy {
		u.father[fy] = fx
	} else {
		u.father[fx] = fy
	}

}
func Solution(accounts [][]string) [][]string {
	u := unionFind721{father: make([]int, len(accounts))}
	for i := 0; i < len(accounts); i++ {
		u.father[i] = i
	}
	email2index := make(map[string]int)
	tmp := make([]map[string]struct{}, len(accounts))
	for idx, emails := range accounts {
		temp := map[string]struct{}{}
		for i := 1; i < len(emails); i++ {
			temp[emails[i]] = struct{}{}
			if targetIndex, ok := email2index[emails[i]]; ok {
				u.union(idx, targetIndex)
				continue
			}
			email2index[emails[i]] = idx
		}
		tmp[idx] = temp
	}
	for i := range accounts {
		f := u.find(i)
		if f != i {
			for key := range tmp[i] {
				tmp[f][key] = struct{}{}
			}
		}
	}
	result := make([][]string, 0)
	for i := range accounts {
		f := u.find(i)
		if f == i {
			result = append(result, []string{accounts[i][0]})
			next := make([]string, 0)
			for key := range tmp[i] {
				next = append(next, key)
			}
			sort.Strings(next)
			idx := len(result) - 1
			result[idx] = append(result[idx], next...)
		}
	}
	return result
}
