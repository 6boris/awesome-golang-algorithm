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
		eventTime          int
		startTime, endTime []int
		expect             int
	}{
		{"TestCase1", 5, []int{1, 3}, []int{2, 5}, 2},
		{"TestCase2", 10, []int{0, 7, 9}, []int{1, 8, 10}, 7},
		{"TestCase3", 10, []int{0, 3, 7, 9}, []int{1, 4, 8, 10}, 6},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.eventTime, c.startTime, c.endTime)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.eventTime, c.startTime, c.endTime)
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
