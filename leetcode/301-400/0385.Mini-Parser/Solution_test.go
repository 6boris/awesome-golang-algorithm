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
		inputs string
		expect *NestedInteger
	}{
		{"TestCase1", "324", &NestedInteger{Value: 324, IsNum: true}},
		{"TestCase2", "[123,[456,[789]]]", &NestedInteger{
			Child: []NestedInteger{
				{Value: 123, IsNum: true},
				{Child: []NestedInteger{
					{Value: 456, IsNum: true},
					{Child: []NestedInteger{{Value: 789, IsNum: true}}},
				}},
			},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			s1 := c.expect.serialize()
			s2 := got.serialize()
			if !reflect.DeepEqual(s1, s2) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					s1, s2, c.inputs)
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
