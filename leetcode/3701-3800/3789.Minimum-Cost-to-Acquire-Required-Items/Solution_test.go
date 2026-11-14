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
		c1, c2, cb, n1, n2 int
		expect             int64
	}{
		{"TestCase1", 3, 2, 1, 3, 2, 3},
		{"TestCase2", 5, 4, 15, 2, 3, 22},
		{"TestCase3", 5, 4, 15, 0, 0, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.c1, c.c2, c.cb, c.n1, c.n2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v %v",
					c.expect, got, c.c1, c.c2, c.cb, c.n1, c.n2)

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
