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
		k, p   int
		expect int
	}{
		{"TestCase1", []int{2, 3, 3, 2, 2}, 2, 2, 11},
		{"TestCase2", []int{1, 2, 3, 4}, 4, 1, 10},
		{"TestCase3", []int{6, 20, 5, 18}, 3, 14, 10},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.k, c.p)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.nums, c.k, c.p)
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
