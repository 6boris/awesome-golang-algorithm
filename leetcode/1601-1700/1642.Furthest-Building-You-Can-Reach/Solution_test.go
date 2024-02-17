package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name            string
		heights         []int
		bricks, ladders int
		expect          int
	}{
		{"TestCase1", []int{4, 2, 7, 6, 9, 14, 12}, 5, 1, 4},
		{"TestCase2", []int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2, 7},
		{"TestCase3", []int{14, 3, 19, 3}, 17, 0, 3},
		{"TestCase4", []int{7, 5, 13}, 0, 0, 1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.heights, c.bricks, c.ladders)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.heights, c.bricks, c.ladders)
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
