package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name        string
		skill, mana []int
		expect      int64
	}{
		{"TestCase1", []int{1, 5, 2, 4}, []int{5, 1, 4, 2}, 110},
		{"TestCase2", []int{1, 1, 1}, []int{1, 1, 1}, 5},
		{"TestCase3", []int{1, 2, 3, 4}, []int{1, 2}, 21},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.skill, c.mana)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.skill, c.mana)
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
