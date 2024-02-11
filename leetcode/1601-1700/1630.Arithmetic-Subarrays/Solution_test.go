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
		inputs []int
		l, r   []int
		expect []bool
	}{
		{"TestCase1", []int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}, []bool{true, false, true}},
		{"TestCase2", []int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}, []int{0, 1, 6, 4, 8, 7}, []int{4, 4, 9, 7, 9, 10}, []bool{false, true, false, false, true, true}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.l, c.r)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
