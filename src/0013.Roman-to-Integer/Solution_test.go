package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		got := romanToInt("III")
		want := 3
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		got := romanToInt("IV")
		want := 4
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-3", func(t *testing.T) {
		got := romanToInt("IX")
		want := 9
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-4", func(t *testing.T) {
		got := romanToInt("LVIII")
		want := 58
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-5", func(t *testing.T) {
		got := romanToInt("MCMXCIV")
		want := 1994
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

}
