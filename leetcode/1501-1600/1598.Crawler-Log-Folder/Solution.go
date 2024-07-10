package Solution

func Solution(logs []string) int {
	deep := 0
	for _, op := range logs {
		if op == "./" {
			continue
		}
		if op == "../" {
			deep = max(0, deep-1)
			continue
		}
		deep++
	}
	return deep
}
