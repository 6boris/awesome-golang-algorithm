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
		messages []string
		senders  []string
		expect   string
	}{
		{"TestCase1", []string{"Hello userTwooo", "Hi userThree", "Wonderful day Alice", "Nice day userThree"}, []string{"Alice", "userTwo", "userThree", "Alice"}, "Alice"},
		{"TestCase2", []string{"How is leetcode for everyone", "Leetcode is useful for practice"}, []string{"Bob", "Charlie"}, "Charlie"},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.messages, c.senders)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.messages, c.senders)
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
