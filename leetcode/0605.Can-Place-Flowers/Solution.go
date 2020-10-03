package Solution

func Solution(flowerbed []int, n int) bool {
	l, count := len(flowerbed), 0
	for i := 0; i < l; i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == l-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
	}
	return count >= n
}
