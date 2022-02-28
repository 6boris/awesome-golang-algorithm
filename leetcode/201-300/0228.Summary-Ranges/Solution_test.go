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
		inputs []int
		expect []string
	}{
		{"TestCase1", []int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{"TestCase2", []int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{"TestCase3", []int{1, 2, 3}, []string{"1->3"}},
		{"TestCase4", []int{0, 1, 2, 4, 5, 6, 8}, []string{"0->2", "4->6", "8"}},
		{"TestCase5", []int{1, 2, 3, 5}, []string{"1->3", "5"}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
