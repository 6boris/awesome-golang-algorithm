package Solution

import "fmt"

func Solution(secret string, guess string) string {
	a, b := 0, 0
	bucket := make([]int, 10)
	missMatch := make([]int, 10)

	for idx := 0; idx < len(secret); idx++ {
		if secret[idx] == guess[idx] {
			a++
			continue
		}

		missMatch[secret[idx]-'0']++
		bucket[guess[idx]-'0']++
	}

	for idx := 0; idx < 10; idx++ {
		if bucket[idx] != 0 && missMatch[idx] != 0 {
			add := bucket[idx]
			if missMatch[idx] < add {
				add = missMatch[idx]
			}
			b += add
		}
	}
	ans := fmt.Sprintf("%dA%dB", a, b)
	return ans
}
