package Solution

func maxArea(height []int) int {
	max, l, r := 0, 0, len(height)-1
	for l < r {
		//	获取当前的坐标形成的面积
		/*
			1.Min(height[l], height[r]) : 获取当前左右最小高度【面积需要最小高】
			2. Min(height[l], height[r])*(r-l) : 获取当前可以形成的面积
			3.Max(max, Min(height[l], height[r])*(r-l)): 获取目前最大面积
		*/
		max = Max(max, Min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
