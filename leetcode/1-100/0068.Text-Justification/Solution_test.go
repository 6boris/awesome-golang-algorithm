package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		words []string
		k int
		expect []string
	}{
		{"TestCase", []string{"This", "is", "an", "example", "of", "text", "justification."}, 16, []string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		}},
		{"TestCase", []string{"What","must","be","acknowledgment","shall","be"}, 16, []string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		}},
		{"TestCase", []string{"Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"}, 20, 
		[]string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  ",
		}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := Solution(c.words, c.k)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, ret, c.words, c.k)
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
