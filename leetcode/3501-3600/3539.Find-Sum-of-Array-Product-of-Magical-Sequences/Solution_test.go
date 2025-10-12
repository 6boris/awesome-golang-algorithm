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
		m, k   int
		nums   []int
		expect int
	}{
		{"TestCase1", 5, 5, []int{1, 10, 100, 10000, 1000000}, 991600007},
		{"TestCase2", 2, 2, []int{5, 4, 3, 2, 1}, 170},
		{"TestCase3", 1, 1, []int{28}, 28},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.m, c.k, c.nums)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.m, c.k, c.nums)
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
