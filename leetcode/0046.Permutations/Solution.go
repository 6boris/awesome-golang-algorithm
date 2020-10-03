package Solution

var ans [][]int
var visit []bool

func permute(nums []int) [][]int {

	ans = [][]int{}
	visit = make([]bool, len(nums))
	dfs(nums, []int{})
	return ans
}

func dfs(nums []int, curr []int) {
	if len(curr) == len(nums) {
		temp := make([]int, len(curr))
		copy(temp, curr)
		ans = append(ans, temp)
		return
	}

	for idx, num := range nums {
		if !visit[idx] {
			curr = append(curr, num)
			visit[idx] = true

			dfs(nums, curr)

			visit[idx] = false
			curr = curr[:len(curr)-1]
		}
	}
}
