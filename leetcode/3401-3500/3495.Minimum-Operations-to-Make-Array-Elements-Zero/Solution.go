package Solution

func get(num int) int64 {
	var cnt int64
	i := 1
	base := 1

	for base <= num {
		end := base*2 - 1
		if end > num {
			end = num
		}
		cnt += int64((i+1)/2) * int64(end-base+1)
		i++
		base *= 2
	}
	return cnt
}

func Solution(queries [][]int) int64 {
	var res int64
	for _, q := range queries {
		count1 := get(q[1])
		count2 := get(q[0] - 1)
		res += (count1 - count2 + 1) / 2
	}
	return res
}
