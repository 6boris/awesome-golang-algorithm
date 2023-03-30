package Solution

type windows struct {
	start, end, sum int
}

func Solution(nums []int, firstLen int, secondLen int) int {
	ans := 0
	tsum, start := 0, 0
	firstWindows := make([]windows, 0)
	for i := 0; i < len(nums); i++ {
		tsum += nums[i]
		if i == firstLen-1 {
			firstWindows = append(firstWindows, windows{start, i, tsum})
			continue
		} else if i >= firstLen {
			tsum -= nums[start]
			start++
			firstWindows = append(firstWindows, windows{start, i, tsum})
		}
	}
	tsum, start = 0, 0
	secondWindows := make([]windows, 0)
	for i := 0; i < len(nums); i++ {
		tsum += nums[i]
		if i == secondLen-1 {
			secondWindows = append(secondWindows, windows{start, i, tsum})
			continue
		} else if i >= secondLen {
			tsum -= nums[start]
			start++
			secondWindows = append(secondWindows, windows{start, i, tsum})
		}
	}

	for _, first := range firstWindows {
		for _, second := range secondWindows {
			if (first.end < second.start || second.end < first.start) && first.sum+second.sum > ans {
				ans = first.sum + second.sum
			}
		}
	}
	return ans
}
