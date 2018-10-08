package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		got := reverse(123)
		want := 321
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		got := reverse(-123)
		want := -321
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-3", func(t *testing.T) {
		got := reverse(120)
		want := 21
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})
}
