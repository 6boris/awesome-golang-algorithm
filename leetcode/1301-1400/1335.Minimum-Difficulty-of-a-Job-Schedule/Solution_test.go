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
		jobs   []int
		day    int
		expect int
	}{
		{"TestCase1", []int{6, 5, 4, 3, 2, 1}, 2, 7},
		{"TestCase2", []int{9, 9, 9}, 4, -1},
		{"TestCase3", []int{1, 1, 1}, 3, 3},
		{"TestCase4", []int{100, 23, 48, 96, 72, 33, 67, 300, 235, 786, 999, 1, 0, 3, 5}, 10, 1308},
		{"TestCase5", []int{100, 23, 48, 96, 72, 33, 67, 300, 235, 786, 999, 1, 0, 3, 5}, 5, 1008},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.jobs, c.day)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.jobs, c.day)
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
