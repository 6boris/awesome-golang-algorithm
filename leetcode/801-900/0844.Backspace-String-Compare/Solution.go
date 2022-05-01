package Solution

func Solution(s string, t string) bool {
	sbs := []byte(s)
	tbs := []byte(t)
	sidx := -1
	for _, b := range sbs {
		if b == '#' {
			if sidx != -1 {
				sidx--
			}
			continue
		}
		sidx++
		sbs[sidx] = b
	}

	tidx := -1
	for _, b := range tbs {
		if b == '#' {
			if tidx != -1 {
				tidx--
			}
			continue
		}

		tidx++
		tbs[tidx] = b
	}

	return string(sbs[:sidx+1]) == string(tbs[:tidx+1])
}
