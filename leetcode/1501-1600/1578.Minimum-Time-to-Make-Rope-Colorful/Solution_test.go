package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name     string
		colors   string
		needTime []int
		expect   int
	}{
		{"TestCase1", "abaac", []int{1, 2, 3, 4, 5}, 3},
		{"TestCase2", "abc", []int{1, 2, 3}, 0},
		{"TestCase3", "aabaa", []int{1, 2, 3, 4, 1}, 2},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.colors, c.needTime)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.colors, c.needTime)
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
