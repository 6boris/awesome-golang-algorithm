package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name    string
		ss, tt  string
		maxCost int
		expect  int
	}{
		{"TestCase1", "abcd", "bcdf", 3, 3},
		{"TestCase2", "abcd", "cdef", 3, 1},
		{"TestCase3", "abcd", "acde", 0, 1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.ss, c.tt, c.maxCost)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.ss, c.tt, c.maxCost)
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
