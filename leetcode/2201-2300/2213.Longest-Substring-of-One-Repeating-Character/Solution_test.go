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
		s, queryCharacter string
		queryIndex        []int
		expect            []int
	}{
		{"TestCase1", "babacc", "bcb", []int{1, 3, 3}, []int{3, 3, 4}},
		{"TestCase2", "abyzz", "aa", []int{2, 1}, []int{2, 3}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.s, c.queryCharacter, c.queryIndex)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.s, c.queryCharacter, c.queryIndex)
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
