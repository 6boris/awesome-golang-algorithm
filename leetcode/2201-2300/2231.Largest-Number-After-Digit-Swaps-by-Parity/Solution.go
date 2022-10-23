package Solution

import "fmt"

func Solution(num int) int {
	odd := make([]byte, 10)
	oddPos := make([]int, 0)
	even := make([]byte, 10)
	evenPos := make([]int, 0)
	numString := fmt.Sprintf("%d", num)
	for idx, b := range numString {
		if b&1 == 1 {
			odd[b-'0']++
			oddPos = append(oddPos, idx)
			continue
		}
		even[b-'0']++
		evenPos = append(evenPos, idx)
	}

	ans := make([]byte, len(numString))
	if len(oddPos) > 0 {
		start := 0
		for idx := 9; idx >= 0; idx-- {
			for times := odd[idx]; times > 0; times-- {
				ans[oddPos[start]] = byte(idx + '0')
				start++
			}
		}
	}
	if len(evenPos) > 0 {
		start := 0
		for idx := 9; idx >= 0; idx-- {
			for times := even[idx]; times > 0; times-- {
				ans[evenPos[start]] = byte(idx + '0')
				start++
			}
		}
	}
	var n int
	fmt.Sscanf(string(ans), "%d", &n)
	return n
}
