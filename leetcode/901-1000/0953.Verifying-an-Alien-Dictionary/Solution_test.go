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
		words  []string
		order  string
		expect bool
	}{
		{"TestCase1", []string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
		{"TestCase2", []string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false},
		{"TestCase3", []string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.words, c.order)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.words, c.order)
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
