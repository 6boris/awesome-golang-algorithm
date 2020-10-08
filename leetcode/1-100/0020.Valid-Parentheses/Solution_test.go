package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := "()"
		got := isValid(data)
		want := true
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		data := "()[]{}"
		got := isValid(data)
		want := true
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-3", func(t *testing.T) {
		data := "(]"
		got := isValid(data)
		want := false
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-4", func(t *testing.T) {
		data := "([)]"
		got := isValid(data)
		want := false
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-5", func(t *testing.T) {
		data := "{[]}"
		got := isValid(data)
		want := true
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})
}
