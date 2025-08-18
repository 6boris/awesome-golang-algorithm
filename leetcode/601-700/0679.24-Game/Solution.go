package Solution

import "math"

func Solution(cards []int) bool {
	nums := make([]float64, len(cards))
	for i, v := range cards {
		nums[i] = float64(v)
	}
	return dfs(nums)
}

// 递归函数：判断当前数字数组是否可以通过操作得到24
func dfs(nums []float64) bool {
	if len(nums) == 1 {
		// 终止条件：只有一个数字，判断是否接近24
		return math.Abs(nums[0]-24) < 1e-6
	}
	// 遍历所有两两组合
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// 取出两个数
			a, b := nums[i], nums[j]
			// 剩余的数字
			rest := make([]float64, 0, len(nums)-1)
			for k := 0; k < len(nums); k++ {
				if k != i && k != j {
					rest = append(rest, nums[k])
				}
			}

			// 所有可能的运算
			operations := [][]float64{
				{a + b},
				{a - b},
				{b - a},
				{a * b},
			}
			// 除法要判断除数不为零
			if math.Abs(b) > 1e-6 {
				operations = append(operations, []float64{a / b})
			}
			if math.Abs(a) > 1e-6 {
				operations = append(operations, []float64{b / a})
			}

			// 递归
			for _, op := range operations {
				nextNums := append(rest, op[0])
				if dfs(nextNums) {
					return true
				}
			}
		}
	}
	return false
}
