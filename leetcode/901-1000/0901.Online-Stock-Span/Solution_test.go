package Solution

import (
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect []int
	}{
		{"TestCase1", []int{100, 80, 60, 70, 60, 75, 85}, []int{1, 1, 1, 2, 1, 4, 6}},
		{"TestCase2", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			u := Constructor()
			for idx, n := range c.inputs {
				if r := u.Next(n); r != c.expect[idx] {
					t.Fatalf("expected: %v, but got: %v, with inputs: %v", c.expect[idx], r, n)
				}
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
