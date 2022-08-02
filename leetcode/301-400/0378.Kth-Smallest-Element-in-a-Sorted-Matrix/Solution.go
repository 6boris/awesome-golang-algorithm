package Solution

// 不是最优解法
func Solution(matrix [][]int, k int) int {
	length := len(matrix)

	ans := matrix[0]

	for row := 1; row < length; row++ {
		ans = kthMerge(ans, matrix[row])
	}
	return ans[k-1]
}

func kthMerge(left, right []int) []int {
	ll, lr := len(left), len(right)
	ans := make([]int, ll+lr)
	l, r, idx := 0, 0, -1
	for l < ll || r < lr {
		idx++
		if l >= ll {
			ans[idx] = right[r]
			r++
			continue
		}
		if r >= lr {
			ans[idx] = left[l]
			l++
			continue
		}
		if left[l] < right[r] {
			ans[idx] = left[l]
			l++
			continue
		}
		ans[idx] = right[r]
		r++
	}

	return ans
}
