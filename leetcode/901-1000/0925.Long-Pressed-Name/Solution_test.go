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
		username, typed string
		expect          bool
	}{
		{"TestCase1", "alex", "aaleex", true},
		{"TestCase2", "alex", "alex", true},
		{"TestCase3", "saeed", "ssaaedd", false},
		{"TestCase4", "alex", "aaleexa", false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.username, c.typed)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, %v",
					c.expect, got, c.username, c.typed)
			}
			got = Solution2(c.username, c.typed)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, %v",
					c.expect, got, c.username, c.typed)
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
