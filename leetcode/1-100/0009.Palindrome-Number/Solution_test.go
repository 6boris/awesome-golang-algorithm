package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		got := isPalindrome(121121)
		want := true
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		got := isPalindrome(-121)
		want := false
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-3", func(t *testing.T) {
		got := isPalindrome(10)
		want := false
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})
}
