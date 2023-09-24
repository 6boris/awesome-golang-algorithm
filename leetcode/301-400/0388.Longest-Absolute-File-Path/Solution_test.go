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
		expect int
	}{
		{"TestCase1", "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext", 20},
		{"TestCase2", "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext", 32},
		{"TestCase3", "a", 0},
		{"TestCase4", "a.txt", 5},
		{"TestCase5", "file1.txt\nfile2.txt\tlongfile.txt", 22},
		{"TestCase6", "a\n\tb.txt\na2\n\tb2.txt", 9},
		{"TestCase7", "a\n\taa\n\t\taaa\n\t\t\tfile1.txt\naaaaaaaaaaaaaaaaaaaaa\n\tsth.png", 29},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
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
