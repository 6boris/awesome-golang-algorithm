package Solution

func Solution(equations [][]string, values []float64, queries [][]string) []float64 {
	r := make(map[string]map[string]float64)
	for i, eq := range equations {
		a, b := eq[0], eq[1]
		if _, ok := r[a]; !ok {
			r[a] = make(map[string]float64)
		}
		if _, ok := r[b]; !ok {
			r[b] = make(map[string]float64)
		}
		r[a][b] = values[i]
		r[b][a] = float64(1) / values[i]
	}

	var dfs func(float64, *float64, string, string, map[string]struct{})
	dfs = func(i float64, ans *float64, a, b string, used map[string]struct{}) {
		if *ans != -1.0 {
			return
		}
		if a == b {
			*ans = i
			return
		}
		if _, ok := used[a]; ok {
			return
		}
		used[a] = struct{}{}
		for rel, v := range r[a] {
			dfs(i*v, ans, rel, b, used)
		}
	}
	ans := make([]float64, 0)
	for _, q := range queries {
		a, b := q[0], q[1]
		_, ok1 := r[a]
		_, ok2 := r[b]
		if !ok1 || !ok2 {
			ans = append(ans, -1.0)
			continue
		}
		if a == b {
			ans = append(ans, 1.0)
			continue
		}
		// dfs
		tmp := float64(-1.0)
		dfs(float64(1), &tmp, a, b, map[string]struct{}{})
		ans = append(ans, tmp)
	}
	return ans
}
