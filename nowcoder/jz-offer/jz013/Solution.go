package Solution

import "fmt"

func reOrderArray(arr []int) []int {
	ans := make([]int, 0, 0)
	for _, v := range arr {
		if v&1 != 0 {
			fmt.Println(v)
			ans = append(ans, v)
		}
	}
	for _, v := range arr {
		if v&1 == 0 {
			fmt.Println(v)

			ans = append(ans, v)
		}
	}
	return ans
}

func reOrderArray2(arr []int) []int {
	i := 0

	for j := 0; j < len(arr); j++ {
		if arr[j]&1 == 1 {
			tmp := arr[j]
			for k := j - 1; k >= i; k-- {
				arr[k+1] = arr[k]
			}
			arr[i] = tmp
			i += 1
		}

	}

	return arr
}
