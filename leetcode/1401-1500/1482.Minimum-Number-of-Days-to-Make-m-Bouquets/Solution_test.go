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
		b      []int
		m, k   int
		expect int
	}{
		{"TestCase1", []int{1, 10, 3, 10, 2}, 3, 1, 3},
		{"TestCase2", []int{1, 10, 3, 10, 2}, 3, 2, -1},
		{"TestCase3", []int{7, 7, 7, 7, 12, 7, 7}, 2, 3, 12},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.b, c.m, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.b, c.m, c.k)
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
