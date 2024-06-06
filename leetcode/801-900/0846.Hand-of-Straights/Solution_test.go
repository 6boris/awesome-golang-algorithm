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
		hand      []int
		groupSize int
		expect    bool
	}{
		{"TestCase1", []int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3, true},
		{"TestCase2", []int{1, 2, 3, 4, 5}, 4, false},
		{"TestCase3", []int{1, 1, 2, 2, 3, 3}, 2, false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.hand, c.groupSize)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.hand, c.groupSize)
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
