package Solution

func Solution(n int, edges [][]int) []int {
	selectNodes := make([]bool, n)
	for _, edge := range edges {
		selectNodes[edge[1]] = true
	}
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		if !selectNodes[i] {
			ans = append(ans, i)
		}
	}
	return ans
}
