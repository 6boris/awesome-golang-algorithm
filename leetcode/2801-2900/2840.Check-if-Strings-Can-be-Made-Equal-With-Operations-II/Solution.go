package Solution

import (
	"slices"
)

func Solution(s1 string, s2 string) bool {
	a1, b1 := make([]byte, 0), make([]byte, 0)
	n := len(s1)
	for i := 0; i < n; i += 2 {
		a1 = append(a1, s1[i])
		b1 = append(b1, s2[i])
	}

	slices.SortFunc(a1, func(i, j byte) int {
		if i < j {
			return -1
		}
		if i == j {
			return 0
		}
		return 1
	})
	slices.SortFunc(b1, func(i, j byte) int {
		if i < j {
			return -1
		}
		if i == j {
			return 0
		}
		return 1

	})

	for i := range a1 {
		if a1[i] != b1[i] {
			return false
		}
	}

	a2, b2 := make([]byte, 0), make([]byte, 0)
	for i := 1; i < n; i += 2 {
		a2 = append(a2, s1[i])
		b2 = append(b2, s2[i])
	}
	slices.SortFunc(a2, func(i, j byte) int {
		if i < j {
			return -1
		}
		if i == j {
			return 0
		}
		return 1

	})
	slices.SortFunc(b2, func(i, j byte) int {
		if i < j {
			return -1
		}
		if i == j {
			return 0
		}
		return 1

	})
	for i := range a2 {
		if a2[i] != b2[i] {
			return false
		}
	}
	return true
}
