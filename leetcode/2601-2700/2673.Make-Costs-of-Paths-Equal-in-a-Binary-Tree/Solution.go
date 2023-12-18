package Solution

func findBaseN(n int) int {
	x := 0
	for n > 0 {
		n >>= 1
		x++
	}
	return x
}

func Solution(n int, cost []int) int {
	dist := make([]int, len(cost))
	dist[0] = cost[0]
	for start := 0; start < n/2; start++ {
		dist[start*2+1] = dist[start] + cost[start*2+1]
		dist[start*2+2] = dist[start] + cost[start*2+2]

	}

	maxPathSum := 0
	for i := n / 2; i < n; i++ {
		if dist[i] > maxPathSum {
			maxPathSum = dist[i]
		}
	}
	baseN := findBaseN(n)

	ans := 0
	start, end := n/2, n
	for i := start; i < end; i++ {
		dist[i] = maxPathSum - dist[i]
	}
	baseN--
	start, end = start/2, start
	for baseN > 0 {
		for i := start; i < end; i++ {
			left, right := i*2+1, i*2+2
			dl := dist[left]
			dr := dist[right]
			minDiff := dl
			add := dr - dl
			if minDiff > dr {
				minDiff = dr
				add = dl - dr
			}
			ans += add
			dist[i] = minDiff
		}
		start, end = start/2, start
		baseN--
	}
	return ans
}
