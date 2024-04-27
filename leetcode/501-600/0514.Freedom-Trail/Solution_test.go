package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name      string
		ring, key string
		expect    int
	}{
		{"TestCase1", "godding", "gd", 4},
		{"TestCase2", "godding", "godding", 13},
		{"TestCase3", "pqwcx", "cpqwx", 13},
		{"TestCase4", "caotmcaataijjxi", "oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx", 137},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.ring, c.key)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.ring, c.key)
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
