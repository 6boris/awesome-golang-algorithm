package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name                                   string
		ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int
		expect                                 int
	}{
		{"TestCase1", -3, 0, 3, 4, 0, -1, 9, 2, 45},
		{"TestCase2", -2, -2, 2, 2, 3, 3, 4, 4, 17},
		{"TestCase3", -5, -5, 0, -4, -3, -3, 3, 3, 41},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.ax1, c.ay1, c.ax2, c.ay2, c.bx1, c.by1, c.bx2, c.by2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v",
					c.expect, got)
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
