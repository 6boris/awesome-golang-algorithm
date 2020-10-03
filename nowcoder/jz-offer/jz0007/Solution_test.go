package Solution

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	//	test cases
	cases := []struct {
		name   string
		inputs int
		expect int
	}{
		{"TestCase", 1, 1},
		{"TestCase", 2, 1},
		{"TestCase", 3, 2},
		{"TestCase", 4, 3},
		{"TestCase", 6, 8},
		{"TestCase", 7, 13},
		{"TestCase", 8, 21},
		{"TestCase", 13, 233},
		{"TestCase", 21, 10946},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := Fibonacci(c.inputs)
			ast.Equal(actual, c.expect, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}

func TestSolution2(t *testing.T) {
	ast := assert.New(t)

	//	test cases
	cases := []struct {
		name   string
		inputs int
		expect int
	}{
		{"TestCase", 1, 1},
		{"TestCase", 2, 1},
		{"TestCase", 3, 2},
		{"TestCase", 4, 3},
		{"TestCase", 6, 8},
		{"TestCase", 7, 13},
		{"TestCase", 8, 21},
		{"TestCase", 13, 233},
		{"TestCase", 21, 10946},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := Fibonacci2(c.inputs)
			ast.Equal(actual, c.expect, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}

func TestSolution3(t *testing.T) {
	ast := assert.New(t)

	//	test cases
	cases := []struct {
		name   string
		inputs int
		expect int
	}{
		{"TestCase", 1, 1},
		{"TestCase", 2, 1},
		{"TestCase", 3, 2},
		{"TestCase", 4, 3},
		{"TestCase", 6, 8},
		{"TestCase", 7, 13},
		{"TestCase", 8, 21},
		{"TestCase", 13, 233},
		{"TestCase", 21, 10946},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := Fibonacci3(c.inputs)
			ast.Equal(actual, c.expect, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}

//	Benchmark Test
func BenchmarkSolution(b *testing.B) {

}

//	Example Test
func ExampleSolution() {

}
