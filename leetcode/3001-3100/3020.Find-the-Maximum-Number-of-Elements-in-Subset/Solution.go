package Solution

func Solution(nums []int) int {
	cnt := make(map[int]int)
	for i := range nums {
		cnt[nums[i]]++
	}

	var (
		ret int = 1
		tmp int = 0
	)
	skip := make(map[int]struct{})
	for ele := range cnt {
		if _, ok := skip[ele]; ok {
			continue
		}

		tmp = 0
		if ele == 1 {
			if cnt[1]&1 == 0 {
				cnt[1]--
			}
			ret = max(ret, cnt[1])
			skip[1] = struct{}{}
			continue
		}
		for cnt[ele] >= 2 {
			skip[ele] = struct{}{}
			ele *= ele
			tmp += 2
		}
		if cnt[ele] == 0 {
			tmp--
		} else {
			skip[ele] = struct{}{}
			tmp++
		}
		ret = max(ret, tmp)
	}
	return ret
}
