package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name      string
		startTime []int
		endTime   []int
		queryTime int
		expect    int
	}{
		{"TestCase", []int{1, 2, 3}, []int{3, 2, 7}, 4, 1},
		{"TestCase", []int{4}, []int{4}, 4, 1},
		{"TestCase", []int{4}, []int{4}, 5, 0},
		{"TestCase", []int{1, 1, 1, 1}, []int{1, 3, 2, 4}, 7, 0},
		{"TestCase", []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10}, 5, 5},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.startTime, c.endTime, c.queryTime)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.startTime, c.endTime, c.queryTime)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
