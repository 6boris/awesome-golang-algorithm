package Solution

type item624 struct {
	v, i int
}

func Solution(arrays [][]int) int {
	a, b := [2]item624{{arrays[0][0], 0}, {arrays[1][0], 1}}, [2]item624{{arrays[0][len(arrays[0])-1], 0}, {arrays[1][len(arrays[1])-1], 1}}
	if a[0].v < a[1].v {
		a[0], a[1] = a[1], a[0]
	}
	if b[0].v > b[1].v {
		b[0], b[1] = b[1], b[0]
	}
	for i := 2; i < len(arrays); i++ {
		arr := arrays[i]
		l := len(arr)
		if arr[0] <= a[1].v {
			a[0] = a[1]
			a[1] = item624{v: arr[0], i: i}
		} else if arr[0] < a[0].v {
			a[0] = item624{v: arr[0], i: i}
		}

		if arr[l-1] >= b[1].v {
			b[0] = b[1]
			b[1] = item624{v: arr[l-1], i: i}
		} else if arr[l-1] > b[0].v {
			b[0] = item624{v: arr[l-1], i: i}
		}
	}
	// 最大的情况
	if b[1].i != a[1].i {
		return b[1].v - a[1].v
	}
	ans := 0
	if b[1].i != a[0].i {
		ans = max(ans, b[1].v-a[0].v)
	}
	if b[0].i != a[1].i {
		ans = max(ans, b[0].v-a[1].v)
	}
	if b[0].i != a[0].i {
		ans = max(ans, b[0].v-a[0].v)
	}

	return ans
}
