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
		k      int
		expect []string
	}{
		{"TestCase1", []string{"i", "love", "leetcode", "i", "love", "coding"}, 2, []string{"i", "love"}},
		{"TestCase2", []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4, []string{"the", "is", "sunny", "day"}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.words, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, %d",
					c.expect, got, c.words, c.k)
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
