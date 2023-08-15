package Solution

func Solution(numCourses int, prerequisites [][]int) []int {
	length := len(prerequisites)
	ans := make([]int, numCourses)
	if length == 0 {
		for i := 0; i < numCourses; i++ {
			ans[i] = i
		}
		return ans
	}

	in := make([]int, numCourses)
	dep := make(map[int][]int)
	for _, p := range prerequisites {
		wantLean, firstLearn := p[0], p[1]
		in[wantLean]++
		if _, ok := dep[firstLearn]; !ok {
			dep[firstLearn] = make([]int, 0)
		}
		dep[firstLearn] = append(dep[firstLearn], wantLean)
	}
	queue := make([]int, 0)
	left := numCourses
	index := 0
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			queue = append(queue, i)
			ans[index] = i
			index++
			left--
		}
	}
	for left > 0 {
		next := make([]int, 0)
		for _, done := range queue {
			for _, rel := range dep[done] {
				in[rel]--
				if in[rel] == 0 {
					next = append(next, rel)
					ans[index] = rel
					index++
					left--
				}
			}
		}
		if len(next) == 0 && left > 0 {
			return nil
		}
		queue = next
	}

	return ans
}
