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
		x, y   []int
		expect int
	}{
		{"TestCase1", []int{1, 2, 1, 3, 2}, []int{5, 3, 4, 6, 2}, 14},
		{"TestCase2", []int{1, 2, 1, 2}, []int{4, 5, 6, 7}, -1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.x, c.y)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.x, c.y)
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
