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
		target []int
		n      int
		expect []string
	}{
		{"TestCase1", []int{1, 3}, 3, []string{"Push", "Push", "Pop", "Push"}},
		{"TestCase2", []int{1, 2, 3}, 3, []string{"Push", "Push", "Push"}},
		{"TestCase3", []int{1, 2}, 4, []string{"Push", "Push"}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.target, c.n)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.target, c.n)
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
