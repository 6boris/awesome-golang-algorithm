package Solution

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect []int
	}{
		{"TestCase1", []int{0, 1, 1, 0}, []int{0, 1}},
		{"TestCase2", []int{0, 3, 2, 1, 3, 2}, []int{2, 3}},
		{"TestCase3", []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}, []int{4, 5}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			sort.Ints(got)
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
