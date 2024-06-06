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
		inputs []string
		target string
		expect int
	}{
		{"TestCase1", []string{"0201", "0101", "0102", "1212", "2002"}, "0202", 6},
		{"TestCase2", []string{"8888"}, "0009", 1},
		{"TestCase3", []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888", -1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.target)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.target)
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
