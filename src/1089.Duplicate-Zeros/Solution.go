package Solution

func Solution1(arr []int) {
	n := len(arr)
	if n > 1 {
		for i := 0; i < n; i++ {
			if arr[i] == 0 {
				for j := n - 1; j > i; j-- {
					arr[j] = arr[j-1]
				}
				i++
			}
		}
	}
}

func Solution2(arr []int) {
	n := len(arr)
	newArr := make([]int, 0)
	c := 0
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			newArr = append(newArr, 0)
			newArr = append(newArr, 0)
			c += 2
		} else {
			newArr = append(newArr, arr[i])
			c++
		}
		if c == n {
			break
		}
	}
	arr = newArr
}
