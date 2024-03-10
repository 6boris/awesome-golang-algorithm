package Solution

func Solution(dominoes string) string {
	bs := []byte(dominoes)
	source := []byte(dominoes)
	l := len(dominoes)
	ok := true
	for ok {
		ok = false
		i := 0
		for i < l {
			if bs[i] == '.' {
				i++
				continue
			}
			if bs[i] == 'L' {
				if i == 0 {
					i++
					continue
				}
				if bs[i-1] == '.' {
					if i >= 2 && source[i-2] == 'R' {
						i++
						continue
					}
					ok = true
					bs[i-1] = 'L'
				}
				i++
				continue
			}
			if bs[i] == 'R' {
				if i == l-1 {
					i++
					continue
				}
				if bs[i+1] == '.' {
					if i <= l-3 && source[i+2] == 'L' {
						i++
						continue
					}

					ok = true
					bs[i+1] = 'R'
					i += 2
					continue
				}
				i++
			}
		}
		copy(source, bs)
	}
	return string(bs)
}
