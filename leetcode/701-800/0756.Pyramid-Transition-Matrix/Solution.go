package Solution

func Solution(bottom string, allowed []string) bool {
	allowedMap := make(map[[2]byte][]byte)
	for _, pattern := range allowed {
		key := [2]byte{pattern[0], pattern[1]}
		allowedMap[key] = append(allowedMap[key], pattern[2])
	}

	var canBuild func(string) bool
	canBuild = func(current string) bool {
		if len(current) == 1 {
			return true
		}

		var nextLayers []string

		var backtrack func(int, []byte)
		backtrack = func(index int, path []byte) {
			if index == len(current)-1 {
				nextLayers = append(nextLayers, string(path))
				return
			}
			pair := [2]byte{current[index], current[index+1]}
			if tops, exists := allowedMap[pair]; exists {
				for _, top := range tops {
					backtrack(index+1, append(path, top))
				}
			}
		}

		backtrack(0, []byte{})

		for _, nextLayer := range nextLayers {
			if canBuild(nextLayer) {
				return true
			}
		}
		return false
	}

	return canBuild(bottom)
}
