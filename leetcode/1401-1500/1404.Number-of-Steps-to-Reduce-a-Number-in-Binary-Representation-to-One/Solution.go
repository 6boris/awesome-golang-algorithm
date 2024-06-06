package Solution

func Solution(s string) int {
	bs := []byte(s)
	steps := 0
	idx := len(s) - 1
	for idx >= 0 {
		if string(bs[:idx+1]) == "1" {
			break
		}
		if bs[idx] == '0' {
			idx--
			steps++
			continue
		}
		steps += 1
		for ; idx >= 0 && bs[idx] != '0'; idx-- {
			steps++
		}
		if idx >= 0 {
			bs[idx] = '1'
		}
	}

	return steps
}
