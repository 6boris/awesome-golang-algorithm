package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name            string
		operations, ans []string
		nums            []int
		expect          bool
	}{
		{"TestCase1",
			[]string{
				"add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "search", "search",
				"erase", "erase", "add", "search", "erase", "search", "search", "add", "search", "add", "erase",
				"erase", "add", "erase", "erase", "erase", "search", "search", "erase", "erase", "erase", "search",
				"add", "add", "add", "add", "erase", "erase", "add", "add", "erase", "erase", "erase", "search",
				"search", "add", "add", "erase", "erase", "search", "add", "search", "search", "search", "search",
				"search", "search", "search", "search", "search", "search", "search", "search"},
			[]string{
				"null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "false",
				"false", "true", "false", "null", "false", "true", "true", "false", "null", "false", "null",
				"false", "false", "null", "true", "false", "true", "false", "false", "false", "false", "true",
				"true", "null", "null", "null", "null", "false", "false", "null", "null", "false", "true", "false",
				"false", "false", "null", "null", "true", "false", "false", "null", "false", "false", "false",
				"false", "false", "true", "true", "false", "false", "false", "true", "true"},
			[]int{
				3, 14, 3, 18, 1, 18, 17, 8, 5, 16, 3, 10, 9, 14, 13, 18, 11, 16,
				1, 4, 7, 0, 5, 12, 11, 20, 17, 2, 1, 4, 17, 12, 1, 20, 7, 6, 21,
				2, 13, 16, 17, 6, 5, 10, 13, 10, 17, 0, 1, 8, 7, 4, 15, 14, 13, 0,
				13, 0, 19, 18, 1, 20, 15, 0, 21, 18}, true},
		{"TestCase2", []string{
			"add", "add", "add", "add", "add", "add",
			"erase", "erase", "erase", "erase", "erase", "erase", "erase",
			"search", "search", "search", "search", "search", "search",
		}, []string{
			"null", "null", "null", "null", "null", "null",
			"true", "true", "true", "true", "true", "true", "false",
			"false", "false", "false", "false", "false", "false",
		}, []int{5, 3, 1, 2, 2, 6, 5, 3, 1, 2, 2, 6, 7, 1, 2, 3, 4, 5, 6}, true},
		{"testCase3", []string{
			"add", "add", "add", "search", "add", "search", "erase", "erase", "search"}, []string{
			"null", "null", "null", "false", "null", "true", "false", "true", "false",
		}, []int{1, 2, 3, 0, 4, 1, 0, 1, 1}, true},
		{"TestCase4", []string{
			"add", "add", "add",
			"search", "erase", "erase", "search", "erase", "search",
		}, []string{"null", "null", "null", "true", "true", "true", "true", "true", "false"}, []int{1, 1, 1, 1, 1, 1, 1, 1, 1}, true},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.operations, c.ans, c.nums)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.operations, c.ans, c.nums)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
