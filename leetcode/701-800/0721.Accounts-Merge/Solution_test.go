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
		name   string
		inputs [][]string
		expect [][]string
	}{
		{"TestCase1", [][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}},
			[][]string{{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"}, {"John", "johnnybravo@mail.com"}, {"Mary", "mary@mail.com"}}},
		{"TestCase2", [][]string{
			{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
			{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
			{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
			{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
			{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
		}, [][]string{
			{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"},
			{"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"},
			{"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"},
			{"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"},
			{"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"},
		}},
		{"TestCase3", [][]string{
			{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"},
			{"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"},
			{"Kevin", "Kevin4@m.co", "Kevin2@m.co", "Kevin2@m.co"},
			{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe2@m.co"},
			{"Gabe", "Gabe3@m.co", "Gabe4@m.co", "Gabe2@m.co"},
		}, [][]string{
			{"Alex", "Alex0@m.co", "Alex4@m.co", "Alex5@m.co"},
			{"Ethan", "Ethan0@m.co", "Ethan3@m.co"},
			{"Gabe", "Gabe0@m.co", "Gabe2@m.co", "Gabe3@m.co", "Gabe4@m.co"},
			{"Kevin", "Kevin2@m.co", "Kevin4@m.co"},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			sort.Slice(got, func(i, j int) bool {
				a := got[i][0]
				b := got[j][0]
				if a == b {
					return len(got[i]) > len(got[j])
				}
				return a < b
			})
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
