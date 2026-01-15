package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name         string
		n, m         int
		hBars, vBars []int
		expect       int
	}{
		{"TestCase1", 2, 1, []int{2, 3}, []int{2}, 4},
		{"TestCase2", 1, 1, []int{2}, []int{2}, 4},
		{"TestCase3", 2, 3, []int{2, 3}, []int{2, 4}, 4},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.m, c.hBars, c.vBars)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.n, c.m, c.hBars, c.vBars)
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
