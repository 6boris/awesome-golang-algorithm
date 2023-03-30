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
		homepage string
		ops      []Op
		expect   []string
	}{
		{"TestCase1", "leetcode.com", []Op{
			{name: "visit", url: "google.com"},
			{name: "visit", url: "facebook.com"},
			{name: "visit", url: "youtube.com"},
			{name: "back", steps: 1},
			{name: "back", steps: 1},
			{name: "forward", steps: 1},
			{name: "visit", url: "linkedin.com"},
			{name: "forward", steps: 2},
			{name: "back", steps: 2},
			{name: "back", steps: 7},
		}, []string{"facebook.com", "google.com", "facebook.com", "linkedin.com", "google.com", "leetcode.com"}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.homepage, c.ops)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.homepage, c.ops)
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
