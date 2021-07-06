package Solution

import "sort"

func Solution1(nums []int) int {
	counter := make(map[int]int);
	for _, num := range nums {
		if v, ok := counter[num]; ok {
			counter[num] = v + 1;
		} else {
			counter[num] = 1;
		}
	}
	values := make([]int, 0, len(counter))
	for _, val := range counter {
		values = append(values, val);
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)));
	n := len(nums);
	res, total := 0, 0
	for _, val := range values {
		total += val;
		res += 1;
		if total >= n / 2 {
			return res;
		}
	}
	return n;
}

func Solution2(nums []int) int {
	n, values := len(nums), make([]int, 100001);
	for _, num := range nums {
		values[num]++;
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j];
	})
	res, total := 0, 0;
	for i := 0; i < len(values) && values[i] > 0; i++ {
		total += values[i];
		res += 1;
		if total >= n / 2 {
			return res;	
		}
	}
	return n;
}