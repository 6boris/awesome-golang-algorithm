package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name             string
		zero, one, limit int
		expect           int
	}{
		{"TestCase1", 1, 1, 2, 2},
		{"TestCase2", 1, 2, 1, 1},
		{"TestCase3", 3, 3, 2, 14},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.zero, c.one, c.limit)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.zero, c.one, c.limit)
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
