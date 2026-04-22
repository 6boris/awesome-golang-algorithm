package Solution

func Solution(queries []string, dictionary []string) []string {
	var ret []string
	dict := make(map[string]struct{})
	for i := range dictionary {
		dict[dictionary[i]] = struct{}{}
	}
	for i := range queries {
		for j := range dictionary {
			diff := 0
			for k := range queries[i] {
				if queries[i][k] != dictionary[j][k] {
					diff++
					if diff > 2 {
						break
					}
				}
			}
			if diff <= 2 {
				ret = append(ret, queries[i])
				break
			}
		}
	}
	return ret
}
