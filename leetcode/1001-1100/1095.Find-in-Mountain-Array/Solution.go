package Solution

type MountainArray struct {
	data []int
}

func (m *MountainArray) get(index int) int {
	return m.data[index]
}
func (m *MountainArray) length() int {
	return len(m.data)
}
func Solution(target int, mountainArr *MountainArray) int {
	length := mountainArr.length()
	l, r := 0, length-1
	// 二分找到index
	top := 0
	for l < r {
		mid := l + (r-l)/2
		if mid == 0 || mid == length-1 {
			return -1
		}
		a := mountainArr.get(mid - 1)
		b := mountainArr.get(mid)
		c := mountainArr.get(mid + 1)
		if a < b && c < b {
			top = mid
			break
		}
		if a > b {
			r = mid
			continue
		}
		l = mid + 1
	}
	//fmt.Printf("find top: %d\n", top)
	// 二分搜索左侧

	var (
		sl func(int, int, int) int
		sr func(int, int, int) int
	)
	sl = func(left, right, target int) int {
		// 0, 1
		for left < right {
			mid := left + (right-left)/2
			if r := mountainArr.get(mid); r == target {
				return mid
			} else if r < target {
				left = mid + 1
			} else {
				right = mid
			}
		}
		return -1
	}
	sr = func(left, right, target int) int {
		for left < right {
			mid := left + (right-left)/2
			if r := mountainArr.get(mid); r == target {
				return mid
			} else if r > target {
				// 5, 4, 3, 2,1
				left = mid + 1
			} else {
				right = mid
			}
		}
		return -1

	}
	if r := sl(0, top+1, target); r != -1 {
		return r
	}
	if r := sr(top, length, target); r != -1 {
		return r
	}
	return -1
}
