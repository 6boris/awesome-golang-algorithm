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
		banned    []int
		n, maxSum int
		expect    int
	}{
		{"TestCase1", []int{1, 6, 5}, 5, 6, 2},
		{"TestCase2", []int{1, 2, 3, 4, 5, 6, 7}, 8, 1, 0},
		{"TestCase3", []int{11}, 7, 50, 7},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.banned, c.n, c.maxSum)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.banned, c.n, c.maxSum)
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
