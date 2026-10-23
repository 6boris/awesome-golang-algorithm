package Solution

func Solution(maxChoosableInteger int, desiredTotal int) bool {
	sum := (1 + maxChoosableInteger) * maxChoosableInteger / 2
	if sum < desiredTotal {
		return false
	}
	if desiredTotal <= 0 {
		return true
	}

	// memo 用于记录不同状态下的胜负结果
	// key: 状态掩码（state），value: 当前状态下先手是否必胜
	memo := make(map[int]bool)
	var dfs func(state int, currentTotal int) bool
	dfs = func(state int, currentTotal int) bool {
		// 如果该状态已经计算过，直接返回缓存结果
		if res, exists := memo[state]; exists {
			return res
		}

		// 尝试选择每一个未被使用的数字（从 1 到 maxChoosableInteger）
		for i := 1; i <= maxChoosableInteger; i++ {
			// 判断第 i 个数字是否已经被使用过（对应二进制第 i-1 位）
			mask := 1 << (i - 1)
			if (state & mask) == 0 {
				// 如果当前玩家选择数字 i 能够直接达到或超过 desiredTotal，或者
				// 递归下去让对手面临必败状态（即 dfs 返回 false），那么当前玩家就赢了
				if currentTotal+i >= desiredTotal || !dfs(state|mask, currentTotal+i) {
					memo[state] = true
					return true
				}
			}
		}

		// 如果遍历了所有可能，自己都无法获胜，说明当前状态必败
		memo[state] = false
		return false
	}

	// 初始状态：所有数字都未被使用（state = 0），当前累加和为 0
	return dfs(0, 0)
}
