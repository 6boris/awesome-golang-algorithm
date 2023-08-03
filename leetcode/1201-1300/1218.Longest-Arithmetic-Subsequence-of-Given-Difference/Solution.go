package Solution

func Solution(arr []int, difference int) int {
	l := len(arr)
	cache := make(map[int]int)
	cache[arr[0]] = 1

	ans := 1
	// 4,12,10,0,-2,7,-8,9,-9,-12,-12,8,8
	// diff = 0
	for i := 1; i < l; i++ {
		target := arr[i] - difference
		if v, ok := cache[target]; !ok {
			cache[arr[i]] = 1
		} else if v+1 > cache[arr[i]] {
			cache[arr[i]] = v + 1
		}
		if cache[arr[i]] > ans {
			ans = cache[arr[i]]
		}
	}
	return ans
}
