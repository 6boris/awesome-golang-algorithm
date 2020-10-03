package Solution

func Solution(x bool) bool {
	return x
}

func isSubsequence(sub, start string) bool {
	startSub := 0
	startMain := 0

	for startMain < len(start) && startSub < len(sub) {
		if start[startMain] == sub[startSub] {
			startSub++
		}
		startMain++
	}

	return startSub == len(sub)
}
