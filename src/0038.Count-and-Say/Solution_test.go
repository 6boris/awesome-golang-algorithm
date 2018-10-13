package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		got := countAndSay(1)
		want := "1"
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		got := countAndSay(4)
		want := "1211"
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

}
