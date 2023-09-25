package Solution

func Solution(arr []int) int {
	l := len(arr)
	indies := make(map[int]int)
	for i, n := range arr {
		indies[n] = i
	}
	sum := make([]int, l)
	sum2 := make([]int, l)
	sum2[0] = arr[0]
	for i := 1; i < l; i++ {
		sum[i] = sum[i-1] + i
		sum2[i] = sum2[i-1] + arr[i]
	}
	group := 0
	target := l - 1
	targetIndex := indies[target] // 理想情况是是这样 5,4,3最小合

	for {
		if targetIndex == 0 {
			group++
			break
		}
		a := sum[target] - sum[targetIndex-1]
		b := sum2[target] - sum2[targetIndex-1]
		if a == b {
			group++
			target = targetIndex - 1
			targetIndex = indies[target]
		} else {
			exp := make(map[int]struct{})
			for i := target; i >= targetIndex; i-- {
				exp[arr[i]] = struct{}{}
			}
			which := -1
			for i := target; i >= targetIndex; i-- {
				if _, ok := exp[i]; !ok {
					if which == -1 || indies[i] < which {
						which = indies[i]
					}
					continue
				}
				if arr[i] < targetIndex {
					if which == -1 || indies[arr[i]] < which {
						which = indies[arr[i]]
					}
				}
			}

			targetIndex = which
		}
	}

	return group
}
