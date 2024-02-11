package Solution

func Solution(nums []int, k int) int {
	zeros := make([]int, 0)
	ones := make(map[int]int)
	count := 0
	ans := 0
	for i, n := range nums {
		if n == 1 {
			count++
			if k == 0 {
				if count > ans {
					ans = count
				}
				continue
			}

			idx := 0
			if len(zeros) > k {
				idx = zeros[len(zeros)-k]
			}
			add := ones[idx]
			if diff := i - idx + 1; diff+add > ans {
				ans = diff + add
			}
		} else {
			if k == 0 {
				count = 0
				continue
			}
			if k == 1 {
				diff := 1
				if len(zeros) > 0 {
					diff = max(i - zeros[len(zeros)-1])
				}
				if diff > ans {
					ans = diff
				}
			} else {
				idx := 0
				if len(zeros) > k-1 {
					idx = zeros[len(zeros)-k+1]
				}
				add := ones[idx]
				if diff := i - idx + 1; diff+add > ans {
					ans = diff + add
				}
			}
			zeros = append(zeros, i)
			ones[i] = count
			count = 0
		}
	}
	return ans
}
