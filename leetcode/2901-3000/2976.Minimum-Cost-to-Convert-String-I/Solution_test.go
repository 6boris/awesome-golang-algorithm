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
		original, changed []byte
		cost              []int
		expect            int64
	}{
		{"TestCase1", "abcd", "acbe", []byte("abcced"), []byte("bcbebe"), []int{2, 5, 5, 1, 2, 20}, int64(28)},
		{"TestCase2", "aaaa", "bbbb", []byte("ac"), []byte("cb"), []int{1, 2}, int64(12)},
		{"TestCase3", "abcd", "abce", []byte("a"), []byte("e"), []int{10000}, -1},
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
