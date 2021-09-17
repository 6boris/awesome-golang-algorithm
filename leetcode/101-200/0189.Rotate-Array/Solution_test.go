package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		nums   []int
		k      int
		expect []int
	}{
		{"TestCase", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{"TestCase", []int{1, 2, 3, 4, 5, 6, 7}, 7, []int{1, 2, 3, 4, 5, 6, 7}},
		{"TestCase", []int{1, 2, 3, 4, 5, 6, 7}, 21, []int{1, 2, 3, 4, 5, 6, 7}},
		{"TestCase", []int{1, 2, 3, 4, 5, 6, 7}, 6, []int{2, 3, 4, 5, 6, 7, 1}},
		{"TestCase", []int{1, 2, 3, 4, 5, 6, 7}, 20, []int{2, 3, 4, 5, 6, 7, 1}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			tmp := c.nums
			Solution(c.nums, c.k)
			if !reflect.DeepEqual(c.nums, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, c.nums, tmp, c.k)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
