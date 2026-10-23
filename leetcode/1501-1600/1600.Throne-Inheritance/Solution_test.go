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
		kingName string
		inputs   []operation
		expect   [][]string
	}{
		{"TestCase1", "king", []operation{
			{"birth", "king", "andy"},
			{"birth", "king", "bob"},
			{"birth", "king", "catherine"},
			{"birth", "andy", "matthew"},
			{"birth", "bob", "alex"},
			{"birth", "bob", "asha"},
			{"getInheritanceOrder", "", ""},
			{"death", "", "bob"},
			{"getInheritanceOrder", "", ""},
		}, [][]string{
			{"king", "andy", "matthew", "bob", "alex", "asha", "catherine"},
			{"king", "andy", "matthew", "alex", "asha", "catherine"},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.kingName, c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.kingName, c.inputs)
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
