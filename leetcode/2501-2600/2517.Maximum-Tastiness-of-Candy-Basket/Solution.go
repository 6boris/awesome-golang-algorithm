package Solution

import "sort"

func canFormBasket(prices []int, k, diff int) bool {
	count := 1                // 选第一个
	lastSelected := prices[0] // 第一个元素
	for _, price := range prices[1:] {
		if price-lastSelected >= diff {
			count++
			lastSelected = price
			if count == k {
				return true
			}
		}
	}
	return false
}

func Solution(price []int, k int) int {
	sort.Ints(price)
	left, right := 0, price[len(price)-1]-price[0]
	result := 0

	for left <= right {
		mid := (left + right) / 2
		if canFormBasket(price, k, mid) {
			result = mid   // 记录当前满足条件的最大差值
			left = mid + 1 // 试图找到更大的差值
		} else {
			right = mid - 1 // 减小差值
		}
	}
	return result
}
