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
		name        string
		x, y, bound int
		expect      []int
	}{
		{"TestCase1", 2, 3, 10, []int{2, 3, 4, 5, 7, 9, 10}},
		{"TestCase2", 3, 5, 15, []int{2, 4, 6, 8, 10, 14}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.x, c.y, c.bound)
			sort.Ints(got)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.x, c.y, c.bound)
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
