package Solution

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name        string
		recipes     []string
		ingredients [][]string
		supplies    []string
		expect      []string
	}{
		{"TestCase1", []string{"bread"}, [][]string{{"yeast", "flour"}}, []string{"yeast", "flour", "conr"}, []string{"bread"}},
		{"TestCase2", []string{"bread", "sandwich"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}}, []string{"yeast", "flour", "meat"}, []string{"bread", "sandwich"}},
		{"TestCase3", []string{"bread", "sandwich", "burger"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}}, []string{"yeast", "flour", "meat"}, []string{"bread", "burger", "sandwich"}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.recipes, c.ingredients, c.supplies)
			sort.Strings(got)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.recipes, c.ingredients, c.supplies)
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
