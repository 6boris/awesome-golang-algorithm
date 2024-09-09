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
		w1, w2 []string
		expect []string
	}{
		{"TestCase1", []string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}, []string{"facebook", "google", "leetcode"}},
		{"TestCase2", []string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"}, []string{"apple", "google", "leetcode"}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.w1, c.w2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.w1, c.w2)
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
