package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name               string
		eventTime, k       int
		startTime, endTime []int
		expect             int
	}{
		{"TestCase1", 5, 1, []int{1, 3}, []int{2, 5}, 2},
		{"TestCase2", 10, 1, []int{0, 2, 9}, []int{1, 4, 10}, 6},
		{"TestCase3", 5, 2, []int{0, 1, 2, 3, 4}, []int{1, 2, 3, 4, 5}, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.eventTime, c.k, c.startTime, c.endTime)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.eventTime, c.k, c.startTime, c.endTime)
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
