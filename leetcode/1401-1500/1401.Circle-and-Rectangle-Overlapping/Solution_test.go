package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name                                     string
		radius, xCenter, yCenter, x1, y1, x2, y2 int
		expect                                   bool
	}{
		{"TestCase1", 1, 0, 0, 1, -1, 3, 1, true},
		{"TestCase2", 1, 1, 1, 1, -3, 2, -1, false},
		{"TestCase3", 1, 0, 0, -1, 0, 0, 1, true},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.radius, c.xCenter, c.yCenter, c.x1, c.y1, c.x2, c.y2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v %v %v %v",
					c.expect, got, c.radius, c.xCenter, c.yCenter, c.x1, c.y1, c.x2, c.y2)
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
