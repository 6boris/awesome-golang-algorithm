package Solution

func Solution(req_skills []string, people [][]string) []int {
	skillOrder := make(map[string]int)
	targetSkill := [16]bool{}
	for idx, r := range req_skills {
		skillOrder[r] = idx
		targetSkill[idx] = true
	}

	l := len(people)
	dp := make(map[[16]bool][]int)
	for idx := 0; idx < l; idx++ {
		if len(people[idx]) == 0 {
			continue
		}

		self := [16]bool{}
		for _, str := range people[idx] {
			self[skillOrder[str]] = true
		}

		tmp := make(map[[16]bool][]int)
		tmp[self] = []int{idx}
		for v, path := range dp {
			if p1, ok := tmp[v]; !ok || len(p1) > len(path) {
				tmp[v] = path
			}
			tmpCopy := [16]bool{}
			// v = 1, 2, 3
			// self = 2, 3 -> v
			// self = 1, 2, 3 -> self(只有1个)
			// self = 1, 2, 3,4 -> self
			// self = 4,5 -> v + 1
			addSelf := false
			addV := false
			for i := 0; i < 16; i++ {
				if self[i] && !v[i] {
					addSelf = true
				}
				if v[i] && !self[i] {
					addV = true
				}
				tmpCopy[i] = self[i] || v[i]
			}
			var n []int
			if !addSelf && !addV {
				n = []int{idx}
			} else if addSelf && addV {
				x := make([]int, len(path))
				copy(x, path)
				x = append(x, idx)
				n = x
			} else if addSelf {
				n = []int{idx}
			} else if addV {
				n = path
			}

			if p1, ok := tmp[tmpCopy]; !ok || len(p1) > len(n) {
				tmp[tmpCopy] = n
			}
		}
		dp = tmp
	}
	return dp[targetSkill]
}
