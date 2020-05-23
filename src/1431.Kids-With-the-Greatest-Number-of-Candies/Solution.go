package Solution

func Solution(candies []int, extraCandies int) []bool {
	max := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	res := make([]bool, 0)
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}
