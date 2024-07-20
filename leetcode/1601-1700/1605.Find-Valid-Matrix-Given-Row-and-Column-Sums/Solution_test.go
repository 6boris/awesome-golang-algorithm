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
		a, b   []int
		expect [][]int
	}{
		{"TestCase1", []int{3, 8}, []int{4, 7}, [][]int{{3, 0}, {1, 7}}},
		{"TestCase2", []int{5, 7, 10}, []int{8, 6, 8}, [][]int{{5, 0, 0}, {3, 4, 0}, {0, 2, 8}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.a, c.b)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.a, c.b)
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
