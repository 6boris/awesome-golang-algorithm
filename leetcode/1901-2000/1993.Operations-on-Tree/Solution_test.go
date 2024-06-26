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
		p      []int
		opts   []op
		expect []bool
	}{
		{"TestCase1", []int{-1, 0, 0, 1, 1, 2, 2}, []op{
			{"l", 2, 2},
			{"u", 2, 3},
			{"u", 2, 2},
			{"l", 4, 5},
			{"", 0, 1},
			{"l", 0, 1},
		}, []bool{true, false, true, true, true, false}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.p, c.opts)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.p, c.opts)
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
