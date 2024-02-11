package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name        string
		dist, speed []int
		expect      int
	}{
		{"TestCase1", []int{1, 3, 4}, []int{1, 1, 1}, 3},
		{"TestCase2", []int{1, 1, 2, 3}, []int{1, 1, 1, 1}, 1},
		{"TestCase3", []int{3, 2, 4}, []int{5, 3, 2}, 1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.dist, c.speed)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.dist, c.speed)
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
