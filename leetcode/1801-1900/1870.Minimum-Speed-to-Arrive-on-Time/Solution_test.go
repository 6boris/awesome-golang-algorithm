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
		dist   []int
		hour   float64
		expect int
	}{
		{"TestCase1", []int{1, 3, 2}, float64(6), 1},
		{"TestCase2", []int{1, 3, 2}, 2.7, 3},
		{"TestCase3", []int{1, 3, 2}, 1.9, -1},
		{"TestCase4", []int{1, 1, 100000}, 2.01, 10000000},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.dist, c.hour)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.dist, c.hour)
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
