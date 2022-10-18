package Solution

import "fmt"

func Solution(n int) string {
	nString := fmt.Sprintf("%d", n)
	length := len(nString)
	dotCount := length / 3
	if length%3 == 0 {
		dotCount--
	}
	if dotCount == 0 {
		return nString
	}

	target := make([]byte, length+dotCount)
	loop, targetIdx := 0, length+dotCount-1

	for idx := length - 1; idx >= 0; idx-- {
		target[targetIdx] = nString[idx]
		targetIdx, loop = targetIdx-1, loop+1
		if loop == 3 && targetIdx >= 0 {
			target[targetIdx] = '.'
			targetIdx--
			loop = 0
		}
	}
	return string(target)
}
