package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name          string
		nums          []int
		k, multiplier int
		expect        []int
	}{
		{"TestCase1", []int{2, 1, 3, 5, 6}, 5, 2, []int{8, 4, 6, 5, 6}},
		{"TestCase2", []int{1, 2}, 3, 4, []int{16, 8}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.k, c.multiplier)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.nums, c.k, c.multiplier)
			}
		})
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
