package Solution

func Solution(s string) int {
	bs := []byte(s)
	stackIndex := -1
	for index := 0; index < len(s); index++ {
		if s[index] == ']' {
			if stackIndex != -1 && bs[stackIndex] == '[' {
				stackIndex--
				continue
			}
		}
		stackIndex++
		bs[stackIndex] = s[index]
	}
	if stackIndex == -1 {
		return 0
	}
	// 最终就是]]]]]][[[[ 这情况
	count := (stackIndex + 1) / 2
	need := count / 2
	if count&1 == 1 {
		need++
	}
	return need
}
