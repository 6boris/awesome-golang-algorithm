package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name    string
		k       int
		arrival []int
		load    []int
		expect  []int
	}{
		{"TestCase1", 3, []int{1, 2, 3, 4, 5}, []int{5, 2, 3, 3, 3}, []int{1}},
		{"TestCase2", 3, []int{1, 2, 3, 4}, []int{1, 2, 1, 2}, []int{0}},
		{"TestCase3", 3, []int{1, 2, 3}, []int{10, 12, 11}, []int{0, 1, 2}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.k, c.arrival, c.load)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.k, c.arrival, c.load)
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
