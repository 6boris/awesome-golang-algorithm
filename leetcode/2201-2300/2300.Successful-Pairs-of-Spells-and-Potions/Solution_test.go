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
		s, p    []int
		success int64
		expect  []int
	}{
		{"TestCase1", []int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7, []int{4, 0, 3}},
		{"TestCase2", []int{3, 1, 2}, []int{8, 5, 8}, 16, []int{2, 0, 2}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.s, c.p, c.success)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.s, c.p, c.success)
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
