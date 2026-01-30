package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name              string
		source, target    string
		original, changed []string
		cost              []int
		expect            int64
	}{
		{"TestCase1", "abcd", "acbe", []string{"a", "b", "c", "c", "e", "d"}, []string{"b", "c", "b", "e", "b", "e"}, []int{2, 5, 5, 1, 2, 20}, 28},
		{"TestCase2", "abcdefgh", "acdeeghh", []string{"bcd", "fgh", "thh"}, []string{"cde", "thh", "ghh"}, []int{1, 3, 5}, 9},
		{"TestCase3", "abcdefgh", "addddddd", []string{"bcd", "defgh"}, []string{"ddd", "ddddd"}, []int{100, 1578}, -1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.source, c.target, c.original, c.changed, c.cost)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v %v",
					c.expect, got, c.source, c.target, c.original, c.changed, c.cost)
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
