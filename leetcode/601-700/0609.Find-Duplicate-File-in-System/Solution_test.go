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
		inputs []string
		expect [][]string
	}{
		{"TestCase1", []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}, [][]string{
			{"root/a/1.txt", "root/c/3.txt"},
			{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"},
		}},
		{"TestCase2", []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)"}, [][]string{
			{"root/a/1.txt", "root/c/3.txt"},
			{"root/a/2.txt", "root/c/d/4.txt"},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			sort.Slice(got, func(i, j int) bool {
				a, b := len(got[i]), len(got[j])
				if a == b {
					for ii := 0; ii < a; ii++ {
						if got[i][ii] == got[j][ii] {
							continue
						}
						return got[i][ii] < got[j][ii]
					}
					return true
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
