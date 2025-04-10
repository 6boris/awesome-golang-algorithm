package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name          string
		start, finish int64
		limit         int
		s             string
		expect        int64
	}{
		{"TestCase1", 1, 6000, 4, "124", 5},
		{"TestCase2", 15, 215, 6, "10", 2},
		{"TestCase3", 1000, 2000, 4, "3000", 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.start, c.finish, c.limit, c.s)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.start, c.finish, c.limit, c.s)
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
